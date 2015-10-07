#!/bin/bash

echo "--- PYANG ---"
pyang --format=tree *.yang

echo ""
echo "--- GOYANG ---"
goyang --format=tree *.yang

echo ""
echo "--- GOYANG (proto) --"
goyang --format=proto *.yang

echo 'syntax = "proto2";' > bridges.proto
echo 'syntax = "proto2";' > vxlan.proto
echo 'syntax = "proto2";' > subnets.proto
goyang --format=proto bridges.yang >> bridges.proto
goyang --format=proto vxlan.yang >> vxlan.proto
goyang --format=proto subnets.yang >> subnets.proto

