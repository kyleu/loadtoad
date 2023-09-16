package app

import (
	"context"

	"github.com/kyleu/loadtoad/app/lib/exec"
	"github.com/kyleu/loadtoad/app/lib/filesystem"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/lib/scripting"
	"github.com/kyleu/loadtoad/app/lib/websocket"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
)

type Services struct {
	Exec     *exec.Service
	Socket   *websocket.Service
	Script   *scripting.Service
	Har      *har.Service
	LoadToad *loadtoad.Service
}

func NewServices(_ context.Context, _ *State, logger util.Logger) (*Services, error) {
	p := util.GetEnv("loadtoad_path", "./data")
	fs, err := filesystem.NewFileSystem(p, false, "")
	if err != nil {
		return nil, err
	}
	if !fs.Exists(".") {
		logger.Infof("creating data directory at [%s]", p)
		_ = fs.CreateDirectory(".")
	}

	ex := exec.NewService()
	ws := websocket.NewService(nil, nil, nil)
	sc := scripting.NewService(fs, "scripts")
	hs := har.NewService(fs)
	lt := loadtoad.NewService(fs, hs, ws, sc)
	return &Services{Exec: ex, Socket: ws, Script: sc, Har: hs, LoadToad: lt}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
