package config


import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/go-bongo/bongo"
	"log"
	"os"
)

var DatabaseHostURL string
var DatabaseName string

var connectDB *bongo.Connection
var CertManager *bongo.Collection

func InitDBEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}
	DatabaseHostURL = os.Getenv("DATABASE_MONGODB_HOST_URL")
	DatabaseName = os.Getenv("DATABASE_NAME")
}

// Connect Database
func InitDBConnection() {
	// DB Connect
	connection, err := CreateConnectionDB()
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return
	}
	connectDB = connection
}

// Initialize Database Collections
func InitDBCollections() {
	CertManager = connectDB.Collection("certManager")
}

func CreateConnectionDB() (*bongo.Connection, error) {
	config := &bongo.Config{
		ConnectionString: DatabaseHostURL,
		Database:         DatabaseName,
	}
	connection, err := bongo.Connect(config)
	return connection, err
}

func CloseConnectionDB(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Connection to MongoDB closed.")
	return nil
}
