package main

import (
	"fmt"
	"os/exec"
)

func main() {
	lsCmd := exec.Command("bash", "-c", "ls", "-a", "-l", "-h")
	lsOut, _ := lsCmd.Output()

	fmt.Println(string(lsOut))

}
