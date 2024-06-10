// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vexec/Detail.html:2
package vexec

//line views/vexec/Detail.html:2
import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/exec"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vexec/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vexec/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vexec/Detail.html:11
type Detail struct {
	layout.Basic
	Exec *exec.Exec
}

//line views/vexec/Detail.html:16
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexec/Detail.html:16
	qw422016.N().S(`
  `)
//line views/vexec/Detail.html:17
	StreamExecDetail(qw422016, p.Exec, ps)
//line views/vexec/Detail.html:17
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vexec/Detail.html:19
	components.StreamSVGIcon(qw422016, "file", ps)
//line views/vexec/Detail.html:19
	qw422016.N().S(`Output</h3>
    <div class="mt">`)
//line views/vexec/Detail.html:20
	components.StreamTerminal(qw422016, "console-list", p.Exec.Buffer.String())
//line views/vexec/Detail.html:20
	qw422016.N().S(`</div>
  </div>
  `)
//line views/vexec/Detail.html:22
	StreamExecScript(qw422016, "console-list", p.Exec.WebPath()+"/connect", ps)
//line views/vexec/Detail.html:22
	qw422016.N().S(`
`)
//line views/vexec/Detail.html:23
}

//line views/vexec/Detail.html:23
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexec/Detail.html:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexec/Detail.html:23
	p.StreamBody(qw422016, as, ps)
//line views/vexec/Detail.html:23
	qt422016.ReleaseWriter(qw422016)
//line views/vexec/Detail.html:23
}

//line views/vexec/Detail.html:23
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vexec/Detail.html:23
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexec/Detail.html:23
	p.WriteBody(qb422016, as, ps)
//line views/vexec/Detail.html:23
	qs422016 := string(qb422016.B)
//line views/vexec/Detail.html:23
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexec/Detail.html:23
	return qs422016
//line views/vexec/Detail.html:23
}

//line views/vexec/Detail.html:25
func StreamExecDetail(qw422016 *qt422016.Writer, ex *exec.Exec, ps *cutil.PageState) {
//line views/vexec/Detail.html:25
	qw422016.N().S(`
  <div class="card">
`)
//line views/vexec/Detail.html:27
	if ex.Completed == nil {
//line views/vexec/Detail.html:27
		qw422016.N().S(`    <div class="right">
      <a href="`)
//line views/vexec/Detail.html:29
		qw422016.E().S(ex.WebPath())
//line views/vexec/Detail.html:29
		qw422016.N().S(`/kill"><button>Kill</button></a>
    </div>
`)
//line views/vexec/Detail.html:31
	}
//line views/vexec/Detail.html:31
	qw422016.N().S(`    <h3>`)
//line views/vexec/Detail.html:32
	components.StreamSVGIcon(qw422016, "desktop", ps)
//line views/vexec/Detail.html:32
	qw422016.N().S(`Process [`)
//line views/vexec/Detail.html:32
	qw422016.E().S(ex.String())
//line views/vexec/Detail.html:32
	qw422016.N().S(`]</h3>
    <div class="overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink">Key</th>
            <td>`)
//line views/vexec/Detail.html:38
	if len(ex.Link) > 0 {
//line views/vexec/Detail.html:38
		qw422016.N().S(`<a href="`)
//line views/vexec/Detail.html:38
		qw422016.E().S(ex.Link)
//line views/vexec/Detail.html:38
		qw422016.N().S(`">`)
//line views/vexec/Detail.html:38
		qw422016.E().S(ex.Key)
//line views/vexec/Detail.html:38
		qw422016.N().S(`</a>`)
//line views/vexec/Detail.html:38
	} else {
//line views/vexec/Detail.html:38
		qw422016.E().S(ex.Key)
//line views/vexec/Detail.html:38
	}
//line views/vexec/Detail.html:38
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Index</th>
            <td>`)
//line views/vexec/Detail.html:42
	qw422016.N().D(ex.Idx)
//line views/vexec/Detail.html:42
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>PID</th>
            <td>`)
//line views/vexec/Detail.html:46
	qw422016.N().D(ex.PID)
//line views/vexec/Detail.html:46
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Command</th>
            <td>`)
//line views/vexec/Detail.html:50
	qw422016.E().S(ex.Cmd)
//line views/vexec/Detail.html:50
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Path</th>
            <td>`)
//line views/vexec/Detail.html:54
	qw422016.E().S(ex.Path)
//line views/vexec/Detail.html:54
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Environment</th>
            <td>
              <ul>
`)
//line views/vexec/Detail.html:60
	for _, x := range ex.Env {
//line views/vexec/Detail.html:60
		qw422016.N().S(`                <li>`)
//line views/vexec/Detail.html:61
		qw422016.E().S(x)
//line views/vexec/Detail.html:61
		qw422016.N().S(`</li>
