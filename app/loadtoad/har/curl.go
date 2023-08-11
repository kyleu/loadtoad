package har

import (
	"github.com/kyleu/loadtoad/app/util"
	"github.com/samber/lo"
	"net/url"
	"strings"
)

func (e *Entry) Curl() string {
	command := "curl -X " + e.Request.Method + " --compressed "
	if e.Request.HTTPVersion == "HTTP/1.0" {
		command += " -0"
	}
	var cookies []string
	if len(e.Request.Cookies) > 0 {
		for _, cookie := range e.Request.Cookies {
			cookies = append(cookies, url.QueryEscape(cookie.Name)+"="+url.QueryEscape(cookie.Value))
		}
		command += " \\\n  -b \"" + strings.Join(cookies[:], "&") + "\""
	}
	for _, h := range e.Request.Headers {
		if h.Name == "Accept-Encoding" && strings.Contains(h.Value, "br") {
			encs := lo.Filter(util.StringSplitAndTrim(h.Value, ","), func(x string, _ int) bool {
				return x == "deflate" || x == "gzip"
			})
			h = &NVP{Name: h.Name, Value: strings.Join(encs, ", "), Comment: h.Comment}
		}
		command += " \\\n  -H \"" + h.Name + ": " + h.Value + "\""
	}
	if e.Request.Method == "POST" && e.Request.PostData != nil && len(e.Request.PostData.Text) > 0 {
		command += " \\\n  -d \"" + e.Request.PostData.Text + "\""
	}
	command += " \\\n  \"" + e.Request.URL + "\""
	return command
}
