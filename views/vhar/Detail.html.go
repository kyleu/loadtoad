// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vhar/Detail.html:2
package vhar

//line views/vhar/Detail.html:2
import (
	"fmt"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vhar/Detail.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vhar/Detail.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vhar/Detail.html:13
type Detail struct {
	layout.Basic
	Har *har.Log
}

//line views/vhar/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vhar/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="`)
//line views/vhar/Detail.html:21
	qw422016.E().S(p.Har.WebPath())
//line views/vhar/Detail.html:21
	qw422016.N().S(`/delete" class="link-confirm" data-message="Are you sure?"><button>Delete</button></a>
    </div>
    <h3>`)
//line views/vhar/Detail.html:23
	components.StreamSVGRefIcon(qw422016, `book`, ps)
//line views/vhar/Detail.html:23
	qw422016.N().S(` `)
//line views/vhar/Detail.html:23
	qw422016.E().S(p.Har.Key)
//line views/vhar/Detail.html:23
	qw422016.N().S(`</h3>
    <div class="mt">
      `)
//line views/vhar/Detail.html:25
	qw422016.E().S(util.MicrosToMillis(p.Har.Entries.TotalDuration()))
//line views/vhar/Detail.html:25
	qw422016.N().S(` elapsed, downloaded `)
//line views/vhar/Detail.html:25
	qw422016.E().S(util.ByteSizeSI(int64(p.Har.Entries.TotalResponseBodySize())))
//line views/vhar/Detail.html:25
	qw422016.N().S(`
    </div>
    <!-- $PF_SECTION_START(actions)$ -->
    <div class="mt">
      <a href="`)
//line views/vhar/Detail.html:29
	qw422016.E().S(p.Har.WebPath())
//line views/vhar/Detail.html:29
	qw422016.N().S(`/run"><button>Run</button></a>
      <a href="`)
//line views/vhar/Detail.html:30
	qw422016.E().S(p.Har.WebPath())
//line views/vhar/Detail.html:30
	qw422016.N().S(`/bench"><button>Benchmark</button></a>
      <a href="`)
//line views/vhar/Detail.html:31
	qw422016.E().S(p.Har.WebPath())
//line views/vhar/Detail.html:31
	qw422016.N().S(`/trim"><button>Trim</button></a>
    </div>
    <!-- $PF_SECTION_END(actions)$ -->
  </div>

  <div class="card">
    <h3>`)
//line views/vhar/Detail.html:37
	qw422016.E().S(util.StringPlural(len(p.Har.Entries), "Entry"))
//line views/vhar/Detail.html:37
	qw422016.N().S(`</h3>
    <div class="mts">
      <ul class="accordion">
`)
//line views/vhar/Detail.html:40
	for i, e := range p.Har.Entries {
//line views/vhar/Detail.html:42
		key := fmt.Sprintf("%d", i)
		e = e.Cleaned()

//line views/vhar/Detail.html:44
		qw422016.N().S(`        <li>
          <input id="accordion-entry-`)
//line views/vhar/Detail.html:46
		qw422016.E().S(key)
//line views/vhar/Detail.html:46
		qw422016.N().S(`" type="checkbox" hidden />
          <label title="`)
//line views/vhar/Detail.html:47
		qw422016.E().S(e.Request.URL)
//line views/vhar/Detail.html:47
		qw422016.N().S(`" for="accordion-entry-`)
//line views/vhar/Detail.html:47
		qw422016.E().S(key)
//line views/vhar/Detail.html:47
		qw422016.N().S(`">
            `)
//line views/vhar/Detail.html:48
		StreamRenderEntryOptions(qw422016, key, e, false)
//line views/vhar/Detail.html:48
		qw422016.N().S(`
            `)
//line views/vhar/Detail.html:49
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vhar/Detail.html:49
		qw422016.N().S(` `)
//line views/vhar/Detail.html:49
		qw422016.E().S(e.String())
//line views/vhar/Detail.html:49
		qw422016.N().S(`
            <div class="clear"></div>
          </label>
          <div class="bd-animated"><div><div>
          `)
//line views/vhar/Detail.html:53
		StreamRenderEntry(qw422016, key, e, ps)
//line views/vhar/Detail.html:53
		qw422016.N().S(`
          </div></div></div>
          `)
//line views/vhar/Detail.html:55
		StreamRenderEntryModals(qw422016, key, e, false)
//line views/vhar/Detail.html:55
		qw422016.N().S(`
        </li>
`)
//line views/vhar/Detail.html:57
	}
//line views/vhar/Detail.html:57
	qw422016.N().S(`      </ul>
    </div>
  </div>
  <script>
    function clip(k) {
      if (!navigator.clipboard) {
        return;
      }
      const el = document.getElementById("modal-entry-" + k + "-curl-data");
      navigator.clipboard.writeText(el.innerText);
    }
  </script>
`)
//line views/vhar/Detail.html:70
}

//line views/vhar/Detail.html:70
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vhar/Detail.html:70
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Detail.html:70
	p.StreamBody(qw422016, as, ps)
//line views/vhar/Detail.html:70
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Detail.html:70
}

//line views/vhar/Detail.html:70
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vhar/Detail.html:70
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Detail.html:70
	p.WriteBody(qb422016, as, ps)
//line views/vhar/Detail.html:70
	qs422016 := string(qb422016.B)
//line views/vhar/Detail.html:70
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Detail.html:70
	return qs422016
//line views/vhar/Detail.html:70
}
