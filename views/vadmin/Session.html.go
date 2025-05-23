// Code generated by qtc from "Session.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vadmin/Session.html:1
package vadmin

//line views/vadmin/Session.html:1
import (
	"fmt"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vadmin/Session.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Session.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Session.html:11
type Session struct{ layout.Basic }

//line views/vadmin/Session.html:13
func (p *Session) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Session.html:13
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vadmin/Session.html:15
	components.StreamSVGIcon(qw422016, `desktop`, ps)
//line views/vadmin/Session.html:15
	qw422016.N().S(` Session</h3>
    <em>`)
//line views/vadmin/Session.html:16
	qw422016.N().D(len(ps.Session))
//line views/vadmin/Session.html:16
	qw422016.N().S(` values</em>
  </div>
`)
//line views/vadmin/Session.html:18
	if len(ps.Session) > 0 {
//line views/vadmin/Session.html:18
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vadmin/Session.html:20
		components.StreamSVGIcon(qw422016, `list`, ps)
//line views/vadmin/Session.html:20
		qw422016.N().S(` Values</h3>
    <div class="overflow full-width">
      <table class="mt expanded">
        <tbody>
`)
//line views/vadmin/Session.html:24
		for _, k := range util.MapKeysSorted(ps.Session) {
//line views/vadmin/Session.html:25
			v := ps.Session[k]

//line views/vadmin/Session.html:25
			qw422016.N().S(`          <tr>
            <th class="shrink">`)
//line views/vadmin/Session.html:27
			qw422016.E().S(k)
//line views/vadmin/Session.html:27
			qw422016.N().S(`</th>
            <td>`)
//line views/vadmin/Session.html:28
			qw422016.E().S(fmt.Sprint(v))
//line views/vadmin/Session.html:28
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vadmin/Session.html:30
		}
//line views/vadmin/Session.html:30
		qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/Session.html:35
	} else {
//line views/vadmin/Session.html:35
		qw422016.N().S(`  <div class="card">
    <em>Empty session</em>
  </div>
`)
//line views/vadmin/Session.html:39
	}
//line views/vadmin/Session.html:39
	qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vadmin/Session.html:41
	components.StreamSVGIcon(qw422016, `profile`, ps)
//line views/vadmin/Session.html:41
	qw422016.N().S(` Profile</h3>
    <div class="mt">`)
//line views/vadmin/Session.html:42
	components.StreamJSON(qw422016, ps.Profile)
//line views/vadmin/Session.html:42
	qw422016.N().S(`</div>
  </div>
`)
//line views/vadmin/Session.html:44
}

//line views/vadmin/Session.html:44
func (p *Session) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Session.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Session.html:44
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Session.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Session.html:44
}

//line views/vadmin/Session.html:44
func (p *Session) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Session.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Session.html:44
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Session.html:44
	qs422016 := string(qb422016.B)
//line views/vadmin/Session.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Session.html:44
	return qs422016
//line views/vadmin/Session.html:44
}
