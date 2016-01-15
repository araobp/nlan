#!/bin/bash

DIR=$GOPATH/src/github.com/araobp/nlan/master

$DIR/master -reset
$DIR/master -state $GOPATH/src/github.com/araobp/nlan/etc/$1.yaml 
