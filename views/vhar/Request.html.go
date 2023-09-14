// Code generated by qtc from "Request.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vhar/Request.html:2
package vhar

//line views/vhar/Request.html:2
import (
	"fmt"
	"strings"

	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
)

//line views/vhar/Request.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vhar/Request.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vhar/Request.html:11
func StreamRenderRequest(qw422016 *qt422016.Writer, key string, r *har.Request, ps *cutil.PageState) {
//line views/vhar/Request.html:11
	qw422016.N().S(`
  <table class="min-200 expanded">
    <tbody>
      <tr>
        <th class="shrink">Method</th>
        <td>`)
//line views/vhar/Request.html:16
	qw422016.E().S(r.Method)
//line views/vhar/Request.html:16
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th class="shrink">URL</th>
        <td><a target="_blank" rel="noopener noreferrer" href="`)
//line views/vhar/Request.html:20
	qw422016.E().S(r.URL)
//line views/vhar/Request.html:20
	qw422016.N().S(`">`)
//line views/vhar/Request.html:20
	qw422016.E().S(r.URL)
//line views/vhar/Request.html:20
	qw422016.N().S(`</a></td>
      </tr>
      <tr>
        <th class="shrink">Headers </th>
        <td>`)
//line views/vhar/Request.html:24
	streamrenderNVPsHidden(qw422016, fmt.Sprintf("request-header-%s", key), "Header", r.Headers, r.HeadersSize, ps)
//line views/vhar/Request.html:24
	qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Request.html:26
	if len(r.QueryString) > 0 {
//line views/vhar/Request.html:26
		qw422016.N().S(`      <tr>
        <th class="shrink">Query String</th>
        <td>`)
//line views/vhar/Request.html:29
		streamrenderNVPsHidden(qw422016, fmt.Sprintf("request-query-string-%s", key), "Query String", r.QueryString, 0, ps)
//line views/vhar/Request.html:29
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Request.html:31
	}
//line views/vhar/Request.html:32
	if len(r.Cookies) > 0 {
//line views/vhar/Request.html:32
		qw422016.N().S(`      <tr>
        <th class="shrink">Cookies</th>
        <td>`)
//line views/vhar/Request.html:35
		streamrenderCookiesHidden(qw422016, fmt.Sprintf("request-cookies-%s", key), r.Cookies, ps)
//line views/vhar/Request.html:35
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Request.html:37
	}
//line views/vhar/Request.html:38
	if r.PostData != nil {
//line views/vhar/Request.html:38
		qw422016.N().S(`      <tr>
        <th class="shrink">Body</th>
        <td>
          <div class="right"><em>`)
//line views/vhar/Request.html:42
		if r.PostData.MimeType != "" {
//line views/vhar/Request.html:42
			qw422016.E().S(r.PostData.MimeType)
//line views/vhar/Request.html:42
			qw422016.N().S(`, `)
//line views/vhar/Request.html:42
		}
//line views/vhar/Request.html:42
		qw422016.E().S(util.ByteSizeSI(int64(r.BodySize)))
//line views/vhar/Request.html:42
		qw422016.N().S(`</em></div>
`)
//line views/vhar/Request.html:43
		if len(r.PostData.Params) > 0 {
//line views/vhar/Request.html:43
			qw422016.N().S(`          <ul>
`)
//line views/vhar/Request.html:45
			for _, param := range r.PostData.Params {
//line views/vhar/Request.html:45
				qw422016.N().S(`            <li><strong>`)
//line views/vhar/Request.html:46
				qw422016.E().S(param.Name)
//line views/vhar/Request.html:46
				qw422016.N().S(`</strong>: `)
//line views/vhar/Request.html:46
				qw422016.E().S(param.Value)
//line views/vhar/Request.html:46
				qw422016.N().S(`</li>
`)
//line views/vhar/Request.html:47
			}
//line views/vhar/Request.html:47
			qw422016.N().S(`          </ul>
`)
//line views/vhar/Request.html:49
		} else if r.PostData.Text != "" {
//line views/vhar/Request.html:49
			qw422016.N().S(`          <ul class="accordion">
            <li>
              <input id="accordion-request-body-`)
//line views/vhar/Request.html:52
			qw422016.E().S(key)
//line views/vhar/Request.html:52
			qw422016.N().S(`" type="checkbox" hidden />
              <label class="no-padding" for="accordion-request-body-`)
//line views/vhar/Request.html:53
			qw422016.E().S(key)
//line views/vhar/Request.html:53
			qw422016.N().S(`"><em>(click to show)</em></label>
              <div class="bd"><pre>`)
//line views/vhar/Request.html:54
			qw422016.E().S(r.PostData.Text)
//line views/vhar/Request.html:54
			qw422016.N().S(`</pre></div>
            </li>
          </ul>
`)
//line views/vhar/Request.html:57
		}
//line views/vhar/Request.html:57
		qw422016.N().S(`        </td>
      </tr>
`)
//line views/vhar/Request.html:60
	}
//line views/vhar/Request.html:60
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vhar/Request.html:63
}

