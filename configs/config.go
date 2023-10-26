package configs

import (
	"github.com/spf13/viper"
)

type SecretsConfig struct {
	SecretKey           string `mapstructure:"SECRET_KEY"`
	PostgresDb          string `mapstructure:"POSTGRES_DB"`
	PostgresUser        string `mapstructure:"POSTGRES_USER"`
	PostgresPassword    string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresHost        string `mapstructure:"POSTGRES_HOST"`
	PostgresPort        string `mapstructure:"POSTGRES_PORT"`
	Sslmode             string `mapstructure:"SSLMODE"`
	BaseUrl             string `mapstructure:"BASE_URL"`
	EmailSenderName     string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

type MainConfig struct {
	ServerPort       string `mapstructure:"server_port"`
	LogFile          string `mapstructure:"log_file"`
	FeatureEnabled   string `mapstructure:"feature_enabled"`
	MaxItemsPerPage  string `mapstructure:"max_items_per_page"`
	StaticPath       string `mapstructure:"static_path"`
	StaticImagesPath string `mapstructure:"static_images_path"`
}

func LoadAppConfig(path string) (mainConfig MainConfig, err error) {
	appViper := viper.New()
	setupViper(appViper, path, "config", "yaml")

	// Load the configurations
	if err = appViper.ReadInConfig(); err != nil {
		return
	}

	// Unmarshal the configurations into respective structs
	if err = appViper.Unmarshal(&mainConfig); err != nil {
		return
	}

	return mainConfig, nil
}

// LoadSecretConfig reads the secret configuration from a file or environment variables.
func LoadSecretConfig(path string) (secretConfig SecretsConfig, err error) {
	secretsViper := viper.New()
	setupViper(secretsViper, path, ".env", "env")

	// Load the configurations
	if err = secretsViper.ReadInConfig(); err != nil {
		return
	}

	// Unmarshal the configurations into respective structs
	if err = secretsViper.Unmarshal(&secretConfig); err != nil {
		return
	}

	return secretConfig, nil
}

// Helper function to set up a Viper instance
func setupViper(v *viper.Viper, path string, configname string, configtype string) {
	v.AddConfigPath(path)
	v.SetConfigName(configname) // Register config file name (no extension)
	v.SetConfigType(configtype) // Look for specific type
	v.AutomaticEnv()
}
