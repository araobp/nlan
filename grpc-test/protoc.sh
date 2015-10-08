#!/bin/bash

protoc -I ./api ./api/agent_api.proto --go_out=plugins=grpc:api
