package main // import github.com/kyleu/loadtoad

import (
	"context"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/cmd"
)

var (
	version = "0.3.4" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(context.Background(), &app.BuildInfo{Version: version, Commit: commit, Date: date})
}
