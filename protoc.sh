#!/bin/bash

#protoc -I ./helloworld/helloworld ./helloworld/helloworld/helloworld.proto --go_out=plugins=grpc:helloworld/helloworld
protoc -I ./grpc-test/api ./grpc-test/api/agent_api.proto --go_out=plugins=grpc:grpc-test/api
