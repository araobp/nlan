#!/bin/bash

cp `which agent` .
docker build -f $1 -t nlan/agent:ver0.1 .

