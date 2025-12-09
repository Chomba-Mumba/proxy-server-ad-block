package configs

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type resource struct {
	Name           string
	Endpoint       string
	DestinationURL string
}

type configuration struct {
	Server struct {
		Host       string
		ListenPort string
	}
	Resources []resource
}

func NewConfiguration() (*configuration, error) {

	viper.AddConfigPath("../settings")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("no config file was found")
		} else {
			return nil, fmt.Errorf("error loading config file: %s", err)
		}
	}

	var Config configuration

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to decode config into struct %v", err)
	}

	return &Config, nil
}
