export GOPATH := $(shell pwd)
export PATH := ${PATH}:${GOPATH}/bin
export GOBIN := ${GOPATH}/bin
export BRANCH := $(shell git branch | grep '*' | tr -d '* ')
.PHONY: dependence watch build watch tpl model controller flow fig
UNAME_S := $(shell uname -s)
DOCKER_RUN_GO := fig run --rm goapp
FIG_VERSION := $(shell fig --version 2>/dev/null)

ifeq ($(UNAME_S),Linux)
	XARGS := xargs --no-run-if-empty
else
	XARGS := xargs
endif
export XARGS

all:clear-pkg mysql

	
clear-pkg:
	rm -rf pkg

build:
	time go install xrom.go

get-martini:
	go get github.com/go-martini/martini

get-xrom:
	go get github.com/go-xorm/xorm

get-mysql:
	go get github.com/go-sql-driver/mysql

get-goquery:
	go get github.com/PuerkitoBio/goquery

dependence:get-xrom get-mysql get-martini get-goquery
	
install:
	go install github.com/go-sql-driver/mysql	
	go install github.com/go-xorm/xorm
	go install github.com/PuerkitoBio/goquery
	go install github.com/go-martini/martini

init:dependence install

server:
	go run server.go

xrom:
	time go run xrom.go
	
mysql:
	go run mysql.go

main:
	go run main.go

start:clear-pkg install build run

gotest:
	go run gotest.go


tool:
	go run tools.go

json:
	go run json.go

test:
	go run src/github.com/go-xorm/xorm/examples/goroutine.go

fig: 
ifdef FIG_VERSION
	@echo "Found fig version $(FIG_VERSION)"
else
	@echo fig Not found try to install it
	curl -L https://github.com/docker/fig/releases/download/1.0.1/fig-`uname -s`-`uname -m` > /usr/local/bin/fig; chmod +x /usr/local/bin/fig
endif

initc:fig
	$(DOCKER_RUN_GO) chmod +x dep.sh && ./dep.sh	

shell: fig
	$(DOCKER_RUN_GO) bash

mysqlc: fig
	$(DOCKER_RUN_GO) go run mysql.go

serverc: fig
	$(DOCKER_RUN_GO) go run server.go

xromc: fig
	$(DOCKER_RUN_GO) go run xrom.go

toolc: fig
	$(DOCKER_RUN_GO) go run tools.go

jsonc: fig
	$(DOCKER_RUN_GO) go run json.go




