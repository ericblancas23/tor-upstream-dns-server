package config

import "github.com/spf13/viper"

type TorConfig struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
}

type Config struct {
	Tor      TorConfig         `mapstructure:"tor"`
	Manual   map[string]string `mapstructure:"manual"`
	CacheTTL int               `mapstructure:"cache_ttl"`
	DnsTTL   int               `mapstructure:"dns_ttl"`
}

func ReadConfig(path string) (*Config, error) {
	var config Config
	viper.SetConfigName("config")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
