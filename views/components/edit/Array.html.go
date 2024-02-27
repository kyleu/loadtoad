// Code generated by qtc from "Array.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/edit/Array.html:2
package edit

//line views/components/edit/Array.html:2
import (
	"github.com/samber/lo"

	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/views/components"
)

//line views/components/edit/Array.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/Array.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/Array.html:9
func StreamSelect(qw422016 *qt422016.Writer, key string, id string, value string, opts []string, titles []string, indent int) {
//line views/components/edit/Array.html:9
	qw422016.N().S(`<select name="`)
//line views/components/edit/Array.html:10
	qw422016.E().S(key)
//line views/components/edit/Array.html:10
	qw422016.N().S(`"`)
//line views/components/edit/Array.html:10
	if id != `` {
//line views/components/edit/Array.html:10
		qw422016.N().S(` `)
//line views/components/edit/Array.html:10
		qw422016.N().S(`id="`)
//line views/components/edit/Array.html:10
		qw422016.E().S(id)
//line views/components/edit/Array.html:10
		qw422016.N().S(`"`)
//line views/components/edit/Array.html:10
	}
//line views/components/edit/Array.html:10
	qw422016.N().S(`>`)
//line views/components/edit/Array.html:11
	for idx, opt := range opts {
//line views/components/edit/Array.html:13
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/edit/Array.html:18
		components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:19
		if opt == value {
//line views/components/edit/Array.html:19
			qw422016.N().S(`<option selected="selected" value="`)
//line views/components/edit/Array.html:20
			qw422016.E().S(opt)
//line views/components/edit/Array.html:20
			qw422016.N().S(`">`)
//line views/components/edit/Array.html:20
			qw422016.E().S(title)
//line views/components/edit/Array.html:20
			qw422016.N().S(`</option>`)
//line views/components/edit/Array.html:21
		} else {
//line views/components/edit/Array.html:21
			qw422016.N().S(`<option value="`)
//line views/components/edit/Array.html:22
			qw422016.E().S(opt)
//line views/components/edit/Array.html:22
			qw422016.N().S(`">`)
//line views/components/edit/Array.html:22
			qw422016.E().S(title)
//line views/components/edit/Array.html:22
			qw422016.N().S(`</option>`)
//line views/components/edit/Array.html:23
		}
//line views/components/edit/Array.html:24
	}
//line views/components/edit/Array.html:25
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:25
	qw422016.N().S(`</select>`)
//line views/components/edit/Array.html:27
}

//line views/components/edit/Array.html:27
func WriteSelect(qq422016 qtio422016.Writer, key string, id string, value string, opts []string, titles []string, indent int) {
//line views/components/edit/Array.html:27
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:27
	StreamSelect(qw422016, key, id, value, opts, titles, indent)
//line views/components/edit/Array.html:27
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:27
}

//line views/components/edit/Array.html:27
func Select(key string, id string, value string, opts []string, titles []string, indent int) string {
//line views/components/edit/Array.html:27
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:27
	WriteSelect(qb422016, key, id, value, opts, titles, indent)
//line views/components/edit/Array.html:27
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:27
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:27
	return qs422016
//line views/components/edit/Array.html:27
}

//line views/components/edit/Array.html:29
func StreamSelectVertical(qw422016 *qt422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:30
	id = cutil.CleanID(key, id)

//line views/components/edit/Array.html:30
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/Array.html:32
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:32
	qw422016.N().S(`<label for="`)
//line views/components/edit/Array.html:33
	qw422016.E().S(id)
//line views/components/edit/Array.html:33
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/Array.html:33
	qw422016.E().S(title)
//line views/components/edit/Array.html:33
	qw422016.N().S(`</em></label>`)
//line views/components/edit/Array.html:34
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:34
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/Array.html:35
	StreamSelect(qw422016, key, id, value, opts, titles, indent)
//line views/components/edit/Array.html:35
	qw422016.N().S(`</div>`)
//line views/components/edit/Array.html:36
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:36
	qw422016.N().S(`</div>`)
//line views/components/edit/Array.html:38
}

