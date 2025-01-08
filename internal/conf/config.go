package conf

import (
	"github.com/spf13/viper"
)

func LoadConfig() AppConfig {
	configPaths := []string{
		"config/", "../config/", "../../config/", "../../../config/",
		"../../../../config/",
	}

	for _, path := range configPaths {
		viper.SetConfigName("env.json")
		viper.SetConfigType("json")
		viper.AddConfigPath(path)
		err := viper.ReadInConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				continue
			} else {
				panic(err)
			}
		}
		var config AppConfig
		err = viper.Unmarshal(&config)
		if err != nil {
			panic(err)
		}
		return config
	}

	panic("config file not found")

}
