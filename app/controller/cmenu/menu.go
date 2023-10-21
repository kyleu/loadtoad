// Package cmenu - Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import (
	"context"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/lib/filter"
	"github.com/kyleu/loadtoad/app/lib/menu"
	"github.com/kyleu/loadtoad/app/lib/user"
	"github.com/kyleu/loadtoad/app/util"
)

func MenuFor(
	ctx context.Context, isAuthed bool, isAdmin bool, profile *user.Profile, params filter.ParamSet, as *app.State, logger util.Logger, //nolint:revive
) (menu.Items, any, error) {
	var ret menu.Items
	var data any
	// $PF_SECTION_START(routes)$
	ret = append(ret, harMenu(as.Services.Har), menu.Separator, loadtoadMenu(as.Services.LoadToad, logger))
	admin := &menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"}
	if len(as.Services.Script.ListScripts(logger)) > 0 {
		ret = append(ret, menu.Separator, scriptingMenu(as.Services.Script, logger))
	}
	if len(as.Services.Exec.Execs) > 0 {
		ret = append(ret, menu.Separator, processMenu(as.Services.Exec.Execs))
	}
	ret = append(ret, menu.Separator, admin)
	const aboutDesc = "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(routes)$
	return ret, data, nil
}
