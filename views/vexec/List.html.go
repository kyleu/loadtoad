// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vexec/List.html:2
package vexec

//line views/vexec/List.html:2
import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/exec"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vexec/List.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vexec/List.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vexec/List.html:11
type List struct {
	layout.Basic
	Execs exec.Execs
}

//line views/vexec/List.html:16
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexec/List.html:16
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/admin/exec/new" title="start a new process">`)
//line views/vexec/List.html:18
	components.StreamSVGIcon(qw422016, `plus`, ps)
//line views/vexec/List.html:18
	qw422016.N().S(`</a></div>
    <h3>`)
//line views/vexec/List.html:19
	components.StreamSVGIcon(qw422016, `desktop`, ps)
//line views/vexec/List.html:19
	qw422016.E().S(util.StringPlural(len(p.Execs), "Process"))
//line views/vexec/List.html:19
	qw422016.N().S(`</h3>
    `)
//line views/vexec/List.html:20
	StreamTable(qw422016, p.Execs, as, ps)
//line views/vexec/List.html:20
	qw422016.N().S(`
  </div>
`)
//line views/vexec/List.html:22
}

//line views/vexec/List.html:22
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexec/List.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexec/List.html:22
	p.StreamBody(qw422016, as, ps)
//line views/vexec/List.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/vexec/List.html:22
}

//line views/vexec/List.html:22
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vexec/List.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexec/List.html:22
	p.WriteBody(qb422016, as, ps)
//line views/vexec/List.html:22
	qs422016 := string(qb422016.B)
//line views/vexec/List.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexec/List.html:22
	return qs422016
//line views/vexec/List.html:22
}
