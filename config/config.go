package config

import (
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	EndPointServer1 string `mapstructure:"END_POINT_SERVER_1"`
	EndPointServer2 string `mapstructure:"END_POINT_SERVER_2"`
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
