package util

import (
	"log"
	"os/exec"
	"strings"
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
	out, err := exec.Command(name, arg...).CombinedOutput()
	c.Logger.Printf("cmd: %s %s", name, strings.Join(arg, " "))
	if len(out) > 0 {
		c.Logger.Println(string(out))
	}
	if err != nil {
		c.Logger.Println(err)
	}
	return err
}

// This function just skips executing the command.
func (c *Command) CmdSkip(name string, arg ...string) error {
	c.Logger.Printf("cmd skipped: %s %s", name, strings.Join(arg, " "))
	return nil
}

// This function returns Cmd or CmdSkip function.
// When restarting NLAN agent, restart must be true.
func GetCmd(logger *log.Logger, mode int, panicMode bool) (func(string, ...string) error, func(string, ...string) error) {
	var c = Command{Logger: logger}

	var f1 func(string, ...string) error
	switch mode {
	case DEBUG:
		f1 = func(name string, arg ...string) error {
			err := c.CmdSkip(name, arg...)
			if panicMode == true && err != nil {
				panic(err)
			} else {
				return err
			}

		}
	default:
		f1 = func(name string, arg ...string) error {
			err := c.Cmd(name, arg...)
			if panicMode == true && err != nil {
				panic(err)
			} else {
				return err
			}
		}
	}

	var f2 func(string, ...string) error
	switch mode {
	// TODO: RESTART mode
	//case RESTART, DEBUG:
	case DEBUG:
		f2 = func(name string, arg ...string) error {
			err := c.CmdSkip(name, arg...)
			if panicMode == true && err != nil {
				panic(err)
			} else {
				return err
			}
		}
	default:
		f2 = func(name string, arg ...string) error {
			err := c.Cmd(name, arg...)
			if panicMode == true && err != nil {
				panic(err)
			} else {
				return err
			}
		}
	}

	return f1, f2
}
