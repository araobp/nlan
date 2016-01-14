package context

type Context struct {
	Cmd  func(string, ...string) error
	CmdP func(string, ...string) error
}

func (c *Context) GetCmd() (func(string, ...string) error, func(string, ...string) error) {
	return c.Cmd, c.CmdP
}
