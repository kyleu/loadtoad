// Code generated by qtc from "Menu.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/layout/Menu.html:1
package layout

//line views/layout/Menu.html:1
import (
	"fmt"
	"strings"

	"github.com/kyleu/loadtoad/app/controller/cmenu"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/menu"
	"github.com/kyleu/loadtoad/views/components"
)

//line views/layout/Menu.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Menu.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Menu.html:11
func StreamMenu(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/layout/Menu.html:12
	if len(ps.Menu.Visible()) > 0 {
//line views/layout/Menu.html:12
		qw422016.N().S(`<div class="menu-container">`)
//line views/layout/Menu.html:14
		components.StreamIndent(qw422016, true, 2)
//line views/layout/Menu.html:14
		qw422016.N().S(`<div class="menu">`)
//line views/layout/Menu.html:16
		components.StreamIndent(qw422016, true, 3)
//line views/layout/Menu.html:16
		qw422016.N().S(`<menu class="level-0">`)
//line views/layout/Menu.html:18
		for _, i := range ps.Menu {
//line views/layout/Menu.html:19
			StreamMenuItem(qw422016, i, []string{}, ps.Breadcrumbs, 3, ps)
//line views/layout/Menu.html:20
		}
//line views/layout/Menu.html:21
		components.StreamIndent(qw422016, true, 3)
//line views/layout/Menu.html:21
		qw422016.N().S(`</menu>`)
//line views/layout/Menu.html:23
		components.StreamIndent(qw422016, true, 2)
//line views/layout/Menu.html:23
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:25
		components.StreamIndent(qw422016, true, 1)
//line views/layout/Menu.html:25
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:27
	}
//line views/layout/Menu.html:28
}

//line views/layout/Menu.html:28
func WriteMenu(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/layout/Menu.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:28
	StreamMenu(qw422016, ps)
//line views/layout/Menu.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:28
}

//line views/layout/Menu.html:28
func Menu(ps *cutil.PageState) string {
//line views/layout/Menu.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:28
	WriteMenu(qb422016, ps)
//line views/layout/Menu.html:28
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:28
	return qs422016
//line views/layout/Menu.html:28
}

