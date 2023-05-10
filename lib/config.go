package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config connection
type Config struct {
	Host string
	Name string
	User string
	Password string
	Port string
}

// GetEnv load env
func GetEnv() (conf *Config) {
	if _, err := os.Stat("lib/.env"); err == nil {
		err := godotenv.Load("lib/.env")
		if err != nil {
		 	log.Fatal("Error loading .env file")
		}
  } 

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	dbPort := os.Getenv("DB_PORT")

	data := &Config {
		Host: dbHost,
		Name: dbName,
		User: dbUser,
		Password: dbPwd,
		Port: dbPort,
	}
	return data
 }

// NewConfig connection
func NewConfig() *Config {
	conf := GetEnv()
	return &Config {
		Host: conf.Host,
		Name: conf.Name,
		User: conf.User,
		Password: conf.Password,
		Port: conf.Port,
	}
}

// DsnString postgresql driver
// func (c *Config) DsnString() (conn *sql.DB) {
// 	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable user=%s password=%s port=%s", c.Host, c.Name, c.User, c.Password, c.Port)
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }

// DsnStringGorm postgresql driver
func (c *Config) DsnStringGorm() (conn *gorm.DB) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable user=%s password=%s port=%s", c.Host, c.Name, c.User, c.Password, c.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}