//line views/vhar/Request.html:63
func WriteRenderRequest(qq422016 qtio422016.Writer, key string, r *har.Request, ps *cutil.PageState) {
//line views/vhar/Request.html:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Request.html:63
	StreamRenderRequest(qw422016, key, r, ps)
//line views/vhar/Request.html:63
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Request.html:63
}

//line views/vhar/Request.html:63
func RenderRequest(key string, r *har.Request, ps *cutil.PageState) string {
//line views/vhar/Request.html:63
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Request.html:63
	WriteRenderRequest(qb422016, key, r, ps)
//line views/vhar/Request.html:63
	qs422016 := string(qb422016.B)
//line views/vhar/Request.html:63
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Request.html:63
	return qs422016
//line views/vhar/Request.html:63
}

//line views/vhar/Request.html:65
func streamrenderCookiesHidden(qw422016 *qt422016.Writer, key string, cookies har.Cookies, ps *cutil.PageState) {
//line views/vhar/Request.html:65
	qw422016.N().S(`
  <ul class="accordion">
    <li>
      <input id="accordion-cookies-`)
//line views/vhar/Request.html:68
	qw422016.E().S(key)
//line views/vhar/Request.html:68
	qw422016.N().S(`" type="checkbox" hidden />
      <label class="no-padding" for="accordion-cookies-`)
//line views/vhar/Request.html:69
	qw422016.E().S(key)
//line views/vhar/Request.html:69
	qw422016.N().S(`">`)
//line views/vhar/Request.html:69
	qw422016.E().S(util.StringPlural(len(cookies), "Cookie"))
//line views/vhar/Request.html:69
	qw422016.N().S(` <em>(click to show)</em></label>
      <div class="bd">`)
//line views/vhar/Request.html:70
	streamrenderCookies(qw422016, key, cookies, ps)
//line views/vhar/Request.html:70
	qw422016.N().S(`</div>
    </li>
  </ul>
`)
//line views/vhar/Request.html:73
}

//line views/vhar/Request.html:73
func writerenderCookiesHidden(qq422016 qtio422016.Writer, key string, cookies har.Cookies, ps *cutil.PageState) {
//line views/vhar/Request.html:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Request.html:73
	streamrenderCookiesHidden(qw422016, key, cookies, ps)
//line views/vhar/Request.html:73
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Request.html:73
}

//line views/vhar/Request.html:73
func renderCookiesHidden(key string, cookies har.Cookies, ps *cutil.PageState) string {
//line views/vhar/Request.html:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Request.html:73
	writerenderCookiesHidden(qb422016, key, cookies, ps)
//line views/vhar/Request.html:73
	qs422016 := string(qb422016.B)
//line views/vhar/Request.html:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Request.html:73
	return qs422016
//line views/vhar/Request.html:73
}

