package util

import "github.com/spf13/viper"

type Config struct {
	DBSource      string `mapstructure:"DB_SOURCE"`
	ApiKey        string `mapstructure:"API_KEY"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path, env string) (Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName(env)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
