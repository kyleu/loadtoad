// Code generated by qtc from "Sockets.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Sockets.html:2
package vadmin

//line views/vadmin/Sockets.html:2
import (
	"strings"

	"github.com/google/uuid"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/websocket"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vadmin/Sockets.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Sockets.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Sockets.html:14
type Sockets struct {
	layout.Basic
	Channels    []string
	Connections []*websocket.Connection
	Taps        []uuid.UUID
}

//line views/vadmin/Sockets.html:21
func (p *Sockets) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:21
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/admin/sockets/tap"><button>Tap Traffic</button></a></div>
    <h3>Activity Taps</h3>
    <div class="mt">
`)
//line views/vadmin/Sockets.html:26
	if len(p.Taps) == 0 {
//line views/vadmin/Sockets.html:26
		qw422016.N().S(`      <em>no active taps</em>
`)
//line views/vadmin/Sockets.html:28
	} else {
//line views/vadmin/Sockets.html:28
		qw422016.N().S(`      <ul>
`)
//line views/vadmin/Sockets.html:30
		for _, x := range p.Taps {
//line views/vadmin/Sockets.html:30
			qw422016.N().S(`        <li>`)
//line views/vadmin/Sockets.html:31
			qw422016.E().S(x.String())
//line views/vadmin/Sockets.html:31
			qw422016.N().S(`</li>
`)
//line views/vadmin/Sockets.html:32
		}
//line views/vadmin/Sockets.html:32
		qw422016.N().S(`      </ul>
`)
//line views/vadmin/Sockets.html:34
	}
//line views/vadmin/Sockets.html:34
	qw422016.N().S(`    </div>
  </div>
  <div class="card">
    <h3>Channels</h3>
    <div class="mt">
`)
//line views/vadmin/Sockets.html:40
	if len(p.Channels) == 0 {
//line views/vadmin/Sockets.html:40
		qw422016.N().S(`      <em>no active channels</em>
`)
//line views/vadmin/Sockets.html:42
	} else {
//line views/vadmin/Sockets.html:42
		qw422016.N().S(`      <ul>
`)
//line views/vadmin/Sockets.html:44
		for _, x := range p.Channels {
//line views/vadmin/Sockets.html:44
			qw422016.N().S(`        <li><a href="/admin/sockets/chan/`)
//line views/vadmin/Sockets.html:45
			qw422016.E().S(x)
//line views/vadmin/Sockets.html:45
			qw422016.N().S(`">`)
//line views/vadmin/Sockets.html:45
			qw422016.E().S(x)
//line views/vadmin/Sockets.html:45
			qw422016.N().S(`</a></li>
`)
//line views/vadmin/Sockets.html:46
		}
//line views/vadmin/Sockets.html:46
		qw422016.N().S(`      </ul>
`)
//line views/vadmin/Sockets.html:48
	}
//line views/vadmin/Sockets.html:48
	qw422016.N().S(`    </div>
  </div>
  <div class="card">
    <h3>Active Connections</h3>
    <div class="mt">
`)
//line views/vadmin/Sockets.html:54
	if len(p.Connections) == 0 {
//line views/vadmin/Sockets.html:54
		qw422016.N().S(`      <em>no active connections</em>
`)
//line views/vadmin/Sockets.html:56
	} else {
//line views/vadmin/Sockets.html:56
		qw422016.N().S(`      <table class="expanded">
        <thead>
          <tr>
            <th class="shrink">ID</th>
            <th>Profile Name</th>
            <th>Service</th>
            <th>Channels</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vadmin/Sockets.html:68
		for _, x := range p.Connections {
//line views/vadmin/Sockets.html:68
			qw422016.N().S(`          <tr>
            <td class="shrink"><a href="/admin/sockets/conn/`)
//line views/vadmin/Sockets.html:70
			qw422016.E().S(x.ID.String())
//line views/vadmin/Sockets.html:70
			qw422016.N().S(`">`)
//line views/vadmin/Sockets.html:70
			qw422016.E().S(x.ID.String())
//line views/vadmin/Sockets.html:70
			qw422016.N().S(`</a></td>
            <td>`)
//line views/vadmin/Sockets.html:71
			qw422016.E().S(x.Profile.String())
//line views/vadmin/Sockets.html:71
			qw422016.N().S(`</td>
            <td>`)
//line views/vadmin/Sockets.html:72
			qw422016.E().S(x.Svc)
//line views/vadmin/Sockets.html:72
			qw422016.N().S(`</td>
            <td>`)
//line views/vadmin/Sockets.html:73
			qw422016.E().S(strings.Join(x.Channels, ", "))
//line views/vadmin/Sockets.html:73
			qw422016.N().S(`</td>
            <td class="shrink"><a href="#modal-conn-`)
//line views/vadmin/Sockets.html:74
			qw422016.E().S(x.ID.String())
//line views/vadmin/Sockets.html:74
			qw422016.N().S(`"><button type="button">JSON</button></a></td>
          </tr>
`)
//line views/vadmin/Sockets.html:76
		}
//line views/vadmin/Sockets.html:76
		qw422016.N().S(`        </tbody>
      </table>
`)
//line views/vadmin/Sockets.html:79
	}
