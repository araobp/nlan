package main

import (
	"fmt"
	_ "os"
	"os/exec"
)

func main() {

	cmd := exec.Command("date")
	out, _ := cmd.Output()
	fmt.Printf("%s\n", out)

}
