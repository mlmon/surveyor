package main

import (
	"errors"
	"os/exec"
)

var Which = which

func which(binary string) bool {
	err := exec.Command("which", binary).Run()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode() == 0
		}
	}
	return true
}
