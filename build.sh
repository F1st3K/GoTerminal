#!/bin/bash

mkdir -p bin
rm -rf bin/*
for file in *.go; do
    if [ -f "$file" ]; then
        go build -o "bin/${file%.go}" "$file"
    fi

done