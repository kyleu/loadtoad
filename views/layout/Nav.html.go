// Code generated by qtc from "Nav.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Nav.html:2
package layout

//line views/layout/Nav.html:2
import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/menu"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
)

//line views/layout/Nav.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Nav.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Nav.html:10
func StreamNav(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:10
	qw422016.N().S(`
<nav id="navbar">
  <a class="logo" href="`)
//line views/layout/Nav.html:12
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:12
	qw422016.N().S(`" title="`)
//line views/layout/Nav.html:12
	qw422016.E().S(util.AppName)
//line views/layout/Nav.html:12
	qw422016.N().S(` `)
//line views/layout/Nav.html:12
	qw422016.E().S(as.BuildInfo.String())
//line views/layout/Nav.html:12
	qw422016.N().S(`">`)
//line views/layout/Nav.html:12
	components.StreamSVGSimple(qw422016, ps.RootIcon, 32, ps)
//line views/layout/Nav.html:12
	qw422016.N().S(`</a>
  <div class="breadcrumbs">
    <a href="`)
//line views/layout/Nav.html:14
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:14
	qw422016.N().S(`" class="nav-root-icon" title="`)
//line views/layout/Nav.html:14
	qw422016.E().S(util.AppName)
//line views/layout/Nav.html:14
	qw422016.N().S(`">`)
//line views/layout/Nav.html:14
	components.StreamSVGBreadcrumb(qw422016, ps.RootIcon, ps)
//line views/layout/Nav.html:14
	qw422016.N().S(`</a>
    <a class="link nav-root-item" href="`)
//line views/layout/Nav.html:15
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:15
	qw422016.N().S(`">`)
//line views/layout/Nav.html:15
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:15
	qw422016.N().S(`</a>`)
//line views/layout/Nav.html:15
	StreamNavItems(qw422016, ps)
//line views/layout/Nav.html:15
	qw422016.N().S(`
  </div>
`)
//line views/layout/Nav.html:17
	if ps.SearchPath != "-" {
//line views/layout/Nav.html:17
		qw422016.N().S(`  <form action="`)
//line views/layout/Nav.html:18
		qw422016.E().S(ps.SearchPath)
//line views/layout/Nav.html:18
		qw422016.N().S(`" class="search" title="search">
    <input id="search-input" type="search" name="q" placeholder=" " />
    <div class="search-image" style="display: none;"><svg><use xlink:href="#svg-searchbox" /></svg></div>
  </form>
`)
//line views/layout/Nav.html:22
	}
//line views/layout/Nav.html:22
	qw422016.N().S(`  `)
//line views/layout/Nav.html:23
	StreamProfileLink(qw422016, as, ps)
//line views/layout/Nav.html:23
	qw422016.N().S(`
`)
//line views/layout/Nav.html:24
	if !ps.HideMenu {
//line views/layout/Nav.html:24
		qw422016.N().S(`  <input type="checkbox" id="menu-toggle-input" style="display: none;" />
  <label class="menu-toggle" for="menu-toggle-input"><div class="spinner diagonal part-1"></div><div class="spinner horizontal"></div><div class="spinner diagonal part-2"></div></label>
  `)
//line views/layout/Nav.html:27
		StreamMenu(qw422016, ps)
//line views/layout/Nav.html:27
		qw422016.N().S(`
`)
//line views/layout/Nav.html:28
	}
//line views/layout/Nav.html:28
	qw422016.N().S(`</nav>`)
//line views/layout/Nav.html:29
}

//line views/layout/Nav.html:29
func WriteNav(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:29
	StreamNav(qw422016, as, ps)
//line views/layout/Nav.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:29
}

//line views/layout/Nav.html:29
func Nav(as *app.State, ps *cutil.PageState) string {
//line views/layout/Nav.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:29
	WriteNav(qb422016, as, ps)
//line views/layout/Nav.html:29
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:29
	return qs422016
//line views/layout/Nav.html:29
}

