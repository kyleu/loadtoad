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

func NewService(ws *websocket.Service) *Service {
	fs := filesystem.NewFileSystem(util.GetEnv("loadtoad_path", "./tmp"))
	ret := &Service{FS: fs, Socket: ws}
	return ret
}
