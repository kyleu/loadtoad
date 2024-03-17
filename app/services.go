package app // import github.com/kyleu/loadtoad

import (
	"context"

	"github.com/kyleu/loadtoad/app/lib/filesystem"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
)

type Services struct {
	CoreServices

	LoadToad *loadtoad.Service
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	p := util.GetEnv("loadtoad_path", "./data")
	fs, err := filesystem.NewFileSystem(p, false, "")
	if err != nil {
		return nil, err
	}
	if !fs.Exists(".") {
		logger.Infof("creating data directory at [%s]", p)
		_ = fs.CreateDirectory(".")
	}

	core := initCoreServices(ctx, st, logger)

	lt := loadtoad.NewService(fs, core.Har, core.Socket, core.Script)
	return &Services{CoreServices: core, LoadToad: lt}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
