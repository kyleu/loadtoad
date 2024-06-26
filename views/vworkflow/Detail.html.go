// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkflow/Detail.html:1
package vworkflow

//line views/vworkflow/Detail.html:1
import (
	"fmt"
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
	"github.com/kyleu/loadtoad/views/vhar"
)

//line views/vworkflow/Detail.html:17
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkflow/Detail.html:17
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkflow/Detail.html:17
type Detail struct {
	layout.Basic
	Workflow *loadtoad.Workflow
	Entries  har.Entries
}

//line views/vworkflow/Detail.html:23
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkflow/Detail.html:23
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="/workflow/`)
//line views/vworkflow/Detail.html:26
	qw422016.N().U(p.Workflow.ID)
//line views/vworkflow/Detail.html:26
	qw422016.N().S(`/edit"><button type="button">Edit</button></a>
      <a href="#modal-workflow"><button type="button">JSON</button></a>
    </div>
    `)
//line views/vworkflow/Detail.html:29
	components.StreamJSONModal(qw422016, "workflow", "Workflow", p.Workflow, 3)
//line views/vworkflow/Detail.html:29
	qw422016.N().S(`
    <h3>`)
//line views/vworkflow/Detail.html:30
	components.StreamSVGIcon(qw422016, `sitemap`, ps)
//line views/vworkflow/Detail.html:30
	qw422016.N().S(` `)
//line views/vworkflow/Detail.html:30
	qw422016.E().S(p.Workflow.Title())
//line views/vworkflow/Detail.html:30
	qw422016.N().S(`</h3>
    <div class="mt">
      <a href="`)
//line views/vworkflow/Detail.html:32
	qw422016.E().S(p.Workflow.WebPath())
//line views/vworkflow/Detail.html:32
	qw422016.N().S(`/run?ok=true"><button>Run</button></a>
`)
//line views/vworkflow/Detail.html:33
	if len(p.Workflow.Replacements) > 0 {
//line views/vworkflow/Detail.html:33
		qw422016.N().S(`      <a href="`)
//line views/vworkflow/Detail.html:34
		qw422016.E().S(p.Workflow.WebPath())
//line views/vworkflow/Detail.html:34
		qw422016.N().S(`/run"><button>Run w/ Options</button></a>
`)
//line views/vworkflow/Detail.html:35
	}
//line views/vworkflow/Detail.html:35
	qw422016.N().S(`      <a href="`)
//line views/vworkflow/Detail.html:36
	qw422016.E().S(p.Workflow.WebPath())
//line views/vworkflow/Detail.html:36
	qw422016.N().S(`/bench"><button>Benchmark</button></a>
    </div>
  </div>
`)
//line views/vworkflow/Detail.html:39
	if len(p.Workflow.Scripts) > 0 {
//line views/vworkflow/Detail.html:39
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vworkflow/Detail.html:41
		components.StreamSVGIcon(qw422016, `file-code`, ps)
//line views/vworkflow/Detail.html:41
		qw422016.N().S(` `)
//line views/vworkflow/Detail.html:41
		qw422016.E().S(util.StringPlural(len(p.Workflow.Scripts), "Script"))
//line views/vworkflow/Detail.html:41
		qw422016.N().S(`</h3>
    <div class="mt">
      <table class="min-200">
        <tbody>
`)
//line views/vworkflow/Detail.html:45
		for _, scr := range p.Workflow.Scripts {
//line views/vworkflow/Detail.html:45
			qw422016.N().S(`          <tr>
`)
//line views/vworkflow/Detail.html:48
			link := scr
			if !strings.HasSuffix(link, ".js") {
				link += ".js"
			}

//line views/vworkflow/Detail.html:52
			qw422016.N().S(`            <td class="shrink"><a href="/admin/scripting/`)
//line views/vworkflow/Detail.html:53
			qw422016.N().U(link)
//line views/vworkflow/Detail.html:53
			qw422016.N().S(`">`)
//line views/vworkflow/Detail.html:53
			qw422016.E().S(scr)
//line views/vworkflow/Detail.html:53
			qw422016.N().S(`</a></td>
          </tr>
`)
//line views/vworkflow/Detail.html:55
		}
//line views/vworkflow/Detail.html:55
		qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vworkflow/Detail.html:60
	}
