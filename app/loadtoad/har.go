package loadtoad

import (
	"github.com/kyleu/loadtoad/app/lib/filesystem"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/pkg/errors"
	"path"
	"strings"
)

func (s *Service) ListHars(logger util.Logger) []string {
	return s.FS.ListExtension("./har", "har", nil, false, logger)
}

func (s *Service) LoadHar(fn string) (*har.Log, error) {
	key := fn
	if !strings.HasSuffix(fn, ".har") {
		fn += ".har"
	}
	if !strings.Contains(fn, "har/") {
		fn = path.Join("har", fn)
	}
	if !s.FS.Exists(fn) {
		return nil, errors.Errorf("missing file [%s]", fn)
	}
	b, err := s.FS.ReadFile(fn)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading file [%s]", fn)
	}
	ret := &har.HarWrapper{}
	err = util.FromJSON(b, ret)
	if err != nil {
		return nil, errors.Wrapf(err, "error decoding file [%s]", fn)
	}
	ret.Log.Key = key
	ret.Log.Entries = ret.Log.Entries.Trimmed()
	return ret.Log, nil
}

func (s *Service) SaveHar(fn string, b []byte) error {
	if !strings.HasSuffix(fn, ".har") {
		fn += ".har"
	}
	if !strings.Contains(fn, "har/") {
		fn = path.Join("har", fn)
	}
	return s.FS.WriteFile(fn, b, filesystem.DefaultMode, true)
}
