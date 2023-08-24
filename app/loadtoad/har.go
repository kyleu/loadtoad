package loadtoad

import (
	har2 "github.com/kyleu/loadtoad/app/lib/har"
	"path"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/loadtoad/app/lib/filesystem"
	"github.com/kyleu/loadtoad/app/util"
)

func (s *Service) ListHars(logger util.Logger) []string {
	return s.FS.ListExtension("./har", "har", nil, true, logger)
}

func (s *Service) LoadHar(fn string) (*har2.Log, error) {
	key := fn
	if !strings.HasSuffix(fn, har2.Ext) {
		fn += har2.Ext
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
	ret := &har2.Wrapper{}
	err = util.FromJSON(b, ret)
	if err != nil {
		return nil, errors.Wrapf(err, "error decoding file [%s]", fn)
	}
	ret.Log.Key = strings.TrimSuffix(key, har2.Ext)
	ret.Log.Entries = ret.Log.Entries.Trimmed()
	return ret.Log, nil
}

func (s *Service) DeleteHar(key string, logger util.Logger) error {
	fn := key
	if !strings.HasSuffix(fn, har2.Ext) {
		fn += har2.Ext
	}
	if !strings.Contains(fn, "har/") {
		fn = path.Join("har", fn)
	}
	if !s.FS.Exists(fn) {
		return errors.Errorf("missing file [%s]", fn)
	}
	err := s.FS.Remove(fn, logger)
	if err != nil {
		return errors.Wrapf(err, "error deleting file [%s]", fn)
	}
	return nil
}

func (s *Service) SaveHar(fn string, b []byte) error {
	if !strings.HasSuffix(fn, har2.Ext) {
		fn += har2.Ext
	}
	if !strings.Contains(fn, "har/") {
		fn = path.Join("har", fn)
	}
	return s.FS.WriteFile(fn, b, filesystem.DefaultMode, true)
}
