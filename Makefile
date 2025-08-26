APP=$(shell basename $(shell git remote get-url origin))
REGISTRY=s4myr4y
VERSION=$(shell git describe --tags --always --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGTOS=linux
TARGETARCH=amd64

format:
	go fmt ./

lint:
	golint

test:
	go test -v 

get: 
	go get

build: format get
	CGO_ENABLED=0 GOOS=$(TARGTOS) GOARCH=$(TARGETARCH) go build -v -o kbot -ldflags "-X"=https://github.com/S4MYR4Y/kbot/cmd.AppVersion=${VERSION} 

image:
	docker build . -t $(REGISTRY)/$(APP):$(VERSION)-$(TARGETARCH)

push: 
	docker push $(REGISTRY)/$(APP):$(VERSION)-$(TARGETARCH)

clean:
	rm -rf kbot