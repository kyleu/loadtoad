// Code generated by qtc from "JSON.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/JSON.html:2
package components

//line views/components/JSON.html:2
import (
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
)

//line views/components/JSON.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/JSON.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/JSON.html:7
func StreamJSONModal(qw422016 *qt422016.Writer, key string, title string, item any, indent int) {
//line views/components/JSON.html:7
	qw422016.N().S(`<div id="modal-`)
//line views/components/JSON.html:8
	qw422016.E().S(key)
//line views/components/JSON.html:8
	qw422016.N().S(`" class="modal" style="display: none;"><a class="backdrop" href="#"></a><div class="modal-content"><div class="modal-header"><a href="#" class="modal-close">×</a><h2>`)
//line views/components/JSON.html:13
	qw422016.E().S(title)
//line views/components/JSON.html:13
	qw422016.N().S(`</h2></div><div class="modal-body"><div id="modal-`)
//line views/components/JSON.html:16
	qw422016.E().S(key)
//line views/components/JSON.html:16
	qw422016.N().S(`-data" hidden="hidden" style="display:none;">`)
//line views/components/JSON.html:16
	qw422016.E().S(util.ToJSON(item))
//line views/components/JSON.html:16
	qw422016.N().S(`</div><button onclick="clip('`)
//line views/components/JSON.html:17
	qw422016.E().S(key)
//line views/components/JSON.html:17
	qw422016.N().S(`');">Copy to clipboard</button>`)
//line views/components/JSON.html:18
	StreamJSON(qw422016, item)
//line views/components/JSON.html:18
	qw422016.N().S(`</div></div></div><script>function clip(k) {if (!navigator.clipboard) {return;}const el = document.getElementById("modal-" + k + "-data");navigator.clipboard.writeText(el.innerText);}</script>`)
//line views/components/JSON.html:31
}

//line views/components/JSON.html:31
func WriteJSONModal(qq422016 qtio422016.Writer, key string, title string, item any, indent int) {
//line views/components/JSON.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/JSON.html:31
	StreamJSONModal(qw422016, key, title, item, indent)
//line views/components/JSON.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/components/JSON.html:31
}

//line views/components/JSON.html:31
func JSONModal(key string, title string, item any, indent int) string {
//line views/components/JSON.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/JSON.html:31
	WriteJSONModal(qb422016, key, title, item, indent)
//line views/components/JSON.html:31
	qs422016 := string(qb422016.B)
//line views/components/JSON.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/JSON.html:31
	return qs422016
//line views/components/JSON.html:31
}

//line views/components/JSON.html:33
func StreamJSON(qw422016 *qt422016.Writer, v any) {
//line views/components/JSON.html:35
	b, ok := v.([]byte)
	if ok {
		_ = util.FromJSON(b, &v)
	}

//line views/components/JSON.html:40
	json := util.ToJSONBytes(v, true)

//line views/components/JSON.html:41
	if len(json) > (1024 * 256) {
//line views/components/JSON.html:41
		qw422016.N().S(`<div><em>This JSON is too large (<strong>`)
//line views/components/JSON.html:42
		qw422016.E().S(util.ByteSizeSI(int64(len(json))))
//line views/components/JSON.html:42
		qw422016.N().S(`</strong>), highlighting is disabled</em></div><div class="overflow full-width"><pre>`)
//line views/components/JSON.html:44
		qw422016.E().S(string(json))
//line views/components/JSON.html:44
		qw422016.N().S(`</pre></div>`)
//line views/components/JSON.html:46
	} else {
//line views/components/JSON.html:47
		out, err := cutil.FormatLang(string(json), "json")

//line views/components/JSON.html:48
		if err == nil {
//line views/components/JSON.html:48
			qw422016.N().S(`<div class="overflow full-width">`)
//line views/components/JSON.html:49
			qw422016.N().S(out)
//line views/components/JSON.html:49
			qw422016.N().S(`</div>`)
//line views/components/JSON.html:50
		} else {
//line views/components/JSON.html:50
			qw422016.N().S(`<div class="error">`)
//line views/components/JSON.html:51
			qw422016.E().S(err.Error())
//line views/components/JSON.html:51
			qw422016.N().S(`</div>`)
//line views/components/JSON.html:52
		}
//line views/components/JSON.html:53
	}
//line views/components/JSON.html:54
}

//line views/components/JSON.html:54
func WriteJSON(qq422016 qtio422016.Writer, v any) {
//line views/components/JSON.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/JSON.html:54
	StreamJSON(qw422016, v)
//line views/components/JSON.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/components/JSON.html:54
}

//line views/components/JSON.html:54
func JSON(v any) string {
//line views/components/JSON.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/JSON.html:54
	WriteJSON(qb422016, v)
//line views/components/JSON.html:54
	qs422016 := string(qb422016.B)
//line views/components/JSON.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/JSON.html:54
	return qs422016
//line views/components/JSON.html:54
}
