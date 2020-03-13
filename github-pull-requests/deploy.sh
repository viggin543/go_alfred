#!/usr/bin/env bash

function fail() {
  local reason=$1
  echo "cant proceed $reason failed!"
  exit 1
}

find . -name '*test.go' | go test || fail "tests"
go build -o Pullrequests main.go || fail "build"
aws s3 --profile personal_s3 cp Pullrequests  s3://opentikva/ --acl public-read
mv Pullrequests /Users/domrevigor/Dropbox/alfred/Alfred.alfredpreferences/workflows/user.workflow.7B1E72CE-CBC3-4366-B2A1-0D91E1C573ED/