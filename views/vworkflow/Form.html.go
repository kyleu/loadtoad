// Code generated by qtc from "Form.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkflow/Form.html:1
package vworkflow

//line views/vworkflow/Form.html:1
import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/controller/cutil"
	"github.com/kyleu/loadtoad/app/loadtoad"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/kyleu/loadtoad/views/components"
	"github.com/kyleu/loadtoad/views/layout"
)

//line views/vworkflow/Form.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkflow/Form.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkflow/Form.html:10
type Form struct {
	layout.Basic
	Workflow *loadtoad.Workflow
	Archives []string
}

//line views/vworkflow/Form.html:16
func (p *Form) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkflow/Form.html:16
	qw422016.N().S(`
`)
//line views/vworkflow/Form.html:18
	w := p.Workflow
	nameHelp := "The name of the workflow"
	testsHelp := "An array of tests, see example"
	replsHelp := "String replacements, applied throughout the workflow"
	varsHelp := "Custom variable keys and values used by the workflow"
	scriptsHelp := "JavaScript files used by the workflow"

//line views/vworkflow/Form.html:24
	qw422016.N().S(`  <div class="card">
`)
//line views/vworkflow/Form.html:26
	if w.ID != "" {
//line views/vworkflow/Form.html:26
		qw422016.N().S(`    <div class="right">
      <a href="`)
//line views/vworkflow/Form.html:28
		qw422016.E().S(w.WebPath())
//line views/vworkflow/Form.html:28
		qw422016.N().S(`/delete" class="link-confirm" data-message="Are you sure?"><button>Delete</button></a>
    </div>
`)
//line views/vworkflow/Form.html:30
	}
//line views/vworkflow/Form.html:30
	qw422016.N().S(`    <h3>`)
//line views/vworkflow/Form.html:31
	components.StreamSVGRefIcon(qw422016, `sitemap`, ps)
//line views/vworkflow/Form.html:31
	qw422016.N().S(` `)
//line views/vworkflow/Form.html:31
	qw422016.E().S(ps.Title)
//line views/vworkflow/Form.html:31
	qw422016.N().S(`</h3>
    <div class="mt">
      <form action="" method="post">
        <table class="expanded min-200">
          <tbody>
            `)
//line views/vworkflow/Form.html:36
	components.StreamTableInput(qw422016, "name", "", "Name", w.Name, 6, nameHelp)
//line views/vworkflow/Form.html:36
	qw422016.N().S(`
            <tr>
              <td class="shrink">
                <label for="edit-tests" title="`)
//line views/vworkflow/Form.html:39
	qw422016.E().S(testsHelp)
//line views/vworkflow/Form.html:39
	qw422016.N().S(`"><strong>Tests</strong></label>
                <div>(<a href="#modal-tests">example</a>)</div>
`)
//line views/vworkflow/Form.html:41
	if len(p.Archives) > 0 {
//line views/vworkflow/Form.html:41
		qw422016.N().S(`                <div class="mt">
                  <select id="archive-select">
                    <option value="">choose an archive</option>
`)
//line views/vworkflow/Form.html:45
		for _, a := range p.Archives {
//line views/vworkflow/Form.html:45
			qw422016.N().S(`                    <option>`)
//line views/vworkflow/Form.html:46
			qw422016.E().S(a)
//line views/vworkflow/Form.html:46
			qw422016.N().S(`</option>
`)
//line views/vworkflow/Form.html:47
		}
//line views/vworkflow/Form.html:47
		qw422016.N().S(`                  </select>
                </div>
                <div class="mt">
                  <a href="" onclick="addArchive();return false;"><button>Add Archive</button></a>
                </div>
`)
//line views/vworkflow/Form.html:53
	}
//line views/vworkflow/Form.html:53
	qw422016.N().S(`              </td>
              <td>`)
//line views/vworkflow/Form.html:55
	components.StreamFormTextarea(qw422016, "tests", "edit-tests", 12, util.ToJSON(w.Tests), testsHelp)
