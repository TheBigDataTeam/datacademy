package util

import "github.com/spf13/viper"

// Config stores all configuration
// The values are read by viper from config files or env variables
type Config struct {
	DBDriver    string `mapstructure:"DB_DRIVER"`
	DBSource    string `mapstructure:"DB_SOURCE"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	MongoSource string `mapstructure:"MONGO_SOURCE"`
	MongoDBname string `mapstructure:"MONGO_DB_NAME"`
	MongoCname  string `mapstructure:"MONGO_C_NAME"`
	AppURL      string `mapstructure:"APP_URL"`
}

// LoadConfig reads configuration from a file or env variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("api")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
