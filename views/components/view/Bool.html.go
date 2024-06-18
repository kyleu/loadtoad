// Code generated by qtc from "Bool.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/view/Bool.html:1
package view

//line views/components/view/Bool.html:1
import (
	"fmt"

	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
)

//line views/components/view/Bool.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Bool.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Bool.html:9
func StreamBool(qw422016 *qt422016.Writer, b bool) {
//line views/components/view/Bool.html:10
	if b {
//line views/components/view/Bool.html:10
		qw422016.N().S(`true`)
//line views/components/view/Bool.html:10
	} else {
//line views/components/view/Bool.html:10
		qw422016.N().S(`false`)
//line views/components/view/Bool.html:10
	}
//line views/components/view/Bool.html:11
}

//line views/components/view/Bool.html:11
func WriteBool(qq422016 qtio422016.Writer, b bool) {
//line views/components/view/Bool.html:11
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Bool.html:11
	StreamBool(qw422016, b)
//line views/components/view/Bool.html:11
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Bool.html:11
}

//line views/components/view/Bool.html:11
func Bool(b bool) string {
//line views/components/view/Bool.html:11
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Bool.html:11
	WriteBool(qb422016, b)
//line views/components/view/Bool.html:11
	qs422016 := string(qb422016.B)
//line views/components/view/Bool.html:11
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Bool.html:11
	return qs422016
//line views/components/view/Bool.html:11
}

//line views/components/view/Bool.html:13
func StreamBoolIcon(qw422016 *qt422016.Writer, b bool, size int, cls string, ps *cutil.PageState, titles ...string) {
//line views/components/view/Bool.html:15
	if cls == "" {
		cls = "inline"
	}
	extra := ""
	switch len(titles) {
	case 0:
		if b {
			extra = "true"
		} else {
			extra = "false"
		}
	case 1:
		l, r := util.StringSplitLast(titles[0], '|', true)
		extra = util.Choose(b || r == "", l, r)
	case 2:
		extra = util.Choose(b, titles[0], titles[1])
	}
	if extra != "" {
		extra = fmt.Sprintf(" title=%q", extra)
	}
	icon := util.Choose(b, "check", "times")

//line views/components/view/Bool.html:36
	qw422016.N().S(`<span`)
//line views/components/view/Bool.html:37
	qw422016.N().S(extra)
//line views/components/view/Bool.html:37
	qw422016.N().S(`>`)
//line views/components/view/Bool.html:37
	components.StreamSVGRef(qw422016, icon, size, size, cls, ps)
//line views/components/view/Bool.html:37
	qw422016.N().S(`</span>`)
//line views/components/view/Bool.html:38
}

//line views/components/view/Bool.html:38
func WriteBoolIcon(qq422016 qtio422016.Writer, b bool, size int, cls string, ps *cutil.PageState, titles ...string) {
//line views/components/view/Bool.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Bool.html:38
	StreamBoolIcon(qw422016, b, size, cls, ps, titles...)
//line views/components/view/Bool.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Bool.html:38
}

//line views/components/view/Bool.html:38
func BoolIcon(b bool, size int, cls string, ps *cutil.PageState, titles ...string) string {
//line views/components/view/Bool.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Bool.html:38
	WriteBoolIcon(qb422016, b, size, cls, ps, titles...)
//line views/components/view/Bool.html:38
	qs422016 := string(qb422016.B)
//line views/components/view/Bool.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Bool.html:38
	return qs422016
//line views/components/view/Bool.html:38
}
