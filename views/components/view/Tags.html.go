// Code generated by qtc from "Tags.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/view/Tags.html:1
package view

//line views/components/view/Tags.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Tags.html:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Tags.html:1
func StreamTags(qw422016 *qt422016.Writer, values []string) {
//line views/components/view/Tags.html:1
	qw422016.N().S(`<div class="tag-editor"><div class="tags">`)
//line views/components/view/Tags.html:4
	for _, x := range values {
//line views/components/view/Tags.html:4
		qw422016.N().S(`<span class="item">`)
//line views/components/view/Tags.html:5
		qw422016.E().S(x)
//line views/components/view/Tags.html:5
		qw422016.N().S(`</span>`)
//line views/components/view/Tags.html:6
	}
//line views/components/view/Tags.html:6
	qw422016.N().S(`</div></div>`)
//line views/components/view/Tags.html:9
}

//line views/components/view/Tags.html:9
func WriteTags(qq422016 qtio422016.Writer, values []string) {
//line views/components/view/Tags.html:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Tags.html:9
	StreamTags(qw422016, values)
//line views/components/view/Tags.html:9
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Tags.html:9
}

//line views/components/view/Tags.html:9
func Tags(values []string) string {
//line views/components/view/Tags.html:9
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Tags.html:9
	WriteTags(qb422016, values)
//line views/components/view/Tags.html:9
	qs422016 := string(qb422016.B)
//line views/components/view/Tags.html:9
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Tags.html:9
	return qs422016
//line views/components/view/Tags.html:9
}
