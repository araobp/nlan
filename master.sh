#!/bin/bash

export ETCD_ADDRESS="http://localhost:2379"
./nlan -clear -state $GOPATH/src/github.com/araobp/nlan/etc/ptn.yaml 
./nlan -state $GOPATH/src/github.com/araobp/nlan/etc/ptn.yaml 
