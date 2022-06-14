package config

import "github.com/spf13/viper"

type database struct {
	DbDriver string `mapstructure:"DB_DRIVER"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
	DBReload bool   `mapstructure:"DB_RELOAD"`
}
type server struct {
	Host string `mapstructure:"SERVER_HOST"`
	Port string `mapstructure:"SERVER_PORT"`
}
type Config struct {
	Database database `mapstructure:",squash"`
	Server   server   `mapstructure:",squash"`
}

func LoadConfig(configFile string, paths ...string) (Config, error) {
	config := Config{}
	viper.SetConfigName(configFile)
	viper.SetConfigType("env")
	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
