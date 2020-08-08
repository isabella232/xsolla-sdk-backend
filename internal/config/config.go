package config

import (
	"github.com/mashinapetro/cleanenv"
)

type Config struct {
	ServerHost string `env:"X_SERVER_HOST" env-default:"0.0.0.0" env-description:"Address used for serving incoming requests"`
	ServerPort int    `env:"X_SERVER_PORT" env-default:"8081" env-description:"Port used for serving incoming requests"`

	LogPath     string `env:"X_LOGS_FILE_PATH" env-description:"Absolute path to file, where logs will be stored"`
	ProgramName string `env:"X_PROGRAM_NAME" env-description:"Program name"`

	MerchantID      int    `env:"X_MERCHANT_ID" env-required:"true" env-description:"Merchant ID"`
	ProjectID       int    `env:"X_PROJECT_ID" env-required:"true" env-description:"Project ID"`
	PublisherAPIKey string `env:"X_PUBLISHER_API_KEY" env-required:"true" env-description:"Publisher Api Key"`
}

func GetConfigDescription() (string, error) {
	return cleanenv.GetDescription(&Config{}, nil)
}

func Init() (Config, error) {
	c := Config{}
	err := cleanenv.ReadEnv(&c)
	return c, err
}
