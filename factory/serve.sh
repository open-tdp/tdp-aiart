#!/bin/sh
#

export TDP_DEBUG=1

export CGO_ENABLED=0
export GO111MODULE=on

####################################################################

[ -d ../var ] || mkdir -p ../var

go mod tidy
go run main.go server -c ../var/server.yml
