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
  <a class="profile" href="`)
//line views/layout/Nav.html:19
	qw422016.E().S(ps.ProfilePath)
//line views/layout/Nav.html:19
	qw422016.N().S(`">
    `)
//line views/layout/Nav.html:20
	components.StreamSVGRef(qw422016, `profile`, 24, 24, ``, ps)
//line views/layout/Nav.html:20
	qw422016.N().S(`
  </a>`)
//line views/layout/Nav.html:21
	if !ps.HideMenu {
//line views/layout/Nav.html:21
		qw422016.N().S(`

  <input type="checkbox" id="menu-toggle-input" style="display: none;" />
  <label class="menu-toggle" for="menu-toggle-input"><div class="spinner diagonal part-1"></div><div class="spinner horizontal"></div><div class="spinner diagonal part-2"></div></label>
  `)
//line views/layout/Nav.html:25
		StreamMenu(qw422016, ps)
//line views/layout/Nav.html:25
	}
//line views/layout/Nav.html:25
	qw422016.N().S(`
</nav>`)
//line views/layout/Nav.html:26
}

//line views/layout/Nav.html:26
func WriteNav(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:26
	StreamNav(qw422016, as, ps)
//line views/layout/Nav.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:26
}

//line views/layout/Nav.html:26
func Nav(as *app.State, ps *cutil.PageState) string {
//line views/layout/Nav.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:26
	WriteNav(qb422016, as, ps)
//line views/layout/Nav.html:26
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:26
	return qs422016
//line views/layout/Nav.html:26
}

//line views/layout/Nav.html:28
func StreamNavItems(qw422016 *qt422016.Writer, m menu.Items, breadcrumbs cmenu.Breadcrumbs) {
//line views/layout/Nav.html:29
	for idx, bc := range breadcrumbs {
//line views/layout/Nav.html:31
		i := m.GetByPath(breadcrumbs[:idx+1])

//line views/layout/Nav.html:33
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:33
		qw422016.N().S(`<span class="separator">/</span>`)
//line views/layout/Nav.html:35
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:36
		if i == nil {
//line views/layout/Nav.html:38
			bcLink := ""
			if strings.Contains(bc, "||") {
				bci := strings.Index(bc, "||")
				bcLink = bc[bci+2:]
				bc = bc[:bci]
			}

//line views/layout/Nav.html:44
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:45
			qw422016.E().S(bcLink)
//line views/layout/Nav.html:45
			qw422016.N().S(`">`)
//line views/layout/Nav.html:45
			qw422016.E().S(bc)
//line views/layout/Nav.html:45
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:46
		} else {
//line views/layout/Nav.html:46
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:47
			qw422016.E().S(i.Route)
//line views/layout/Nav.html:47
			qw422016.N().S(`">`)
//line views/layout/Nav.html:47
			qw422016.E().S(i.Title)
//line views/layout/Nav.html:47
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:48
		}
//line views/layout/Nav.html:49
	}
//line views/layout/Nav.html:50
}

//line views/layout/Nav.html:50
func WriteNavItems(qq422016 qtio422016.Writer, m menu.Items, breadcrumbs cmenu.Breadcrumbs) {
//line views/layout/Nav.html:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:50
	StreamNavItems(qw422016, m, breadcrumbs)
//line views/layout/Nav.html:50
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:50
}

//line views/layout/Nav.html:50
func NavItems(m menu.Items, breadcrumbs cmenu.Breadcrumbs) string {
//line views/layout/Nav.html:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:50
	WriteNavItems(qb422016, m, breadcrumbs)
//line views/layout/Nav.html:50
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:50
	return qs422016
//line views/layout/Nav.html:50
}
