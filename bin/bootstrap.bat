@ECHO OFF

rem Downloads and installs the Go libraries and tools needed in other scripts

@ECHO ON
go install github.com/air-verse/air@latest
go install github.com/valyala/quicktemplate/qtc@latest
go install gotest.tools/gotestsum@latest
go install mvdan.cc/gofumpt@latest
