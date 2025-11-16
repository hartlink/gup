package cmd

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	ui "gup/internal"
	"gup/pkg/i18n"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: i18n.T("update.short"),
	Long:  i18n.T("update.long"),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		// Verificar si no somos root y necesitamos sudo
		if os.Geteuid() != 0 {
			// Verificar permisos sudo ANTES de iniciar Bubble Tea
			fmt.Printf("🔐 %s\n", i18n.T("update.checking_sudo"))
			sudoCheck := exec.Command("sudo", "-v")
			sudoCheck.Stdin = os.Stdin
			sudoCheck.Stdout = os.Stdout
			sudoCheck.Stderr = os.Stderr
			if err := sudoCheck.Run(); err != nil {
				fmt.Printf("❌ %s: %v\n", i18n.T("update.sudo_required"), err)
				os.Exit(1)
			}
		}

		// Comando a ejecutar (usar -n para no-interactive sudo)
		command := "sudo -n apt update"
		if os.Geteuid() == 0 {
			command = "apt update"
		}

		description := i18n.T("update.description")

		if verbose {
			fmt.Printf("%s: %s\n", i18n.T("error.executing"), command)
		}

		// Crear el modelo de Bubble Tea
		model := ui.NewCommandModel(command, description)

		// Ejecutar la interfaz
		p := tea.NewProgram(model)
		if _, err := p.Run(); err != nil {
			fmt.Printf("%s: %v\n", i18n.T("error.interface"), err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
