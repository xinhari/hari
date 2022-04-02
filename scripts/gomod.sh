#!/bin/bash

set -x

## Submit an update for go mod
git branch -D go-mod
git branch go-mod
git checkout go-mod
git pull origin main
GOPROXY=direct go get xinhari.com/xinhari@main
go fmt
go mod tidy
git add go.mod go.sum
git commit -m "Update go.mod"
git push origin go-mod
git checkout main
git branch -D go-mod
