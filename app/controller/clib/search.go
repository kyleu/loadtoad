// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/search"
	"github.com/kyleu/loadtoad/views/vsearch"
)

func Search(rc *fasthttp.RequestCtx) {
	controller.Act("search", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := string(rc.URI().QueryArgs().Peek("q"))
		params := &search.Params{Q: q, PS: ps.Params}
		results, errs := search.Search(ps.Context, params, as, ps)
		ps.Title = "Search Results"
		if q != "" {
			ps.Title = fmt.Sprintf("[%s] %s", q, ps.Title)
		}
		ps.Data = results
		ps.DefaultNavIcon = "search"
		bc := []string{"Search||/search"}
		if q != "" {
			bc = append(bc, q)
		}
		return controller.Render(rc, as, &vsearch.Results{Params: params, Results: results, Errors: errs}, ps, bc...)
	})
}
