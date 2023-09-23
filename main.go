// Content managed by Project Forge, see [projectforge.md] for details.
package main // import github.com/kyleu/loadtoad

import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/cmd"
)

var (
	version = "0.0.15" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(&app.BuildInfo{Version: version, Commit: commit, Date: date})
}
