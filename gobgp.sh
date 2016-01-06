#!/bin/bash

ROUTER=rr
ip=`etcdctl get /nlan/hosts/$ROUTER | cut -d '/' -f1`
gobgp -u $ip $@

