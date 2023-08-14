package app

import (
	"context"

	"github.com/kyleu/loadtoad/app/lib/websocket"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
)

type Services struct {
	LoadToad *loadtoad.Service
	Socket   *websocket.Service
}

func NewServices(_ context.Context, _ *State, logger util.Logger) (*Services, error) {
	ws := websocket.NewService(nil, nil, nil)
	lt := loadtoad.NewService(ws, logger)
	return &Services{LoadToad: lt, Socket: ws}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
