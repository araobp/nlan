#!/bin/bash

echo "--- PYANG ---"
pyang --format=tree *.yang

echo ""
echo "--- GOYANG ---"
goyang --format=tree *.yang

echo ""
echo "--- GOYANG (proto) --"
goyang --format=proto *.yang

goyang --format=proto bridges.yang > bridges.proto
goyang --format=proto vxlan.yang > vxlan.proto
goyang --format=proto subnets.yang > subnets.proto

