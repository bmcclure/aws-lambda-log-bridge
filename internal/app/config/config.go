package config

import "github.com/spf13/viper"

type Config struct {

}

func (Config) setDefaults() {
	viper.SetDefault("apps", []AppConfig{})
}
