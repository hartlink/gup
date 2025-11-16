package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"gup/pkg/i18n"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: i18n.T("update.short"),
	Long:  i18n.T("update.long"),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		// Mostrar título
		fmt.Printf("\n🚀 %s\n", i18n.T("ui.title"))
		fmt.Println()

		// Verificar si no somos root y necesitamos sudo
		if os.Geteuid() != 0 {
			fmt.Printf("🔐 %s\n", i18n.T("update.checking_sudo"))
		}

		// Preparar comando
		var cmdName string
		var cmdArgs []string

		if os.Geteuid() == 0 {
			cmdName = "apt"
			cmdArgs = []string{"update"}
		} else {
			cmdName = "sudo"
			cmdArgs = []string{"apt", "update"}
		}

		if verbose {
			fmt.Printf("🔧 %s: %s %v\n", i18n.T("error.executing"), cmdName, cmdArgs)
		}

		// Ejecutar comando con output en tiempo real
		fmt.Printf("⏳ %s...\n\n", i18n.T("update.description"))

		aptCmd := exec.Command(cmdName, cmdArgs...)
		aptCmd.Stdout = os.Stdout
		aptCmd.Stderr = os.Stderr
		aptCmd.Stdin = os.Stdin

		err := aptCmd.Run()

		fmt.Println()
		if err != nil {
			// apt update puede retornar exit codes no-cero por warnings
			// pero aún así completar la actualización de la lista de paquetes
			if exitErr, ok := err.(*exec.ExitError); ok {
				exitCode := exitErr.ExitCode()
				if exitCode == 100 {
					// Exit code 100 significa que hubo algunos errores pero la lista se actualizó
					fmt.Printf("⚠️  %s\n", i18n.T("update.partial_success"))
				} else {
					fmt.Printf("❌ %s: %v\n", i18n.T("ui.error"), err)
					os.Exit(1)
				}
			} else {
				fmt.Printf("❌ %s: %v\n", i18n.T("ui.error"), err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("✅ %s\n", i18n.T("ui.success"))
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
