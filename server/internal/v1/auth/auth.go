package auth

// "golang.org/x/oauth2/google"

// // const (
// // 	key    = "rand"
// // 	maxAge = 86400 * 30
// // 	isProd = false
// // )

// type Config struct {
// 	GoogleLoginConfig oauth2.Config
// }

// var AppConfig Config

// func NewAuth2() oauth2.Config {

// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No .env file found, assuming production settings")
// 	}
// 	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
// 	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
// 	redirectURL := os.Getenv("REDIRECT_URL")

// 	log.Println("loaded auth variables")
// 	AppConfig.GoogleLoginConfig = oauth2.Config{
// 		RedirectURL:  redirectURL,
// 		ClientID:     googleClientID,
// 		ClientSecret: googleClientSecret,
// 		Scopes: []string{
// 			"https://www.googleapis.com/auth/userinfo.email",
// 			"https://www.googleapis.com/auth/userinfo.profile",
// 		},
// 		Endpoint: google.Endpoint,
// 	}

// 	return AppConfig.GoogleLoginConfig

// }
