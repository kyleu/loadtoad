package main // import github.com/kyleu/loadtoad

import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/cmd"
)

var (
	version = "0.2.12" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(&app.BuildInfo{Version: version, Commit: commit, Date: date})
}