//line views/vworkflow/Form.html:55
	qw422016.N().S(`</td>
            </tr>
            <tr>
              <td class="shrink">
                <label for="edit-replacements" title="`)
//line views/vworkflow/Form.html:59
	qw422016.E().S(replsHelp)
//line views/vworkflow/Form.html:59
	qw422016.N().S(`"><strong>Replacements</strong></label>
                <div>(<a href="#modal-repls">example</a>)</div>
              </td>
              <td>`)
//line views/vworkflow/Form.html:62
	components.StreamFormTextarea(qw422016, "replacements", "edit-replacements", 12, util.ToJSON(w.Replacements), replsHelp)
//line views/vworkflow/Form.html:62
	qw422016.N().S(`</td>
            </tr>
            <tr>
              <td class="shrink">
                <label for="edit-variables" title="`)
//line views/vworkflow/Form.html:66
	qw422016.E().S(varsHelp)
//line views/vworkflow/Form.html:66
	qw422016.N().S(`"><strong>Variables</strong></label>
                <div>(<a href="#modal-vars">example</a>)</div>
              </td>
              <td>`)
//line views/vworkflow/Form.html:69
	components.StreamFormTextarea(qw422016, "variables", "edit-variables", 12, util.ToJSON(w.Variables), varsHelp)
//line views/vworkflow/Form.html:69
	qw422016.N().S(`</td>
            </tr>
            <tr>
              <td class="shrink">
                <label for="edit-scripts" title="`)
//line views/vworkflow/Form.html:73
	qw422016.E().S(scriptsHelp)
//line views/vworkflow/Form.html:73
	qw422016.N().S(`"><strong>Scripts</strong></label>
                <div>(<a href="#modal-scripts">example</a>)</div>
              </td>
              <td>`)
//line views/vworkflow/Form.html:76
	components.StreamFormTextarea(qw422016, "scripts", "edit-scripts", 12, util.ToJSON(w.Scripts), scriptsHelp)
//line views/vworkflow/Form.html:76
	qw422016.N().S(`</td>
            </tr>
            <tr>
              <td colspan="2"><button type="submit">Save Workflow</button></td>
            </tr>
          </tbody>
        </table>
      </form>
    </div>
  </div>
  `)
//line views/vworkflow/Form.html:86
	streamformModals(qw422016, w, as, ps)
//line views/vworkflow/Form.html:86
	qw422016.N().S(`
  <script>
    function addArchive() {
      const har = document.querySelector("#archive-select").value;
      if (har === "") {
        alert("Please select an archive");
        return;
      }
      const ta = document.querySelector("#edit-tests");
      const curr = ta.value;
      const currArray = JSON.parse(curr);
      currArray.push({"har": har});
      ta.value = JSON.stringify(currArray, null, 2);
      console.log(har, curr);
    }
  </script>
`)
//line views/vworkflow/Form.html:102
}

//line views/vworkflow/Form.html:102
func (p *Form) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkflow/Form.html:102
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkflow/Form.html:102
	p.StreamBody(qw422016, as, ps)
//line views/vworkflow/Form.html:102
	qt422016.ReleaseWriter(qw422016)
//line views/vworkflow/Form.html:102
}

//line views/vworkflow/Form.html:102
func (p *Form) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkflow/Form.html:102
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkflow/Form.html:102
	p.WriteBody(qb422016, as, ps)
//line views/vworkflow/Form.html:102
	qs422016 := string(qb422016.B)
//line views/vworkflow/Form.html:102
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkflow/Form.html:102
	return qs422016
//line views/vworkflow/Form.html:102
}

