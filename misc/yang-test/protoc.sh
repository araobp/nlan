#!/bin/bash

MODELDIR="model"

cd $MODELDIR
protoc -I ./bridges ./bridges/bridges.proto --go_out=plugins=grpc:./bridges
protoc -I ./vxlan ./vxlan/vxlan.proto --go_out=plugins=grpc:./vxlan
protoc -I ./subnets ./subnets/subnets.proto --go_out=plugins=grpc:./subnets
