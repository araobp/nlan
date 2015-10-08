#!/bin/bash

MODELDIR="model"

cd $MODELDIR

echo "--- PYANG ---"
pyang --format=tree */*.yang

echo ""
echo "--- GOYANG ---"
goyang --format=tree */*.yang

echo ""
echo "--- GOYANG (proto) --"
goyang --format=proto */*.yang

echo 'syntax = "proto2";' > bridges/bridges.proto
echo 'syntax = "proto2";' > vxlan/vxlan.proto
echo 'syntax = "proto2";' > subnets/subnets.proto
echo 'syntax = "proto2";' > hello/hello.proto
goyang --format=proto bridges/bridges.yang >> bridges/bridges.proto
goyang --format=proto vxlan/vxlan.yang >> vxlan/vxlan.proto
goyang --format=proto subnets/subnets.yang >> subnets/subnets.proto
goyang --format=proto hello/hello.yang >> hello/hello.proto

