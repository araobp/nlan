#!/bin/bash

go install github.com/araobp/go-nlan/nlan/agent
./docker_mng.py stop
./docker_mng.py rm
./build.sh
./docker_mng.py run
