package config

import "github.com/spf13/viper"

type DbConfig struct {
	MongoUri string `mapstructure:"uri"`
}

type Config struct {
	MongoDB DbConfig `mapstructure:"mongo"`
}

var config Config

func LoadConfig() (Config, error) {

	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("pkg/config/")

	if err := vp.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := vp.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

