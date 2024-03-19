// Code generated by qtc from "Int.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/view/Int.html:2
package view

//line views/components/view/Int.html:2
import "github.com/kyleu/loadtoad/app/util"

//line views/components/view/Int.html:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Int.html:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Int.html:4
func StreamInt(qw422016 *qt422016.Writer, i any) {
//line views/components/view/Int.html:5
	qw422016.E().V(i)
//line views/components/view/Int.html:6
}

//line views/components/view/Int.html:6
func WriteInt(qq422016 qtio422016.Writer, i any) {
//line views/components/view/Int.html:6
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Int.html:6
	StreamInt(qw422016, i)
//line views/components/view/Int.html:6
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Int.html:6
}

//line views/components/view/Int.html:6
func Int(i any) string {
//line views/components/view/Int.html:6
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Int.html:6
	WriteInt(qb422016, i)
//line views/components/view/Int.html:6
	qs422016 := string(qb422016.B)
//line views/components/view/Int.html:6
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Int.html:6
	return qs422016
//line views/components/view/Int.html:6
}

//line views/components/view/Int.html:8
func StreamIntArray(qw422016 *qt422016.Writer, value []any) {
//line views/components/view/Int.html:9
	StreamStringArray(qw422016, util.ArrayToStringArray(value))
//line views/components/view/Int.html:10
}

//line views/components/view/Int.html:10
func WriteIntArray(qq422016 qtio422016.Writer, value []any) {
//line views/components/view/Int.html:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Int.html:10
	StreamIntArray(qw422016, value)
//line views/components/view/Int.html:10
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Int.html:10
}

//line views/components/view/Int.html:10
func IntArray(value []any) string {
//line views/components/view/Int.html:10
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Int.html:10
	WriteIntArray(qb422016, value)
//line views/components/view/Int.html:10
	qs422016 := string(qb422016.B)
//line views/components/view/Int.html:10
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Int.html:10
	return qs422016
//line views/components/view/Int.html:10
}
