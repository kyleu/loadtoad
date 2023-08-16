// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vscripting/List.html:2
package vscripting

//line views/vscripting/List.html:2
import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vscripting/List.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vscripting/List.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vscripting/List.html:9
type List struct {
	layout.Basic
	Scripts []string
}

//line views/vscripting/List.html:14
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vscripting/List.html:14
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/admin/scripting/new"><button>New</button></a></div>
    <h3>`)
//line views/vscripting/List.html:17
	components.StreamSVGRefIcon(qw422016, `sitemap`, ps)
//line views/vscripting/List.html:17
	qw422016.N().S(` Scripts</h3>
    <div class="mt">
      <table class="min-200">
        <thead>
          <tr>
            <th class="shrink">Name</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vscripting/List.html:26
	for _, x := range p.Scripts {
//line views/vscripting/List.html:26
		qw422016.N().S(`          <tr>
            <td><a href="/admin/scripting/`)
//line views/vscripting/List.html:28
		qw422016.E().S(x)
//line views/vscripting/List.html:28
		qw422016.N().S(`">`)
//line views/vscripting/List.html:28
		qw422016.E().S(x)
//line views/vscripting/List.html:28
		qw422016.N().S(`</a></td>
          </tr>
`)
//line views/vscripting/List.html:30
	}
//line views/vscripting/List.html:30
	qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vscripting/List.html:35
}

//line views/vscripting/List.html:35
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vscripting/List.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vscripting/List.html:35
	p.StreamBody(qw422016, as, ps)
//line views/vscripting/List.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/vscripting/List.html:35
}

//line views/vscripting/List.html:35
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vscripting/List.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vscripting/List.html:35
	p.WriteBody(qb422016, as, ps)
//line views/vscripting/List.html:35
	qs422016 := string(qb422016.B)
//line views/vscripting/List.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vscripting/List.html:35
	return qs422016
//line views/vscripting/List.html:35
}
