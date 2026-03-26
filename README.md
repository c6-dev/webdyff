# webdyff

A thin wrapper for [dyff](https://github.com/homeport/dyff) that builds a WebAssembly binary and provides a static webpage to run the diff, available [here](https://c6-dev.github.io/webdyff/).

## Build

          GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o web/app.wasm ./cmd/wasm
          cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" web/wasm_exec.js
The /web directory will contain all the necessary files to deploy to any static site hosting.  

The binary can be further optimized with [wasm-opt](https://github.com/WebAssembly/binaryen), if desired:  
          
          wasm-opt -Oz web/app.wasm -o web/app.wasm --all-features
