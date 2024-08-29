package essential

import (
	"errors"
	"os"
	"os/exec"
)

func ClearTerminal() error {
	var cmd *exec.Cmd = nil

	if OS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	if OS == "linux" || OS == "darwin" {
		cmd = exec.Command("clear")
	}
	if cmd == nil {
		return errors.New("can't clear terminal: OS is not supported")
	}

	cmd.Stdout = os.Stdout

	return cmd.Run()
}
