package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sync-file-tool",
	Short: "sync-file-tool is a CLI tool to sync files between main website and local files",
	Long: `sync-file-tool is a CLI tool to sync files between main website and local files
For using this tool, you need to have a config file in the root of your project.
The config file should be named sync-file-tool.yaml and should have the following structure:
account_name: your_account_name
security_key: your_security_key

This project was created by @HJYup. (danyil butov)`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
