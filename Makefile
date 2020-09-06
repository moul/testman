GOPKG ?=	moul.io/testman
DOCKER_IMAGE ?=	moul/testman
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

integration: install
	cd examples && make integration
