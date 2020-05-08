# ./Makefile
VERSION := $(shell date +'%Y%m%d%H').$(shell git rev-parse --short=8 HEAD)
NAME := $(shell echo core)

GOPWD := $(shell pwd)
GOBASEDIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../..)
GOHOMEDIR := $(abspath $(HOME)/.go)

GOPATH = $(GOBASEDIR):$(GOHOMEDIR)
GOBIN = $(GOHOMEDIR)/bin

$(info root makefile GOPATH=$(GOPATH))
$(info root makefile GOBIN=$(GOBIN))

all:
	@echo "Project:" $(NAME) $(VERSION)

builds: $(NAME)
$(NAME): *.go
	go build -o ./deploy/build/$(NAME) -v

run:
	go run main.go

release:
	mkdir -p deploy/releases/$(NAME)-"$(VERSION)"
	/src/$(NAME)

	rsync -avzr --delete \
		--filter='- $(NAME)-*' \
		--filter='- /$(NAME)' \
		--filter='+ /.git/' \
		--filter='+ /.gitignore/' \
		--filter='+ /releases/' \
		--filter='+ /glide.lock/' \
		--filter='+ /README.MD/' \
		--filter='- .*' \
		--filter='- *~' \
		--filter='- *.org' \
		. deploy/releases/$(NAME)-"$(VERSION)"/src/$(NAME)

	tar czf deploy/releases/$(NAME)-"$(VERSION)".tgz $(NAME)-"$(VERSION)"
