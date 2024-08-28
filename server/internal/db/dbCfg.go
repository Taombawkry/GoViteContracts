package databaseCfg

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/NhyiraAmofaSekyi/go-webserver/internal/db/database"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	DB *database.Queries
}

// var instance *DBConfig
// var once sync.Once

// // GetDBConfig provides a global access point to DBConfig and initializes it if not already done.
// func GetDBConfig() (*DBConfig, error) {
// 	var initErr error
// 	once.Do(func() {
// 		instance, initErr = newDBConfig()
// 	})
// 	return instance, initErr
// }

func NewDBConfig(env string) (*DBConfig, error) {
	var dbURL string
	if env == "development" {
		dbURL = os.Getenv("DB_URL")
	} else {
		dbURL = os.Getenv("DB_URL")
	}

	if dbURL == "" {
		return nil, fmt.Errorf("%s not found in environment", dbURL)
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to the database: %v", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't ping database: %v", err)
	}

	db := database.New(conn)

	dbConfig := &DBConfig{
		DB: db,
	}
	fmt.Printf("DBCfg.DB: %v\n", dbConfig.DB)

	return dbConfig, nil
}
