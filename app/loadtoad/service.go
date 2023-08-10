package loadtoad

import (
	"github.com/kyleu/loadtoad/app/lib/filesystem"
	"github.com/kyleu/loadtoad/app/lib/websocket"
)

type Service struct {
	FS     *filesystem.FileSystem
	Socket *websocket.Service
}

func NewService(ws *websocket.Service) *Service {
	fs := filesystem.NewFileSystem("./tmp")
	return &Service{FS: fs, Socket: ws}
}
