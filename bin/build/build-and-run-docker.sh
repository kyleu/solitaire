#!/bin/bash

## Builds the Docker image and runs it

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

echo "Building [linux amd64]..."
GOOS=linux GOARCH=amd64 make build
mv ./build/debug/solitaire .
docker build -t=solitaire -f=./tools/release/Dockerfile.release .
rm ./solitaire
docker run -it solitaire