`)
//line views/vexec/Detail.html:62
	}
//line views/vexec/Detail.html:62
	qw422016.N().S(`              </ul>
            </td>
          </tr>
          <tr>
            <th>Started</th>
            <td title="`)
//line views/vexec/Detail.html:68
	qw422016.E().S(util.TimeToFull(ex.Started))
//line views/vexec/Detail.html:68
	qw422016.N().S(`">`)
//line views/vexec/Detail.html:68
	qw422016.E().S(util.TimeRelative(ex.Started))
//line views/vexec/Detail.html:68
	qw422016.N().S(`</td>
          </tr>
`)
//line views/vexec/Detail.html:70
	if ex.Completed != nil {
//line views/vexec/Detail.html:70
		qw422016.N().S(`          <tr>
            <th>Completed</th>
            <td title="`)
//line views/vexec/Detail.html:73
		qw422016.E().S(util.TimeToFull(ex.Completed))
//line views/vexec/Detail.html:73
		qw422016.N().S(`">`)
//line views/vexec/Detail.html:73
		qw422016.E().S(util.TimeRelative(ex.Completed))
//line views/vexec/Detail.html:73
		qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Exit Code</th>
            <td>`)
//line views/vexec/Detail.html:77
		qw422016.N().D(ex.ExitCode)
//line views/vexec/Detail.html:77
		qw422016.N().S(`</td>
          </tr>
`)
//line views/vexec/Detail.html:79
	}
//line views/vexec/Detail.html:79
	qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vexec/Detail.html:84
}

//line views/vexec/Detail.html:84
func WriteExecDetail(qq422016 qtio422016.Writer, ex *exec.Exec, ps *cutil.PageState) {
//line views/vexec/Detail.html:84
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexec/Detail.html:84
	StreamExecDetail(qw422016, ex, ps)
//line views/vexec/Detail.html:84
	qt422016.ReleaseWriter(qw422016)
//line views/vexec/Detail.html:84
}

//line views/vexec/Detail.html:84
func ExecDetail(ex *exec.Exec, ps *cutil.PageState) string {
//line views/vexec/Detail.html:84
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexec/Detail.html:84
	WriteExecDetail(qb422016, ex, ps)
//line views/vexec/Detail.html:84
	qs422016 := string(qb422016.B)
//line views/vexec/Detail.html:84
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexec/Detail.html:84
	return qs422016
//line views/vexec/Detail.html:84
}

//line views/vexec/Detail.html:86
func StreamExecScript(qw422016 *qt422016.Writer, id string, connectURL string, ps *cutil.PageState) {
//line views/vexec/Detail.html:86
	qw422016.N().S(`
  <script>
    function open() {
      console.log("[socket]: open");
    }
    function recv(m) {
      const tbody = document.getElementById("`)
//line views/vexec/Detail.html:92
	qw422016.E().S(id)
//line views/vexec/Detail.html:92
	qw422016.N().S(`");
      const h = m.param["html"].split("\n");
      for (x in h) {
        const row = document.createElement("tr");
        const numTH = document.createElement("th");
        numTH.innerText = (tbody.children.length).toString();
        const textTD = document.createElement("td");
        textTD.innerHTML = h[x];
        row.append(numTH, textTD);
        tbody.append(row);
      }
      const c = document.getElementById("content");
      c.scrollTo(0, c.scrollHeight);
    }
    function err(e) {
      console.log("[socket error]: " + e);
    }
    window.addEventListener('load', () => {
      new loadtoad.Socket(true, open, recv, err, "`)
//line views/vexec/Detail.html:110
	qw422016.E().S(connectURL)
//line views/vexec/Detail.html:110
	qw422016.N().S(`");
    })
  </script>
`)
//line views/vexec/Detail.html:113
}

//line views/vexec/Detail.html:113
func WriteExecScript(qq422016 qtio422016.Writer, id string, connectURL string, ps *cutil.PageState) {
//line views/vexec/Detail.html:113
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexec/Detail.html:113
	StreamExecScript(qw422016, id, connectURL, ps)
//line views/vexec/Detail.html:113
	qt422016.ReleaseWriter(qw422016)
//line views/vexec/Detail.html:113
}

//line views/vexec/Detail.html:113
func ExecScript(id string, connectURL string, ps *cutil.PageState) string {
//line views/vexec/Detail.html:113
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexec/Detail.html:113
	WriteExecScript(qb422016, id, connectURL, ps)
//line views/vexec/Detail.html:113
	qs422016 := string(qb422016.B)
//line views/vexec/Detail.html:113
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexec/Detail.html:113
	return qs422016
//line views/vexec/Detail.html:113
}
