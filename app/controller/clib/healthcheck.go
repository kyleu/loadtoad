// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
)

func Healthcheck(rc *fasthttp.RequestCtx) {
	x := util.ValueMap{"status": "OK"}
	_, _ = cutil.RespondJSON(rc, "", x)
}
