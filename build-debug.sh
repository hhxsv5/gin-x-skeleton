#!/bin/bash
GIN="gin"
if [ -n "$GOPATH" ]; then
    GIN="$GOPATH/bin/gin"
fi
CGO_ENABLED=0 $GIN -p=8000 -a=5200 -b="runtime/main-debug" -i --all --excludeDir="runtime" --buildArgs="-ldflags '-s -w'" run main.go