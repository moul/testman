GOPKG ?=	moul.io/testman
DOCKER_IMAGE ?=	moul/testman
GOMOD_DIRS ?= .
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

integration: install
	cd examples && make integration

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp
	testman -h 2> .tmp/root-usage.txt
	testman test -h 2> .tmp/test-usage.txt
	testman list -h 2> .tmp/list-usage.txt
	embedmd -w README.md