//line views/vworkflow/Detail.html:61
	if len(p.Workflow.Replacements) > 0 {
//line views/vworkflow/Detail.html:61
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vworkflow/Detail.html:63
		components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vworkflow/Detail.html:63
		qw422016.N().S(` `)
//line views/vworkflow/Detail.html:63
		qw422016.E().S(util.StringPlural(len(p.Workflow.Replacements), "Replacement"))
//line views/vworkflow/Detail.html:63
		qw422016.N().S(`</h3>
    <div class="mt">
      <table class="min-200">
        <tbody>
`)
//line views/vworkflow/Detail.html:67
		for _, k := range util.ArraySorted(lo.Keys(p.Workflow.Replacements)) {
//line views/vworkflow/Detail.html:69
			v := strings.Split(p.Workflow.Replacements[k], "||")
			if len(k) > 50 {
				k = k[:50] + "..."
			}

//line views/vworkflow/Detail.html:73
			qw422016.N().S(`          <tr>
            <td class="shrink">`)
//line views/vworkflow/Detail.html:75
			qw422016.E().S(k)
//line views/vworkflow/Detail.html:75
			qw422016.N().S(`</td>
            <td>
`)
//line views/vworkflow/Detail.html:77
			if len(v) == 1 {
//line views/vworkflow/Detail.html:77
				qw422016.N().S(`              `)
//line views/vworkflow/Detail.html:78
				qw422016.E().S(v[0])
//line views/vworkflow/Detail.html:78
				qw422016.N().S(`
`)
//line views/vworkflow/Detail.html:79
			} else {
//line views/vworkflow/Detail.html:79
				qw422016.N().S(`              <ul>
`)
//line views/vworkflow/Detail.html:81
				for _, x := range v {
//line views/vworkflow/Detail.html:81
					qw422016.N().S(`                <li>`)
//line views/vworkflow/Detail.html:82
					qw422016.E().S(x)
//line views/vworkflow/Detail.html:82
					qw422016.N().S(`</li>
`)
//line views/vworkflow/Detail.html:83
				}
//line views/vworkflow/Detail.html:83
				qw422016.N().S(`              </ul>
`)
//line views/vworkflow/Detail.html:85
			}
//line views/vworkflow/Detail.html:85
			qw422016.N().S(`            </td>
          </tr>
`)
//line views/vworkflow/Detail.html:88
		}
//line views/vworkflow/Detail.html:88
		qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vworkflow/Detail.html:93
	}
//line views/vworkflow/Detail.html:94
	if len(p.Workflow.Variables) > 0 {
//line views/vworkflow/Detail.html:94
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vworkflow/Detail.html:96
		components.StreamSVGIcon(qw422016, `sitemap`, ps)
//line views/vworkflow/Detail.html:96
		qw422016.N().S(` `)
//line views/vworkflow/Detail.html:96
		qw422016.E().S(util.StringPlural(len(p.Workflow.Variables), "Variable"))
//line views/vworkflow/Detail.html:96
		qw422016.N().S(`</h3>
    <div class="mt">
      <table class="min-200">
        <tbody>
`)
//line views/vworkflow/Detail.html:100
		for k, v := range p.Workflow.Variables {
//line views/vworkflow/Detail.html:100
			qw422016.N().S(`          <tr>
            <th class="shrink">`)
//line views/vworkflow/Detail.html:102
			qw422016.E().S(k)
//line views/vworkflow/Detail.html:102
			qw422016.N().S(`</th>
            <td>`)
//line views/vworkflow/Detail.html:103
			qw422016.E().V(v)
//line views/vworkflow/Detail.html:103
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vworkflow/Detail.html:105
		}
//line views/vworkflow/Detail.html:105
		qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vworkflow/Detail.html:110
	}
//line views/vworkflow/Detail.html:110
	qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vworkflow/Detail.html:112
	components.StreamSVGIcon(qw422016, `play`, ps)
//line views/vworkflow/Detail.html:112
	qw422016.N().S(` `)
//line views/vworkflow/Detail.html:112
	qw422016.E().S(util.StringPlural(len(p.Entries), "Request"))
//line views/vworkflow/Detail.html:112
	qw422016.N().S(`</h3>
    <div class="mt">
      `)
//line views/vworkflow/Detail.html:114
	streamshowEntries(qw422016, p.Entries, ps)
//line views/vworkflow/Detail.html:114
	qw422016.N().S(`
    </div>
  </div>
`)
//line views/vworkflow/Detail.html:117
}

//line views/vworkflow/Detail.html:117
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkflow/Detail.html:117
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkflow/Detail.html:117
	p.StreamBody(qw422016, as, ps)
//line views/vworkflow/Detail.html:117
	qt422016.ReleaseWriter(qw422016)
//line views/vworkflow/Detail.html:117
}

//line views/vworkflow/Detail.html:117
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkflow/Detail.html:117
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkflow/Detail.html:117
	p.WriteBody(qb422016, as, ps)
