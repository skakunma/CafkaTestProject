package config

import (
	storage2 "github.com/skakunma/CafkaTestProject/internal/storage"
	"go.uber.org/zap"
)

type Config struct {
	Sugar      *zap.SugaredLogger
	FlagAdress string
	Store      *storage2.Storage
	Salt       string
}

func NewConfig() (*Config, error) {
	config := &Config{}
	config.FlagAdress = ":8080"

	logger, err := zap.NewDevelopment()

	if err != nil {
		return nil, err
	}

	config.Sugar = logger.Sugar()

	store, err := storage2.NewStorage("host=localhost user=postgres password=example dbname=kafkaProject sslmode=disable")
	if err != nil {
		return nil, err
	}

	config.Store = &store
	return config, nil
}
