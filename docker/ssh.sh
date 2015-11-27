#!/bin/bash

ip=`etcdctl get /hosts/$1 | cut -d '/' -f1`
ssh root@ip

