package databases

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/lib/pq"
	"github.com/matheuscdd/gamescrever/api/environment"
	"github.com/semeqpd/api_semeq/common/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/cd mongo-driver/mongo/readpref"
)

var (
	mongoClient *mongo.Client
	postgresClient *sql.DB
	once sync.Once
)

func initClients() {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	env, err := environment.LoadEnv()
	if err != nil {
		log.Fatalf("error on loading: %v", err)
		return
	}

	readPref := readpref.SecondaryPreferred()

	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(env.MongoURI).SetReadPreference(readPref).SetMaxPoolSize(1024))
	if err != nil {
		log.Fatalf("cannot load MongoDB: %v", err)
		return
	}

	postgresClient, err = sql.Open("postgres", env.PostgresURI)
	if err != nil {
		log.Fatalf("cannot load PostgreSQL: %v", err)
		return
	}

	_ = GetAWSClient()
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}

func GetPostgresClient() *sql.DB {
	return postgresClient
}

func GetAWSClient() *aws.Config {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	env, err := environment.LoadEnv()
	if err != nil {
		log.Fatalf("error on loading: %v", err)
		return nil
	}

	awsClient, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(env.AWSAccessKey, env.AWSSecretKey, "")),
		config.WithRegion(env.AWSRegion),
	)
	if err != nil {
		log.Fatalf("Failed to create AWS session => %v", err)
		return nil
	}

	return awsClient
}

func ConnectDatabases() {
	once.Do(initClients)
}