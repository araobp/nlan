package main

import (
	"bytes"
	"flag"
	"log"
	"net"
	"os"

	config_ptn "github.com/araobp/go-nlan/nlan/agent/config/ptn"
	con "github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/common"
	env "github.com/araobp/go-nlan/nlan/env"
	nlan "github.com/araobp/go-nlan/nlan/model/nlan"
	st "github.com/araobp/go-nlan/nlan/state"
	"github.com/araobp/go-nlan/nlan/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type agent struct {
	con *con.Context
}

func (a *agent) route(crud int, in *nlan.Request, configMode int) (exit uint32) {
	model := in.Model
	ptn := model.GetPtn()
	dvr := model.GetDvr()

	exit = 0
	defer func() {
		if r := recover(); r != nil {
			c := a.con
			logger := c.Logger
			logger.Println(r)
			exit = 1
		}
	}()

	if ptn != nil {
		config_ptn.Crud(crud, ptn, a.con)
	}
	if dvr != nil {
		//
	}
	return exit
}

// Returns a log message as a string
func (a *agent) logMessage() string {
	c := a.con
	buf := c.Logbuf
	return buf.String()
}

// Returns a pointer to a log buffer
func (a *agent) logBuf() *bytes.Buffer {
	c := a.con
	return c.Logbuf
}

// Returns a log file name
func logFile() string {
	target := os.Getenv("HOSTNAME")
	return "nlan-agent-" + target + ".log"
}

// gRPC Add method
func (a *agent) Add(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	exit := a.route(env.ADD, in, util.CONFIG)
	response := nlan.Response{Exit: exit, LogMessage: a.logMessage()}
	common.WriteLog(logFile(), a.logBuf())
	return &response, nil
}

// gRPC Update method
func (a *agent) Update(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	exit := a.route(env.UPDATE, in, util.CONFIG)
	response := nlan.Response{Exit: exit, LogMessage: a.logMessage()}
	common.WriteLog(logFile(), a.logBuf())
	return &response, nil
}

// gRPC Delete method
func (a *agent) Delete(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	exit := a.route(env.DELETE, in, util.CONFIG)
	response := nlan.Response{Exit: exit, LogMessage: a.logMessage()}
	common.WriteLog(logFile(), a.logBuf())
	return &response, nil
}

// gRPC Hello method
func (a *agent) Hello(ctx context.Context, cp *nlan.Capabilities) (*nlan.Capabilities, error) {
	// TODO: impl
	c := []string{"hello"}
	cs := nlan.Capabilities{Capability: c}
	return &cs, nil
}

func main() {
	target := os.Getenv("HOSTNAME")
	logPrefix := "[" + target + "] "
	ope := flag.String("ope", "ADD", "CRUD operation")
	filename := flag.String("state", "", "state file")
	roster := flag.String("roster", "", "roster file")
	mode := flag.String("mode", "config", "config mode")
	flag.Parse()

	var configMode int
	switch *mode {
	case "debug":
		configMode = util.DEBUG
	case "config":
		configMode = util.CONFIG
	case "restart":
		configMode = util.RESTART
	}

	var logbuf bytes.Buffer
	logger := log.New(&logbuf, logPrefix, log.LstdFlags)
	cmd, cmdp := util.GetCmd(logger, configMode, true)
	c := &con.Context{Cmd: cmd, CmdP: cmdp, Logger: logger, Logbuf: &logbuf}
	a := agent{con: c}

	defer func() {
		log.Print(logbuf.String())
		common.WriteLog(logFile(), &logbuf)
	}()

	var states *[]st.State
	switch {
	case *filename != "":
		logger.Print("### Direct config mode ###")

		switch *roster {
		case "":
			states, _ = common.ReadState(filename, nil)
		default:
			states, _ = common.ReadState(filename, roster)
		}
		for _, v := range *states {
			router := v.Router
			model := v.Model
			log.Printf("target: %s -- router: %s\n", target, router)
			if router == target {
				request := nlan.Request{Model: &model}
				var ope_ int
				switch *ope {
				case "add":
					ope_ = env.ADD
				case "update":
					ope_ = env.UPDATE
				case "delete":
					ope_ = env.DELETE
				}
				a.route(ope_, &request, configMode)
				log.Print(logbuf.String())
			}
		}
	default:
		logger.Print("### gRPC server mode ###")
		util.RegisterHost()
		listen, err := net.Listen("tcp", env.PORT)
		defer listen.Close()
		if err != nil {
			logger.Print(err)
		}
		s := grpc.NewServer()
		nlan.RegisterNlanAgentServer(s, &a)
		s.Serve(listen)
	}
}