//line views/layout/Menu.html:30
func StreamMenuItem(qw422016 *qt422016.Writer, i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:32
	path = append(path, i.Key)
	active, final := breadcrumbs.Active(i, path)

//line views/layout/Menu.html:35
	if i.Key == "" {
//line views/layout/Menu.html:36
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:36
		qw422016.N().S(`<li class="separator"></li>`)
//line views/layout/Menu.html:38
	} else if len(i.Children.Visible()) > 0 {
//line views/layout/Menu.html:39
		itemID := strings.Join(path, "--")

//line views/layout/Menu.html:40
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:41
		if active {
//line views/layout/Menu.html:41
			qw422016.N().S(`<li class="active" data-menu-key="`)
//line views/layout/Menu.html:41
			qw422016.E().S(i.Key)
//line views/layout/Menu.html:41
			qw422016.N().S(`">`)
//line views/layout/Menu.html:41
		} else {
//line views/layout/Menu.html:41
			qw422016.N().S(`<li data-menu-key="`)
//line views/layout/Menu.html:41
			qw422016.E().S(i.Key)
//line views/layout/Menu.html:41
			qw422016.N().S(`">`)
//line views/layout/Menu.html:41
		}
//line views/layout/Menu.html:42
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:42
		qw422016.N().S(`<input id="`)
//line views/layout/Menu.html:43
		qw422016.E().S(itemID)
//line views/layout/Menu.html:43
		qw422016.N().S(`-input" type="checkbox"`)
//line views/layout/Menu.html:43
		if active {
//line views/layout/Menu.html:43
			qw422016.N().S(` `)
//line views/layout/Menu.html:43
			qw422016.N().S(`checked="checked"`)
//line views/layout/Menu.html:43
		}
//line views/layout/Menu.html:43
		qw422016.N().S(` `)
//line views/layout/Menu.html:43
		qw422016.N().S(`hidden="hidden" />`)
//line views/layout/Menu.html:44
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:45
		if final {
//line views/layout/Menu.html:45
			qw422016.N().S(`<label class="final" for="`)
//line views/layout/Menu.html:45
			qw422016.E().S(itemID)
//line views/layout/Menu.html:45
			qw422016.N().S(`-input" title="`)
//line views/layout/Menu.html:45
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:45
			qw422016.N().S(`">`)
//line views/layout/Menu.html:45
		} else {
//line views/layout/Menu.html:45
			qw422016.N().S(`<label for="`)
//line views/layout/Menu.html:45
			qw422016.E().S(itemID)
//line views/layout/Menu.html:45
			qw422016.N().S(`-input" title="`)
//line views/layout/Menu.html:45
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:45
			qw422016.N().S(`">`)
//line views/layout/Menu.html:45
		}
//line views/layout/Menu.html:46
		components.StreamExpandCollapse(qw422016, indent+3, ps)
//line views/layout/Menu.html:47
		streammenuBadge(qw422016, i, indent+3, ps)
//line views/layout/Menu.html:48
		components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:49
		if i.Icon != "" {
//line views/layout/Menu.html:50
			if i.Warning != "" {
//line views/layout/Menu.html:50
				qw422016.N().S(`<a class="link-confirm" data-message="`)
//line views/layout/Menu.html:51
				qw422016.E().S(i.Warning)
//line views/layout/Menu.html:51
				qw422016.N().S(`" href="`)
//line views/layout/Menu.html:51
				qw422016.E().S(i.Route)
//line views/layout/Menu.html:51
				qw422016.N().S(`">`)
//line views/layout/Menu.html:51
				components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:51
				qw422016.N().S(`</a>`)
//line views/layout/Menu.html:51
				qw422016.N().S(` `)
//line views/layout/Menu.html:52
			} else {
//line views/layout/Menu.html:52
				qw422016.N().S(`<a href="`)
//line views/layout/Menu.html:53
				qw422016.E().S(i.Route)
//line views/layout/Menu.html:53
				qw422016.N().S(`">`)
//line views/layout/Menu.html:53
				components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:53
				qw422016.N().S(`</a>`)
//line views/layout/Menu.html:53
				qw422016.N().S(` `)
//line views/layout/Menu.html:54
			}
//line views/layout/Menu.html:55
		}
//line views/layout/Menu.html:56
		if i.Route != "" {
//line views/layout/Menu.html:57
			if i.Warning != "" {
//line views/layout/Menu.html:57
				qw422016.N().S(`<a class="link-confirm" data-message="`)
//line views/layout/Menu.html:58
				qw422016.E().S(i.Warning)
//line views/layout/Menu.html:58
				qw422016.N().S(`" href="`)
//line views/layout/Menu.html:58
				qw422016.E().S(i.Route)
//line views/layout/Menu.html:58
				qw422016.N().S(`">`)
//line views/layout/Menu.html:58
				qw422016.E().S(i.Title)
//line views/layout/Menu.html:58
				qw422016.N().S(`</a>`)
//line views/layout/Menu.html:59
			} else {
//line views/layout/Menu.html:59
				qw422016.N().S(`<a href="`)
//line views/layout/Menu.html:60
				qw422016.E().S(i.Route)
//line views/layout/Menu.html:60
				qw422016.N().S(`">`)
//line views/layout/Menu.html:60
				qw422016.E().S(i.Title)
//line views/layout/Menu.html:60
				qw422016.N().S(`</a>`)
//line views/layout/Menu.html:61
			}
//line views/layout/Menu.html:62
		} else {
//line views/layout/Menu.html:63
			qw422016.E().S(i.Title)
//line views/layout/Menu.html:64
		}
//line views/layout/Menu.html:65
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:66
		if final {
//line views/layout/Menu.html:66
			qw422016.N().S(`</label>`)
//line views/layout/Menu.html:66
		} else {
//line views/layout/Menu.html:66
			qw422016.N().S(`</label>`)
//line views/layout/Menu.html:66
		}
//line views/layout/Menu.html:67
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:67
		qw422016.N().S(`<div class="menu-content level-`)
//line views/layout/Menu.html:68
		qw422016.N().D(len(path))
//line views/layout/Menu.html:68
		qw422016.N().S(`">`)
//line views/layout/Menu.html:69
		components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:69
		qw422016.N().S(`<menu>`)
//line views/layout/Menu.html:71
		for _, i := range i.Children.Visible() {
//line views/layout/Menu.html:72
			StreamMenuItem(qw422016, i, path, breadcrumbs, indent+3, ps)
//line views/layout/Menu.html:73
		}
//line views/layout/Menu.html:74
		components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:74
		qw422016.N().S(`</menu>`)
//line views/layout/Menu.html:76
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:76
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:78
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:78
		qw422016.N().S(`</li>`)
//line views/layout/Menu.html:80
	} else {
//line views/layout/Menu.html:82
		finalClass := "item"
		if active {
			finalClass += " active"
		}
		if final {
			finalClass += " final"
		}
		if i.Warning != "" {
			finalClass += " link-confirm"
		}

//line views/layout/Menu.html:93
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:93
		qw422016.N().S(`<li data-menu-key="`)
//line views/layout/Menu.html:94
		qw422016.E().S(i.Key)
//line views/layout/Menu.html:94
		qw422016.N().S(`">`)
//line views/layout/Menu.html:95
		if i.Warning != "" {
//line views/layout/Menu.html:95
			qw422016.N().S(`<a class="`)
//line views/layout/Menu.html:96
			qw422016.E().S(finalClass)
//line views/layout/Menu.html:96
			qw422016.N().S(`" data-message="`)
//line views/layout/Menu.html:96
			qw422016.E().S(i.Warning)
//line views/layout/Menu.html:96
			qw422016.N().S(`" href="`)
//line views/layout/Menu.html:96
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:96
			qw422016.N().S(`" title="`)
//line views/layout/Menu.html:96
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:96
			qw422016.N().S(`">`)
//line views/layout/Menu.html:97
		} else {
//line views/layout/Menu.html:97
			qw422016.N().S(`<a class="`)
//line views/layout/Menu.html:98
			qw422016.E().S(finalClass)
//line views/layout/Menu.html:98
			qw422016.N().S(`" href="`)
//line views/layout/Menu.html:98
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:98
			qw422016.N().S(`" title="`)
//line views/layout/Menu.html:98
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:98
			qw422016.N().S(`">`)
//line views/layout/Menu.html:99
		}
//line views/layout/Menu.html:100
		streammenuBadge(qw422016, i, indent+3, ps)
//line views/layout/Menu.html:101
		if i.Icon != "" {
//line views/layout/Menu.html:102
			components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:102
			qw422016.N().S(` `)
//line views/layout/Menu.html:103
		}
//line views/layout/Menu.html:103
		qw422016.N().S(`<span>`)
//line views/layout/Menu.html:104
		qw422016.E().S(i.Title)
//line views/layout/Menu.html:104
		qw422016.N().S(`</span>`)
//line views/layout/Menu.html:105
		if i.Warning != "" {
//line views/layout/Menu.html:105
			qw422016.N().S(`</a>`)
//line views/layout/Menu.html:107
		} else {
//line views/layout/Menu.html:107
			qw422016.N().S(`</a>`)
//line views/layout/Menu.html:109
		}
//line views/layout/Menu.html:109
		qw422016.N().S(`</li>`)
//line views/layout/Menu.html:111
	}
//line views/layout/Menu.html:112
}

