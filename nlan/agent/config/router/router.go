package router

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/env"
	"github.com/araobp/go-nlan/nlan/model/nlan"
)

func Crud(crud int, in *nlan.Router, con *context.Context) {

	loopback := in.Loopback
	logger := con.Logger
	logger.Print("Router called...")
	var crudRouter func(string, *context.Context)

	switch crud {
	case env.ADD:
		crudRouter = addRouter
	case env.UPDATE:
		crudRouter = updateRouter
	case env.DELETE:
		crudRouter = deleteRouter
	default:
		logger.Fatal("CRUD unidentified")
	}

	logger.Printf("Loopback: %s", loopback)
	crudRouter(loopback, con)
	logger.Print("crudRouter() completed")
}

func addRouter(loopback string, con *context.Context) {

	cmd, _ := con.GetCmd()
	_ = con.Logger
	cmd("ip", "addr", "add", "dev", "lo", loopback)
}

func updateRouter(loopback string, con *context.Context) {
	//
}

func deleteRouter(loopback string, con *context.Context) {
	//
}
