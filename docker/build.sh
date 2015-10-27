#!/bin/bash

cp `which agent` .
docker build -t nlan/agent:ver0.1 .

