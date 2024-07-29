package config

import (
	"github.com/spf13/viper"
)

// InitConfig initializes the configuration using Viper.
// It reads the .env file in the current directory and sets the configuration values.
func InitConfig() {
	// Create loggers
	log := CreateLoggers(nil)

	// Set the configuration type to dotenv
	viper.SetConfigType("dotenv")

	// Add the current directory as a configuration path
	viper.AddConfigPath(".")

	// Set the configuration name to .env
	viper.SetConfigName(".env")

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		// Log the error and return if there was an error reading the configuration file
		log.Fatal(err)
		return
	}
}
