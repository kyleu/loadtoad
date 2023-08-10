package har

import (
	"net/url"
	"strings"
)

func (e *Entry) CurlFromEntry() (string, error) {
	command := "curl -X " + e.Request.Method
	if e.Request.HTTPVersion == "HTTP/1.0" {
		command += " -0"
	}
	var cookies []string
	if len(e.Request.Cookies) > 0 {
		for _, cookie := range e.Request.Cookies {
			cookies = append(cookies, url.QueryEscape(cookie.Name)+"="+url.QueryEscape(cookie.Value))
		}
		command += " -b \"" + strings.Join(cookies[:], "&") + "\" "
	}
	for _, h := range e.Request.Headers {
		command += " -H \"" + h.Name + ": " + h.Value + "\" "
	}
	if e.Request.Method == "POST" && len(e.Request.PostData.Text) > 0 {
		command += "-d \"" + e.Request.PostData.Text + "\""
	}
	command += " \"" + e.Request.URL + "\""
	return command, nil
}
