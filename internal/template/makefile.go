package template

var (
	Makefile = `
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=xinhari.com/xinhari/api/proto

{{if ne .Type "web"}}
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go install xinhari.com/hari/cmd/protoc-gen-xinhari@latest
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/cosmtrek/air@latest
	go mod tidy

.PHONY: proto
proto:
    {{if eq .UseGoPath true}}
	protoc --proto_path=${GOPATH}/src:. --xinhari_out=${MODIFY}:. --go_out=${MODIFY}. proto/{{.Alias}}/{{.Alias}}.proto
    {{else}}
	protoc --proto_path=. --xinhari_out=${MODIFY}:. --go_out=${MODIFY}:. proto/{{.Alias}}/{{.Alias}}.proto
    {{end}}

.PHONY: build
build: proto
{{else}}
.PHONY: build
build:
{{end}}
	go build -o {{.Alias}}-{{.Type}} *.go
.PHONY: dev
dev:
	air
.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t {{.Alias}}-{{.Type}}:latest
`

	GenerateFile = `package main
//go:generate make proto
`
)
