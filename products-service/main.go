package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kolan92/producsts-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	env := os.Getenv("environment")
	if env == "" {
		env = "local"
	}

	log.Println("Starting service...")
	log.Printf("Env is: %s", env)

	config := readConfiguration(env)

	db := setupDB(config.ConnectionString)

	router := gin.Default()

	server := NewServer(db)

	server.RegisterRouter(router)

	router.Run(fmt.Sprintf("0.0.0.0:%d", config.Port))
}

func setupDB(addr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(addr))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Customer{}, &models.Order{}, &models.Product{}); err != nil {
		panic(err)
	}

	return db
}

type Config struct {
	ConnectionString string `json:"connectionString"`
	Port             int    `json:"port"`
}

func readConfiguration(env string) *Config {
	config := &Config{}

	f, err := os.Open("config." + env + ".json")
	if err != nil {
		log.Fatal("Can not find config file, error: ", err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		log.Fatal("Can not process config file, error: ", err)
	}

	return config
}
