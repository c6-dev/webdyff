# webdyff

A thin wrapper for [dyff](https://github.com/homeport/dyff) that builds a WebAssembly binary and provides a static webpage to run the diff, available [here](https://c6-dev.github.io/webdyff/).

## Build

          GOOS=js GOARCH=wasm go build -o web/app.wasm ./cmd/wasm
          cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" web/wasm_exec.js