//line views/vadmin/Sockets.html:79
	qw422016.N().S(`    </div>
  </div>
`)
//line views/vadmin/Sockets.html:82
	for _, x := range p.Connections {
//line views/vadmin/Sockets.html:82
		qw422016.N().S(`  `)
//line views/vadmin/Sockets.html:83
		components.StreamJSONModal(qw422016, "conn-"+x.ID.String(), "Connection ["+x.ID.String()+"] JSON", x, 1)
//line views/vadmin/Sockets.html:83
		qw422016.N().S(`
`)
//line views/vadmin/Sockets.html:84
	}
//line views/vadmin/Sockets.html:85
}

//line views/vadmin/Sockets.html:85
func (p *Sockets) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:85
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Sockets.html:85
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Sockets.html:85
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Sockets.html:85
}

//line views/vadmin/Sockets.html:85
func (p *Sockets) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Sockets.html:85
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Sockets.html:85
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Sockets.html:85
	qs422016 := string(qb422016.B)
//line views/vadmin/Sockets.html:85
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Sockets.html:85
	return qs422016
//line views/vadmin/Sockets.html:85
}

//line views/vadmin/Sockets.html:87
type Channel struct {
	layout.Basic
	Channel *websocket.Channel
	Members []*websocket.Connection
}

//line views/vadmin/Sockets.html:93
func (p *Channel) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:93
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-channel"><button type="button">JSON</button></a></div>
    <h3>Channel [`)
//line views/vadmin/Sockets.html:96
	qw422016.E().S(p.Channel.Key)
//line views/vadmin/Sockets.html:96
	qw422016.N().S(`]</h3>
  </div>
  `)
//line views/vadmin/Sockets.html:98
	components.StreamJSONModal(qw422016, "channel", "Channel ["+p.Channel.Key+"] JSON", p.Channel, 1)
//line views/vadmin/Sockets.html:98
	qw422016.N().S(`
`)
//line views/vadmin/Sockets.html:99
	for _, m := range p.Members {
//line views/vadmin/Sockets.html:99
		qw422016.N().S(`  `)
//line views/vadmin/Sockets.html:100
		StreamConnectionCard(qw422016, m, as, ps)
//line views/vadmin/Sockets.html:100
		qw422016.N().S(`
`)
//line views/vadmin/Sockets.html:101
	}
//line views/vadmin/Sockets.html:102
}

//line views/vadmin/Sockets.html:102
func (p *Channel) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:102
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Sockets.html:102
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Sockets.html:102
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Sockets.html:102
}

//line views/vadmin/Sockets.html:102
func (p *Channel) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Sockets.html:102
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Sockets.html:102
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Sockets.html:102
	qs422016 := string(qb422016.B)
//line views/vadmin/Sockets.html:102
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Sockets.html:102
	return qs422016
//line views/vadmin/Sockets.html:102
}

//line views/vadmin/Sockets.html:104
type Connection struct {
	layout.Basic
	Connection *websocket.Connection
}

//line views/vadmin/Sockets.html:109
func (p *Connection) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:109
	qw422016.N().S(`
  `)
//line views/vadmin/Sockets.html:110
	StreamConnectionCard(qw422016, p.Connection, as, ps)
//line views/vadmin/Sockets.html:110
	qw422016.N().S(`
  <div class="card">
    <h3>Send Message</h3>
    <form action="/admin/sockets/conn/`)
//line views/vadmin/Sockets.html:113
	qw422016.E().S(p.Connection.ID.String())
