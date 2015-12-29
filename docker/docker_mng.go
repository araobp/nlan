package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	IMAGE = "nlan/agent:ver0.1" // Docker image name
)

func main() {
	ope := os.Args[1]
	nodes := os.Args[2:]
	var cmd func(string) ([]byte, error)
	switch ope {
	case "run":
		cmd = func(node string) ([]byte, error) {
			arg := []string{"run", "-i", "-t", "-d", "-v", "/tmp:/var/volume:rw", "--privileged", "--name", node, "--env", "HOSTNAME=" + node, IMAGE}
			return exec.Command("docker", arg...).Output()
		}
	case "rm":
		cmd = func(node string) ([]byte, error) {
			return exec.Command("docker", "rm", node).Output()
		}
	case "start":
		cmd = func(node string) ([]byte, error) {
			return exec.Command("docker", "start", node).Output()
		}
	case "stop":
		cmd = func(node string) ([]byte, error) {
			return exec.Command("docker", "stop", node).Output()
		}
	default:
		fmt.Println("Operation unidentified")
		os.Exit(1)
	}
	for _, node := range nodes {
		out, err := cmd(node)
		fmt.Printf("%s", out)
		if err != nil {
			fmt.Println("Operation failure")
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
