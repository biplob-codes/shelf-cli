#!/bin/bash
# build.sh
set -e  # exit immediately if any command fails
VERSION=${1:-dev}
BINARY_NAME="shelf"
LDFLAGS_PKG="github.com/biplob-codes/shelf-cli/cmd"
mkdir -p dist
targets=(
  "linux amd64"
  "linux arm64"
  "windows amd64"
  "darwin amd64"
  "darwin arm64"
)
for target in "${targets[@]}"; do
  read -r GOOS GOARCH <<< "$target"
  output="dist/${BINARY_NAME}-${GOOS}-${GOARCH}"
  [ "$GOOS" = "windows" ] && output="${output}.exe"
  echo "Building $output..."
  GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-X ${LDFLAGS_PKG}.version=$VERSION" -o "$output" .
done
echo "Done. Binaries in dist/"