//line views/layout/Nav.html:31
func StreamNavItem(qw422016 *qt422016.Writer, link string, title string, icon string, last bool, ps *cutil.PageState) {
//line views/layout/Nav.html:32
	if link != "" || last {
//line views/layout/Nav.html:32
		qw422016.N().S(`<a class="link`)
//line views/layout/Nav.html:33
		if last {
//line views/layout/Nav.html:33
			qw422016.N().S(` `)
//line views/layout/Nav.html:33
			qw422016.N().S(`last`)
//line views/layout/Nav.html:33
		}
//line views/layout/Nav.html:33
		qw422016.N().S(`" href="`)
//line views/layout/Nav.html:33
		qw422016.E().S(link)
//line views/layout/Nav.html:33
		qw422016.N().S(`">`)
//line views/layout/Nav.html:34
	}
//line views/layout/Nav.html:34
	qw422016.N().S(`<span title="`)
//line views/layout/Nav.html:35
	qw422016.E().S(title)
//line views/layout/Nav.html:35
	qw422016.N().S(`">`)
//line views/layout/Nav.html:35
	components.StreamSVGBreadcrumb(qw422016, icon, ps)
//line views/layout/Nav.html:35
	qw422016.N().S(`</span><span class="nav-item-title">`)
//line views/layout/Nav.html:36
	qw422016.E().S(title)
//line views/layout/Nav.html:36
	qw422016.N().S(`</span>`)
//line views/layout/Nav.html:37
	if link != "" || last {
//line views/layout/Nav.html:37
		qw422016.N().S(`</a>`)
//line views/layout/Nav.html:39
	}
//line views/layout/Nav.html:40
}

//line views/layout/Nav.html:40
func WriteNavItem(qq422016 qtio422016.Writer, link string, title string, icon string, last bool, ps *cutil.PageState) {
//line views/layout/Nav.html:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:40
	StreamNavItem(qw422016, link, title, icon, last, ps)
//line views/layout/Nav.html:40
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:40
}

//line views/layout/Nav.html:40
func NavItem(link string, title string, icon string, last bool, ps *cutil.PageState) string {
//line views/layout/Nav.html:40
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:40
	WriteNavItem(qb422016, link, title, icon, last, ps)
//line views/layout/Nav.html:40
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:40
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:40
	return qs422016
//line views/layout/Nav.html:40
}

//line views/layout/Nav.html:42
func StreamNavItems(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/layout/Nav.html:43
	for idx, bc := range ps.Breadcrumbs {
//line views/layout/Nav.html:45
		i := ps.Menu.GetByPath(ps.Breadcrumbs[:idx+1])
		if i == nil {
			i = menu.ItemFromString(bc, ps.DefaultNavIcon)
		}

//line views/layout/Nav.html:50
		components.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:50
		qw422016.N().S(`<span class="separator">/</span>`)
//line views/layout/Nav.html:52
		components.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:53
		StreamNavItem(qw422016, i.Route, i.Title, i.Icon, idx == len(ps.Breadcrumbs)-1, ps)
//line views/layout/Nav.html:54
	}
//line views/layout/Nav.html:55
}

//line views/layout/Nav.html:55
func WriteNavItems(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/layout/Nav.html:55
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:55
	StreamNavItems(qw422016, ps)
//line views/layout/Nav.html:55
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:55
}

//line views/layout/Nav.html:55
func NavItems(ps *cutil.PageState) string {
//line views/layout/Nav.html:55
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:55
	WriteNavItems(qb422016, ps)
//line views/layout/Nav.html:55
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:55
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:55
	return qs422016
//line views/layout/Nav.html:55
}

//line views/layout/Nav.html:57
func StreamProfileLink(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:57
	qw422016.N().S(`<a class="profile" title="Settings" href="`)
//line views/layout/Nav.html:58
	qw422016.E().S(ps.ProfilePath)
//line views/layout/Nav.html:58
	qw422016.N().S(`">`)
//line views/layout/Nav.html:59
	components.StreamSVGSimple(qw422016, `profile`, 24, ps)
//line views/layout/Nav.html:59
	qw422016.N().S(`</a>`)
//line views/layout/Nav.html:61
}

//line views/layout/Nav.html:61
func WriteProfileLink(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:61
	StreamProfileLink(qw422016, as, ps)
//line views/layout/Nav.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:61
}

//line views/layout/Nav.html:61
func ProfileLink(as *app.State, ps *cutil.PageState) string {
//line views/layout/Nav.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:61
	WriteProfileLink(qb422016, as, ps)
//line views/layout/Nav.html:61
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:61
	return qs422016
//line views/layout/Nav.html:61
}
