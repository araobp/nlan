package main

import (
	"net"

	pb "github.com/araobp/grpc-study/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":8282"
)
