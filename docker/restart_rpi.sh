#!/bin/bash

go install github.com/araobp/nlan/agent
./docker_mng stop $@
./docker_mng rm $@
./build.sh Dockerfile_rpi
./docker_mng run $@
