package config

import (
	"errors"
	"path"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

type AppConfig struct {
	NotionToken string
	PageID      string
}

func LoadConfig() (AppConfig, error) {
	if err := setupViper(); err != nil {
		return AppConfig{}, err
	}
	return AppConfig{
		NotionToken: viper.GetString("notionToken"),
		PageID:      viper.GetString("pageID"),
	}, nil
}

func setupViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path.Join(xdg.ConfigHome, "notion-cli"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("missing config file, use cfg option to setup")
		}
		return err
	}
	return nil
}

func SaveConfig(cfg AppConfig) error {
	viper.Set("notionToken", cfg.NotionToken)
	viper.Set("pageID", cfg.PageID)
	err := viper.WriteConfigAs(path.Join(xdg.ConfigHome, "notion-cli", "config.yaml"))
	return err
}
