// Code generated by qtc from "Request.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vadmin/Request.html:1
package vadmin

//line views/vadmin/Request.html:1
import (
	"net/http"

	"github.com/samber/lo"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vadmin/Request.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Request.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Request.html:13
type Request struct {
	layout.Basic
	Rsp http.ResponseWriter
	Req *http.Request
}

//line views/vadmin/Request.html:19
func (p *Request) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Request.html:19
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-ps"><button type="button">Page State</button></a></div>
    <h3>`)
//line views/vadmin/Request.html:22
	components.StreamSVGIcon(qw422016, `link`, ps)
//line views/vadmin/Request.html:22
	qw422016.N().S(` Request Debug</h3>
    <div class="overflow full-width">
      <table>
        <thead>
          <tr>
            <th>Key</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>URL</td>
            <td>`)
//line views/vadmin/Request.html:34
	qw422016.E().S(p.Req.URL.String())
//line views/vadmin/Request.html:34
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Protocol</td>
            <td>`)
//line views/vadmin/Request.html:38
	qw422016.E().S(p.Req.URL.Scheme)
//line views/vadmin/Request.html:38
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Host</td>
            <td>`)
//line views/vadmin/Request.html:42
	qw422016.E().S(p.Req.URL.Host)
//line views/vadmin/Request.html:42
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Path</td>
            <td>`)
//line views/vadmin/Request.html:46
	qw422016.E().S(p.Req.URL.Path)
//line views/vadmin/Request.html:46
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Query String</td>
            <td>`)
//line views/vadmin/Request.html:50
	qw422016.E().S(p.Req.URL.RawQuery)
//line views/vadmin/Request.html:50
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Body Size</td>
            <td>`)
//line views/vadmin/Request.html:54
	qw422016.E().S(util.ByteSizeSI(int64(len(ps.RequestBody))))
//line views/vadmin/Request.html:54
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Browser</td>
            <td>`)
//line views/vadmin/Request.html:58
	qw422016.E().S(ps.Browser)
//line views/vadmin/Request.html:58
	qw422016.N().S(` `)
//line views/vadmin/Request.html:58
	qw422016.E().S(ps.BrowserVersion)
//line views/vadmin/Request.html:58
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>OS</td>
            <td>`)
//line views/vadmin/Request.html:62
	qw422016.E().S(ps.OS)
//line views/vadmin/Request.html:62
	qw422016.N().S(` `)
//line views/vadmin/Request.html:62
	qw422016.E().S(ps.OSVersion)
//line views/vadmin/Request.html:62
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/Request.html:68
	if len(p.Req.Header) > 0 {
//line views/vadmin/Request.html:69
		hd := cutil.RequestHeadersMap(p.Req)

//line views/vadmin/Request.html:69
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vadmin/Request.html:71
		components.StreamSVGIcon(qw422016, `code`, ps)
//line views/vadmin/Request.html:71
		qw422016.N().S(` Headers</h3>
    <div class="overflow full-width">
      <table>
        <thead>
          <tr>
            <th>Key</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vadmin/Request.html:81
		for _, k := range util.ArraySorted(lo.Keys(hd)) {
//line views/vadmin/Request.html:81
			qw422016.N().S(`          <tr>
            <td class="nowrap">`)
//line views/vadmin/Request.html:83
			qw422016.E().S(k)
//line views/vadmin/Request.html:83
			qw422016.N().S(`</td>
            <td>`)
//line views/vadmin/Request.html:84
			qw422016.E().S(hd.GetStringOpt(k))
//line views/vadmin/Request.html:84
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vadmin/Request.html:86
		}
//line views/vadmin/Request.html:86
		qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/Request.html:91
	}
//line views/vadmin/Request.html:91
	qw422016.N().S(`  `)
//line views/vadmin/Request.html:92
	components.StreamJSONModal(qw422016, "ps", "Page State", ps, 1)
//line views/vadmin/Request.html:92
	qw422016.N().S(`
`)
//line views/vadmin/Request.html:93
}

//line views/vadmin/Request.html:93
func (p *Request) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Request.html:93
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Request.html:93
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Request.html:93
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Request.html:93
}

//line views/vadmin/Request.html:93
func (p *Request) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Request.html:93
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Request.html:93
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Request.html:93
	qs422016 := string(qb422016.B)
//line views/vadmin/Request.html:93
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Request.html:93
	return qs422016
//line views/vadmin/Request.html:93
}
