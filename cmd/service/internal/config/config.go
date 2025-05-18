package config

import (
	"github.com/skakunma/CafkaTestProject/cmd/service/internal/storage"
	"go.uber.org/zap"
)

type Config struct {
	Sugar      *zap.SugaredLogger
	FlagAdress string
	Store      *storage.Storage
}

func NewConfig() (*Config, error) {
	config := &Config{}
	config.FlagAdress = ":8080"

	logger, err := zap.NewDevelopment()

	if err != nil {
		return nil, err
	}

	config.Sugar = logger.Sugar()

	store, err := storage.NewStorage("host=localhost user=postgres password=example dbname=kafkaProject sslmode=disable")
	if err != nil {
		return nil, err
	}

	config.Store = &store
	return config, nil

}
