#!/bin/bash

echo "Compiling NLAN model..."
cd nlan/model
./yang.sh
./protoc.sh
cd ../..

echo "Buidling NLAN master..."
go build

echo "Building containers with NLAN agent embedded..."
cd docker
./etcd.sh >/dev/null 2>&1 &
./restart.sh
cd ..

echo "Done!"
echo ""
