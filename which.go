package main

import (
	"os/exec"
)

var Which = which

func which(binary string) bool {
	err := exec.Command("which", binary).Run()
	if err != nil {
		return false
	}
	return true
}
