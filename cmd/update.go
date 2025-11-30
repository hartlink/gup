package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"gup/internal/apt"
	"gup/internal/i18n"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: i18n.T("update.short"),
	Long:  i18n.T("update.long"),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		if err := apt.Update(verbose); err != nil {
			fmt.Printf("❌ %s: %v\n", i18n.T("ui.error"), err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
