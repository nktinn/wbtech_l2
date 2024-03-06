package config

type Config struct {
	HttpServerPort string
}

func GetConfig() Config {
	return Config{
		HttpServerPort: ":8080",
	}
}
