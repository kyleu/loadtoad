// Code generated by qtc from "String.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/edit/String.html:2
package edit

//line views/components/edit/String.html:2
import (
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/views/components"
)

//line views/components/edit/String.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/String.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/String.html:7
func StreamString(qw422016 *qt422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/edit/String.html:8
	if id == "" {
//line views/components/edit/String.html:8
		qw422016.N().S(`<input name="`)
//line views/components/edit/String.html:9
		qw422016.E().S(key)
//line views/components/edit/String.html:9
		qw422016.N().S(`" value="`)
//line views/components/edit/String.html:9
		qw422016.E().S(value)
//line views/components/edit/String.html:9
		qw422016.N().S(`"`)
//line views/components/edit/String.html:9
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:9
		qw422016.N().S(`/>`)
//line views/components/edit/String.html:10
	} else {
//line views/components/edit/String.html:10
		qw422016.N().S(`<input id="`)
//line views/components/edit/String.html:11
		qw422016.E().S(id)
//line views/components/edit/String.html:11
		qw422016.N().S(`" name="`)
//line views/components/edit/String.html:11
		qw422016.E().S(key)
//line views/components/edit/String.html:11
		qw422016.N().S(`" value="`)
//line views/components/edit/String.html:11
		qw422016.E().S(value)
//line views/components/edit/String.html:11
		qw422016.N().S(`"`)
//line views/components/edit/String.html:11
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:11
		qw422016.N().S(`/>`)
//line views/components/edit/String.html:12
	}
//line views/components/edit/String.html:13
}

//line views/components/edit/String.html:13
func WriteString(qq422016 qtio422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/edit/String.html:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:13
	StreamString(qw422016, key, id, value, placeholder...)
//line views/components/edit/String.html:13
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:13
}

//line views/components/edit/String.html:13
func String(key string, id string, value string, placeholder ...string) string {
//line views/components/edit/String.html:13
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:13
	WriteString(qb422016, key, id, value, placeholder...)
//line views/components/edit/String.html:13
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:13
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:13
	return qs422016
//line views/components/edit/String.html:13
}

//line views/components/edit/String.html:15
func StreamStringVertical(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:16
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:16
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/String.html:18
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:18
	qw422016.N().S(`<label for="`)
//line views/components/edit/String.html:19
	qw422016.E().S(id)
//line views/components/edit/String.html:19
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/String.html:19
	qw422016.E().S(title)
//line views/components/edit/String.html:19
	qw422016.N().S(`</em></label>`)
//line views/components/edit/String.html:20
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:20
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/String.html:21
	StreamString(qw422016, key, id, value, help...)
//line views/components/edit/String.html:21
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:22
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:22
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:24
}

