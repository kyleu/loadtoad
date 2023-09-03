// Code generated by qtc from "Response.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vhar/Response.html:2
package vhar

//line views/vhar/Response.html:2
import (
	"fmt"

	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
)

//line views/vhar/Response.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vhar/Response.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vhar/Response.html:11
func StreamRenderResponse(qw422016 *qt422016.Writer, key string, r *har.Response, ps *cutil.PageState) {
//line views/vhar/Response.html:11
	qw422016.N().S(`
  <table class="min-200 expanded">
    <tbody>
      <tr>
        <th class="shrink">Status</th>
        <td>`)
//line views/vhar/Response.html:16
	qw422016.N().D(r.Status)
//line views/vhar/Response.html:16
	qw422016.N().S(`: `)
//line views/vhar/Response.html:16
	qw422016.E().S(r.StatusText)
//line views/vhar/Response.html:16
	qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Response.html:18
	if r.RedirectURL != "" {
//line views/vhar/Response.html:18
		qw422016.N().S(`      <tr>
        <th class="shrink">Redirect URL</th>
        <td>`)
//line views/vhar/Response.html:21
		qw422016.E().S(r.RedirectURL)
//line views/vhar/Response.html:21
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Response.html:23
	}
//line views/vhar/Response.html:23
	qw422016.N().S(`      <tr>
        <th class="shrink">Headers </th>
        <td>`)
//line views/vhar/Response.html:26
	streamrenderNVPsHidden(qw422016, fmt.Sprintf("response-header-%s", key), "Header", r.Headers, r.HeadersSize, ps)
//line views/vhar/Response.html:26
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th class="shrink">Content Type</th>
        <td>`)
//line views/vhar/Response.html:30
	qw422016.E().S(r.ContentType())
//line views/vhar/Response.html:30
	qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Response.html:32
	if len(r.Cookies) > 0 {
//line views/vhar/Response.html:32
		qw422016.N().S(`      <tr>
        <th class="shrink">Cookies</th>
        <td>`)
//line views/vhar/Response.html:35
		streamrenderCookiesHidden(qw422016, fmt.Sprintf("response-cookies-%s", key), r.Cookies, ps)
//line views/vhar/Response.html:35
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Response.html:37
	}
//line views/vhar/Response.html:38
	if r.Content != nil && r.Content.Text != "" {
//line views/vhar/Response.html:38
		qw422016.N().S(`      <tr>
        <th class="shrink">Body</th>
        <td>
          <div class="right"><em>`)
//line views/vhar/Response.html:42
		qw422016.E().S(util.ByteSizeSI(int64(r.Content.Size)))
//line views/vhar/Response.html:42
		qw422016.N().S(`</em></div>
          <ul class="accordion">
            <li>
              <input id="accordion-response-body-`)
//line views/vhar/Response.html:45
		qw422016.E().S(key)
//line views/vhar/Response.html:45
		qw422016.N().S(`" type="checkbox" hidden />
              <label class="no-padding" for="accordion-response-body-`)
//line views/vhar/Response.html:46
		qw422016.E().S(key)
//line views/vhar/Response.html:46
		qw422016.N().S(`"><em>(click to show)</em></label>
              <div class="bd">
`)
//line views/vhar/Response.html:48
		if r.Content.JSON != nil {
//line views/vhar/Response.html:48
			qw422016.N().S(`                `)
//line views/vhar/Response.html:49
			components.StreamJSON(qw422016, r.Content.JSON)
//line views/vhar/Response.html:49
			qw422016.N().S(`
`)
//line views/vhar/Response.html:50
		} else {
//line views/vhar/Response.html:50
			qw422016.N().S(`                <pre>`)
//line views/vhar/Response.html:51
			qw422016.E().S(r.Content.Text)
//line views/vhar/Response.html:51
			qw422016.N().S(`</pre>
`)
//line views/vhar/Response.html:52
		}
//line views/vhar/Response.html:52
		qw422016.N().S(`              </div>
            </li>
          </ul>
        </td>
      </tr>
`)
//line views/vhar/Response.html:58
	}
//line views/vhar/Response.html:58
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vhar/Response.html:61
}

//line views/vhar/Response.html:61
func WriteRenderResponse(qq422016 qtio422016.Writer, key string, r *har.Response, ps *cutil.PageState) {
//line views/vhar/Response.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Response.html:61
	StreamRenderResponse(qw422016, key, r, ps)
//line views/vhar/Response.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Response.html:61
}

//line views/vhar/Response.html:61
func RenderResponse(key string, r *har.Response, ps *cutil.PageState) string {
//line views/vhar/Response.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Response.html:61
	WriteRenderResponse(qb422016, key, r, ps)
//line views/vhar/Response.html:61
	qs422016 := string(qb422016.B)
//line views/vhar/Response.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Response.html:61
	return qs422016
//line views/vhar/Response.html:61
}
