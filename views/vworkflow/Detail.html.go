// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkflow/Detail.html:1
package vworkflow

//line views/vworkflow/Detail.html:1
import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
	"github.com/kyleu/loadtoad/views/vhar"
)

//line views/vworkflow/Detail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkflow/Detail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkflow/Detail.html:12
type Detail struct {
	layout.Basic
	Workflow *loadtoad.Workflow
	Entries  har.Entries
}

//line views/vworkflow/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkflow/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-workflow"><button type="button">JSON</button></a></div>
    `)
//line views/vworkflow/Detail.html:21
	components.StreamJSONModal(qw422016, "workflow", "Workflow", p.Workflow, 3)
//line views/vworkflow/Detail.html:21
	qw422016.N().S(`
    <h3>`)
//line views/vworkflow/Detail.html:22
	components.StreamSVGRefIcon(qw422016, `sitemap`, ps)
//line views/vworkflow/Detail.html:22
	qw422016.N().S(` `)
//line views/vworkflow/Detail.html:22
	qw422016.E().S(p.Workflow.Title())
//line views/vworkflow/Detail.html:22
	qw422016.N().S(`</h3>
    <div class="mt">
      <a href="`)
//line views/vworkflow/Detail.html:24
	qw422016.E().S(p.Workflow.WebPath())
//line views/vworkflow/Detail.html:24
	qw422016.N().S(`/run?ok=true"><button>Run</button></a>
`)
//line views/vworkflow/Detail.html:25
	if len(p.Workflow.Replacements) > 0 {
//line views/vworkflow/Detail.html:25
		qw422016.N().S(`      <a href="`)
//line views/vworkflow/Detail.html:26
		qw422016.E().S(p.Workflow.WebPath())
//line views/vworkflow/Detail.html:26
		qw422016.N().S(`/run"><button>Run w/ Options</button></a>
`)
//line views/vworkflow/Detail.html:27
	}
//line views/vworkflow/Detail.html:27
	qw422016.N().S(`    </div>
  </div>
  <div class="card">
    <h3>`)
//line views/vworkflow/Detail.html:31
	components.StreamSVGRefIcon(qw422016, `play`, ps)
//line views/vworkflow/Detail.html:31
	qw422016.N().S(` `)
//line views/vworkflow/Detail.html:31
	qw422016.N().D(len(p.Entries))
//line views/vworkflow/Detail.html:31
	qw422016.N().S(` `)
//line views/vworkflow/Detail.html:31
	qw422016.E().S(util.StringPluralMaybe("Request", len(p.Entries)))
//line views/vworkflow/Detail.html:31
	qw422016.N().S(`</h3>
    <div class="mt">
      <ul class="accordion">
`)
//line views/vworkflow/Detail.html:34
	for i, e := range p.Entries {
//line views/vworkflow/Detail.html:35
		e = e.Cleaned()

//line views/vworkflow/Detail.html:35
		qw422016.N().S(`        <li>
          <input id="accordion-entry-`)
//line views/vworkflow/Detail.html:37
		qw422016.N().D(i)
//line views/vworkflow/Detail.html:37
		qw422016.N().S(`" type="checkbox" hidden />
          <label for="accordion-entry-`)
//line views/vworkflow/Detail.html:38
		qw422016.N().D(i)
//line views/vworkflow/Detail.html:38
		qw422016.N().S(`">
            `)
//line views/vworkflow/Detail.html:39
		vhar.StreamRenderEntryOptions(qw422016, i, e, false)
//line views/vworkflow/Detail.html:39
		qw422016.N().S(`
            `)
//line views/vworkflow/Detail.html:40
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vworkflow/Detail.html:40
		qw422016.N().S(` `)
//line views/vworkflow/Detail.html:40
		qw422016.E().S(e.String())
//line views/vworkflow/Detail.html:40
		qw422016.N().S(`
            <div class="clear"></div>
          </label>
          <div class="bd">
            `)
//line views/vworkflow/Detail.html:44
		vhar.StreamRenderEntry(qw422016, i, e, ps)
//line views/vworkflow/Detail.html:44
		qw422016.N().S(`
          </div>
          `)
//line views/vworkflow/Detail.html:46
		vhar.StreamRenderEntryModals(qw422016, i, e, false)
//line views/vworkflow/Detail.html:46
		qw422016.N().S(`
        </li>
`)
//line views/vworkflow/Detail.html:48
	}
//line views/vworkflow/Detail.html:48
	qw422016.N().S(`      </ul>
    </div>
  </div>
`)
//line views/vworkflow/Detail.html:52
}

//line views/vworkflow/Detail.html:52
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkflow/Detail.html:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkflow/Detail.html:52
	p.StreamBody(qw422016, as, ps)
//line views/vworkflow/Detail.html:52
	qt422016.ReleaseWriter(qw422016)
//line views/vworkflow/Detail.html:52
}

//line views/vworkflow/Detail.html:52
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkflow/Detail.html:52
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkflow/Detail.html:52
	p.WriteBody(qb422016, as, ps)
//line views/vworkflow/Detail.html:52
	qs422016 := string(qb422016.B)
//line views/vworkflow/Detail.html:52
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkflow/Detail.html:52
	return qs422016
//line views/vworkflow/Detail.html:52
}
