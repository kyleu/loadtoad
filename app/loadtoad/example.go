package loadtoad

import (
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
)

var ExampleWorkflow = &Workflow{
	ID:   "example",
	Name: "Example Workflow",
	Tests: har.Selectors{
		{Har: "foo", Comment: "this will include all requests in the archive named \"foo\""},
		{Har: "foo", URL: "http://google.com/*", Comment: "this will include all requests matching the URL pattern in the archive named \"foo\""},
		{Har: "foo", Idx: 1, Comment: "this will the include the first request in the archive named \"foo\""},
		{Har: "foo", Mime: "application/json", Comment: "this will the include all requests matching the MIME type in the archive named \"foo\""},
		{Har: "foo", URL: "*domain.com*", Mime: "application/json", Idx: 1, Comment: "this will the include the first request matching the MIME type and URL"},
	},
	Replacements: map[string]string{
		"https://google.com": "https://test.google.com",
		"domain.com":         "domain.com||domain-2.com||domain-3.com",
	},
	Variables: util.ValueMap{
		"username": "admin",
		"password": "foo",
	},
	Scripts: []string{"test"},
}
