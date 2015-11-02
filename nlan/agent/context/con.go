package context

import (
	"log"
)

type Context struct {
	Cmd    func(string, ...string) error
	CmdP   func(string, ...string) error
	Logger *log.Logger
}

func (c *Context) GetCmd() (func(string, ...string) error, func(string, ...string) error) {
	return c.Cmd, c.CmdP
}
