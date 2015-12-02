#!/bin/bash

protoc -I ./nlan ./nlan/nlan.proto --go_out=plugins=grpc:./nlan
