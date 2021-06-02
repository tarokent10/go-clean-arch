package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	filename string = "config"
	filetype string = "yml"
)

// Config General Config
type Config struct {
	LOGLEVEL   int
	LOGFILE    string
	DBHOST     string
	DBPORT     string
	DBUSER     string
	DBPASSWORD string
	DATABASE   string
}

var (
	config Config
)

func init() {
	// default settings
	viper.SetDefault("LOGLEVEL", "0")
	viper.SetDefault("LOGFILE", "")
	viper.SetDefault("DBHOST", "localhost")
	viper.SetDefault("DBPORT", "3306")
	viper.SetDefault("DBUSER", "kentaro")
	viper.SetDefault("DBPASSWORD", "kentaro")
	viper.SetDefault("DATABASE", "picture")

	// override by environment valuables(no prefix)
	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()

	viper.SetConfigName(filename)
	viper.SetConfigType(filetype)
	current, _ := os.Getwd()
	viper.AddConfigPath(current + "/infrastructure/config/conf/")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("failed to read config. %s", err.Error())
	}

	var C Config
	if err := viper.Unmarshal(&C); err != nil {
		panic(err.Error())
	}
	config = C
	fmt.Println("conf file has read!")
}

// Get Config instance
func Get() Config {
	return config
}
