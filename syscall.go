package main

import (
	"os"
	"syscall"
)

func main() {
	syscall.Exec("/bin/ls", []string{"ls"}, os.Environ())
}
