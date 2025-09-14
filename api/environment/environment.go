package environment

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI     string
	PostgresURI  string
	JWTSecretKey string
	AWSAccessKey string
	AWSSecretKey string
	AWSRegion    string
	AWSBucket    string
}

var (
	err error
	config *Config
	once sync.Once
)

func LoadEnv() (*Config, error) {
	once.Do(func() {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Cannot load .env file")
			return
		} 

		config = &Config{
			MongoURI: os.Getenv("MONGO_URI"),
			PostgresURI: os.Getenv("POSTGRES_URI"),
			JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
			AWSAccessKey: os.Getenv("MINIO_ROOT_USER"),
			AWSSecretKey: os.Getenv("MINIO_ROOT_PASSWORD"),
			AWSRegion: os.Getenv("MINIO_REGION"),
			AWSBucket: os.Getenv("MINIO_BUCKET"),
		}
	})

	if err != nil {
		return nil, err
	}

	return config, nil


}