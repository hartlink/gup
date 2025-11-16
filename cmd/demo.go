package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	ui "gup/internal"
	"gup/pkg/i18n"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: i18n.T("demo.short"),
	Long:  i18n.T("demo.long"),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		command := fmt.Sprintf("echo '%s' && sleep 2 && echo '%s'", i18n.T("demo.output"), i18n.T("demo.success"))
		description := i18n.T("demo.description")

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
	rootCmd.AddCommand(demoCmd)
}
