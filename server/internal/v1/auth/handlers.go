package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	config "github.com/NhyiraAmofaSekyi/go-webserver/internal/config"

	"github.com/NhyiraAmofaSekyi/go-webserver/internal/db/database"
	"github.com/NhyiraAmofaSekyi/go-webserver/internal/middleware"
	utils "github.com/NhyiraAmofaSekyi/go-webserver/utils"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var hmacSampleSecret = []byte("sample")

func Check(w http.ResponseWriter, r *http.Request) {

	// Respond with a success message
	utils.RespondWithJSON(w, http.StatusCreated, config.AppConfig.ClientURL)

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	dbConfig := config.AppConfig.DBConfig
	var creds struct {
		Username string `json:"username,omitempty"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if creds.Username != "" {
		_, err = dbConfig.DB.GetUserByUsername(r.Context(), creds.Username)
		if err == nil {
			utils.RespondWithJSON(w, http.StatusConflict, map[string]string{"message": "username already exists"})
			return
		} else if !errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to check username")
			return
		}

	} else if creds.Username == "" {
		creds.Username = uuid.New().String()
	}

	if creds.Email != "" {
		_, err = dbConfig.DB.GetUserEmail(r.Context(), sql.NullString{String: creds.Email, Valid: true})
		if err == nil {
			utils.RespondWithJSON(w, http.StatusConflict, map[string]string{"message": "email already exists"})
			return
		} else if !errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(w, http.StatusInternalServerError, "email to check username")
			return
		}
	}

	hashedPassword, err := HashPassword(creds.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Store the new user in the database
	newUser := database.CreateUserEmailPasswordParams{
		Username: creds.Username,
		Email:    sql.NullString{String: creds.Email, Valid: creds.Email != ""},
		Password: sql.NullString{String: hashedPassword, Valid: creds.Password != ""},
	}

	user, err := dbConfig.DB.CreateUserEmailPassword(r.Context(), newUser)
	if err != nil {
		log.Printf("error : %v", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	var TokenParams struct {
		UserID string `json:"userId"`
		Action string `json:"action"`
	}
	TokenParams.UserID = user.ID.String()
	TokenParams.Action = "verify_account"
	token, err := CreateEmailAction(dbConfig, TokenParams)
	if err != nil {
		log.Printf("error : %v", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed create emal action")
		return
	}

	url := config.AppConfig.ClientURL + "/verification/email/" + token

	err = SendTokenEmail(creds.Email, url)
	if err != nil {
		log.Printf("error : %v", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to send email")
		return
	}

	// Respond with a success message
	utils.RespondWithJSON(w, http.StatusCreated, user.ID)
}
func CheckUsername(w http.ResponseWriter, r *http.Request) {
	dbConfig := config.AppConfig.DBConfig

	Username := r.PathValue("username")

	if Username == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, "please provie username")
		return
	}

	// Check if the username exists in the database
	_, err := dbConfig.DB.GetUserByUsername(r.Context(), Username)
	if err == nil {
		// Username exists
		utils.RespondWithJSON(w, http.StatusConflict, map[string]string{"message": "username already exists"})
		return
	} else if !errors.Is(err, sql.ErrNoRows) {
		// An error occurred while checking the username
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to check username")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "username is available"})
}
func FarcasterSignUp(w http.ResponseWriter, r *http.Request) {
	dbConfig := config.AppConfig.DBConfig
	var user FarcasterUserData
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	sessionToken, err := CreateFarcasterSocialLogin(r.Context(), user, *dbConfig)
	if err != nil {
		log.Printf("error: %v", err)
		utils.RespondWithError(w, http.StatusBadRequest, "failes to create user and session")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token", // Name of the cookie
		Value:    sessionToken,    // Value of the cookie, which is the session token
		Path:     "/",             // Cookie path
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false, // Secure flag, should be set to true if using HTTPS
		HttpOnly: true,  // HttpOnly flag, browser is not allowed to access the cookie via JavaScript
	})

	// Respond with a success message
	utils.RespondWithJSON(w, http.StatusCreated, user)

}
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	state := generateSecureToken()
	setSessionCookie(w, "session_state", state)
	log.Printf("set state %v", state)

	opts := []oauth2.AuthCodeOption{oauth2.AccessTypeOffline, oauth2.ApprovalForce}
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL(state, opts...)
	log.Printf("url %v", url)
	http.Redirect(w, r, url, http.StatusFound)
}
func Callback2(w http.ResponseWriter, r *http.Request) {

	deleteCookie := http.Cookie{
		Name:     "session_state",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &deleteCookie)

	dbConfig := config.AppConfig.DBConfig

	errorParam := r.URL.Query().Get("error")
	if errorParam != "" {
		log.Printf("OAuth error: %s", errorParam)
		http.Redirect(w, r, config.AppConfig.ClientURL+"?error="+url.QueryEscape(errorParam), http.StatusFound)
		return
	}
	// Retrieve state from cookie
	opts := []oauth2.AuthCodeOption{oauth2.AccessTypeOffline}

	state, err := getSessionFromCookie(r, "session_state")
	if err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	log.Printf("recieved state: %s ", state)
	if r.URL.Query().Get("state") != state {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	provider := "google"
	log.Printf("provider: %s ", provider)
	code := r.URL.Query().Get("code")
	if code == "" {
		log.Println("No code received")
		http.Error(w, "Authorization code not found", http.StatusBadRequest)
		return
	}

	log.Printf("code: %s ", code)

	googlecon := config.AppConfig.GoogleLoginConfig

	token, err := googlecon.Exchange(context.Background(), code, opts...)
	log.Printf("access token: %v ", token.AccessToken)
	log.Printf("refresh token: %v ", token.RefreshToken)

	if err != nil {
		log.Printf("auth error %v: ", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "invalid token"})
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Printf("auth error %v: ", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "invalid request"})
		return
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("auth error %v: ", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "invalid request"})
		return
	}

	existingUserID, err := CheckIfUserExists(dbConfig, userData)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "user exists"})
		return
	}

	var sessionToken string

	if existingUserID == nil {
		sessionToken, err = CreateSocialLogin(r.Context(), userData, *dbConfig, token)
		if err != nil {
			log.Printf("social login error %v: ", err)
			utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "invalid request"})
			return
		}
	} else {
		err = CheckUserSessions(r.Context(), dbConfig, existingUserID)
		if err != nil {
			log.Printf("sesssion error  %v: ", err)
			utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "invalid request"})
			return
		}

		session, err := dbConfig.DB.CreateUserSession(r.Context(), database.CreateUserSessionParams{
			UserID:    *existingUserID, // Use existing user's ID
			ExpiresAt: time.Now().Add(24 * time.Hour),
			IpAddress: sql.NullString{String: "user's IP address", Valid: true},
			UserAgent: sql.NullString{String: "user's User-Agent header", Valid: true},
		})
		if err != nil {
			log.Printf("Error creating session for existing user: %v", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Failed to create session"})
			return
		}

		// Generate JWT token for the session
		sessionToken, err = generateJWT(session.ID.String())
		if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Failed to generate token"})
			return
		}

	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token", // Name of the cookie
		Value:    sessionToken,    // Value of the cookie, which is the session token
		Path:     "/",             // Cookie path
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false, // Secure flag, should be set to true if using HTTPS
		HttpOnly: true,  // HttpOnly flag, browser is not allowed to access the cookie via JavaScript
		//SameSite: http.SameSiteStrictMode,        // SameSite attribute, use Strict mode for CSRF protection
	})

	log.Println(string(userData))
	http.Redirect(w, r, config.AppConfig.ClientURL, http.StatusFound)

}

func FarcasterSignIn(w http.ResponseWriter, r *http.Request) {

	var user FarcasterUserData
	dbConfig := config.AppConfig.DBConfig
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		return
	}

	sessionToken, err := CreateFarcasterSocialLogin(r.Context(), user, *dbConfig)
	if err != nil {
		log.Printf("error: %v", err)
		utils.RespondWithError(w, http.StatusBadRequest, "failes to create user and session")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token", // Name of the cookie
		Value:    sessionToken,    // Value of the cookie, which is the session token
		Path:     "/",             // Cookie path
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false, // Secure flag, should be set to true if using HTTPS
		HttpOnly: true,  // HttpOnly flag, browser is not allowed to access the cookie via JavaScript
	})

	// Respond with a success message
	utils.RespondWithJSON(w, http.StatusCreated, user)

}
func SignIn(w http.ResponseWriter, r *http.Request) {

	// Decode request body to extract login credentials
	var creds struct {
		Username string `json:"username,omitempty"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password"`
	}
	dbConfig := config.AppConfig.DBConfig
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := dbConfig.DB.GetUserEmail(r.Context(), sql.NullString{String: creds.Email, Valid: true})
	if err != nil {
		utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": "no user found"})
		return
	}

	err = CheckPassword(user.Password.String, creds.Password)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": "wrong password"})
		return
	}

	session, err := dbConfig.DB.GetSessionByUserID(r.Context(), user.ID)
	if err == nil && session.ExpiresAt.After(time.Now()) {
		// Valid session exists, redirect to main page
		http.Redirect(w, r, config.AppConfig.ClientURL, http.StatusFound)
		return
	}

	newSession, err := dbConfig.DB.CreateUserSession(r.Context(), database.CreateUserSessionParams{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		IpAddress: sql.NullString{String: "user's IP address", Valid: true},
		UserAgent: sql.NullString{String: "user's User-Agent header", Valid: true},
	})
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	// Generate JWT for the session
	sessionToken, err := generateJWT(newSession.ID.String())
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Set the session token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false, // Set to true in production when using HTTPS
		HttpOnly: true,  // Prevent JavaScript access to cookie
	})

	// Redirect to the main page after successful login
	http.Redirect(w, r, "http://localhost:8080/api/v1/auth/dashboard", http.StatusFound)

}

