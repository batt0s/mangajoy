package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf *viper.Viper

// Load config to config.Conf from config.json
func LoadConfig() error {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	vp.AddConfigPath("..")
	vp.AddConfigPath("./config")
	if err := vp.ReadInConfig(); err != nil {
		return err
	}
	Conf = vp
	log.Println("Loaded config.")
	return nil
}
