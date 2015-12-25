#!/bin/bash

export ETCD_ADDRESS="http://localhost:2379"
./nlan -reset
./nlan -clear -state $GOPATH/src/github.com/araobp/nlan/etc/$1.yaml 
./nlan -state $GOPATH/src/github.com/araobp/nlan/etc/$1.yaml 
