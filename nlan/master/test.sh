#!/bin/bash

export ETCD_ADDRESS="http://localhost:2379"
./main -state ../etc/ptn.yaml
