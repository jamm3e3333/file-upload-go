package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type GcStorage struct {
	BucketName                   string `env:"GOOGLE_STORAGE_BUCKET_NAME"`
	BasePath                     string `env:"GOOGLE_STORAGE_BASE_PATH"`
	ServiceAccountKeyAbsFilePath string `env:"GOOGLE_STORAGE_SERVICE_ACCOUNT_FILE_PATH"`
	UploadFileTimeoutMinutes     uint32 `env:"GOOGLE_STORAGE_UPLOAD_FILE_TIMEOUT_MINUTES" env-default:"3"`
	SignedUrlExpirationHours     uint32 `env:"GOOGLE_STORAGE_SIGNED_URL_EXPIRATION_HOURS" env-default:"168"`
}

func NewGcsCfg() *GcStorage {
	cfg := &GcStorage{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("failed to create google cloud storage config: %s", err.Error())
	}

	return cfg
}