//line views/vadmin/Sockets.html:113
	qw422016.N().S(`/send" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vadmin/Sockets.html:116
	components.StreamTableInput(qw422016, "from", "", "From", p.Connection.ID.String(), 5, "The connection this message is from")
//line views/vadmin/Sockets.html:116
	qw422016.N().S(`
          `)
//line views/vadmin/Sockets.html:117
	components.StreamTableInput(qw422016, "channel", "", "Channel", "", 5, "The channel this message is for")
//line views/vadmin/Sockets.html:117
	qw422016.N().S(`
          `)
//line views/vadmin/Sockets.html:118
	components.StreamTableInput(qw422016, "cmd", "", "Command", "", 5, "The command for this message")
//line views/vadmin/Sockets.html:118
	qw422016.N().S(`
          `)
//line views/vadmin/Sockets.html:119
	components.StreamTableTextarea(qw422016, "param", "", "Parameter", 8, "", 5, "JSON object message payload")
//line views/vadmin/Sockets.html:119
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Send</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
`)
//line views/vadmin/Sockets.html:125
}

//line views/vadmin/Sockets.html:125
func (p *Connection) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:125
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Sockets.html:125
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Sockets.html:125
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Sockets.html:125
}

//line views/vadmin/Sockets.html:125
func (p *Connection) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Sockets.html:125
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Sockets.html:125
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Sockets.html:125
	qs422016 := string(qb422016.B)
//line views/vadmin/Sockets.html:125
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Sockets.html:125
	return qs422016
//line views/vadmin/Sockets.html:125
}

//line views/vadmin/Sockets.html:127
func StreamConnectionCard(qw422016 *qt422016.Writer, c *websocket.Connection, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:127
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-connection-`)
//line views/vadmin/Sockets.html:129
	qw422016.E().S(c.ID.String())
//line views/vadmin/Sockets.html:129
	qw422016.N().S(`"><button type="button">JSON</button></a></div>
    <h3>`)
//line views/vadmin/Sockets.html:130
	qw422016.E().S(c.ID.String())
//line views/vadmin/Sockets.html:130
	qw422016.N().S(` (`)
//line views/vadmin/Sockets.html:130
	qw422016.E().S(c.Username())
//line views/vadmin/Sockets.html:130
	qw422016.N().S(`)</h3>
    <table class="mt expanded">
      <tbody>
        <tr>
          <th>Connection ID</th>
          <td>`)
//line views/vadmin/Sockets.html:135
	qw422016.E().S(c.ID.String())
//line views/vadmin/Sockets.html:135
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th>Name</th>
          <td>`)
//line views/vadmin/Sockets.html:139
	qw422016.E().S(c.Username())
//line views/vadmin/Sockets.html:139
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th>Theme</th>
          <td>`)
//line views/vadmin/Sockets.html:143
	qw422016.E().S(c.Profile.Theme)
//line views/vadmin/Sockets.html:143
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th>Service</th>
          <td>`)
//line views/vadmin/Sockets.html:147
	qw422016.E().S(c.Svc)
//line views/vadmin/Sockets.html:147
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th>Channels</th>
          <td>`)
//line views/vadmin/Sockets.html:151
	qw422016.E().S(strings.Join(c.Channels, ", "))
//line views/vadmin/Sockets.html:151
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
  `)
//line views/vadmin/Sockets.html:156
	components.StreamJSONModal(qw422016, "connection-"+c.ID.String(), "Connection ["+c.ID.String()+"] JSON", c, 1)
//line views/vadmin/Sockets.html:156
	qw422016.N().S(`
`)
//line views/vadmin/Sockets.html:157
}

//line views/vadmin/Sockets.html:157
func WriteConnectionCard(qq422016 qtio422016.Writer, c *websocket.Connection, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:157
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Sockets.html:157
	StreamConnectionCard(qw422016, c, as, ps)
//line views/vadmin/Sockets.html:157
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Sockets.html:157
}

//line views/vadmin/Sockets.html:157
func ConnectionCard(c *websocket.Connection, as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Sockets.html:157
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Sockets.html:157
	WriteConnectionCard(qb422016, c, as, ps)
//line views/vadmin/Sockets.html:157
	qs422016 := string(qb422016.B)
//line views/vadmin/Sockets.html:157
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Sockets.html:157
	return qs422016
//line views/vadmin/Sockets.html:157
}
