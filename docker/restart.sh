#!/bin/bash

NODES=( pe1 pe2 pe3 pe4 rr ce1 ce2 ce3 ce4 )

go install github.com/araobp/nlan/agent
./docker_mng stop ${NODES[@]}
./docker_mng rm ${NODES[@]}
./build.sh
./docker_mng run ${NODES[@]}
