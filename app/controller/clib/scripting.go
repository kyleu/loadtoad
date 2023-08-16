// Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/views/vscripting"
)

func ScriptingList(rc *fasthttp.RequestCtx) {
	controller.Act("scripting.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Scripting"
		ret := as.Services.Script.ListScripts(ps.Logger)
		ps.Data = ret
		return controller.Render(rc, as, &vscripting.List{Scripts: ret}, ps, "scripting")
	})
}

var examples = []string{"a", "b", "c"}

func ScriptingDetail(rc *fasthttp.RequestCtx) {
	controller.Act("scripting.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		src, err := as.Services.Script.LoadScript(key, ps.Logger)
		if err != nil {
			return "", err
		}
		res := make(map[string]string, len(examples))
		for _, ex := range examples {
			x, err := as.Services.Script.RunScript(src, "test", ex)
			if err != nil {
				return "", err
			}
			res[ex] = fmt.Sprintf("%v", x)
		}
		ps.Title = "Scripting"
		ps.Data = map[string]any{"script": src, "results": res}
		return controller.Render(rc, as, &vscripting.Detail{Path: key, Script: src, Results: res}, ps, "scripting", key)
	})
}
