#!/bin/bash

echo "--- PYANG ---"
pyang --format=tree */*.yang

echo ""
echo "--- GOYANG ---"
goyang --format=tree */*.yang

echo ""
echo "--- GOYANG (proto) --"
goyang --format=proto */*.yang

# Add "syntex", since goyang does not add it
# Add "rpc", since goyang does not seem to support YANG rpc
cat nlan/rpc.proto > nlan/nlan.proto
# Remove "optional", since protobuf3 does not allow it
goyang --format=proto nlan/nlan.yang | awk '{gsub("optional ", "");print}' >> nlan/nlan.proto


