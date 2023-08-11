// Code generated by qtc from "Entry.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vhar/Entry.html:1
package vhar

//line views/vhar/Entry.html:1
import (
	"fmt"

	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad/har"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
)

//line views/vhar/Entry.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vhar/Entry.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vhar/Entry.html:10
func StreamRenderEntryOptions(qw422016 *qt422016.Writer, i int, e *har.Entry, js bool) {
//line views/vhar/Entry.html:10
	qw422016.N().S(`
  <div class="right">
    <a href="#modal-entry-`)
//line views/vhar/Entry.html:12
	qw422016.N().D(i)
//line views/vhar/Entry.html:12
	qw422016.N().S(`-curl"><button type="button">cURL</button></a>
`)
//line views/vhar/Entry.html:13
	if js {
//line views/vhar/Entry.html:13
		qw422016.N().S(`    <a href="#modal-entry-`)
//line views/vhar/Entry.html:14
		qw422016.N().D(i)
//line views/vhar/Entry.html:14
		qw422016.N().S(`"><button type="button">JSON</button></a>
`)
//line views/vhar/Entry.html:15
	}
//line views/vhar/Entry.html:15
	qw422016.N().S(`  </div>
`)
//line views/vhar/Entry.html:17
	if js {
//line views/vhar/Entry.html:17
		qw422016.N().S(`  `)
//line views/vhar/Entry.html:18
		components.StreamJSONModal(qw422016, fmt.Sprintf("entry-%d", i), e.String(), e.Cleaned(), 3)
//line views/vhar/Entry.html:18
		qw422016.N().S(`
`)
//line views/vhar/Entry.html:19
	}
//line views/vhar/Entry.html:19
	qw422016.N().S(`  <div id="modal-entry-`)
//line views/vhar/Entry.html:20
	qw422016.N().D(i)
//line views/vhar/Entry.html:20
	qw422016.N().S(`-curl" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>cURL `)
//line views/vhar/Entry.html:25
	qw422016.E().S(e.String())
//line views/vhar/Entry.html:25
	qw422016.N().S(`</h2>
      </div>
      <div class="modal-body">
        <div id="modal-entry-`)
//line views/vhar/Entry.html:28
	qw422016.N().D(i)
//line views/vhar/Entry.html:28
	qw422016.N().S(`-curl-data" hidden="hidden" style="display:none;">`)
//line views/vhar/Entry.html:28
	qw422016.E().S(e.Curl())
//line views/vhar/Entry.html:28
	qw422016.N().S(`</div>
        <button onclick="clip('`)
//line views/vhar/Entry.html:29
	qw422016.N().D(i)
//line views/vhar/Entry.html:29
	qw422016.N().S(`');">Copy to clipboard</button>
        <div class="mt">
          <pre>`)
//line views/vhar/Entry.html:31
	qw422016.E().S(e.Curl())
//line views/vhar/Entry.html:31
	qw422016.N().S(`</pre>
        </div>
      </div>
    </div>
  </div>
`)
//line views/vhar/Entry.html:36
}

//line views/vhar/Entry.html:36
func WriteRenderEntryOptions(qq422016 qtio422016.Writer, i int, e *har.Entry, js bool) {
//line views/vhar/Entry.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Entry.html:36
	StreamRenderEntryOptions(qw422016, i, e, js)
//line views/vhar/Entry.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Entry.html:36
}

//line views/vhar/Entry.html:36
func RenderEntryOptions(i int, e *har.Entry, js bool) string {
//line views/vhar/Entry.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Entry.html:36
	WriteRenderEntryOptions(qb422016, i, e, js)
//line views/vhar/Entry.html:36
	qs422016 := string(qb422016.B)
//line views/vhar/Entry.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Entry.html:36
	return qs422016
//line views/vhar/Entry.html:36
}

//line views/vhar/Entry.html:38
func StreamRenderEntry(qw422016 *qt422016.Writer, i int, e *har.Entry, ps *cutil.PageState) {
//line views/vhar/Entry.html:38
	qw422016.N().S(`
  <table>
    <tbody>
      <tr>
        <th class="shrink">Duration</th>
        <td>`)
//line views/vhar/Entry.html:43
	qw422016.E().S(util.MicrosToMillis(e.Duration()))
//line views/vhar/Entry.html:43
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th class="shrink">Request Size</th>
        <td>`)
//line views/vhar/Entry.html:47
	qw422016.E().S(util.ByteSizeSI(int64(e.Request.Size())))
//line views/vhar/Entry.html:47
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th class="shrink">Request</th>
        <td>`)
//line views/vhar/Entry.html:51
	streamrenderRequest(qw422016, i, e.Request, ps)
//line views/vhar/Entry.html:51
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th class="shrink">Response Size</th>
        <td>`)
//line views/vhar/Entry.html:55
	qw422016.E().S(util.ByteSizeSI(int64(e.Response.Size())))
//line views/vhar/Entry.html:55
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th class="shrink">Response</th>
        <td>`)
//line views/vhar/Entry.html:59
	streamrenderResponse(qw422016, i, e.Response, ps)
//line views/vhar/Entry.html:59
	qw422016.N().S(`</td>
      </tr>
    </tbody>
  </table>
`)
//line views/vhar/Entry.html:63
}

//line views/vhar/Entry.html:63
func WriteRenderEntry(qq422016 qtio422016.Writer, i int, e *har.Entry, ps *cutil.PageState) {
//line views/vhar/Entry.html:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Entry.html:63
	StreamRenderEntry(qw422016, i, e, ps)
//line views/vhar/Entry.html:63
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Entry.html:63
}

//line views/vhar/Entry.html:63
func RenderEntry(i int, e *har.Entry, ps *cutil.PageState) string {
//line views/vhar/Entry.html:63
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Entry.html:63
	WriteRenderEntry(qb422016, i, e, ps)
//line views/vhar/Entry.html:63
	qs422016 := string(qb422016.B)
//line views/vhar/Entry.html:63
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Entry.html:63
	return qs422016
//line views/vhar/Entry.html:63
}
