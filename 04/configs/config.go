package configs

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"log"
)

var ProviderSet = wire.NewSet(NewConfig)

type Config struct {
	ServerHttpAddr string
	DatabaseDriver string
	DatabaseDsn    string
}

func NewConfig(configFile string) *Config {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Load config file error: %v\n", err)
	}

	return &Config{
		ServerHttpAddr: viper.GetString("server.http.addr"),
		DatabaseDriver: viper.GetString("data.database.driver"),
		DatabaseDsn:    viper.GetString("data.database.dsn"),
	}
}
