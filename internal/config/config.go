package config

import (
	"github.com/knadh/koanf/v2"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DatabaseURL  string `koanf:"DATABASE_URL"`
	Port         string `koanf:"PORT"`
	S3AccessKey  string `koanf:"S3_ACCESS_KEY"`
	S3SecretKey  string `koanf:"S3_SECRET_KEY"`
	S3BucketName string `koanf:"S3_BUCKET_NAME"`
	S3Endpoint   string `koanf:"S3_ENDPOINT"`
	Env          string `koanf:"ENV"`
	SecretKey    string `koanf:"SECRET_KEY"`
}

var Conf *Config

var k = koanf.New(".")

func Load() {
	if err := k.Load(file.Provider(".env"), dotenv.Parser()); err != nil {
		log.Infof("unable to find env file: %v", err)
		log.Info("falling back to env variables")
	}

	if err := k.Load(env.Provider("", ".", nil), nil); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	Conf = &Config{}
	err := k.Unmarshal("", Conf)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
}
