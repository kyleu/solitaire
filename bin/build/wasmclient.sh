#!/bin/bash

## Builds the WebAssembly library located at ./app/wasm

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

echo "building Solitaire WASM client library..."
mkdir -p build/wasm
GOOS=js GOARCH=wasm go build -o ./assets/wasm/solitaire.wasm ./app/wasm/main.go
