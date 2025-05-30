package database

import (
	"DigitalExchange/src/config"
	databaseDrivers "DigitalExchange/src/database/drivers"
	"database/sql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"sync"
)

var (
	connectOnce     sync.Once
	getInstanceOnce sync.Once
	instance        *Database
)

type IDatabaseDriver interface {
	Connect() error
	Close() error
	GetClient() *gorm.DB
	GetDB() *sql.DB
}

type Database struct {
	driver IDatabaseDriver
}

// Init initializes the database by establishing a connection.
func Init() (err error) {
	return GetInstance().Connect()
}

// Connect establishes a connection to the database if not already connected.
func (database *Database) Connect() (err error) {
	connectOnce.Do(func() {
		configs := config.GetInstance()
		dbPort, _ := strconv.Atoi(configs.Get("DB_PORT"))
		// Initialize the driver with configuration values
		database.driver = &databaseDrivers.Postgres{
			Username: configs.Get("DB_USERNAME"),
			Password: configs.Get("DB_PASSWORD"),
			Host:     configs.Get("DB_HOST"),
			Port:     dbPort,
			Database: configs.Get("DB_DATABASE"),
			SSLMode:  configs.Get("DB_SSL_MODE"),
			Timezone: configs.Get("APP_TZ"),
		}
		// Establish connection
		err = database.driver.Connect()
	})
	return
}

// Close terminates the database connection.
func (database *Database) Close() (err error) {
	err = database.driver.Close()
	log.Println("Database Service: Disconnected Successfully.")
	return
}

// GetClient retrieves the gorm.DB client from the database driver.
func (database *Database) GetClient() *gorm.DB {
	return database.driver.GetClient()
}

// GetDB retrieves the sql.DB client from the database driver.
func (database *Database) GetDB() *sql.DB {
	return database.driver.GetDB()
}

// GetInstance returns the singleton instance of the Database.
func GetInstance() *Database {
	getInstanceOnce.Do(func() {
		instance = &Database{} // Initialize the singleton instance
	})
	return instance
}
