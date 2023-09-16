package loadtoad

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"slices"
	"strings"

	"github.com/dop251/goja"
	"github.com/samber/lo"

	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
)

func (s *Service) RunWorkflow(
	ctx context.Context, w *Workflow, repls map[string]string, logger util.Logger,
	logF func(i int, s string), errF func(i int, e error), okF func(i int, result *WorkflowResult),
) (WorkflowResults, error) {
	var ret WorkflowResults
	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}} //nolint:gosec
	entries, err := s.LoadEntries(w.Tests...)
	if err != nil {
		return nil, err
	}
	vms, err := s.LoadScripts(w.Scripts, logger)
	if err != nil {
		return nil, err
	}
	repls, err = extractReplacements(entries, vms, repls)
	if err != nil {
		return nil, err
	}
	vars := w.Variables.Clone()

	var hot []string
	for i, e := range entries {
		e, err = updatedEntry(e, vms)
		if err != nil {
			return nil, err
		}
		var wr *WorkflowResult
		wr, err = s.RunEntry(ctx, w.ID, i, e, client, jar, hot, repls, vars, logF)
		if err == nil {
			for _, vm := range vms {
				newVars, err2 := scriptExtractVariables(vm, wr)
				if err2 != nil {
					return nil, err2
				}
				if len(newVars) > 0 {
					title := util.StringPlural(len(newVars), "new variable")
					keys := strings.Join(lo.Keys(newVars), ", ")
					logF(i, fmt.Sprintf("observed %s: %s", title, keys))
					vars = vars.Merge(newVars)
				}
			}
			if okF != nil {
				okF(i, wr)
			}
			if !slices.Contains(hot, wr.Domain) {
				hot = append(hot, wr.Domain)
			}
			ret = append(ret, wr)
		} else {
			if errF == nil {
				return nil, err
			}
			errF(i, err)
		}
	}
	return ret, nil
}

func extractReplacements(entries har.Entries, vms map[string]*goja.Runtime, repls map[string]string) (map[string]string, error) {
	if repls == nil {
		repls = map[string]string{}
	}
	for _, e := range entries {
		for _, vm := range vms {
			newRepls, err2 := scriptExtractReplacements(vm, e)
			if err2 != nil {
				return nil, err2
			}
			if len(newRepls) > 0 {
				for k, v := range newRepls {
					repls[k] = v
				}
			}
		}
	}
	return repls, nil
}
