package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"gup/internal/i18n"
)

var rootCmd = &cobra.Command{
	Use:   "gup",
	Short: i18n.T("root.short"),
	Long:  i18n.T("root.long"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("🚀 %s\n", i18n.T("app.welcome"))
		fmt.Printf("   %s\n", i18n.T("app.description"))
		fmt.Println()
		fmt.Printf("📋 %s\n", i18n.T("cmd.available"))
		fmt.Printf("   • gup update    - %s\n", i18n.T("cmd.update.desc"))
		fmt.Printf("   • gup demo      - %s\n", i18n.T("cmd.demo.desc"))
		fmt.Printf("   • gup version   - %s\n", i18n.T("cmd.version.desc"))
		fmt.Println()
		fmt.Printf("💡 %s\n", i18n.T("cmd.help"))
		fmt.Printf("⚠️  %s\n", i18n.T("cmd.permissions"))
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	i18n.Init()
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, i18n.T("flag.verbose"))
	rootCmd.PersistentFlags().StringP("lang", "l", "", i18n.T("flag.lang"))

	// Hook para cambiar idioma antes de ejecutar cualquier comando
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if lang, _ := cmd.Flags().GetString("lang"); lang != "" {
			i18n.SetLanguage(lang)
			// Actualizar descripciones después de cambiar idioma
			updateCommandDescriptions()
		}
	}
}

// updateCommandDescriptions actualiza las descripciones de los comandos
func updateCommandDescriptions() {
	rootCmd.Short = i18n.T("root.short")
	rootCmd.Long = i18n.T("root.long")
}
