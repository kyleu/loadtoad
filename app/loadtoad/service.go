package loadtoad

import (
	"github.com/kyleu/loadtoad/app/lib/filesystem"
	"github.com/kyleu/loadtoad/app/lib/websocket"
	"github.com/kyleu/loadtoad/app/util"
)

type Service struct {
	FS     filesystem.FileLoader
	Socket *websocket.Service
}

func NewService(fs filesystem.FileLoader, ws *websocket.Service, logger util.Logger) *Service {
	ret := &Service{FS: fs, Socket: ws}
	return ret
}
