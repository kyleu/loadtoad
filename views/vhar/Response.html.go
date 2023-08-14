// Code generated by qtc from "Response.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vhar/Response.html:1
package vhar

//line views/vhar/Response.html:1
import (
	"fmt"

	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
)

//line views/vhar/Response.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vhar/Response.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vhar/Response.html:9
func streamrenderResponse(qw422016 *qt422016.Writer, i int, r *har.Response, ps *cutil.PageState) {
//line views/vhar/Response.html:9
	qw422016.N().S(`
  <table class="min-200">
    <tbody>
      <tr>
        <th class="shrink">Status</th>
        <td>`)
//line views/vhar/Response.html:14
	qw422016.N().D(r.Status)
//line views/vhar/Response.html:14
	qw422016.N().S(`: `)
//line views/vhar/Response.html:14
	qw422016.E().S(r.StatusText)
//line views/vhar/Response.html:14
	qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Response.html:16
	if r.RedirectURL != "" {
//line views/vhar/Response.html:16
		qw422016.N().S(`      <tr>
        <th class="shrink">Redirect URL</th>
        <td>`)
//line views/vhar/Response.html:19
		qw422016.E().S(r.RedirectURL)
//line views/vhar/Response.html:19
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Response.html:21
	}
//line views/vhar/Response.html:21
	qw422016.N().S(`      <tr>
        <th class="shrink">Headers </th>
        <td>`)
//line views/vhar/Response.html:24
	streamrenderNVPsHidden(qw422016, fmt.Sprintf("response-header-%d", i), "Header", r.Headers, r.HeadersSize, ps)
//line views/vhar/Response.html:24
	qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Response.html:26
	if len(r.Cookies) > 0 {
//line views/vhar/Response.html:26
		qw422016.N().S(`      <tr>
        <th class="shrink">Cookies</th>
        <td>`)
//line views/vhar/Response.html:29
		streamrenderCookiesHidden(qw422016, fmt.Sprintf("response-cookies-%d", i), r.Cookies, ps)
//line views/vhar/Response.html:29
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vhar/Response.html:31
	}
//line views/vhar/Response.html:32
	if r.Content != nil && r.Content.Text != "" {
//line views/vhar/Response.html:32
		qw422016.N().S(`      <tr>
        <th class="shrink">Body</th>
        <td>
          <div class="right"><em>`)
//line views/vhar/Response.html:36
		qw422016.E().S(util.ByteSizeSI(int64(r.Content.Size)))
//line views/vhar/Response.html:36
		qw422016.N().S(`</em></div>
          <ul class="accordion">
            <li>
              <input id="accordion-response-body-`)
//line views/vhar/Response.html:39
		qw422016.N().D(i)
//line views/vhar/Response.html:39
		qw422016.N().S(`" type="checkbox" hidden />
              <label class="no-padding" for="accordion-response-body-`)
//line views/vhar/Response.html:40
		qw422016.N().D(i)
//line views/vhar/Response.html:40
		qw422016.N().S(`"><em>(click to show)</em></label>
              <div class="bd">
                <pre>`)
//line views/vhar/Response.html:42
		qw422016.E().S(r.Content.Text)
//line views/vhar/Response.html:42
		qw422016.N().S(`</pre>
              </div>
            </li>
          </ul>

        </td>
      </tr>
`)
//line views/vhar/Response.html:49
	}
//line views/vhar/Response.html:49
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vhar/Response.html:52
}

//line views/vhar/Response.html:52
func writerenderResponse(qq422016 qtio422016.Writer, i int, r *har.Response, ps *cutil.PageState) {
//line views/vhar/Response.html:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Response.html:52
	streamrenderResponse(qw422016, i, r, ps)
//line views/vhar/Response.html:52
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Response.html:52
}

//line views/vhar/Response.html:52
func renderResponse(i int, r *har.Response, ps *cutil.PageState) string {
//line views/vhar/Response.html:52
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Response.html:52
	writerenderResponse(qb422016, i, r, ps)
//line views/vhar/Response.html:52
	qs422016 := string(qb422016.B)
//line views/vhar/Response.html:52
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Response.html:52
	return qs422016
//line views/vhar/Response.html:52
}
