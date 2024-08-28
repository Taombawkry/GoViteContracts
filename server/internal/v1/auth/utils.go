package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NhyiraAmofaSekyi/go-webserver/internal/config"
	dbCFG "github.com/NhyiraAmofaSekyi/go-webserver/internal/db"
	"github.com/NhyiraAmofaSekyi/go-webserver/internal/db/database"
	"github.com/NhyiraAmofaSekyi/go-webserver/utils/email"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

type OAuthUserData struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}
type FarcasterUserData struct {
	ID             string `json:"id"`
	CustodyAddress string `json:"custodyAddress"`
}
type EmailActionParams struct {
	UserID string `json:"userId"`
	Action string `json:"action"`
}

func CreateFarcasterSocialLogin(c context.Context, user FarcasterUserData, dbCFG dbCFG.DBConfig) (string, error) {

	existingUserID, err := CheckIfSocialLoginUserExists(&dbCFG, user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to find  user: %w", err)
	}
	if existingUserID != nil {
		return "", fmt.Errorf("user exists: %w", err)
	}

	defualtUserNameID := uuid.New()

	defualtUserName := defualtUserNameID.String()
	newUser, err := dbCFG.DB.CreateUserSocial(c, database.CreateUserSocialParams{
		Username:         defualtUserName,
		IsEmailVerified:  sql.NullBool{Bool: false, Valid: true},
		SignupType:       database.NullAccountType{AccountType: database.AccountTypeSocialLogin, Valid: true}, // Assuming this is a field in your database
		SubscriptionType: "standard",
	})
	log.Printf("user %v", user)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	sl, err := dbCFG.DB.CreateUserSocialLogin(c, database.CreateUserSocialLoginParams{
		UserID:         newUser.ID, // Ensure that newUser.ID is obtained correctly from CreateUserSocial's return
		ProviderID:     2,          // Presumed static assignment, ideally, this should be dynamic or configured
		ProviderUserID: user.ID,    // Presuming your SQL accepts NullTime// Assuming updates are also recorded
	})

	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}
	log.Printf("social login %v", sl)

	wallet, err := dbCFG.DB.CreateUserWallet(c, database.CreateUserWalletParams{
		WalletAddress: user.CustodyAddress,
		WalletIndex:   1,
		UserID:        newUser.ID,
		ChainID:       1,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create wallet login: %w", err)
	}
	log.Printf("user address %v", user.CustodyAddress)
	log.Printf("wallet address %v", wallet.WalletAddress)

	session, err := dbCFG.DB.CreateUserSession(c, database.CreateUserSessionParams{
		UserID:    newUser.ID,                                                      // user.ID should be the UUID of the user for whom the session is being created
		ExpiresAt: time.Now().Add(24 * time.Hour),                                  // Example: session expires in 24 hours
		IpAddress: sql.NullString{String: "user's IP address", Valid: true},        // Obtain from the request context or headers
		UserAgent: sql.NullString{String: "user's User-Agent header", Valid: true}, // Obtain from the request headers
	})
	if err != nil {
		return "", fmt.Errorf("failed to create social login: %w", err)
	}
	log.Printf("session %v", session)

	sessionToken, err := generateJWT(string(session.ID.String()))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return sessionToken, nil

}

func CreateSocialLogin(c context.Context, userdata []byte, dbCFG dbCFG.DBConfig, token *oauth2.Token) (string, error) {
	// Read and decode the user data from the request body
	var user OAuthUserData
	err := json.Unmarshal(userdata, &user)
	if err != nil {
		// Return an error with a specific error message
		return "", fmt.Errorf("failed to unmarshall: %w", err)
	}

	defualtUserNameID := uuid.New()

	defualtUserName := defualtUserNameID.String()
	newUser, err := dbCFG.DB.CreateUserSocial(c, database.CreateUserSocialParams{
		Username:         defualtUserName,
		Email:            sql.NullString{String: user.Email, Valid: true},
		IsEmailVerified:  sql.NullBool{Bool: user.VerifiedEmail, Valid: true},
		SignupType:       database.NullAccountType{AccountType: database.AccountTypeSocialLogin, Valid: true}, // Assuming this is a field in your database
		SubscriptionType: "standard",
		ProfileImageUrl:  sql.NullString{String: user.Picture, Valid: true},
	})
	log.Printf("user %v", user)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	// Assuming CreateUserSocialLogin returns any error
	sl, err := dbCFG.DB.CreateUserSocialLogin(c, database.CreateUserSocialLoginParams{
		UserID:            newUser.ID, // Ensure that newUser.ID is obtained correctly from CreateUserSocial's return
		ProviderID:        1,          // Presumed static assignment, ideally, this should be dynamic or configured
		ProviderUserID:    user.ID,
		AccessToken:       sql.NullString{String: token.AccessToken, Valid: true},
		RefreshToken:      sql.NullString{String: token.RefreshToken, Valid: true},
		AccessTokenExpiry: sql.NullTime{Time: token.Expiry, Valid: true}, // Presuming your SQL accepts NullTime// Assuming updates are also recorded
	})
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	log.Printf("social login %v", sl)

	session, err := dbCFG.DB.CreateUserSession(c, database.CreateUserSessionParams{
		UserID:    newUser.ID,                                                      // user.ID should be the UUID of the user for whom the session is being created
		ExpiresAt: time.Now().Add(24 * time.Hour),                                  // Example: session expires in 24 hours
		IpAddress: sql.NullString{String: "user's IP address", Valid: true},        // Obtain from the request context or headers
		UserAgent: sql.NullString{String: "user's User-Agent header", Valid: true}, // Obtain from the request headers
	})

	log.Printf("session %v", session)

	if err != nil {
		return "", fmt.Errorf("failed to create social login: %w", err)
	}

	sessionToken, err := generateJWT(string(session.ID.String()))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return sessionToken, nil

}

