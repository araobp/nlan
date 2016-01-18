#!/bin/bash

MODELDIR=$GOPATH/src/github.com/araobp/nlan/model/nlan
protoc -I $MODELDIR $MODELDIR/nlan.proto --go_out=plugins=grpc:$MODELDIR
