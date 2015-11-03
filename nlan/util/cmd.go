package util

import (
	"log"
	"os/exec"
)

const (
	CONFIG = iota
	RESTART
	DEBUG
)

type Command struct {
	Logger *log.Logger
}

// This function executes the command and waits for its output.
func (c *Command) Cmd(name string, arg ...string) error {
	out, err := exec.Command(name, arg...).Output()
	c.Logger.Printf("cmd: %s, %v\n", name, arg)
	c.Logger.Println(string(out))
	return err
}

// This function just skips executing the command.
func (c *Command) CmdSkip(name string, arg ...string) error {
	c.Logger.Printf("cmd skipped: %s, %v\n", name, arg)
	return nil
}

// This function returns Cmd or CmdSkip function.
// When restarting NLAN agent, restart must be true.
func GetCmd(logger *log.Logger, mode int) (func(string, ...string) error, func(string, ...string) error) {
	var c = Command{Logger: logger}

	var f1 func(string, ...string) error
	switch mode {
	case DEBUG:
		f1 = func(name string, arg ...string) error {
			return c.CmdSkip(name, arg...)
		}
	default:
		f1 = func(name string, arg ...string) error {
			return c.Cmd(name, arg...)
		}
	}

	var f2 func(string, ...string) error
	switch mode {
	case RESTART, DEBUG:
		f2 = func(name string, arg ...string) error {
			return c.CmdSkip(name, arg...)
		}
	default:
		f2 = func(name string, arg ...string) error {
			return c.Cmd(name, arg...)
		}
	}

	return f1, f2
}
