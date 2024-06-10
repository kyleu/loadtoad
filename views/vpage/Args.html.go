// Code generated by qtc from "Args.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vpage/Args.html:2
package vpage

//line views/vpage/Args.html:2
import (
	"strconv"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components/edit"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vpage/Args.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vpage/Args.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vpage/Args.html:12
type Args struct {
	layout.Basic
	URL        string
	Directions string
	ArgRes     *cutil.ArgResults
	Hidden     map[string]string
}

//line views/vpage/Args.html:20
func (p *Args) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vpage/Args.html:20
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vpage/Args.html:22
	if p.Directions == "" {
//line views/vpage/Args.html:22
		qw422016.N().S(`Enter Data`)
//line views/vpage/Args.html:22
	} else {
//line views/vpage/Args.html:22
		qw422016.E().S(p.Directions)
//line views/vpage/Args.html:22
	}
//line views/vpage/Args.html:22
	qw422016.N().S(`</h3>
    <form action="`)
//line views/vpage/Args.html:23
	qw422016.E().S(p.URL)
//line views/vpage/Args.html:23
	qw422016.N().S(`" method="get">
`)
//line views/vpage/Args.html:24
	for k, v := range p.Hidden {
//line views/vpage/Args.html:24
		qw422016.N().S(`      <input type="hidden" name="`)
//line views/vpage/Args.html:25
		qw422016.E().S(k)
//line views/vpage/Args.html:25
		qw422016.N().S(`" value="`)
//line views/vpage/Args.html:25
		qw422016.E().S(v)
//line views/vpage/Args.html:25
		qw422016.N().S(`" />
`)
//line views/vpage/Args.html:26
	}
//line views/vpage/Args.html:26
	qw422016.N().S(`      <div class="overflow full-width">
        <table class="mt min-200 expanded">
          <tbody>
`)
//line views/vpage/Args.html:30
	for _, arg := range p.ArgRes.Args {
//line views/vpage/Args.html:32
		v := util.OrDefault(p.ArgRes.Values.GetStringOpt(arg.Key), arg.Default)
		title := arg.Title
		if len(title) > 50 {
			title = title[:50] + "..."
		}

//line views/vpage/Args.html:38
		switch arg.Type {
//line views/vpage/Args.html:39
		case "bool":
//line views/vpage/Args.html:39
			qw422016.N().S(`            `)
//line views/vpage/Args.html:40
			edit.StreamBoolTable(qw422016, arg.Key, title, v == "true", 5, arg.Description)
//line views/vpage/Args.html:40
			qw422016.N().S(`
`)
//line views/vpage/Args.html:41
		case "textarea":
//line views/vpage/Args.html:41
			qw422016.N().S(`            `)
//line views/vpage/Args.html:42
			edit.StreamTextareaTable(qw422016, arg.Key, "", title, 12, v, 5, arg.Description)
//line views/vpage/Args.html:42
			qw422016.N().S(`
`)
//line views/vpage/Args.html:43
		case "number", "int":
//line views/vpage/Args.html:44
			i, _ := strconv.ParseInt(v, 10, 32)

//line views/vpage/Args.html:44
			qw422016.N().S(`            `)
//line views/vpage/Args.html:45
			edit.StreamIntTable(qw422016, arg.Key, "", title, int(i), 5, arg.Description)
//line views/vpage/Args.html:45
			qw422016.N().S(`
`)
//line views/vpage/Args.html:46
		default:
//line views/vpage/Args.html:46
			qw422016.N().S(`            `)
//line views/vpage/Args.html:47
			edit.StreamDatalistTable(qw422016, arg.Key, "", title, v, arg.Choices, nil, 5, arg.Description)
//line views/vpage/Args.html:47
			qw422016.N().S(`
`)
//line views/vpage/Args.html:48
		}
//line views/vpage/Args.html:49
	}
//line views/vpage/Args.html:49
	qw422016.N().S(`          </tbody>
        </table>
      </div>
      <button class="mt" type="submit">Submit</button>
    </form>
  </div>
`)
//line views/vpage/Args.html:56
}

//line views/vpage/Args.html:56
func (p *Args) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vpage/Args.html:56
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vpage/Args.html:56
	p.StreamBody(qw422016, as, ps)
//line views/vpage/Args.html:56
	qt422016.ReleaseWriter(qw422016)
//line views/vpage/Args.html:56
}

//line views/vpage/Args.html:56
func (p *Args) Body(as *app.State, ps *cutil.PageState) string {
//line views/vpage/Args.html:56
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vpage/Args.html:56
	p.WriteBody(qb422016, as, ps)
//line views/vpage/Args.html:56
	qs422016 := string(qb422016.B)
//line views/vpage/Args.html:56
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vpage/Args.html:56
	return qs422016
//line views/vpage/Args.html:56
}
