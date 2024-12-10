package cmd

import (
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync local file with main website",
	Long:  `Sync local file with main website. The sync will overwrite the main website file with the local file.`,
	Run: func(cmd *cobra.Command, args []string) {
		println("sync called")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
