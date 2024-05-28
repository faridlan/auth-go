package helper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// func SetupFlags() {
// 	flag.StringVar(&configPath, "config", "../.env", "path to get env")
// 	flag.Parse()
// }

// func init() {
// 	SetupFlags()
// }

func GetEnv() (*viper.Viper, error) {

	config := viper.New()
	configPath := os.Getenv("CONFIG")

	if configPath == "" {
		config.SetConfigFile("../.env")
	} else {
		config.SetConfigFile(configPath)
	}

	err := config.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	return config, nil

}
