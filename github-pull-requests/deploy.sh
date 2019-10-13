#!/usr/bin/env bash
go build -o pull-requests main.go
aws s3 --profile personal_s3 cp pull-requests  s3://opentikva/ --acl public-read
mv pull-requests /Users/domrevigor/Dropbox/alfred/Alfred.alfredpreferences/workflows/user.workflow.7B1E72CE-CBC3-4366-B2A1-0D91E1C573ED/