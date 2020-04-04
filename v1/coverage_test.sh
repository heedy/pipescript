#!/bin/bash
# Copied from https://gist.github.com/bfontaine/370f29e0e0a49376257a

set -e

# Run test coverage on each subdirectories and merge the coverage profile.

echo "mode: count" >| profile.cov

# Standard go tooling behavior is to ignore dirs with leading underscores
for dir in $(find . -maxdepth 10 -not -path './.git*' -not -path '*/_*' -type d);
do
if ls $dir/*.go &> /dev/null; then
    go test -covermode=count -coverprofile=$dir/profile.tmp $dir
    if [ -f $dir/profile.tmp ]
    then
        cat $dir/profile.tmp | tail -n +2 >> profile.cov
        rm $dir/profile.tmp
    fi
fi
done

# go tool cover -func profile.cov

# To submit the test coverage result to coveralls.io,
# use goveralls (https://github.com/mattn/goveralls)
goveralls -coverprofile=profile.cov -service=travis-ci -shallow
