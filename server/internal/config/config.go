package config

import (
	"log"
	"os"
	"sync"

	databaseCfg "github.com/NhyiraAmofaSekyi/go-webserver/internal/db"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// "golang.org/x/oauth2"

var (
	AppConfig *Config
	once      sync.Once
)

type Config struct {
	DBConfig          *databaseCfg.DBConfig
	ClientURL         string
	APIHost           string
	GoogleLoginConfig *oauth2.Config
}

func Initialize(dbConfig *databaseCfg.DBConfig) {
	once.Do(func() {
		if dbConfig == nil {
			log.Fatal("Database configuration is nil")
		}
		AppConfig = &Config{
			DBConfig: dbConfig,
		}
	})

	AppConfig.ClientURL = os.Getenv("CLIENT_URL")
	AppConfig.APIHost = os.Getenv("API_HOST")

	googleCon := NewAuth()
	AppConfig.GoogleLoginConfig = &googleCon
}
func NewAuth() oauth2.Config {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, assuming production settings")
	}
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := os.Getenv("REDIRECT_URL")

	log.Println("loaded auth variables")
	GoogleLoginConfig := oauth2.Config{
		RedirectURL:  redirectURL,
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return GoogleLoginConfig

}
