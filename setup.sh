#!/bin/bash

HOMEDIR=$GOPATH/src/github.com/araobp/nlan 

echo "Compiling tega driver..."
go install github.com/araobp/tega/driver

echo "Compiling NLAN model..."
cd $HOMEDIR
cd model
./yang.sh
./protoc.sh
cd $HOMEDIR

echo "Building NLAN master..."
cd $HOMEDIR
cd master
go build -o master
cd $HOMEDIR

#echo "Starting tega db..."
#./tega/tegadb >/dev/null 2>&1 &

echo "Building containers with NLAN agent embedded..."
cd docker
go build docker_mng.go
./restart.sh
cd ..

echo "Done!"
echo ""