//line views/layout/Menu.html:112
func WriteMenuItem(qq422016 qtio422016.Writer, i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:112
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:112
	StreamMenuItem(qw422016, i, path, breadcrumbs, indent, ps)
//line views/layout/Menu.html:112
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:112
}

//line views/layout/Menu.html:112
func MenuItem(i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) string {
//line views/layout/Menu.html:112
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:112
	WriteMenuItem(qb422016, i, path, breadcrumbs, indent, ps)
//line views/layout/Menu.html:112
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:112
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:112
	return qs422016
//line views/layout/Menu.html:112
}

//line views/layout/Menu.html:114
func streammenuBadge(qw422016 *qt422016.Writer, i *menu.Item, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:115
	if i.Badge != "" {
//line views/layout/Menu.html:117
		badgeTitle := ""
		if idx := strings.Index(i.Badge, "**"); idx > -1 {
			badgeTitle = fmt.Sprintf(" title=%q", i.Badge[idx+2:])
			i.Badge = i.Badge[:idx]
		}

//line views/layout/Menu.html:123
		if strings.HasPrefix(i.Badge, ":") {
//line views/layout/Menu.html:124
			components.StreamIndent(qw422016, true, indent)
//line views/layout/Menu.html:124
			qw422016.N().S(`<div class="badge"`)
//line views/layout/Menu.html:125
			qw422016.N().S(badgeTitle)
//line views/layout/Menu.html:125
			qw422016.N().S(`>`)
//line views/layout/Menu.html:125
			components.StreamSVGSimple(qw422016, i.Badge[1:], 18, ps)
//line views/layout/Menu.html:125
			qw422016.N().S(`</div>`)
//line views/layout/Menu.html:126
		} else {
//line views/layout/Menu.html:127
			components.StreamIndent(qw422016, true, indent)
//line views/layout/Menu.html:127
			qw422016.N().S(`<div class="badge"`)
//line views/layout/Menu.html:128
			qw422016.N().S(badgeTitle)
//line views/layout/Menu.html:128
			qw422016.N().S(`>`)
//line views/layout/Menu.html:128
			qw422016.E().S(i.Badge)
//line views/layout/Menu.html:128
			qw422016.N().S(`</div>`)
//line views/layout/Menu.html:129
		}
//line views/layout/Menu.html:130
	}
//line views/layout/Menu.html:131
}

//line views/layout/Menu.html:131
func writemenuBadge(qq422016 qtio422016.Writer, i *menu.Item, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:131
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:131
	streammenuBadge(qw422016, i, indent, ps)
//line views/layout/Menu.html:131
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:131
}

//line views/layout/Menu.html:131
func menuBadge(i *menu.Item, indent int, ps *cutil.PageState) string {
//line views/layout/Menu.html:131
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:131
	writemenuBadge(qb422016, i, indent, ps)
//line views/layout/Menu.html:131
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:131
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:131
	return qs422016
//line views/layout/Menu.html:131
}
