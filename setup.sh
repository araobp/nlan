#!/bin/bash

echo "Compiling NLAN model..."
cd model
./yang.sh
./protoc.sh
cd ..

echo "Building NLAN master..."
cd master
go build -o master

echo "Starting tega db..."
./tega/tegadb >/dev/null 2>&1 &

echo "Building containers with NLAN agent embedded..."
cd docker
go build docker_mng.go
./restart.sh
cd ..

echo "Done!"
echo ""
