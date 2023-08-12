<!--- Content managed by Project Forge, see [projectforge.md] for details. -->
# Releasing

Load Toad uses `goreleaser` to create build artifacts.

You can release to GitHub using `./bin/build/release.sh`, or test the release by running `./bin/build/release-test.sh`.

Your releases are available at https://github.com/kyleu/loadtoad/releases

### Checksums

All release binaries are checksummed, available in `checksums.txt` in the root of the release

### Docker Images

Multiple Docker images will be created. The main image is `ghcr.io/kyleu/loadtoad/x.x.x`, and a debug image is provided at `ghcr.io/kyleu/loadtoad/x.x.x-debug` that includes `delve` for debugging

### Notarization

Release binaries for macOS and iOS are notarized using Apple Notarization services

### Signing

Release binaries and the checksum file are signed using `gpg`

### Source Code

The source code will be bundled in the release, available as `loadtoad_x.x.x_source.zip`

### Universal Binaries

A universal macOS app will be created, containing the complete app for x86-64 and ARM

