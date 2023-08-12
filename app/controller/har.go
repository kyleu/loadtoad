package controller

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/vhar"
	"github.com/kyleu/loadtoad/views/vworkflow"
)

func HarList(rc *fasthttp.RequestCtx) {
	Act("har.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := as.Services.LoadToad.ListHars(ps.Logger)
		ps.Data = ret
		ps.Title = "Archives"
		return Render(rc, as, &vhar.List{Hars: ret}, ps, "har")
	})
}

func HarDetail(rc *fasthttp.RequestCtx) {
	Act("har.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.LoadToad.LoadHar(key)
		if err != nil {
			return "", err
		}
		ps.Title = "Archive [" + key + "]"
		ps.Data = ret
		return Render(rc, as, &vhar.Detail{Har: ret}, ps, "har", ret.Key)
	})
}

func HarUpload(rc *fasthttp.RequestCtx) {
	Act("har.upload", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		mpfrm, err := rc.MultipartForm()
		if err != nil {
			return "", err
		}
		name := strings.Join(mpfrm.Value["n"], "")
		fileHeaders, ok := mpfrm.File["f"]
		if !ok {
			return "", errors.New("no file uploaded")
		}
		if len(fileHeaders) != 1 {
			return "", errors.New("invalid file uploads")
		}
		fileHeader := fileHeaders[0]
		file, err := fileHeader.Open()
		if err != nil {
			return "", err
		}
		if name == "" {
			name = fileHeader.Filename
			if !strings.HasSuffix(name, ".har") {
				name += ".har"
			}
		}

		ps.Logger.Infof("Uploaded File: %+v\n", fileHeader.Filename)
		ps.Logger.Infof("File Size: %+v\n", fileHeader.Size)
		ps.Logger.Infof("MIME Header: %+v\n", fileHeader.Header)

		defer func() { _ = file.Close() }()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return "", err
		}
		err = as.Services.LoadToad.SaveHar(name, fileBytes)
		msg := fmt.Sprintf("Created [%s] (%s)", name, util.ByteSizeSI(fileHeader.Size))
		redir := "/har/" + name
		return FlashAndRedir(true, msg, redir, rc, ps)
	})
}

func HarStart(rc *fasthttp.RequestCtx) {
	Act("har.start", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.LoadToad.LoadHar(key)
		if err != nil {
			return "", err
		}
		w := &loadtoad.Workflow{ID: ret.Key, Name: ret.Key, Tests: har.Selectors{{Har: ret.Key}}}
		ps.Title = "Archive [" + key + "]"
		ps.Data = ret
		channel := "run-" + util.RandomString(16)
		return Render(rc, as, &vworkflow.Start{Workflow: w, Entries: ret.Entries, Channel: channel, Path: "/har/" + ret.Key + "/connect"}, ps, "har", ret.Key, "Run")
	})
}
