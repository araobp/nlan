#!/bin/bash

ROUTER=router
NUM=3

go install github.com/araobp/go-nlan/nlan/agent
./docker_mng.py $ROUTER stop $NUM
./docker_mng.py $ROUTER rm $NUM
./build.sh
./docker_mng.py $ROUTER run $NUM
