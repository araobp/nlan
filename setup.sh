#!/bin/bash

echo "Compiling NLAN model..."
cd model
./yang.sh
./protoc.sh
cd ..

echo "Building NLAN master..."
go build

echo "Building containers with NLAN agent embedded..."
cd docker
./etcd.sh >/dev/null 2>&1 &
./restart.sh
cd ..

echo "Done!"
echo ""
