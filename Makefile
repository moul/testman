GOPKG ?=	moul.io/testman
DOCKER_IMAGE ?=	moul/testman
GOMOD_DIRS ?= .
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

integration: install
	cd examples && make integration