//line views/vhar/Request.html:75
func streamrenderCookies(qw422016 *qt422016.Writer, key string, cs har.Cookies, ps *cutil.PageState) {
//line views/vhar/Request.html:75
	qw422016.N().S(`
  <table>
    <thead>
      <tr>
        <th class="shrink">Name</th>
        <th class="shrink">Value</th>
        <th class="shrink">Path</th>
        <th class="shrink">Domain</th>
        <th class="shrink">Expires</th>
        <th>Tags</th>
      </tr>
    </thead>
    <tbody>
`)
//line views/vhar/Request.html:88
	for i, c := range cs {
//line views/vhar/Request.html:88
		qw422016.N().S(`      <tr>
        <td>`)
//line views/vhar/Request.html:90
		qw422016.E().S(c.Name)
//line views/vhar/Request.html:90
		qw422016.N().S(`</td>
        <td>
`)
//line views/vhar/Request.html:92
		if len(c.Value) > 64 {
//line views/vhar/Request.html:92
			qw422016.N().S(`          <ul class="accordion">
            <li>
              <input id="accordion-`)
//line views/vhar/Request.html:95
			qw422016.E().S(key)
//line views/vhar/Request.html:95
			qw422016.N().S(`-cookie-`)
//line views/vhar/Request.html:95
			qw422016.N().D(i)
//line views/vhar/Request.html:95
			qw422016.N().S(`" type="checkbox" hidden />
              <label class="no-padding" for="accordion-`)
//line views/vhar/Request.html:96
			qw422016.E().S(key)
//line views/vhar/Request.html:96
			qw422016.N().S(`-cookie-`)
//line views/vhar/Request.html:96
			qw422016.N().D(i)
//line views/vhar/Request.html:96
			qw422016.N().S(`">`)
//line views/vhar/Request.html:96
			qw422016.E().S(c.Value[:64])
//line views/vhar/Request.html:96
			qw422016.N().S(`...</label>
              <div class="bd">`)
//line views/vhar/Request.html:97
			qw422016.E().S(c.Value)
//line views/vhar/Request.html:97
			qw422016.N().S(`</div>
            </li>
          </ul>
`)
//line views/vhar/Request.html:100
		} else {
//line views/vhar/Request.html:100
			qw422016.N().S(`          `)
//line views/vhar/Request.html:101
			qw422016.E().S(c.Value)
//line views/vhar/Request.html:101
			qw422016.N().S(`
`)
//line views/vhar/Request.html:102
		}
//line views/vhar/Request.html:102
		qw422016.N().S(`        </td>
        <td>`)
//line views/vhar/Request.html:104
		qw422016.E().S(c.Path)
//line views/vhar/Request.html:104
		qw422016.N().S(`</td>
        <td>`)
//line views/vhar/Request.html:105
		qw422016.E().S(c.Domain)
//line views/vhar/Request.html:105
		qw422016.N().S(`</td>
        <td title="`)
//line views/vhar/Request.html:106
		qw422016.E().S(util.TimeToFull(c.Exp()))
//line views/vhar/Request.html:106
		qw422016.N().S(`">`)
//line views/vhar/Request.html:106
		qw422016.E().S(c.ExpRelative())
//line views/vhar/Request.html:106
		qw422016.N().S(`</td>
        <td>`)
//line views/vhar/Request.html:107
		qw422016.E().S(strings.Join(c.Tags(), ", "))
//line views/vhar/Request.html:107
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Request.html:109
	}
//line views/vhar/Request.html:109
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vhar/Request.html:112
}

//line views/vhar/Request.html:112
func writerenderCookies(qq422016 qtio422016.Writer, key string, cs har.Cookies, ps *cutil.PageState) {
//line views/vhar/Request.html:112
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Request.html:112
	streamrenderCookies(qw422016, key, cs, ps)
//line views/vhar/Request.html:112
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Request.html:112
}

//line views/vhar/Request.html:112
func renderCookies(key string, cs har.Cookies, ps *cutil.PageState) string {
//line views/vhar/Request.html:112
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Request.html:112
	writerenderCookies(qb422016, key, cs, ps)
//line views/vhar/Request.html:112
	qs422016 := string(qb422016.B)
//line views/vhar/Request.html:112
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Request.html:112
	return qs422016
//line views/vhar/Request.html:112
}

