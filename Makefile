.PHONY: build
.DEFAULT_GOAL := build_push


build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	docker build -t arekmano/restic-secret-store .

push:
	docker push arekmano/restic-secret-store:latest

build_push: build push