//line views/components/edit/String.html:24
func WriteStringVertical(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:24
	StreamStringVertical(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:24
}

//line views/components/edit/String.html:24
func StringVertical(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/edit/String.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:24
	WriteStringVertical(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:24
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:24
	return qs422016
//line views/components/edit/String.html:24
}

//line views/components/edit/String.html:26
func StreamStringTable(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:27
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:27
	qw422016.N().S(`<tr>`)
//line views/components/edit/String.html:29
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:29
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/String.html:30
	qw422016.E().S(id)
//line views/components/edit/String.html:30
	qw422016.N().S(`"`)
//line views/components/edit/String.html:30
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/String.html:30
	qw422016.N().S(`>`)
//line views/components/edit/String.html:30
	qw422016.E().S(title)
//line views/components/edit/String.html:30
	qw422016.N().S(`</label></th>`)
//line views/components/edit/String.html:31
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:31
	qw422016.N().S(`<td>`)
//line views/components/edit/String.html:32
	StreamString(qw422016, key, id, value, help...)
//line views/components/edit/String.html:32
	qw422016.N().S(`</td>`)
//line views/components/edit/String.html:33
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:33
	qw422016.N().S(`</tr>`)
//line views/components/edit/String.html:35
}

//line views/components/edit/String.html:35
func WriteStringTable(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:35
	StreamStringTable(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:35
}

//line views/components/edit/String.html:35
func StringTable(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/edit/String.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:35
	WriteStringTable(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:35
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:35
	return qs422016
//line views/components/edit/String.html:35
}

//line views/components/edit/String.html:37
func StreamPassword(qw422016 *qt422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/edit/String.html:38
	if id == "" {
//line views/components/edit/String.html:38
		qw422016.N().S(`<input name="`)
//line views/components/edit/String.html:39
		qw422016.E().S(key)
//line views/components/edit/String.html:39
		qw422016.N().S(`" type="password" value="`)
//line views/components/edit/String.html:39
		qw422016.E().S(value)
//line views/components/edit/String.html:39
		qw422016.N().S(`"`)
//line views/components/edit/String.html:39
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:39
		qw422016.N().S(`/>`)
//line views/components/edit/String.html:40
	} else {
//line views/components/edit/String.html:40
		qw422016.N().S(`<input id="`)
//line views/components/edit/String.html:41
		qw422016.E().S(id)
//line views/components/edit/String.html:41
		qw422016.N().S(`" name="`)
//line views/components/edit/String.html:41
		qw422016.E().S(key)
//line views/components/edit/String.html:41
		qw422016.N().S(`" type="password" value="`)
//line views/components/edit/String.html:41
		qw422016.E().S(value)
//line views/components/edit/String.html:41
		qw422016.N().S(`"`)
//line views/components/edit/String.html:41
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:41
		qw422016.N().S(`/>`)
//line views/components/edit/String.html:42
	}
//line views/components/edit/String.html:43
}

//line views/components/edit/String.html:43
func WritePassword(qq422016 qtio422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/edit/String.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:43
	StreamPassword(qw422016, key, id, value, placeholder...)
//line views/components/edit/String.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:43
}

//line views/components/edit/String.html:43
func Password(key string, id string, value string, placeholder ...string) string {
//line views/components/edit/String.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:43
	WritePassword(qb422016, key, id, value, placeholder...)
//line views/components/edit/String.html:43
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:43
	return qs422016
//line views/components/edit/String.html:43
}

//line views/components/edit/String.html:45
func StreamPasswordVertical(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:46
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:46
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/String.html:48
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:48
	qw422016.N().S(`<label for="`)
//line views/components/edit/String.html:49
	qw422016.E().S(id)
//line views/components/edit/String.html:49
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/String.html:49
	qw422016.E().S(title)
//line views/components/edit/String.html:49
	qw422016.N().S(`</em></label>`)
//line views/components/edit/String.html:50
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:50
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/String.html:51
	StreamPassword(qw422016, key, id, value, help...)
//line views/components/edit/String.html:51
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:52
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:52
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:54
}

//line views/components/edit/String.html:54
func WritePasswordVertical(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:54
	StreamPasswordVertical(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:54
}

//line views/components/edit/String.html:54
func PasswordVertical(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/edit/String.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:54
	WritePasswordVertical(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:54
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:54
	return qs422016
//line views/components/edit/String.html:54
}

//line views/components/edit/String.html:56
func StreamPasswordTable(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:57
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:57
	qw422016.N().S(`<tr>`)
//line views/components/edit/String.html:59
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:59
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/String.html:60
	qw422016.E().S(id)
//line views/components/edit/String.html:60
	qw422016.N().S(`"`)
//line views/components/edit/String.html:60
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/String.html:60
	qw422016.N().S(`>`)
//line views/components/edit/String.html:60
	qw422016.E().S(title)
//line views/components/edit/String.html:60
	qw422016.N().S(`</label></th>`)
//line views/components/edit/String.html:61
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:61
	qw422016.N().S(`<td>`)
//line views/components/edit/String.html:62
	StreamPassword(qw422016, key, id, value, help...)
//line views/components/edit/String.html:62
	qw422016.N().S(`</td>`)
//line views/components/edit/String.html:63
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:63
	qw422016.N().S(`</tr>`)
//line views/components/edit/String.html:65
}

//line views/components/edit/String.html:65
func WritePasswordTable(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:65
	StreamPasswordTable(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:65
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:65
}

//line views/components/edit/String.html:65
func PasswordTable(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/edit/String.html:65
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:65
	WritePasswordTable(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:65
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:65
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:65
	return qs422016
//line views/components/edit/String.html:65
}

//line views/components/edit/String.html:67
func StreamTextarea(qw422016 *qt422016.Writer, key string, id string, rows int, value string, placeholder ...string) {
//line views/components/edit/String.html:68
	if id == "" {
//line views/components/edit/String.html:68
		qw422016.N().S(`<textarea rows="`)
//line views/components/edit/String.html:69
		qw422016.N().D(rows)
//line views/components/edit/String.html:69
		qw422016.N().S(`" name="`)
//line views/components/edit/String.html:69
		qw422016.E().S(key)
//line views/components/edit/String.html:69
		qw422016.N().S(`"`)
//line views/components/edit/String.html:69
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:69
		qw422016.N().S(`>`)
//line views/components/edit/String.html:69
		qw422016.E().S(value)
//line views/components/edit/String.html:69
		qw422016.N().S(`</textarea>`)
//line views/components/edit/String.html:70
	} else {
//line views/components/edit/String.html:70
		qw422016.N().S(`<textarea rows="`)
//line views/components/edit/String.html:71
		qw422016.N().D(rows)
//line views/components/edit/String.html:71
		qw422016.N().S(`" id="`)
//line views/components/edit/String.html:71
		qw422016.E().S(id)
//line views/components/edit/String.html:71
		qw422016.N().S(`" name="`)
//line views/components/edit/String.html:71
		qw422016.E().S(key)
//line views/components/edit/String.html:71
		qw422016.N().S(`"`)
//line views/components/edit/String.html:71
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:71
		qw422016.N().S(`>`)
//line views/components/edit/String.html:71
		qw422016.E().S(value)
//line views/components/edit/String.html:71
		qw422016.N().S(`</textarea>`)
//line views/components/edit/String.html:72
	}
//line views/components/edit/String.html:73
}

//line views/components/edit/String.html:73
func WriteTextarea(qq422016 qtio422016.Writer, key string, id string, rows int, value string, placeholder ...string) {
//line views/components/edit/String.html:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:73
	StreamTextarea(qw422016, key, id, rows, value, placeholder...)
//line views/components/edit/String.html:73
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:73
}

//line views/components/edit/String.html:73
func Textarea(key string, id string, rows int, value string, placeholder ...string) string {
//line views/components/edit/String.html:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:73
	WriteTextarea(qb422016, key, id, rows, value, placeholder...)
//line views/components/edit/String.html:73
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:73
	return qs422016
//line views/components/edit/String.html:73
}

//line views/components/edit/String.html:75
func StreamTextareaVertical(qw422016 *qt422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/edit/String.html:76
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:76
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/String.html:78
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:78
	qw422016.N().S(`<label for="`)
//line views/components/edit/String.html:79
	qw422016.E().S(id)
//line views/components/edit/String.html:79
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/String.html:79
	qw422016.E().S(title)
//line views/components/edit/String.html:79
	qw422016.N().S(`</em></label>`)
//line views/components/edit/String.html:80
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:80
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/String.html:81
	StreamTextarea(qw422016, key, id, rows, value, help...)
//line views/components/edit/String.html:81
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:82
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:82
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:84
}

//line views/components/edit/String.html:84
func WriteTextareaVertical(qq422016 qtio422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/edit/String.html:84
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:84
	StreamTextareaVertical(qw422016, key, id, title, rows, value, indent, help...)
//line views/components/edit/String.html:84
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:84
}

//line views/components/edit/String.html:84
func TextareaVertical(key string, id string, title string, rows int, value string, indent int, help ...string) string {
//line views/components/edit/String.html:84
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:84
	WriteTextareaVertical(qb422016, key, id, title, rows, value, indent, help...)
//line views/components/edit/String.html:84
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:84
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:84
	return qs422016
//line views/components/edit/String.html:84
}

//line views/components/edit/String.html:86
func StreamTextareaTable(qw422016 *qt422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/edit/String.html:87
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:87
	qw422016.N().S(`<tr>`)
//line views/components/edit/String.html:89
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:89
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/String.html:90
	qw422016.E().S(id)
//line views/components/edit/String.html:90
	qw422016.N().S(`"`)
//line views/components/edit/String.html:90
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/String.html:90
	qw422016.N().S(`>`)
//line views/components/edit/String.html:90
	qw422016.E().S(title)
//line views/components/edit/String.html:90
	qw422016.N().S(`</label></th>`)
//line views/components/edit/String.html:91
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:91
	qw422016.N().S(`<td>`)
//line views/components/edit/String.html:92
	StreamTextarea(qw422016, key, id, rows, value, help...)
//line views/components/edit/String.html:92
	qw422016.N().S(`</td>`)
//line views/components/edit/String.html:93
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:93
	qw422016.N().S(`</tr>`)
//line views/components/edit/String.html:95
}

//line views/components/edit/String.html:95
func WriteTextareaTable(qq422016 qtio422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/edit/String.html:95
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:95
	StreamTextareaTable(qw422016, key, id, title, rows, value, indent, help...)
//line views/components/edit/String.html:95
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:95
}

//line views/components/edit/String.html:95
func TextareaTable(key string, id string, title string, rows int, value string, indent int, help ...string) string {
//line views/components/edit/String.html:95
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:95
	WriteTextareaTable(qb422016, key, id, title, rows, value, indent, help...)
//line views/components/edit/String.html:95
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:95
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:95
	return qs422016
//line views/components/edit/String.html:95
}
