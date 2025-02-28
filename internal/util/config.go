package util

import "github.com/spf13/viper" // use viper for use env 

type Config struct {
	APPPORT  string `mapstructure:"APP_PORT"`
	APPNAME  string `mapstructure:"APP_NAME"`
	APPDEBUG string `mapstructure:"APP_DEBUG"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()
	err = viper.ReadConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&config)

	return
}
