#!/bin/bash

ip=`$GOPATH/src/github.com/araobp/tega/scripts/tegactl get nlan.hosts.$1 | tr -d \" | cut -d '/' -f1`
ssh root@$ip