//line views/components/edit/Array.html:38
func WriteSelectVertical(qq422016 qtio422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:38
	StreamSelectVertical(qw422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:38
}

//line views/components/edit/Array.html:38
func SelectVertical(key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/edit/Array.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:38
	WriteSelectVertical(qb422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:38
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:38
	return qs422016
//line views/components/edit/Array.html:38
}

//line views/components/edit/Array.html:40
func StreamSelectTable(qw422016 *qt422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:41
	id = cutil.CleanID(key, id)

//line views/components/edit/Array.html:41
	qw422016.N().S(`<tr>`)
//line views/components/edit/Array.html:43
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:43
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/Array.html:44
	qw422016.E().S(id)
//line views/components/edit/Array.html:44
	qw422016.N().S(`"`)
//line views/components/edit/Array.html:44
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/Array.html:44
	qw422016.N().S(`>`)
//line views/components/edit/Array.html:44
	qw422016.E().S(title)
//line views/components/edit/Array.html:44
	qw422016.N().S(`</label></th>`)
//line views/components/edit/Array.html:45
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:45
	qw422016.N().S(`<td>`)
//line views/components/edit/Array.html:46
	StreamSelect(qw422016, key, id, value, opts, titles, indent)
//line views/components/edit/Array.html:46
	qw422016.N().S(`</td>`)
//line views/components/edit/Array.html:47
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:47
	qw422016.N().S(`</tr>`)
//line views/components/edit/Array.html:49
}

//line views/components/edit/Array.html:49
func WriteSelectTable(qq422016 qtio422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:49
	StreamSelectTable(qw422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:49
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:49
}

//line views/components/edit/Array.html:49
func SelectTable(key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/edit/Array.html:49
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:49
	WriteSelectTable(qb422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:49
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:49
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:49
	return qs422016
//line views/components/edit/Array.html:49
}

//line views/components/edit/Array.html:51
func StreamDatalist(qw422016 *qt422016.Writer, key string, id string, value string, opts []string, titles []string, indent int, placeholder ...string) {
//line views/components/edit/Array.html:52
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:52
	qw422016.N().S(`<input id="`)
//line views/components/edit/Array.html:53
	qw422016.E().S(id)
//line views/components/edit/Array.html:53
	qw422016.N().S(`" list="`)
//line views/components/edit/Array.html:53
	qw422016.E().S(id)
//line views/components/edit/Array.html:53
	qw422016.N().S(`-list" name="`)
//line views/components/edit/Array.html:53
	qw422016.E().S(key)
//line views/components/edit/Array.html:53
	qw422016.N().S(`" value="`)
//line views/components/edit/Array.html:53
	qw422016.E().S(value)
//line views/components/edit/Array.html:53
	qw422016.N().S(`"`)
//line views/components/edit/Array.html:53
	components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Array.html:53
	qw422016.N().S(`/>`)
//line views/components/edit/Array.html:54
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:55
	if len(opts) > 0 {
//line views/components/edit/Array.html:55
		qw422016.N().S(`<datalist id="`)
//line views/components/edit/Array.html:56
		qw422016.E().S(id)
//line views/components/edit/Array.html:56
		qw422016.N().S(`-list">`)
//line views/components/edit/Array.html:57
		for idx, opt := range opts {
//line views/components/edit/Array.html:59
			title := opt
			if idx < len(titles) {
				title = titles[idx]
			}

//line views/components/edit/Array.html:64
			components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:64
			qw422016.N().S(`<option value="`)
//line views/components/edit/Array.html:65
			qw422016.E().S(opt)
//line views/components/edit/Array.html:65
			qw422016.N().S(`">`)
//line views/components/edit/Array.html:65
			qw422016.E().S(title)
//line views/components/edit/Array.html:65
			qw422016.N().S(`</option>`)
//line views/components/edit/Array.html:66
		}
//line views/components/edit/Array.html:67
		components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:67
		qw422016.N().S(`</datalist>`)
//line views/components/edit/Array.html:69
	}
//line views/components/edit/Array.html:70
}

//line views/components/edit/Array.html:70
func WriteDatalist(qq422016 qtio422016.Writer, key string, id string, value string, opts []string, titles []string, indent int, placeholder ...string) {
//line views/components/edit/Array.html:70
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:70
	StreamDatalist(qw422016, key, id, value, opts, titles, indent, placeholder...)
//line views/components/edit/Array.html:70
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:70
}

//line views/components/edit/Array.html:70
func Datalist(key string, id string, value string, opts []string, titles []string, indent int, placeholder ...string) string {
//line views/components/edit/Array.html:70
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:70
	WriteDatalist(qb422016, key, id, value, opts, titles, indent, placeholder...)
//line views/components/edit/Array.html:70
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:70
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:70
	return qs422016
//line views/components/edit/Array.html:70
}

//line views/components/edit/Array.html:72
func StreamDatalistVertical(qw422016 *qt422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:73
	id = cutil.CleanID(key, id)

//line views/components/edit/Array.html:73
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/Array.html:75
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:75
	qw422016.N().S(`<label for="`)
//line views/components/edit/Array.html:76
	qw422016.E().S(id)
//line views/components/edit/Array.html:76
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/Array.html:76
	qw422016.E().S(title)
//line views/components/edit/Array.html:76
	qw422016.N().S(`</em></label>`)
//line views/components/edit/Array.html:77
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:77
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/Array.html:78
	StreamDatalist(qw422016, key, id, value, opts, titles, indent)
//line views/components/edit/Array.html:78
	qw422016.N().S(`</div>`)
//line views/components/edit/Array.html:79
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:79
	qw422016.N().S(`</div>`)
//line views/components/edit/Array.html:81
}

//line views/components/edit/Array.html:81
func WriteDatalistVertical(qq422016 qtio422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:81
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:81
	StreamDatalistVertical(qw422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:81
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:81
}

//line views/components/edit/Array.html:81
func DatalistVertical(key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/edit/Array.html:81
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:81
	WriteDatalistVertical(qb422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:81
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:81
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:81
	return qs422016
//line views/components/edit/Array.html:81
}

//line views/components/edit/Array.html:83
func StreamDatalistTable(qw422016 *qt422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:84
	id = cutil.CleanID(key, id)

//line views/components/edit/Array.html:84
	qw422016.N().S(`<tr>`)
//line views/components/edit/Array.html:86
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:86
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/Array.html:87
	qw422016.E().S(id)
//line views/components/edit/Array.html:87
	qw422016.N().S(`"`)
//line views/components/edit/Array.html:87
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/Array.html:87
	qw422016.N().S(`>`)
//line views/components/edit/Array.html:87
	qw422016.E().S(title)
//line views/components/edit/Array.html:87
	qw422016.N().S(`</label></th>`)
//line views/components/edit/Array.html:88
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:88
	qw422016.N().S(`<td>`)
//line views/components/edit/Array.html:89
	StreamDatalist(qw422016, key, id, value, opts, titles, indent)
//line views/components/edit/Array.html:89
	qw422016.N().S(`</td>`)
//line views/components/edit/Array.html:90
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:90
	qw422016.N().S(`</tr>`)
//line views/components/edit/Array.html:92
}

//line views/components/edit/Array.html:92
func WriteDatalistTable(qq422016 qtio422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:92
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:92
	StreamDatalistTable(qw422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:92
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:92
}

//line views/components/edit/Array.html:92
func DatalistTable(key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/edit/Array.html:92
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:92
	WriteDatalistTable(qb422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:92
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:92
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:92
	return qs422016
//line views/components/edit/Array.html:92
}

//line views/components/edit/Array.html:94
func StreamRadio(qw422016 *qt422016.Writer, key string, value string, opts []string, titles []string, indent int) {
//line views/components/edit/Array.html:95
	for idx, opt := range opts {
//line views/components/edit/Array.html:97
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/edit/Array.html:102
		components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:103
		if opt == value {
//line views/components/edit/Array.html:103
			qw422016.N().S(`<label class="radio-label"><input type="radio" name="`)
//line views/components/edit/Array.html:104
			qw422016.E().S(key)
//line views/components/edit/Array.html:104
			qw422016.N().S(`" value="`)
//line views/components/edit/Array.html:104
			qw422016.E().S(opt)
//line views/components/edit/Array.html:104
			qw422016.N().S(`" checked="checked" />`)
//line views/components/edit/Array.html:104
			qw422016.N().S(` `)
//line views/components/edit/Array.html:104
			qw422016.E().S(title)
//line views/components/edit/Array.html:104
			qw422016.N().S(`</label>`)
//line views/components/edit/Array.html:105
		} else {
//line views/components/edit/Array.html:105
			qw422016.N().S(`<label class="radio-label"><input type="radio" name="`)
//line views/components/edit/Array.html:106
			qw422016.E().S(key)
//line views/components/edit/Array.html:106
			qw422016.N().S(`" value="`)
//line views/components/edit/Array.html:106
			qw422016.E().S(opt)
//line views/components/edit/Array.html:106
			qw422016.N().S(`" />`)
//line views/components/edit/Array.html:106
			qw422016.N().S(` `)
//line views/components/edit/Array.html:106
			qw422016.E().S(title)
//line views/components/edit/Array.html:106
			qw422016.N().S(`</label>`)
//line views/components/edit/Array.html:107
		}
//line views/components/edit/Array.html:108
	}
//line views/components/edit/Array.html:109
}

//line views/components/edit/Array.html:109
func WriteRadio(qq422016 qtio422016.Writer, key string, value string, opts []string, titles []string, indent int) {
//line views/components/edit/Array.html:109
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:109
	StreamRadio(qw422016, key, value, opts, titles, indent)
//line views/components/edit/Array.html:109
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:109
}

//line views/components/edit/Array.html:109
func Radio(key string, value string, opts []string, titles []string, indent int) string {
//line views/components/edit/Array.html:109
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:109
	WriteRadio(qb422016, key, value, opts, titles, indent)
//line views/components/edit/Array.html:109
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:109
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:109
	return qs422016
//line views/components/edit/Array.html:109
}

//line views/components/edit/Array.html:111
func StreamRadioVertical(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:111
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/Array.html:113
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:113
	qw422016.N().S(`<div>`)
//line views/components/edit/Array.html:114
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/Array.html:114
	qw422016.N().S(`>`)
//line views/components/edit/Array.html:114
	qw422016.E().S(title)
//line views/components/edit/Array.html:114
	qw422016.N().S(`</div>`)
//line views/components/edit/Array.html:115
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:115
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/Array.html:117
	StreamRadio(qw422016, key, value, opts, titles, indent+2)
//line views/components/edit/Array.html:118
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:118
	qw422016.N().S(`</div>`)
//line views/components/edit/Array.html:120
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:120
	qw422016.N().S(`</div>`)
//line views/components/edit/Array.html:122
}

//line views/components/edit/Array.html:122
func WriteRadioVertical(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:122
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:122
	StreamRadioVertical(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:122
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:122
}

//line views/components/edit/Array.html:122
func RadioVertical(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/edit/Array.html:122
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:122
	WriteRadioVertical(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:122
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:122
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:122
	return qs422016
//line views/components/edit/Array.html:122
}

//line views/components/edit/Array.html:124
func StreamRadioTable(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:124
	qw422016.N().S(`<tr>`)
//line views/components/edit/Array.html:126
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:126
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/edit/Array.html:127
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/Array.html:127
	qw422016.E().S(title)
//line views/components/edit/Array.html:127
	qw422016.N().S(`</label></th>`)
//line views/components/edit/Array.html:128
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:128
	qw422016.N().S(`<td>`)
//line views/components/edit/Array.html:130
	StreamRadio(qw422016, key, value, opts, titles, indent+2)
//line views/components/edit/Array.html:131
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:131
	qw422016.N().S(`</td>`)
//line views/components/edit/Array.html:133
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:133
	qw422016.N().S(`</tr>`)
//line views/components/edit/Array.html:135
}

//line views/components/edit/Array.html:135
func WriteRadioTable(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:135
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:135
	StreamRadioTable(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:135
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:135
}

//line views/components/edit/Array.html:135
func RadioTable(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/edit/Array.html:135
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:135
	WriteRadioTable(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/edit/Array.html:135
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:135
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:135
	return qs422016
//line views/components/edit/Array.html:135
}

//line views/components/edit/Array.html:137
func StreamCheckbox(qw422016 *qt422016.Writer, key string, values []string, opts []string, titles []string, indent int) {
//line views/components/edit/Array.html:138
	for idx, opt := range opts {
//line views/components/edit/Array.html:140
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/edit/Array.html:145
		components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:146
		if lo.Contains(values, opt) {
//line views/components/edit/Array.html:146
			qw422016.N().S(`<label><input type="checkbox" name="`)
//line views/components/edit/Array.html:147
			qw422016.E().S(key)
//line views/components/edit/Array.html:147
			qw422016.N().S(`" value="`)
//line views/components/edit/Array.html:147
			qw422016.E().S(opt)
//line views/components/edit/Array.html:147
			qw422016.N().S(`" checked="checked" />`)
//line views/components/edit/Array.html:147
			qw422016.N().S(` `)
//line views/components/edit/Array.html:147
			qw422016.E().S(title)
//line views/components/edit/Array.html:147
			qw422016.N().S(`</label>`)
//line views/components/edit/Array.html:148
		} else {
//line views/components/edit/Array.html:148
			qw422016.N().S(`<label><input type="checkbox" name="`)
//line views/components/edit/Array.html:149
			qw422016.E().S(key)
//line views/components/edit/Array.html:149
			qw422016.N().S(`" value="`)
//line views/components/edit/Array.html:149
			qw422016.E().S(opt)
//line views/components/edit/Array.html:149
			qw422016.N().S(`" />`)
//line views/components/edit/Array.html:149
			qw422016.N().S(` `)
//line views/components/edit/Array.html:149
			qw422016.E().S(title)
//line views/components/edit/Array.html:149
			qw422016.N().S(`</label>`)
//line views/components/edit/Array.html:150
		}
//line views/components/edit/Array.html:151
	}
//line views/components/edit/Array.html:152
}

//line views/components/edit/Array.html:152
func WriteCheckbox(qq422016 qtio422016.Writer, key string, values []string, opts []string, titles []string, indent int) {
//line views/components/edit/Array.html:152
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:152
	StreamCheckbox(qw422016, key, values, opts, titles, indent)
//line views/components/edit/Array.html:152
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:152
}

//line views/components/edit/Array.html:152
func Checkbox(key string, values []string, opts []string, titles []string, indent int) string {
//line views/components/edit/Array.html:152
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:152
	WriteCheckbox(qb422016, key, values, opts, titles, indent)
//line views/components/edit/Array.html:152
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:152
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:152
	return qs422016
//line views/components/edit/Array.html:152
}

//line views/components/edit/Array.html:154
func StreamCheckboxVertical(qw422016 *qt422016.Writer, key string, title string, values []string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:154
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/Array.html:156
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:156
	qw422016.N().S(`<div>`)
//line views/components/edit/Array.html:157
	qw422016.E().S(title)
//line views/components/edit/Array.html:157
	qw422016.N().S(`</div>`)
//line views/components/edit/Array.html:158
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:158
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/Array.html:160
	StreamCheckbox(qw422016, key, values, opts, titles, indent+2)
//line views/components/edit/Array.html:161
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:161
	qw422016.N().S(`</div>`)
//line views/components/edit/Array.html:163
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:163
	qw422016.N().S(`</div>`)
//line views/components/edit/Array.html:165
}

//line views/components/edit/Array.html:165
func WriteCheckboxVertical(qq422016 qtio422016.Writer, key string, title string, values []string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:165
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:165
	StreamCheckboxVertical(qw422016, key, title, values, opts, titles, indent, help...)
//line views/components/edit/Array.html:165
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:165
}

//line views/components/edit/Array.html:165
func CheckboxVertical(key string, title string, values []string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/edit/Array.html:165
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:165
	WriteCheckboxVertical(qb422016, key, title, values, opts, titles, indent, help...)
//line views/components/edit/Array.html:165
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:165
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:165
	return qs422016
//line views/components/edit/Array.html:165
}

//line views/components/edit/Array.html:167
func StreamCheckboxTable(qw422016 *qt422016.Writer, key string, title string, values []string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:167
	qw422016.N().S(`<tr>`)
//line views/components/edit/Array.html:169
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:169
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/edit/Array.html:170
	qw422016.E().S(title)
//line views/components/edit/Array.html:170
	qw422016.N().S(`</label></th>`)
//line views/components/edit/Array.html:171
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:171
	qw422016.N().S(`<td class="checkboxes">`)
//line views/components/edit/Array.html:173
	StreamCheckbox(qw422016, key, values, opts, titles, indent+2)
//line views/components/edit/Array.html:174
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Array.html:174
	qw422016.N().S(`</td>`)
//line views/components/edit/Array.html:176
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Array.html:176
	qw422016.N().S(`</tr>`)
//line views/components/edit/Array.html:178
}

//line views/components/edit/Array.html:178
func WriteCheckboxTable(qq422016 qtio422016.Writer, key string, title string, values []string, opts []string, titles []string, indent int, help ...string) {
//line views/components/edit/Array.html:178
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Array.html:178
	StreamCheckboxTable(qw422016, key, title, values, opts, titles, indent, help...)
//line views/components/edit/Array.html:178
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Array.html:178
}

//line views/components/edit/Array.html:178
func CheckboxTable(key string, title string, values []string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/edit/Array.html:178
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Array.html:178
	WriteCheckboxTable(qb422016, key, title, values, opts, titles, indent, help...)
//line views/components/edit/Array.html:178
	qs422016 := string(qb422016.B)
//line views/components/edit/Array.html:178
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Array.html:178
	return qs422016
//line views/components/edit/Array.html:178
}
