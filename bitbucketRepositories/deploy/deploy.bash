#!/bin/zsh
cd ..
go build .
LOCAL_WF_PATH=/Users/domrevigor/Dropbox/alfred/Alfred.alfredpreferences/workflows/user.workflow.50357F67-777F-43EE-8B8B-AABBC46A1D39
cp bitbucketRepositories ${LOCAL_WF_PATH}
