#!/bin/bash

HOST=localhost
PORT=8739
ROUTER=rr

ip=`curl http://$HOST:$PORT/hosts/$ROUTER?tega_id=gobgp.sh | tr -d \" | cut -d '/' -f1`
gobgp -u $ip $@