func SignOut(w http.ResponseWriter, r *http.Request) {

	// Retrieve the session token from the cookie
	cookie, err := r.Cookie("session_token")
	dbConfig := config.AppConfig.DBConfig
	if err != nil {
		// No cookie found, redirect to the homepage
		http.Redirect(w, r, config.AppConfig.ClientURL, http.StatusFound)
		return
	}

	// Decode the session token to extract the session ID
	claims, err := middleware.ParseJWT(cookie.Value)
	if err != nil {
		// If the token is invalid, just redirect to the home page after clearing the cookie
		middleware.ClearSessionCookie(w)
		http.Redirect(w, r, config.AppConfig.ClientURL, http.StatusFound)
		return
	}

	sessionIDstr, ok := claims["id"].(string)
	if !ok {
		// If the session ID is not found, redirect to the homepage
		middleware.ClearSessionCookie(w)
		http.Redirect(w, r, config.AppConfig.ClientURL, http.StatusFound)
		return
	}

	sessionID, err := uuid.Parse(sessionIDstr)
	if err != nil {
		// If the session ID is not a valid UUID, redirect to the homepage
		middleware.ClearSessionCookie(w)
		http.Redirect(w, r, config.AppConfig.ClientURL, http.StatusFound)
		return
	}

	// Delete the session from the database
	err = dbConfig.DB.DeleteSessionByID(r.Context(), sessionID)
	if err != nil {
		// Log error but still clear the cookie and redirect
		log.Printf("Error deleting session: %v", err)
	}

	// Clear the session cookie
	middleware.ClearSessionCookie(w)

	// Redirect to the homepage
	http.Redirect(w, r, config.AppConfig.ClientURL, http.StatusFound)

}

