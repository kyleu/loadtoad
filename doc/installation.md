<!--- Content managed by Project Forge, see [projectforge.md] for details. -->
# Installation

## Pre-built binaries
Download any package from the [release page](https://github.com/kyleu/loadtoad/releases).

## Running with Docker
```shell
docker run -p 15550:15550 ghcr.io/kyleu/loadtoad:latest
docker run -p 15550:15550 ghcr.io/kyleu/loadtoad:latest-debug
```

## Built from source

### go install
```shell
go install github.com/kyleu/loadtoad@latest
```

### Source code

If you want to contribute to the project, please follow the steps on our [contributing guide](contributing).

If you just want to build from source for whatever reason, follow these steps:

```shell
git clone https://github.com/kyleu/loadtoad
cd loadtoad
make build
./build/debug/loadtoad --help
```

A script has been provided at `./bin/dev.sh` that will auto-reload when the project changes.

Note that you may need to run `./bin/bootstrap.sh` before building the project for the first time.
