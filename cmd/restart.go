package cmd

import (
	"fmt"
	"gup/internal/i18n"
	"gup/internal/system"
	"os"

	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: i18n.T("restart.short"),
	Long:  i18n.T("restart.long"),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		// Ask for confirmation
		if !system.PromptConfirmRestart() {
			fmt.Printf("\n✅ %s\n", i18n.T("restart.cancelled"))
			return
		}

		// Execute restart
		if err := system.Restart(verbose); err != nil {
			fmt.Printf("❌ %s: %v\n", i18n.T("ui.error"), err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}