func Dashboard(w http.ResponseWriter, r *http.Request) {

	dbConfig := config.AppConfig.DBConfig

	// Extract session and user ID from the request context
	sessionID, ok := r.Context().Value(middleware.AuthSessionID).(uuid.UUID)
	if !ok {
		http.Redirect(w, r, config.AppConfig.ClientURL, http.StatusNotFound)

		return
	}

	userID, ok := r.Context().Value(middleware.AuthUserID).(uuid.UUID)
	if !ok {
		http.Redirect(w, r, config.AppConfig.ClientURL, http.StatusNotFound)
		return
	}

	// Fetch the session and user data from the database
	session, err := dbConfig.DB.GetSessionByID(r.Context(), sessionID)
	if err != nil {
		http.Redirect(w, r, config.AppConfig.ClientURL, http.StatusNotFound)
		return
	}

	user, err := dbConfig.DB.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Redirect(w, r, config.AppConfig.ClientURL, http.StatusNotFound)
		return
	}

	// Prepare the response data combining user and session info
	responseData := struct {
		User    database.User
		Session database.Session
	}{
		User:    user,
		Session: session,
	}

	// Marshal data to JSON and write response
	utils.RespondWithJSON(w, 200, responseData)

}

func EmailVerification(w http.ResponseWriter, r *http.Request) {

	dbConfig := config.AppConfig.DBConfig
	var params struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		return
	}
	err := UpdateUserEmailWithToken(r.Context(), dbConfig, params.Token)
	if err != nil {
		log.Printf("cant update: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "email updated"})

}

func EmailAction(w http.ResponseWriter, r *http.Request) {

	// Decode the incoming JSON to extract needed data
	var params struct {
		UserID string `json:"userId"`
		Action string `json:"action"`
	}
	dbConfig := config.AppConfig.DBConfig

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		return
	}

	token, err := CreateEmailAction(dbConfig, params)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Failed to create email action"})
		return
	}

	email, err := GetUserEmailByID(r.Context(), dbConfig, params.UserID)
	if err != nil {
		log.Printf("cant find user id: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "cant find user email"})
		return
	}

	err = SendTokenEmail(email, token)
	if err != nil {
		log.Printf("cant send email: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, token)

}
