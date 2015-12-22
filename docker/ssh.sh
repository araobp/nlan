#!/bin/bash

ip=`etcdctl get /nlan/hosts/$1 | cut -d '/' -f1`
ssh root@$ip

