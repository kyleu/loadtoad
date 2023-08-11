package har

import (
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"strings"
)

type Selector struct {
	Har  string `json:"har,omitempty"`
	URL  string `json:"url,omitempty"`
	Mime string `json:"mime,omitempty"`
	Idx  int    `json:"idx,omitempty"`
}

type Selectors []*Selector

func (e Entries) Find(s *Selector) (Entries, error) {
	matches := func(k string, v string) bool {
		pre, suff := strings.HasPrefix(k, "*"), strings.HasSuffix(k, "*")
		k = strings.TrimSuffix(strings.TrimPrefix(k, "*"), "*")
		if pre && suff {
			return strings.Contains(v, k)
		}
		if pre {
			return strings.HasSuffix(v, k)
		}
		if suff {
			return strings.HasPrefix(v, k)
		}
		return k == v
	}
	ret := slices.Clone(e)
	if s.URL != "" && s.URL != "*" {
		ret = lo.Filter(ret, func(e *Entry, _ int) bool {
			return matches(s.URL, e.Request.URL)
		})
	}
	if s.Mime != "" && s.Mime != "*" {
		ret = lo.Filter(ret, func(e *Entry, _ int) bool {
			return matches(s.Mime, e.Response.Content.MimeType)
		})
	}
	if s.Idx > 0 {
		if s.Idx > len(ret) {
			return nil, errors.Errorf("index [%d] does not exist among [%d] entries", s.Idx-1, len(ret))
		}
		return Entries{ret[s.Idx-1]}, nil
	}
	return ret, nil
}
