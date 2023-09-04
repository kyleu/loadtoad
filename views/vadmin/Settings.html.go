// Code generated by qtc from "Settings.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Settings.html:2
package vadmin

//line views/vadmin/Settings.html:2
import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vadmin/Settings.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Settings.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Settings.html:10
type Settings struct {
	layout.Basic
}

//line views/vadmin/Settings.html:14
func (p *Settings) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:14
	qw422016.N().S(`
  <div class="card">
`)
//line views/vadmin/Settings.html:16
	if util.AppSource != "" {
//line views/vadmin/Settings.html:16
		qw422016.N().S(`    <div class="right"><a target="_blank" rel="noopener noreferrer" href="`)
//line views/vadmin/Settings.html:17
		qw422016.E().S(util.AppSource)
//line views/vadmin/Settings.html:17
		qw422016.N().S(`"><button>Github</button></a></div>
`)
//line views/vadmin/Settings.html:18
	}
//line views/vadmin/Settings.html:18
	qw422016.N().S(`    <h3 title="github:`)
//line views/vadmin/Settings.html:19
	qw422016.E().S(as.BuildInfo.Commit)
//line views/vadmin/Settings.html:19
	qw422016.N().S(`">`)
//line views/vadmin/Settings.html:19
	components.StreamSVGRefIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:19
	qw422016.E().S(util.AppName)
//line views/vadmin/Settings.html:19
	qw422016.N().S(` `)
//line views/vadmin/Settings.html:19
	qw422016.E().S(as.BuildInfo.String())
//line views/vadmin/Settings.html:19
	qw422016.N().S(`</h3>
`)
//line views/vadmin/Settings.html:20
	if util.AppLegal != "" {
//line views/vadmin/Settings.html:20
		qw422016.N().S(`    <div class="mt">`)
//line views/vadmin/Settings.html:21
		qw422016.N().S(util.AppLegal)
//line views/vadmin/Settings.html:21
		qw422016.N().S(`</div>
`)
//line views/vadmin/Settings.html:22
	}
//line views/vadmin/Settings.html:23
	if util.AppURL != "" {
//line views/vadmin/Settings.html:23
		qw422016.N().S(`    <p><a target="_blank" rel="noopener noreferrer" href="`)
//line views/vadmin/Settings.html:24
		qw422016.N().S(util.AppURL)
//line views/vadmin/Settings.html:24
		qw422016.N().S(`">`)
//line views/vadmin/Settings.html:24
		qw422016.N().S(util.AppURL)
//line views/vadmin/Settings.html:24
		qw422016.N().S(`</a></p>
`)
//line views/vadmin/Settings.html:25
	}
//line views/vadmin/Settings.html:25
	qw422016.N().S(`  </div>

  <div class="flex-wrap flex-align-stretch">
    <div class="card flex-grow-1 flex-basis-0">
      <h3>`)
//line views/vadmin/Settings.html:30
	components.StreamSVGRefIcon(qw422016, `archive`, ps)
//line views/vadmin/Settings.html:30
	qw422016.N().S(`Admin Functions</h3>
      <ul class="mt">
        <li><a href="/admin/server">App Information</a></li>
        <li><a href="/admin/modules">View Go modules</a></li>
        <li><a href="/theme">Edit Themes</a></li>
        <li><a href="/admin/exec">Managed Processes</a></li>
        <li><a href="/admin/scripting">Script Files</a></li>
        <li><a href="/admin/logs">Recent Logs</a></li>
      </ul>
    </div>
    <div class="card flex-grow-1 flex-basis-0">
      <h3>`)
//line views/vadmin/Settings.html:41
	components.StreamSVGRefIcon(qw422016, `bolt`, ps)
//line views/vadmin/Settings.html:41
	qw422016.N().S(`HTTP Methods</h3>
      <ul class="mt">
        <li><a href="/admin/sitemap">Sitemap</a></li>
        <li><a href="/admin/routes">View HTTP routes</a></li>
        <li><a href="/admin/session">Parse and display session</a></li>
        <li><a href="/admin/request">Debug HTTP request</a></li>
        <li><a href="/admin/sockets">Active WebSockets</a></li>
      </ul>
    </div>
    <div class="card flex-grow-1 flex-basis-0">
      <h3>`)
//line views/vadmin/Settings.html:51
	components.StreamSVGRefIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:51
	qw422016.N().S(`App Profiling</h3>
      <ul class="mt">
        <li><a href="/admin/memusage">Memory Usage</a></li>
        <li><a href="/admin/gc">Collect garbage</a></li>
        <li><a href="/admin/heap">Write memory dump</a></li>
        <li><a href="/admin/cpu/start">Start CPU profile</a></li>
        <li><a href="/admin/cpu/stop">Stop CPU profile</a></li>
      </ul>
    </div>
  </div>
`)
//line views/vadmin/Settings.html:61
}

//line views/vadmin/Settings.html:61
func (p *Settings) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Settings.html:61
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Settings.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Settings.html:61
}

//line views/vadmin/Settings.html:61
func (p *Settings) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Settings.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Settings.html:61
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Settings.html:61
	qs422016 := string(qb422016.B)
//line views/vadmin/Settings.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Settings.html:61
	return qs422016
//line views/vadmin/Settings.html:61
}
