// Code generated by qtc from "IconPicker.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/IconPicker.html:2
package components

//line views/components/IconPicker.html:2
import (
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
)

//line views/components/IconPicker.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/IconPicker.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/IconPicker.html:7
func StreamIconPicker(qw422016 *qt422016.Writer, key string, selected string, ps *cutil.PageState, indent int) {
//line views/components/IconPicker.html:8
	StreamIndent(qw422016, true, indent)
//line views/components/IconPicker.html:8
	qw422016.N().S(`<div class="choice">`)
//line views/components/IconPicker.html:10
	for _, k := range util.SVGIconKeys {
//line views/components/IconPicker.html:11
		StreamIndent(qw422016, true, indent+1)
//line views/components/IconPicker.html:11
		qw422016.N().S(`<label title="`)
//line views/components/IconPicker.html:12
		qw422016.E().S(k)
//line views/components/IconPicker.html:12
		qw422016.N().S(`">`)
//line views/components/IconPicker.html:13
		if k == selected {
//line views/components/IconPicker.html:13
			qw422016.N().S(`<input type="radio" name="`)
//line views/components/IconPicker.html:14
			qw422016.E().S(key)
//line views/components/IconPicker.html:14
			qw422016.N().S(`" value="`)
//line views/components/IconPicker.html:14
			qw422016.E().S(k)
//line views/components/IconPicker.html:14
			qw422016.N().S(`" checked="checked" />`)
//line views/components/IconPicker.html:15
		} else {
//line views/components/IconPicker.html:15
			qw422016.N().S(`<input type="radio" name="`)
//line views/components/IconPicker.html:16
			qw422016.E().S(key)
//line views/components/IconPicker.html:16
			qw422016.N().S(`" value="`)
//line views/components/IconPicker.html:16
			qw422016.E().S(k)
//line views/components/IconPicker.html:16
			qw422016.N().S(`" />`)
//line views/components/IconPicker.html:17
		}
//line views/components/IconPicker.html:18
		qw422016.N().S(` `)
//line views/components/IconPicker.html:19
		StreamSVGRef(qw422016, k, 48, 48, "", ps)
//line views/components/IconPicker.html:19
		qw422016.N().S(`</label>`)
//line views/components/IconPicker.html:21
	}
//line views/components/IconPicker.html:22
	StreamIndent(qw422016, true, indent+1)
//line views/components/IconPicker.html:22
	qw422016.N().S(`<div class="clear"></div>`)
//line views/components/IconPicker.html:24
	StreamIndent(qw422016, true, indent)
//line views/components/IconPicker.html:24
	qw422016.N().S(`</div>`)
//line views/components/IconPicker.html:26
}

//line views/components/IconPicker.html:26
func WriteIconPicker(qq422016 qtio422016.Writer, key string, selected string, ps *cutil.PageState, indent int) {
//line views/components/IconPicker.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/IconPicker.html:26
	StreamIconPicker(qw422016, key, selected, ps, indent)
//line views/components/IconPicker.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/components/IconPicker.html:26
}

//line views/components/IconPicker.html:26
func IconPicker(key string, selected string, ps *cutil.PageState, indent int) string {
//line views/components/IconPicker.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/IconPicker.html:26
	WriteIconPicker(qb422016, key, selected, ps, indent)
//line views/components/IconPicker.html:26
	qs422016 := string(qb422016.B)
//line views/components/IconPicker.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/IconPicker.html:26
	return qs422016
//line views/components/IconPicker.html:26
}
