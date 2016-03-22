#!/bin/bash

go install github.com/araobp/nlan/agent
./docker_mng.py stop $@
./docker_mng.py rm $@
./build.sh Dockerfile_rpi
./docker_mng.py run $@
