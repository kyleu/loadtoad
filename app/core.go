// Package app - Content managed by Project Forge, see [projectforge.md] for details.
package app

import (
	"context"

	"github.com/kyleu/loadtoad/app/lib/exec"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/lib/scripting"
	"github.com/kyleu/loadtoad/app/lib/websocket"
	"github.com/kyleu/loadtoad/app/util"
)

type CoreServices struct {
	Har    *har.Service
	Exec   *exec.Service
	Script *scripting.Service
	Socket *websocket.Service
}

//nolint:revive
func initCoreServices(ctx context.Context, st *State, logger util.Logger) CoreServices {
	return CoreServices{
		Har:    har.NewService(st.Files),
		Exec:   exec.NewService(),
		Script: scripting.NewService(st.Files, "scripts"),
		Socket: websocket.NewService(nil, nil),
	}
}
