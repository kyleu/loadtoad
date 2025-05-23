@ECHO OFF

rem Downloads and installs the Go libraries and tools needed in other scripts

@ECHO ON
go install github.com/air-verse/air@latest
go install github.com/valyala/quicktemplate/qtc@latest
go install gotest.tools/gotestsum@latest
go install mvdan.cc/gofumpt@latest
go install github.com/nikolaydubina/go-cover-treemap@latest
go install github.com/go-delve/delve/cmd/dlv@latest
go install golang.org/x/mobile/cmd/gomobile@latest
go install golang.org/x/mobile/cmd/gobind@latest
