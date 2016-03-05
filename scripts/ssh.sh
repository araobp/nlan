#!/bin/bash

HOST="localhost"
PORT="8739"

ip=`curl http://$HOST:$PORT/hosts/$1?tega_id=ssh.sh | tr -d \" | cut -d '/' -f1`
ssh root@$ip

