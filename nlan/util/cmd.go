package util

import (
	"log"
	"os/exec"
)

// This function executes the command and waits for its output.
func Cmd(name string, arg ...string) ([]byte, error) {
	out, err := exec.Command(name, arg...).Output()
	log.Printf("cmd: %s, %v\n", name, arg)
	log.Println(string(out))
	return out, err
}

// This function just skips executing the command.
func CmdSkip(name string, arg ...string) ([]byte, error) {
	log.Printf("cmd skipped: %s, %v\n", name, arg)
	return nil, nil
}

// This function returns Cmd or CmdSkip function.
// When restarting NLAN agent, restart must be true.
func GetCmdP(restart bool) func(string, ...string) ([]byte, error) {
	var f func(string, ...string) ([]byte, error)
	switch restart {
	case true:
		f = CmdSkip
	case false:
		f = Cmd
	}
	return f
}
