// Code generated by qtc from "Debug.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/Debug.html:2
package views

//line views/Debug.html:2
import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/Debug.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Debug.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Debug.html:9
type Debug struct{ layout.Basic }

//line views/Debug.html:11
func (p *Debug) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Debug.html:11
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/Debug.html:13
	components.StreamSVGRefIcon(qw422016, `graph`, ps)
//line views/Debug.html:13
	qw422016.E().S(ps.Title)
//line views/Debug.html:13
	qw422016.N().S(`</h3>
`)
//line views/Debug.html:14
	if s, ok := ps.Data.(string); ok {
//line views/Debug.html:14
		qw422016.N().S(`    <div class="overflow full-width"><pre class="mt">`)
//line views/Debug.html:15
		qw422016.E().S(s)
//line views/Debug.html:15
		qw422016.N().S(`</pre></div>
`)
//line views/Debug.html:16
	} else {
//line views/Debug.html:16
		qw422016.N().S(`    <div class="mt">`)
//line views/Debug.html:17
		qw422016.N().S(components.JSON(ps.Data))
//line views/Debug.html:17
		qw422016.N().S(`</div>
`)
//line views/Debug.html:18
	}
//line views/Debug.html:18
	qw422016.N().S(`  </div>
`)
//line views/Debug.html:20
}

//line views/Debug.html:20
func (p *Debug) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Debug.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Debug.html:20
	p.StreamBody(qw422016, as, ps)
//line views/Debug.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/Debug.html:20
}

//line views/Debug.html:20
func (p *Debug) Body(as *app.State, ps *cutil.PageState) string {
//line views/Debug.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Debug.html:20
	p.WriteBody(qb422016, as, ps)
//line views/Debug.html:20
	qs422016 := string(qb422016.B)
//line views/Debug.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Debug.html:20
	return qs422016
//line views/Debug.html:20
}
