// Code generated by qtc from "Start.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkflow/Start.html:1
package vworkflow

//line views/vworkflow/Start.html:1
import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vworkflow/Start.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkflow/Start.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkflow/Start.html:10
type Start struct {
	layout.Basic
	Workflow *loadtoad.Workflow
	Entries  har.Entries
	Channel  string
}

//line views/vworkflow/Start.html:17
func (p *Start) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkflow/Start.html:17
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-workflow"><button type="button">JSON</button></a></div>
    `)
//line views/vworkflow/Start.html:20
	components.StreamJSONModal(qw422016, "workflow", "Workflow", p.Workflow, 3)
//line views/vworkflow/Start.html:20
	qw422016.N().S(`
    <h3>`)
//line views/vworkflow/Start.html:21
	components.StreamSVGRefIcon(qw422016, `sitemap`, ps)
//line views/vworkflow/Start.html:21
	qw422016.N().S(` `)
//line views/vworkflow/Start.html:21
	qw422016.E().S(p.Workflow.Title())
//line views/vworkflow/Start.html:21
	qw422016.N().S(` Results</h3>
    <div class="mt">
`)
//line views/vworkflow/Start.html:23
	if len(p.Entries) == 0 {
//line views/vworkflow/Start.html:23
		qw422016.N().S(`      <div><em>no entries</em></div>
`)
//line views/vworkflow/Start.html:25
	}
//line views/vworkflow/Start.html:25
	qw422016.N().S(`    </div>
    <div id="results" class="mt">
    </div>
  </div>
`)
//line views/vworkflow/Start.html:30
	for idx, e := range p.Entries {
//line views/vworkflow/Start.html:30
		qw422016.N().S(`  `)
//line views/vworkflow/Start.html:31
		streamentryPlaceholder(qw422016, idx, e, ps)
//line views/vworkflow/Start.html:31
		qw422016.N().S(`
`)
//line views/vworkflow/Start.html:32
	}
//line views/vworkflow/Start.html:32
	qw422016.N().S(`  <script>
    var sock;
    function open() {
      console.log("[Open]");
    }
    function recv(m) {
      let el = document.querySelector("#entry-placeholder-" + m.param.idx);
      if (!el) {
        if (m.param.idx === -1) {
          el = document.querySelector("#results");
        } else {
          throw "invalid index [" + m.param.idx + "]";
        }
      }
      switch (m.cmd) {
        case "log":
          const logs = el.querySelector(".logs");
          if (!logs) {
            throw "invalid logs for index [" + m.param.idx + "]";
          }
          const logEl = document.createElement("div");
          logEl.innerText = " - " + m.param.ctx;
          logs.appendChild(logEl);
          break;
        case "error":
          throw "error [" + m.param.ctx + "] encountered for index [" + m.param.idx + "]";
        case "ok":
          const deets = el.querySelector(".details");
          if (!deets) {
            throw "invalid details for index [" + m.param.idx + "]";
          }
          const deetsDiv = document.createElement("div");
          deetsDiv.innerHTML = m.param.ctx;
          deets.replaceChildren(deetsDiv);
          break;
        case "complete":
          const completeEl = document.createElement("div");
          completeEl.innerText = m.param.ctx;
          el.appendChild(completeEl);
          break;
        default:
          console.log("[Recv]", m.param.idx, m.cmd);
          break;
      }
    }
    function err(e) {
      console.log("[Err]", e);
    }
    document.addEventListener("DOMContentLoaded", function() {
      sock = new loadtoad.Socket(false, open, recv, err, "/workflow/`)
//line views/vworkflow/Start.html:82
	qw422016.E().S(p.Workflow.ID)
//line views/vworkflow/Start.html:82
	qw422016.N().S(`/connect?channel=`)
//line views/vworkflow/Start.html:82
	qw422016.E().S(p.Channel)
//line views/vworkflow/Start.html:82
	qw422016.N().S(`");
      console.log("[Load]");
    });
  </script>
`)
//line views/vworkflow/Start.html:86
}

//line views/vworkflow/Start.html:86
func (p *Start) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkflow/Start.html:86
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkflow/Start.html:86
	p.StreamBody(qw422016, as, ps)
//line views/vworkflow/Start.html:86
	qt422016.ReleaseWriter(qw422016)
//line views/vworkflow/Start.html:86
}

//line views/vworkflow/Start.html:86
func (p *Start) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkflow/Start.html:86
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkflow/Start.html:86
	p.WriteBody(qb422016, as, ps)
//line views/vworkflow/Start.html:86
	qs422016 := string(qb422016.B)
//line views/vworkflow/Start.html:86
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkflow/Start.html:86
	return qs422016
//line views/vworkflow/Start.html:86
}

//line views/vworkflow/Start.html:88
func streamentryPlaceholder(qw422016 *qt422016.Writer, idx int, ent *har.Entry, ps *cutil.PageState) {
//line views/vworkflow/Start.html:88
	qw422016.N().S(`
  <div id="entry-placeholder-`)
//line views/vworkflow/Start.html:89
	qw422016.N().D(idx)
//line views/vworkflow/Start.html:89
	qw422016.N().S(`" class="card">
    <h3>`)
//line views/vworkflow/Start.html:90
	components.StreamSVGRefIcon(qw422016, `file`, ps)
//line views/vworkflow/Start.html:90
	qw422016.E().S(ent.String())
//line views/vworkflow/Start.html:90
	qw422016.N().S(`</h3>
    <div class="clear"></div>
    <div class="mts details">
      <em>pending...</em>
    </div>
    <div class="logs">
    </div>
  </div>
`)
//line views/vworkflow/Start.html:98
}

//line views/vworkflow/Start.html:98
func writeentryPlaceholder(qq422016 qtio422016.Writer, idx int, ent *har.Entry, ps *cutil.PageState) {
//line views/vworkflow/Start.html:98
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkflow/Start.html:98
	streamentryPlaceholder(qw422016, idx, ent, ps)
//line views/vworkflow/Start.html:98
	qt422016.ReleaseWriter(qw422016)
//line views/vworkflow/Start.html:98
}

//line views/vworkflow/Start.html:98
func entryPlaceholder(idx int, ent *har.Entry, ps *cutil.PageState) string {
//line views/vworkflow/Start.html:98
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkflow/Start.html:98
	writeentryPlaceholder(qb422016, idx, ent, ps)
//line views/vworkflow/Start.html:98
	qs422016 := string(qb422016.B)
//line views/vworkflow/Start.html:98
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkflow/Start.html:98
	return qs422016
//line views/vworkflow/Start.html:98
}
