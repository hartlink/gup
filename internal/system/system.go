package system

import (
	"bufio"
	"fmt"
	"gup/internal/i18n"
	"os"
	"os/exec"
	"strings"
)

// Restart reboots the system
func Restart(verbose bool) error {
	var cmdName string
	var cmdArgs []string

	if os.Geteuid() == 0 {
		cmdName = "reboot"
		cmdArgs = []string{"now"}
	} else {
		cmdName = "sudo"
		cmdArgs = []string{"reboot", "now"}
	}

	if verbose {
		fmt.Printf("🔧 %s: %s %v\n", i18n.T("error.executing"), cmdName, cmdArgs)
	}

	fmt.Printf("🔄 %s\n", i18n.T("restart.executing"))

	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

// PromptRestart asks the user if they want to restart the server
func PromptRestart() bool {
	fmt.Printf("\n%s ", i18n.T("upgrade.prompt_restart"))

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.ToLower(strings.TrimSpace(response))

	// Accept both Spanish (s/sí) and English (y/yes) responses
	return response == "s" || response == "sí" || response == "si" ||
		response == "y" || response == "yes"
}

// PromptConfirmRestart asks for confirmation before restarting
func PromptConfirmRestart() bool {
	fmt.Printf("\n%s\n", i18n.T("restart.warning"))
	fmt.Printf("%s ", i18n.T("restart.confirm"))

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.ToLower(strings.TrimSpace(response))

	// Accept both Spanish (s/sí) and English (y/yes) responses
	return response == "s" || response == "sí" || response == "si" ||
		response == "y" || response == "yes"
}
