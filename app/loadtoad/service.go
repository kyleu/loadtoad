package loadtoad

import (
	"github.com/kyleu/loadtoad/app/lib/filesystem"
	"github.com/kyleu/loadtoad/app/lib/websocket"
	"github.com/kyleu/loadtoad/app/util"
)

type Service struct {
	FS     *filesystem.FileSystem
	Socket *websocket.Service
}

func NewService(ws *websocket.Service, logger util.Logger) *Service {
	p := util.GetEnv("loadtoad_path", "./data")
	fs := filesystem.NewFileSystem(p)
	if !fs.Exists(".") {
		logger.Infof("creating data directory at [%s]", p)
		_ = fs.CreateDirectory(".")
	}
	ret := &Service{FS: fs, Socket: ws}
	return ret
}
