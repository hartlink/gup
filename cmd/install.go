package cmd

import (
	"fmt"
	"gup/internal/apt"
	"gup/internal/i18n"
	"os"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [packages...]",
	Short: i18n.T("install.short"),
	Long:  i18n.T("install.long"),
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		if err := apt.Install(args, verbose); err != nil {
			fmt.Printf("❌ %s: %v\n", i18n.T("ui.error"), err)
			os.Exit(1)
		}

		fmt.Printf("✅ %s\n", i18n.T("install.success"))
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
