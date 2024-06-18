// Code generated by qtc from "Package.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/view/Package.html:1
package view

//line views/components/view/Package.html:1
import "github.com/kyleu/loadtoad/app/util"

//line views/components/view/Package.html:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Package.html:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Package.html:3
func StreamPackage(qw422016 *qt422016.Writer, v util.Pkg) {
//line views/components/view/Package.html:4
	for idx, x := range v {
//line views/components/view/Package.html:5
		qw422016.E().S(x)
//line views/components/view/Package.html:5
		if len(v) < idx {
//line views/components/view/Package.html:5
			qw422016.N().S(`/`)
//line views/components/view/Package.html:5
		}
//line views/components/view/Package.html:6
	}
//line views/components/view/Package.html:7
}

//line views/components/view/Package.html:7
func WritePackage(qq422016 qtio422016.Writer, v util.Pkg) {
//line views/components/view/Package.html:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Package.html:7
	StreamPackage(qw422016, v)
//line views/components/view/Package.html:7
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Package.html:7
}

//line views/components/view/Package.html:7
func Package(v util.Pkg) string {
//line views/components/view/Package.html:7
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Package.html:7
	WritePackage(qb422016, v)
//line views/components/view/Package.html:7
	qs422016 := string(qb422016.B)
//line views/components/view/Package.html:7
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Package.html:7
	return qs422016
//line views/components/view/Package.html:7
}
