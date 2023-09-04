package loadtoad

import (
	"context"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/loadtoad/app/lib/filesystem"
	"github.com/kyleu/loadtoad/app/lib/filter"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/lib/search/result"
	"github.com/kyleu/loadtoad/app/util"
)

func (s *Service) ListHars(logger util.Logger) []string {
	return s.FS.ListExtension("./har", "har", nil, true, logger)
}

func (s *Service) LoadHar(fn string) (*har.Log, error) {
	key := fn
	if !strings.HasSuffix(fn, har.Ext) {
		fn += har.Ext
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
	ret := &har.Wrapper{}
	err = util.FromJSON(b, ret)
	if err != nil {
		return nil, errors.Wrapf(err, "error decoding file [%s]", fn)
	}
	ret.Log.Key = strings.TrimSuffix(key, har.Ext)
	ret.Log.Entries = ret.Log.Entries.Trimmed().WithJSON()
	return ret.Log, nil
}

func (s *Service) DeleteHar(key string, logger util.Logger) error {
	fn := key
	if !strings.HasSuffix(fn, har.Ext) {
		fn += har.Ext
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
	if !strings.HasSuffix(fn, har.Ext) {
		fn += har.Ext
	}
	if !strings.Contains(fn, "har/") {
		fn = path.Join("har", fn)
	}
	return s.FS.WriteFile(fn, b, filesystem.DefaultMode, true)
}

func (s *Service) SearchHars(ctx context.Context, ps filter.ParamSet, q string, logger util.Logger) (result.Results, error) {
	return lo.FilterMap(s.ListHars(logger), func(fn string, _ int) (*result.Result, bool) {
		log, err := s.LoadHar(fn)
		if err != nil {
			logger.Warnf("error loading har [%s]: %+v", fn, err)
			return nil, false
		}
		res := result.NewResult("archive", log.Key, log.WebPath(), log.Key, "book", log, log, q)
		if len(res.Matches) > 0 {
			return res, true
		}
		return nil, false
	}), nil
}
