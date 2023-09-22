package cmd

import (
	"fmt"
	"os"

	"github.com/imdevinc/notion-cli/internal/config"
	"github.com/spf13/cobra"
)

var token string
var pageID string

func init() {
	cfgCommand.Flags().StringVarP(&token, "notion-token", "t", "", "Notion Integration Token")
	cfgCommand.Flags().StringVarP(&pageID, "page", "p", "", "PageID in Notion to add text to")
	cfgCommand.MarkFlagRequired("notion-token")
	cfgCommand.MarkFlagRequired("page")
	rootCmd.AddCommand(cfgCommand)
}

var cfgCommand = &cobra.Command{
	Use: "cfg",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.ParseFlags(os.Args); err != nil {
			fmt.Println(errorStyle.Style(err.Error()))
			os.Exit(1)
		}
		if err := config.SaveConfig(config.AppConfig{
			NotionToken: token,
			PageID:      pageID,
		}); err != nil {
			fmt.Println(errorStyle.Style(err.Error()))
			os.Exit(1)
		}
	},
}
