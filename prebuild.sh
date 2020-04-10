#!/bin/bash
# go get -u github.com/jteeuwen/go-bindata/...
echo "Generating documentation..."
go-bindata -pkg=resources -nometadata -nocompress -prefix resources/ -o resources/bindata.go resources/docs/...
echo "Generating parser..."
# go get -u golang.org/x/tools/cmd/goyacc
goyacc -o parser.go -p parser parser.y
