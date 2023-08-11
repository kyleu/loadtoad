package har

import (
	"bytes"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/samber/lo"
	"golang.org/x/net/http/httpguts"
	"net/http"
	"net/url"
)

type Entry struct {
	PageRef         string       `json:"pageref,omitempty"`
	StartedDateTime string       `json:"startedDateTime"`
	Time            float32      `json:"time"`
	Request         *Request     `json:"request"`
	Response        *Response    `json:"response"`
	Cache           *Cache       `json:"cache"`
	PageTimings     *PageTimings `json:"timings"`
	ServerIPAddress string       `json:"serverIPAddress,omitempty"`
	Connection      string       `json:"connection,omitempty"`
	Comment         string       `json:"comment,omitempty"`
}

func (e *Entry) String() string {
	if len(e.Request.URL) > 64 {
		return e.Request.URL[:64] + "..."
	}
	return e.Request.URL
}

func (e *Entry) Duration() string {
	return util.MicrosToMillis(e.PageTimings.Wait * 1000)
}

func (e *Entry) Clone() *Entry {
	return &Entry{
		PageRef: e.PageRef, StartedDateTime: e.StartedDateTime, Time: e.Time, Request: e.Request, Response: e.Response,
		Cache: e.Cache, PageTimings: e.PageTimings, ServerIPAddress: e.ServerIPAddress, Connection: e.Connection, Comment: e.Comment,
	}
}

func (e *Entry) Cleaned() *Entry {
	ret := e
	if ret.Request.PostData != nil && len(ret.Request.PostData.Text) > 1024*16 {
		ret = ret.Clone()
		ret.Request.PostData.Text = util.ByteSizeSI(int64(len(ret.Request.PostData.Text)))
	}
	if ret.Response != nil && ret.Response.Content != nil && ret.Response.Content.Size > 1024*16 {
		ret = ret.Clone()
		ret.Response.Content.Text = util.ByteSizeSI(int64(len(ret.Response.Content.Text)))
	}
	return ret
}

func (e *Entry) ToRequest(ignoreCookies bool) (*http.Request, error) {
	body := ""

	if e.Request.PostData != nil {
		if len(e.Request.PostData.Params) == 0 {
			body = e.Request.PostData.Text
		} else {
			form := url.Values{}
			for _, p := range e.Request.PostData.Params {
				form.Add(p.Name, p.Value)
			}
			body = form.Encode()
		}
	}

	req, _ := http.NewRequest(e.Request.Method, e.Request.URL, bytes.NewBuffer([]byte(body)))

	for _, h := range e.Request.Headers {
		if httpguts.ValidHeaderFieldName(h.Name) && httpguts.ValidHeaderFieldValue(h.Value) && h.Name != "Cookie" {
			req.Header.Add(h.Name, h.Value)
		}
	}

	if !ignoreCookies {
		for _, c := range e.Request.Cookies {
			cookie := &http.Cookie{Name: c.Name, Value: url.QueryEscape(c.Value), HttpOnly: false, Domain: c.Domain}
			req.AddCookie(cookie)
		}
	}

	return req, nil
}

type Entries []*Entry

func (e Entries) ForPage(ref string) Entries {
	return lo.Filter(e, func(x *Entry, _ int) bool {
		return x.PageRef == ref
	})
}
