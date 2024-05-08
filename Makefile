// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
.DEFAULT_GOAL=build

build:
	go fmt ./...
	go vet ./...
	go build -o bin/watch-my-ip app/*.go

install:
	cp bin/watch-my-ip /usr/local/sbin/watch-my-ip

golib-latest:
	go get -u github.com/skeletonkey/lib-core-go@latest
	go get -u github.com/skeletonkey/lib-instance-gen-go@latest

	go mod tidy

app-init:
	go generate
