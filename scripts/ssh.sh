#!/bin/bash

ip=`$GOPATH/src/github.com/araobp/tega/scripts/tegactl get hosts.$1 | tr -d \" | cut -d '/' -f1`
ssh root@$ip

