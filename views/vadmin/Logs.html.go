// Code generated by qtc from "Logs.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->
// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Logs.html:3
package vadmin

//line views/vadmin/Logs.html:3
import (
	"go.uber.org/zap/zapcore"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vadmin/Logs.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Logs.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Logs.html:12
type Logs struct {
	layout.Basic
	Logs []*zapcore.Entry
}

//line views/vadmin/Logs.html:17
func (p *Logs) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Logs.html:17
	qw422016.N().S(`
  <div class="card">
    <h3>Recent Logs</h3>
    `)
//line views/vadmin/Logs.html:20
	streamlogTable(qw422016, p.Logs)
//line views/vadmin/Logs.html:20
	qw422016.N().S(`
  </div>
`)
//line views/vadmin/Logs.html:22
}

//line views/vadmin/Logs.html:22
func (p *Logs) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Logs.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Logs.html:22
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Logs.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Logs.html:22
}

//line views/vadmin/Logs.html:22
func (p *Logs) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Logs.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Logs.html:22
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Logs.html:22
	qs422016 := string(qb422016.B)
//line views/vadmin/Logs.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Logs.html:22
	return qs422016
//line views/vadmin/Logs.html:22
}

//line views/vadmin/Logs.html:24
func streamlogTable(qw422016 *qt422016.Writer, logs []*zapcore.Entry) {
//line views/vadmin/Logs.html:24
	qw422016.N().S(`
  <table class="mt">
    <thead>
      <tr>
        <th>Level</th>
        <th>Message</th>
        <th>Occurred</th>
      </tr>
    </thead>
    <tbody>
`)
//line views/vadmin/Logs.html:34
	for _, l := range logs {
//line views/vadmin/Logs.html:34
		qw422016.N().S(`      <tr>
        <td>`)
//line views/vadmin/Logs.html:36
		qw422016.E().S(l.Level.String())
//line views/vadmin/Logs.html:36
		qw422016.N().S(`</td>
        <td>`)
//line views/vadmin/Logs.html:37
		qw422016.E().S(l.Message)
//line views/vadmin/Logs.html:37
		qw422016.N().S(`</td>
        <td>`)
//line views/vadmin/Logs.html:38
		qw422016.E().S(util.TimeRelative(&l.Time))
//line views/vadmin/Logs.html:38
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vadmin/Logs.html:40
	}
//line views/vadmin/Logs.html:40
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vadmin/Logs.html:43
}

//line views/vadmin/Logs.html:43
func writelogTable(qq422016 qtio422016.Writer, logs []*zapcore.Entry) {
//line views/vadmin/Logs.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Logs.html:43
	streamlogTable(qw422016, logs)
//line views/vadmin/Logs.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Logs.html:43
}

//line views/vadmin/Logs.html:43
func logTable(logs []*zapcore.Entry) string {
//line views/vadmin/Logs.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Logs.html:43
	writelogTable(qb422016, logs)
//line views/vadmin/Logs.html:43
	qs422016 := string(qb422016.B)
//line views/vadmin/Logs.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Logs.html:43
	return qs422016
//line views/vadmin/Logs.html:43
}
