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

// This function executes the command and returns the output.
func OutputCmd(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).CombinedOutput()
	strout := string(out)
	if len(out) > 0 {
		log.Println(strout)
	}
	if err != nil {
		log.Println(err)
	}
	return strout, err
}

// This function executes the command.
func Cmd(name string, args ...string) error {
	_, err := OutputCmd(name, args...)
	return err
}

// This function just skips executing the command.
func CmdSkip(name string, args ...string) error {
	log.Printf("cmd skipped: %s %s", name, strings.Join(args, " "))
	return nil
}

// This function returns Cmd or CmdSkip function.
// When restarting NLAN agent, restart must be true.
func GetCmd(panicMode bool) (func(string, ...string) error, func(string, ...string) error) {

	var f1 func(string, ...string) error
	f1 = func(name string, args ...string) error {
		err := Cmd(name, args...)
		if panicMode == true && err != nil {
			panic(err)
		} else {
			return err
		}
	}

	var f2 func(string, ...string) error
	f2 = func(name string, args ...string) error {
		// TODO: Replace Cmd() with CmdSkip()
		err := Cmd(name, args...)
		if panicMode == true && err != nil {
			panic(err)
		} else {
			return err
		}
	}

	return f1, f2
}
