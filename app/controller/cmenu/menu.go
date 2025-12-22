package cmenu

import (
	"context"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/lib/filter"
	"github.com/kyleu/loadtoad/app/lib/menu"
	"github.com/kyleu/loadtoad/app/lib/telemetry"
	"github.com/kyleu/loadtoad/app/lib/user"
	"github.com/kyleu/loadtoad/app/util"
)

func MenuFor(
	ctx context.Context, as *app.State, isAuthed bool, isAdmin bool, profile *user.Profile, params filter.ParamSet, logger util.Logger,
) (menu.Items, any, error) {
	ctx, sp, _ := telemetry.StartSpan(ctx, "menu", logger)
	defer sp.Complete()
	var ret menu.Items
	var data any
	// $PF_SECTION_START(menu)$
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
	// $PF_SECTION_END(menu)$
	return ret, data, nil
}
