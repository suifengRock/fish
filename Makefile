export GOPATH := $(shell pwd)
export PATH := ${PATH}:${GOPATH}/bin
export GOBIN := ${GOPATH}/bin
export BRANCH := $(shell git branch | grep '*' | tr -d '* ')
.PHONY: dependence watch build watch tpl model controller flow
UNAME_S := $(shell uname -s)
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

dependence:
	go install github.com/go-martini/martini

goquery:
	go get github.com/PuerkitoBio/goquery

server:
	go run server.go

xrom:
	go run xrom.go
	
mysql:
	go run mysql.go

main:
	go run main.go

start:clear-pkg dependence build run

gotest:
	go run gotest.go


tool:
	go run tools.go

json:
	go run json.go


