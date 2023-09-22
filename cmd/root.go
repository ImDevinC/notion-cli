package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/imdevinc/notion-cli/internal/config"
	"github.com/imdevinc/notion-cli/internal/notion"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var errorStyle = chalk.Red.NewStyle().WithBackground(chalk.ResetColor)
var cfg config.AppConfig

var rootCmd = &cobra.Command{
	Use:   "notion-cli",
	Short: "notion-cli \"<message>\"",
	Long:  `notion-cli will add your message to the configured pageID.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nclient := notion.NewClient(notion.Config{
			Token: cfg.NotionToken,
		})
		note := strings.Join(args, " ")
		err := nclient.AddNote(context.TODO(), cfg.PageID, note)
		if err != nil {
			fmt.Println(errorStyle.Style(err.Error()))
			os.Exit(1)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		fmt.Println(errorStyle.Style(err.Error()))
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(errorStyle.Style(err.Error()))
		os.Exit(1)
	}
}
