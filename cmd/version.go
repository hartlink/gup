package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"gup/internal/i18n"
)

const (
	version = "1.0.0"
	build   = "dev"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: i18n.T("version.short"),
	Long:  i18n.T("version.long"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%s (build: %s)\n", i18n.T("app.name"), version, build)
		fmt.Println(i18n.T("version.tool"))
		fmt.Println(i18n.T("version.developed"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
