package cmd

import (
	"fmt"
	"gup/internal/apt"
	"gup/internal/i18n"
	"gup/internal/system"
	"os"

	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: i18n.T("upgrade.short"),
	Long:  i18n.T("upgrade.long"),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		if err := apt.Upgrade(verbose); err != nil {
			fmt.Printf("❌ %s: %v\n", i18n.T("ui.error"), err)
			os.Exit(1)
		}

		fmt.Printf("✅ %s\n", i18n.T("upgrade.success"))
		fmt.Println()

		// Prompt for restart
		if system.PromptRestart() {
			if err := system.Restart(verbose); err != nil {
				fmt.Printf("❌ %s: %v\n", i18n.T("ui.error"), err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
