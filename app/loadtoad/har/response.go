package har

import (
	"github.com/samber/lo"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Status      int      `json:"status"`
	StatusText  string   `json:"statusText"`
	HTTPVersion string   `json:"httpVersion"`
	Cookies     Cookies  `json:"cookies"`
	Headers     NVPs     `json:"headers"`
	Content     *Content `json:"content"`
	RedirectURL string   `json:"redirectURL"`
	HeadersSize int      `json:"headersSize"`
	BodySize    int      `json:"bodySize"`
	Comment     string   `json:"comment,omitempty"`
}

func ResponseFromHTTP(r *http.Response) *Response {
	cooks := lo.Map(r.Cookies(), func(c *http.Cookie, _ int) *Cookie {
		exp := c.Expires.Format("2006-01-02T15:04:05.000Z")
		return &Cookie{Name: c.Name, Value: c.Value, Path: c.Path, Domain: c.Domain, Expires: exp, HTTPOnly: c.HttpOnly, Secure: c.Secure}
	})
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(bodyBytes)
	content := &Content{Size: len(body), Text: body}
	ret := &Response{Status: r.StatusCode, StatusText: r.Status, Cookies: cooks, Headers: nil, Content: content, BodySize: content.Size}
	return ret
}
