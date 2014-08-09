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
	time go install server.go

get-martini:
	go get github.com/go-martini/martini

get-xrom:
	go get github.com/go-xorm/xorm

get-mysql:
	go get github.com/go-sql-driver/mysql

dependence:
	go install github.com/go-martini/martini

server:
	${GOBIN}/server

xrom:
	go run xrom.go
	
mysql:
	go run mysql.go

main:
	go run main.go

start:clear-pkg dependence build run