//line views/vworkflow/Form.html:104
func streamformModals(qw422016 *qt422016.Writer, w *loadtoad.Workflow, as *app.State, ps *cutil.PageState) {
//line views/vworkflow/Form.html:104
	qw422016.N().S(`
  <div id="modal-tests" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Tests Example</h2>
      </div>
      <div class="modal-body">
        <div id="modal-tests-data" hidden="hidden" style="display:none;">`)
//line views/vworkflow/Form.html:113
	qw422016.E().S(util.ToJSON(loadtoad.ExampleWorkflow.Tests))
//line views/vworkflow/Form.html:113
	qw422016.N().S(`</div>
        <button onclick="clip('tests');">Copy to clipboard</button>
        <div class="mt">
          <pre>`)
//line views/vworkflow/Form.html:116
	qw422016.E().S(util.ToJSON(loadtoad.ExampleWorkflow.Tests))
//line views/vworkflow/Form.html:116
	qw422016.N().S(`</pre>
        </div>
      </div>
    </div>
  </div>
  <div id="modal-repls" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Replacements Example</h2>
      </div>
      <div class="modal-body">
        <div id="modal-repls-data" hidden="hidden" style="display:none;">`)
//line views/vworkflow/Form.html:129
	qw422016.E().S(util.ToJSON(loadtoad.ExampleWorkflow.Replacements))
//line views/vworkflow/Form.html:129
	qw422016.N().S(`</div>
        <button onclick="clip('repls');">Copy to clipboard</button>
        <div class="mt">
          <pre>`)
//line views/vworkflow/Form.html:132
	qw422016.E().S(util.ToJSON(loadtoad.ExampleWorkflow.Replacements))
//line views/vworkflow/Form.html:132
	qw422016.N().S(`</pre>
        </div>
      </div>
    </div>
  </div>
  <div id="modal-vars" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Variables Example</h2>
      </div>
      <div class="modal-body">
        <div id="modal-vars-data" hidden="hidden" style="display:none;">`)
//line views/vworkflow/Form.html:145
	qw422016.E().S(util.ToJSON(loadtoad.ExampleWorkflow.Variables))
//line views/vworkflow/Form.html:145
	qw422016.N().S(`</div>
        <button onclick="clip('vars');">Copy to clipboard</button>
        <div class="mt">
          <pre>`)
//line views/vworkflow/Form.html:148
	qw422016.E().S(util.ToJSON(loadtoad.ExampleWorkflow.Variables))
//line views/vworkflow/Form.html:148
	qw422016.N().S(`</pre>
        </div>
      </div>
    </div>
  </div>
  <div id="modal-scripts" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Variables Example</h2>
      </div>
      <div class="modal-body">
        <div id="modal-scripts-data" hidden="hidden" style="display:none;">`)
//line views/vworkflow/Form.html:161
	qw422016.E().S(util.ToJSON(loadtoad.ExampleWorkflow.Scripts))
//line views/vworkflow/Form.html:161
	qw422016.N().S(`</div>
        <button onclick="clip('scripts');">Copy to clipboard</button>
        <div class="mt">
          <pre>`)
//line views/vworkflow/Form.html:164
	qw422016.E().S(util.ToJSON(loadtoad.ExampleWorkflow.Scripts))
//line views/vworkflow/Form.html:164
	qw422016.N().S(`</pre>
        </div>
      </div>
    </div>
  </div>
  <script>
    function clip(k) {
      if (!navigator.clipboard) {
        return;
      }
      const el = document.getElementById("modal-" + k + "-data");
      navigator.clipboard.writeText(el.innerText);
    }
  </script>
`)
//line views/vworkflow/Form.html:178
}

//line views/vworkflow/Form.html:178
func writeformModals(qq422016 qtio422016.Writer, w *loadtoad.Workflow, as *app.State, ps *cutil.PageState) {
//line views/vworkflow/Form.html:178
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkflow/Form.html:178
	streamformModals(qw422016, w, as, ps)
//line views/vworkflow/Form.html:178
	qt422016.ReleaseWriter(qw422016)
//line views/vworkflow/Form.html:178
}

//line views/vworkflow/Form.html:178
func formModals(w *loadtoad.Workflow, as *app.State, ps *cutil.PageState) string {
//line views/vworkflow/Form.html:178
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkflow/Form.html:178
	writeformModals(qb422016, w, as, ps)
//line views/vworkflow/Form.html:178
	qs422016 := string(qb422016.B)
//line views/vworkflow/Form.html:178
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkflow/Form.html:178
	return qs422016
//line views/vworkflow/Form.html:178
}
