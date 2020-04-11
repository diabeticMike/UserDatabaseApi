package config

type (
	// Configuration is struct for holding service's configuration info
	Configuration struct {
		MongoURL  string       `json:"MongoURL" validate:"required"`
		ListenURL string       `json:"ListenURL" validate:"required"`
		Log       LoggerConfig `json:"Log" validate:"required"`
	}

	// LoggerConfig is a struct for holding logger configuration
	LoggerConfig struct {
		ServiceName string `json:"ServiceName" validate:"required"`
		FileName    string `json:"FileName" validate:"required"`
		LogLevel    string `json:"LogLevel" validate:"required"`
	}
)
