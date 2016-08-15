#!/bin/bash
# go get -u github.com/jteeuwen/go-bindata/...
echo "Generating documentation..."
go-bindata -pkg=resources -nometadata -nocompress -prefix resources/ -o resources/bindata.go resources/docs/... resources/data/...
echo "Generating parser..."
go tool yacc -o parser.go -p parser parser.y