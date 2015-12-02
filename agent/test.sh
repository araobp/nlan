#!/bin/bash

export ETCD_ADDRESS="http://localhost:2379"
export HOSTNAME="router1"
./main -ope add -state $GOPATH/src/github.com/araobp/go-nlan/nlan/etc/ptn.yaml -roster $GOPATH/src/github.com/araobp/go-nlan/nlan/etc/roster.yaml -mode config 
#./main -ope add -state $GOPATH/src/github.com/araobp/go-nlan/nlan/etc/ptn.yaml 
