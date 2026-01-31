# Contributing

Source code is available at [GitHub](https://github.com/kyleu/loadtoad).

## Setup

```
git clone https://github.com/kyleu/loadtoad
cd loadtoad
./bin/bootstrap.sh
```

## Build and run

```
make build
./build/debug/loadtoad --help
./bin/dev.sh
```

## Before submitting changes

```
./bin/format.sh
./bin/check.sh
./bin/test.sh
make build
```

## Notes

- Go and Node are required for full builds.
- The dev server auto-rebuilds templates and client assets.
