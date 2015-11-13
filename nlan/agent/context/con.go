package context

import (
	"bytes"
	"log"
)

type Context struct {
	Cmd    func(string, ...string) error
	CmdP   func(string, ...string) error
	Logger *log.Logger
	Logbuf *bytes.Buffer
}

func (c *Context) GetCmd() (func(string, ...string) error, func(string, ...string) error) {
	return c.Cmd, c.CmdP
}
