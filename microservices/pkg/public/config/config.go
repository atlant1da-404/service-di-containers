package config

type Config struct {
	Address   string `json:"address" yaml:"address"`
	Database  string `json:"database" yaml:"database"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
}

func GetConfig() *Config {
	// TODO: add file for example
	return &Config{Address: ":3000"}
}
