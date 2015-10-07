#!/bin/bash

protoc -I . ./bridges.proto --go_out=plugins=grpc:.
protoc -I . ./vxlan.proto --go_out=plugins=grpc:.
protoc -I . ./subnets.proto --go_out=plugins=grpc:.
