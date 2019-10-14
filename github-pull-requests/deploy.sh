#!/usr/bin/env bash
go build -o Pullrequests main.go
aws s3 --profile personal_s3 cp Pullrequests  s3://opentikva/ --acl public-read
mv Pullrequests /Users/domrevigor/Dropbox/alfred/Alfred.alfredpreferences/workflows/user.workflow.7B1E72CE-CBC3-4366-B2A1-0D91E1C573ED/