package config

import (
	"adapteris/auth"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Secret []byte
	auth.Config
}

func Read() *Config {
	viper.SetConfigName("app-config")
	viper.SetConfigType("yaml")
	if cfgDir := os.Getenv("ADAPTERIS_CFG_DIR"); cfgDir != "" {
		viper.AddConfigPath(cfgDir)
	}
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Couldn't load the application config, because %+v", err)
	}

	var cfg = Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Couldn't unmarshal the config, because %+v", err)
	}

	return &cfg
}
