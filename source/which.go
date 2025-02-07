package source

import (
	"os/exec"
)

var Which = which

func which(binary string) bool {
	return exec.Command("which", binary).Run() == nil
}
