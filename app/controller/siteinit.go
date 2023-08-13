// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
)

// Initialize system dependencies for the marketing site.
func initSite(*app.State, util.Logger) {
}

// Configure marketing site data for each request.
func initSiteRequest(*app.State, *cutil.PageState) error {
	return nil
}
