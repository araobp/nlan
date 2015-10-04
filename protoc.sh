#!/bin/bash

protoc -I ./helloworld/helloworld ./helloworld/helloworld/helloworld.proto --go_out=plugins=grpc:helloworld/helloworld
