#!/bin/bash

set -e

OS="darwin linux windows"
ARCH="amd64 386"

rm -Rf bin/
mkdir bin/

for GOOS in $OS; do
    for GOARCH in $ARCH; do
        arch="$GOOS-$GOARCH"
        binary="terraform-provider-sakuraiot"
        if [ "$GOOS" = "windows" ]; then
          binary="${binary}.exe"
        fi
        echo "Building $binary $arch"
        GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 govendor build -ldflags "-s -w"  -o $binary builtin/bins/provider-sakuraiot/main.go
        zip -r "bin/terraform-provider-sakuraiot_$arch" $binary
        rm -f $binary
    done
done
