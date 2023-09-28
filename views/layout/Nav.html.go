// Code generated by qtc from "Nav.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Nav.html:2
package layout

//line views/layout/Nav.html:2
import (
	"strings"

	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cmenu"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/lib/menu"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/vutil"
)

//line views/layout/Nav.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Nav.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Nav.html:13
func StreamNav(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:13
	qw422016.N().S(`
<nav id="navbar">
  <a class="logo" href="`)
//line views/layout/Nav.html:15
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:15
	qw422016.N().S(`" title="`)
//line views/layout/Nav.html:15
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:15
	qw422016.N().S(` `)
//line views/layout/Nav.html:15
	qw422016.E().S(as.BuildInfo.String())
//line views/layout/Nav.html:15
	qw422016.N().S(`">`)
//line views/layout/Nav.html:15
	components.StreamSVGRef(qw422016, ps.RootIcon, 32, 32, ``, ps)
//line views/layout/Nav.html:15
	qw422016.N().S(`</a>
  <div class="breadcrumbs">
    <a class="link" href="`)
//line views/layout/Nav.html:17
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:17
	qw422016.N().S(`">`)
//line views/layout/Nav.html:17
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:17
	qw422016.N().S(`</a>`)
//line views/layout/Nav.html:17
	StreamNavItems(qw422016, ps.Menu, ps.Breadcrumbs)
//line views/layout/Nav.html:17
	qw422016.N().S(`
  </div>
`)
//line views/layout/Nav.html:19
	if ps.SearchPath != "-" {
//line views/layout/Nav.html:19
		qw422016.N().S(`  <form action="`)
//line views/layout/Nav.html:20
		qw422016.E().S(ps.SearchPath)
//line views/layout/Nav.html:20
		qw422016.N().S(`" class="search" title="search">
    <input type="search" name="q" placeholder=" " />
    <div class="search-image" style="display: none;"><svg><use xlink:href="#svg-searchbox" /></svg></div>
  </form>
`)
//line views/layout/Nav.html:24
	}
//line views/layout/Nav.html:24
	qw422016.N().S(`  <a class="profile" href="`)
//line views/layout/Nav.html:25
	qw422016.E().S(ps.ProfilePath)
//line views/layout/Nav.html:25
	qw422016.N().S(`">
    `)
//line views/layout/Nav.html:26
	components.StreamSVGRef(qw422016, `profile`, 24, 24, ``, ps)
//line views/layout/Nav.html:26
	qw422016.N().S(`
  </a>
`)
//line views/layout/Nav.html:28
	if !ps.HideMenu {
//line views/layout/Nav.html:28
		qw422016.N().S(`  <input type="checkbox" id="menu-toggle-input" style="display: none;" />
  <label class="menu-toggle" for="menu-toggle-input"><div class="spinner diagonal part-1"></div><div class="spinner horizontal"></div><div class="spinner diagonal part-2"></div></label>
  `)
//line views/layout/Nav.html:31
		StreamMenu(qw422016, ps)
//line views/layout/Nav.html:31
		qw422016.N().S(`
`)
//line views/layout/Nav.html:32
	}
//line views/layout/Nav.html:32
	qw422016.N().S(`</nav>`)
//line views/layout/Nav.html:33
}

//line views/layout/Nav.html:33
func WriteNav(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:33
	StreamNav(qw422016, as, ps)
//line views/layout/Nav.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:33
}

//line views/layout/Nav.html:33
func Nav(as *app.State, ps *cutil.PageState) string {
//line views/layout/Nav.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:33
	WriteNav(qb422016, as, ps)
//line views/layout/Nav.html:33
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:33
	return qs422016
//line views/layout/Nav.html:33
}

//line views/layout/Nav.html:35
func StreamNavItems(qw422016 *qt422016.Writer, m menu.Items, breadcrumbs cmenu.Breadcrumbs) {
//line views/layout/Nav.html:36
	for idx, bc := range breadcrumbs {
//line views/layout/Nav.html:38
		i := m.GetByPath(breadcrumbs[:idx+1])

//line views/layout/Nav.html:40
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:40
		qw422016.N().S(`<span class="separator">/</span>`)
//line views/layout/Nav.html:42
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:43
		if i == nil {
//line views/layout/Nav.html:45
			bcLink := ""
			if strings.Contains(bc, "||") {
				bci := strings.Index(bc, "||")
				bcLink = bc[bci+2:]
				bc = bc[:bci]
			}

//line views/layout/Nav.html:51
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:52
			qw422016.E().S(bcLink)
//line views/layout/Nav.html:52
			qw422016.N().S(`">`)
//line views/layout/Nav.html:52
			qw422016.E().S(bc)
//line views/layout/Nav.html:52
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:53
		} else {
//line views/layout/Nav.html:53
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:54
			qw422016.E().S(i.Route)
//line views/layout/Nav.html:54
			qw422016.N().S(`">`)
//line views/layout/Nav.html:54
			qw422016.E().S(i.Title)
//line views/layout/Nav.html:54
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:55
		}
//line views/layout/Nav.html:56
	}
//line views/layout/Nav.html:57
}

//line views/layout/Nav.html:57
func WriteNavItems(qq422016 qtio422016.Writer, m menu.Items, breadcrumbs cmenu.Breadcrumbs) {
//line views/layout/Nav.html:57
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:57
	StreamNavItems(qw422016, m, breadcrumbs)
//line views/layout/Nav.html:57
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:57
}

//line views/layout/Nav.html:57
func NavItems(m menu.Items, breadcrumbs cmenu.Breadcrumbs) string {
//line views/layout/Nav.html:57
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:57
	WriteNavItems(qb422016, m, breadcrumbs)
//line views/layout/Nav.html:57
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:57
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:57
	return qs422016
//line views/layout/Nav.html:57
}
