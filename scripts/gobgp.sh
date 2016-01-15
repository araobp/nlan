#!/bin/bash

ROUTER=rr
ip=`$GOPATH/src/github.com/araobp/tega/scripts/tegactl get /nlan/hosts/$ROUTER | tr -d \" | cut -d '/' -f1`
gobgp -u $ip $@

