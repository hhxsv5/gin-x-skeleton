#!/bin/bash
GO="go"
if [ -n "$GOROOT" ]; then
    GO="$GOROOT/bin/go"
fi
CGO_ENABLED=0 $GO build -o runtime/main-release -ldflags '-s -w' main.go

echo "Done"