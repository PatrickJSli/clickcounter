package xinput

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func XinputInstalled() bool {
	_, err := exec.LookPath("xinput")
	return err == nil
}

func WaitForClick() {
	commandString := fmt.Sprintf("xinput --test-xi2 --root | grep 'RawButtonPress' -m 1")
	command := exec.Command("bash", "-c", commandString)
	command.Stderr = os.Stdout
	command.Env = os.Environ()

	// Kill the command if the parent process dies
	command.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGTERM,
	}

	err := command.Run()

	if err != nil {
		command.Cancel()
		os.Exit(1)
	}
}
