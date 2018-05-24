# Simple makefile.
REGISTRY := rt:5000
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))

# Simply use the commit hash as tag.
TAG := $(shell git rev-parse --short `git rev-list -n 1 --abbrev HEAD`)

build:
	docker build -t ${REGISTRY}/${PROJECT}:${TAG} .

push:
	docker push ${REGISTRY}/${PROJECT}:${TAG}

exec:
	docker run -it --rm -P --name ${PROJECT} ${REGISTRY}/${PROJECT}:${TAG} 
