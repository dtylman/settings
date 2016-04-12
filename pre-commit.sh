#!/usr/bin/env bash

function check(){
    if [ $? -ne 0 ]
    then
        echo "failed"
        exit 1
    fi
}

echo "gofmt..."
gofmt -w .
echo "goimports..."
goimports -w .
echo "building..."
go build .
echo "building windows..."
GOOS=windows go build .
check
echo "testing..."
go test .

echo "vet..."
go tool vet -composites=false .
check
echo "lint.."
golint .
check

echo "done"