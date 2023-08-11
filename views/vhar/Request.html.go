// Code generated by qtc from "Request.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vhar/Request.html:1
package vhar

//line views/vhar/Request.html:1
import (
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/views/components"
)

//line views/vhar/Request.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vhar/Request.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vhar/Request.html:7
func streamrenderRequest(qw422016 *qt422016.Writer, i int, r *har.Request, ps *cutil.PageState) {
//line views/vhar/Request.html:7
	qw422016.N().S(`
  <ul class="accordion">
    <li>
      <input id="accordion-request-`)
//line views/vhar/Request.html:10
	qw422016.N().D(i)
//line views/vhar/Request.html:10
	qw422016.N().S(`" type="checkbox" hidden />
      <label class="no-padding" for="accordion-request-`)
//line views/vhar/Request.html:11
	qw422016.N().D(i)
//line views/vhar/Request.html:11
	qw422016.N().S(`"><em>(click to show)</em></label>
      <div class="bd">`)
//line views/vhar/Request.html:12
	components.StreamJSON(qw422016, r)
//line views/vhar/Request.html:12
	qw422016.N().S(`</div>
    </li>
  </ul>
`)
//line views/vhar/Request.html:15
}

//line views/vhar/Request.html:15
func writerenderRequest(qq422016 qtio422016.Writer, i int, r *har.Request, ps *cutil.PageState) {
//line views/vhar/Request.html:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Request.html:15
	streamrenderRequest(qw422016, i, r, ps)
//line views/vhar/Request.html:15
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Request.html:15
}

//line views/vhar/Request.html:15
func renderRequest(i int, r *har.Request, ps *cutil.PageState) string {
//line views/vhar/Request.html:15
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Request.html:15
	writerenderRequest(qb422016, i, r, ps)
//line views/vhar/Request.html:15
	qs422016 := string(qb422016.B)
//line views/vhar/Request.html:15
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Request.html:15
	return qs422016
//line views/vhar/Request.html:15
}