//line views/vworkflow/Detail.html:117
	qs422016 := string(qb422016.B)
//line views/vworkflow/Detail.html:117
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkflow/Detail.html:117
	return qs422016
//line views/vworkflow/Detail.html:117
}

//line views/vworkflow/Detail.html:119
func streamshowEntries(qw422016 *qt422016.Writer, ents har.Entries, ps *cutil.PageState) {
//line views/vworkflow/Detail.html:119
	qw422016.N().S(`
`)
//line views/vworkflow/Detail.html:120
	var idx int

//line views/vworkflow/Detail.html:121
	for _, sel := range ents.Selectors() {
//line views/vworkflow/Detail.html:121
		qw422016.N().S(`  <h4 class="mt">`)
//line views/vworkflow/Detail.html:122
		qw422016.E().S(sel.String())
//line views/vworkflow/Detail.html:122
		qw422016.N().S(`</h4>
`)
//line views/vworkflow/Detail.html:123
		selEnts := ents.BySelector(sel)

//line views/vworkflow/Detail.html:124
		if len(selEnts) == 0 {
//line views/vworkflow/Detail.html:124
			qw422016.N().S(`  <em>no requests matched</em>
`)
//line views/vworkflow/Detail.html:126
		} else {
//line views/vworkflow/Detail.html:126
			qw422016.N().S(`  <div class="mt">
    <ul class="accordion">
`)
//line views/vworkflow/Detail.html:129
			for _, e := range selEnts {
//line views/vworkflow/Detail.html:130
				key := fmt.Sprintf("%d", idx)

//line views/vworkflow/Detail.html:130
				qw422016.N().S(`      <li title="`)
//line views/vworkflow/Detail.html:131
				qw422016.E().S(e.Request.URL)
//line views/vworkflow/Detail.html:131
				qw422016.N().S(`">
        <input id="accordion-entry-`)
//line views/vworkflow/Detail.html:132
				qw422016.E().S(key)
//line views/vworkflow/Detail.html:132
				qw422016.N().S(`" type="checkbox" hidden />
        <label for="accordion-entry-`)
//line views/vworkflow/Detail.html:133
				qw422016.E().S(key)
//line views/vworkflow/Detail.html:133
				qw422016.N().S(`">
          `)
//line views/vworkflow/Detail.html:134
				vhar.StreamRenderEntryOptions(qw422016, key, e, false)
//line views/vworkflow/Detail.html:134
				qw422016.N().S(`
          `)
//line views/vworkflow/Detail.html:135
				components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vworkflow/Detail.html:135
				qw422016.N().S(` `)
//line views/vworkflow/Detail.html:135
				qw422016.E().S(e.String())
//line views/vworkflow/Detail.html:135
				qw422016.N().S(`
          <div class="clear"></div>
        </label>
        <div class="bd"><div><div>
          `)
//line views/vworkflow/Detail.html:139
				vhar.StreamRenderEntry(qw422016, key, e, ps)
//line views/vworkflow/Detail.html:139
				qw422016.N().S(`
        </div></div></div>
        `)
//line views/vworkflow/Detail.html:141
				vhar.StreamRenderEntryModals(qw422016, key, e, false)
//line views/vworkflow/Detail.html:141
				qw422016.N().S(`
      </li>
`)
//line views/vworkflow/Detail.html:143
				idx++

//line views/vworkflow/Detail.html:144
			}
//line views/vworkflow/Detail.html:144
			qw422016.N().S(`    </ul>
  </div>
`)
//line views/vworkflow/Detail.html:147
		}
//line views/vworkflow/Detail.html:148
	}
//line views/vworkflow/Detail.html:149
}

//line views/vworkflow/Detail.html:149
func writeshowEntries(qq422016 qtio422016.Writer, ents har.Entries, ps *cutil.PageState) {
//line views/vworkflow/Detail.html:149
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkflow/Detail.html:149
	streamshowEntries(qw422016, ents, ps)
//line views/vworkflow/Detail.html:149
	qt422016.ReleaseWriter(qw422016)
//line views/vworkflow/Detail.html:149
}

//line views/vworkflow/Detail.html:149
func showEntries(ents har.Entries, ps *cutil.PageState) string {
//line views/vworkflow/Detail.html:149
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkflow/Detail.html:149
	writeshowEntries(qb422016, ents, ps)
//line views/vworkflow/Detail.html:149
	qs422016 := string(qb422016.B)
//line views/vworkflow/Detail.html:149
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkflow/Detail.html:149
	return qs422016
//line views/vworkflow/Detail.html:149
}
