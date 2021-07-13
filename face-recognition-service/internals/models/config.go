package models

import "github.com/caarlos0/env"


type Config struct {
	MongoUri string `env:"MONGO_URI" envDefault:"mongodb://face-recognition-mongo-service:27017/test"`
	RedisHost string `env:"REDIS_HOST"`
	RedisTTLSeconds int `env:"REDIS_TTL_SECONDS" envDefault:"60"`
	NumberOfBestMatches int `env:"NUMBER_OF_BEST_MATCHES" envDefault:"3"`
	MaxAmountOfPersons int64 `env:"MAX_AMOUNT_OF_PERSONS" envDefault:"10000"`
	NumberOfWorkers int `env:"NUMBER_OF_WORKERS" envDefault:"10"`
}

func ParseConfig() (Config, error) {
	var out Config
	if err := env.Parse(&out); err != nil {
		return Config{}, err
	}

	return out, nil
}
