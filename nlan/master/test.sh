#!/bin/bash

export ETCD_ADDRESS="http://localhost:2379"
./main -state $GOPATH/src/github.com/araobp/go-nlan/nlan/etc/ptn.yaml 