func CheckIfUserExists(dbConfig *dbCFG.DBConfig, userData []byte) (*uuid.UUID, error) {
	var user OAuthUserData
	err := json.Unmarshal(userData, &user)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling user data: %w", err)
	}

	dbUser, err := dbConfig.DB.GetSocialLoginUserByID(context.Background(), user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			return nil, nil
		}
		return nil, fmt.Errorf("error finding user by email: %w", err)
	}

	return &dbUser, nil
}

func CheckUserSessions(c context.Context, dbConfig *dbCFG.DBConfig, id *uuid.UUID) error {

	err := dbConfig.DB.DeleteSessionByUserID(c, *id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return fmt.Errorf("error finding user session %w", err)
	}

	return nil

}

func CheckIfSocialLoginUserExists(dbConfig *dbCFG.DBConfig, id string) (*uuid.UUID, error) {

	dbUser, err := dbConfig.DB.GetSocialLoginUserByID(context.Background(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			return nil, nil
		}
		return nil, fmt.Errorf("error finding user by email: %w", err)
	}

	return &dbUser, nil
}

// HashPassword hashes a plaintext password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a hashed password with a plaintext password.
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func generateSecureToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func setSessionCookie(w http.ResponseWriter, name, value string) {
	expire := time.Now().Add(24 * time.Hour) // 24 hours expiration
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expire,
		HttpOnly: true,  // Important: make cookie inaccessible to JavaScript
		Secure:   false, // Important: set to true in prod
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}

func getSessionFromCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func generateJWT(id string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour).Unix()
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id, // Include the name in the token
		"nbf": time.Now().Unix(),
		"exp": expirationTime,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RefreshAccessToken(refreshToken string) (*oauth2.Token, error) {
	tokenSource := config.AppConfig.GoogleLoginConfig.TokenSource(context.Background(), &oauth2.Token{
		RefreshToken: refreshToken,
	})
	newToken, err := tokenSource.Token() // This automatically uses the refresh token if the access token is expired
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

func CreateEmailAction(dbConfig *dbCFG.DBConfig, params EmailActionParams) (string, error) {

	userUUID, err := uuid.Parse(params.UserID)
	if err != nil {
		return "", fmt.Errorf("invalid user ID format: %v", err)
	}

	token := uuid.NewString()

	expiryDate := time.Now().Add(5 * time.Hour)

	record, err := dbConfig.DB.CreateUserEmail(context.Background(), database.CreateUserEmailParams{ // Generate a new UUID for the record ID
		UserID:      userUUID,
		Token:       token,
		ActionType:  database.EmailActions(params.Action),
		ExpiryDate:  expiryDate,
		IsCompleted: false,
	})

	if err != nil {
		return "", fmt.Errorf("failed to insert email action record: %w", err)
	}
	return record.Token, nil
}

func SendTokenEmail(recipientEmail, token string) error {
	subject := "Your Verification Token"
	body := fmt.Sprintf("Hello,\n\nHere is your verification token: %s\n\nPlease use this token to complete your verification process.\n\nThank you!", token)

	// Call the existing SendMail function to send the email
	err := email.SendMail(subject, recipientEmail, body)
	if err != nil {
		return fmt.Errorf("failed to send token email: %w", err)
	}

	return nil
}

func GetUserByID(ctx context.Context, dbConfig *dbCFG.DBConfig, id string) (*database.User, error) {

	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	user, err := dbConfig.DB.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	return &user, nil
}

func GetUserEmailByID(ctx context.Context, dbConfig *dbCFG.DBConfig, id string) (string, error) {

	userID, err := uuid.Parse(id)
	if err != nil {
		return "", fmt.Errorf("invalid UUID format: %w", err)
	}

	// Call the SQLC-generated function to fetch the user's email
	Email, err := dbConfig.DB.GetUserEmailById(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve user email: %w", err)
	}

	// Check if the email is valid
	if !Email.Valid {
		return "", fmt.Errorf("no email found for user")
	}

	return Email.String, nil
}

func UpdateUserEmailWithToken(ctx context.Context, dbConfig *dbCFG.DBConfig, token string) error {
	// Retrieve the user ID associated with the token
	userID, err := dbConfig.DB.GetEmailToken(ctx, token)
	if err != nil {
		return fmt.Errorf("failed to retrieve user ID: %w", err)
	}

	err = dbConfig.DB.UpdateUserEmail(ctx, database.UpdateUserEmailParams{
		ID:              userID,
		IsEmailVerified: sql.NullBool{Bool: true, Valid: true},
	})

	if err != nil {
		return fmt.Errorf("failed to update user email: %w", err)
	}

	fmt.Println("User email updated successfully")
	return nil
}

func GetUserByUsername(ctx context.Context, dbConfig *dbCFG.DBConfig, username string) (*database.User, error) {

	user, err := dbConfig.DB.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	return &user, nil
}

func GetUserWallet(ctx context.Context, dbConfig *dbCFG.DBConfig, id string) (*[]database.Wallet, error) {

	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	wallets, err := dbConfig.DB.GetUserWallets(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	return &wallets, nil
}