//line views/vhar/Request.html:114
func streamrenderNVPsHidden(qw422016 *qt422016.Writer, key string, title string, nvps har.NVPs, size int, ps *cutil.PageState) {
//line views/vhar/Request.html:114
	qw422016.N().S(`
  <ul class="accordion">
    <li>
      <input id="accordion-`)
//line views/vhar/Request.html:117
	qw422016.E().S(key)
//line views/vhar/Request.html:117
	qw422016.N().S(`" type="checkbox" hidden />
      <label class="no-padding" for="accordion-`)
//line views/vhar/Request.html:118
	qw422016.E().S(key)
//line views/vhar/Request.html:118
	qw422016.N().S(`">
`)
//line views/vhar/Request.html:119
	if size > 0 {
//line views/vhar/Request.html:119
		qw422016.N().S(`        <div class="right"><em>`)
//line views/vhar/Request.html:120
		qw422016.E().S(util.ByteSizeSI(int64(size)))
//line views/vhar/Request.html:120
		qw422016.N().S(`</em></div>
`)
//line views/vhar/Request.html:121
	}
//line views/vhar/Request.html:121
	qw422016.N().S(`        `)
//line views/vhar/Request.html:122
	qw422016.E().S(util.StringPlural(len(nvps), title))
//line views/vhar/Request.html:122
	qw422016.N().S(`
        <em>(click to show)</em>
      </label>
      <div class="bd">`)
//line views/vhar/Request.html:125
	streamrenderNVPs(qw422016, nvps, ps)
//line views/vhar/Request.html:125
	qw422016.N().S(`</div>
    </li>
  </ul>
`)
//line views/vhar/Request.html:128
}

//line views/vhar/Request.html:128
func writerenderNVPsHidden(qq422016 qtio422016.Writer, key string, title string, nvps har.NVPs, size int, ps *cutil.PageState) {
//line views/vhar/Request.html:128
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Request.html:128
	streamrenderNVPsHidden(qw422016, key, title, nvps, size, ps)
//line views/vhar/Request.html:128
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Request.html:128
}

//line views/vhar/Request.html:128
func renderNVPsHidden(key string, title string, nvps har.NVPs, size int, ps *cutil.PageState) string {
//line views/vhar/Request.html:128
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Request.html:128
	writerenderNVPsHidden(qb422016, key, title, nvps, size, ps)
//line views/vhar/Request.html:128
	qs422016 := string(qb422016.B)
//line views/vhar/Request.html:128
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Request.html:128
	return qs422016
//line views/vhar/Request.html:128
}

//line views/vhar/Request.html:130
func streamrenderNVPs(qw422016 *qt422016.Writer, nvps har.NVPs, ps *cutil.PageState) {
//line views/vhar/Request.html:130
	qw422016.N().S(`
  <table>
    <tbody>
`)
//line views/vhar/Request.html:133
	for _, n := range nvps {
//line views/vhar/Request.html:133
		qw422016.N().S(`      <tr title="`)
//line views/vhar/Request.html:134
		qw422016.E().S(n.Comment)
//line views/vhar/Request.html:134
		qw422016.N().S(`">
        <th class="shrink">`)
//line views/vhar/Request.html:135
		qw422016.E().S(n.Name)
//line views/vhar/Request.html:135
		qw422016.N().S(`</th>
        <td>`)
//line views/vhar/Request.html:136
		qw422016.E().S(n.Value)
//line views/vhar/Request.html:136
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Request.html:138
	}
//line views/vhar/Request.html:138
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vhar/Request.html:141
}

//line views/vhar/Request.html:141
func writerenderNVPs(qq422016 qtio422016.Writer, nvps har.NVPs, ps *cutil.PageState) {
//line views/vhar/Request.html:141
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Request.html:141
	streamrenderNVPs(qw422016, nvps, ps)
//line views/vhar/Request.html:141
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Request.html:141
}

//line views/vhar/Request.html:141
func renderNVPs(nvps har.NVPs, ps *cutil.PageState) string {
//line views/vhar/Request.html:141
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Request.html:141
	writerenderNVPs(qb422016, nvps, ps)
//line views/vhar/Request.html:141
	qs422016 := string(qb422016.B)
//line views/vhar/Request.html:141
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Request.html:141
	return qs422016
//line views/vhar/Request.html:141
}
