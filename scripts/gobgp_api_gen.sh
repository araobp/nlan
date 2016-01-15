#!/bin/bash

cd $GOPATH/src/github.com/osrg/gobgp

protoc -I ./api ./api/gobgp.proto --go_out=plugins=grpc:./api
