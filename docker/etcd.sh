#!/bin/bash

etcd -advertise-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001
