#!/bin/bash

BIN_NAME="chipcom"

SCRIPT_DIR=$(cd $(dirname $0); pwd)
OUTPUT_PATH="${BIN_NAME}"

if [[ "$OSTYPE" == "cygwin" || "$OSTYPE" == "msys" ]]; then
    # Windows
    OUTPUT_PATH="${OUTPUT_PATH}.exe"
fi

rm "${OUTPUT_PATH}"
go build -mod=mod "-ldflags=-s -w" -trimpath -o "${OUTPUT_PATH}" ./main.go
