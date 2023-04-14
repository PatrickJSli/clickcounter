package xinput

import (
	"fmt"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type ClickMsg int
type CommandError int

func XinputInstalled() bool {
	_, err := exec.LookPath("xinput")
	return err == nil
}

func XinputRunning() bool {
	pid, err := exec.Command("pidof", "xinput").Output()
	return err == nil && len(pid) > 0
}

func WaitForClick() tea.Msg {
	commandString := fmt.Sprintf("xinput --test-xi2 --root | grep 'RawButtonPress' -m 1")
	command := exec.Command("bash", "-c", commandString)

	err := command.Run()

	if err != nil {
		return CommandError(1)
	} else {
		return ClickMsg(1)
	}
}

func QuitXinput() {
	exec.Command("pkill", "xinput").Run()
}
