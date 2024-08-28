package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	config "github.com/NhyiraAmofaSekyi/go-webserver/internal/config"

	database "github.com/NhyiraAmofaSekyi/go-webserver/internal/db/database"
	utils "github.com/NhyiraAmofaSekyi/go-webserver/utils"
	aws "github.com/NhyiraAmofaSekyi/go-webserver/utils/aws/awsS3"
	email "github.com/NhyiraAmofaSekyi/go-webserver/utils/email"
	uuid "github.com/google/uuid"
)

type CreateUserParams struct {
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	SignupType       string `json:"signup_type"`
	SubscriptionType string `json:"subscription_type"`
}

type CreateUserWalletParams struct {
	UserID        string `json:"userID"`
	WalletAddress string `json:"walletAddress"`
	WalletName    string `json:"walletName"`
	WalletIndex   string `json:"walletIndex"`
	ChainID       string `json:"chainID"`
}

func isValidEmail(email string) bool {
	// Simple regex for demonstration purposes; consider a more comprehensive one for production
	return regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`).MatchString(email)
}

// func isValidPassword(password string) bool {
// 	// Example: Minimum eight characters, at least one letter and one number
// 	return regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`).MatchString(password)
// }

func MailHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Subject string `json:"subject"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "server error"})
		return
	}

	err = email.SendMail(params.Subject, params.Email, params.Name)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed to send email"})
		return
	}
	fmt.Fprintln(w, "Mail sent successfully")
}

func HtmlMailHandler(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Subject string `json:"subject"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "error parsing json")
		return
	}

	err = email.SendHTML(params.Subject, params.Email, params.Name)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "error sending email")
		return
	}
	fmt.Fprintln(w, "HTML mail sent successfully")
}

func FileForm(w http.ResponseWriter, r *http.Request) {
	// Define the endpoint where the form will submit the data

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	api := "/api/v1/"
	endpoint := "http://" + host + ":" + port + api + "users/upload"

	// Parse the template file
	tmpl, err := template.ParseFiles("./internal/templates/file_form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Data to pass to the template
	data := struct {
		Endpoint string
	}{
		Endpoint: endpoint,
	}

	// Execute the template with the data
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ListObj(w http.ResponseWriter, r *http.Request) {

	err := aws.ListBucketOBJ()

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "success"})
}

func GetObj(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Key string `json:"key"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "server error"})
		return
	}

	_, err = aws.GetObject(params.Key, "arn:aws:s3:eu-north-1:049991758581:accesspoint/test2")
	bucket := os.Getenv("AWS_BUCKET")
	region := os.Getenv("AWS_BUCKET_REGION")
	url := "https://" + bucket + ".s3." + region + ".amazonaws.com/" + params.Key
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, url)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	dbConfig := config.AppConfig.DBConfig

	var params struct {
		UserID string `json:"userId"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request - invalid JSON"})
		return
	}

	r.ParseMultipartForm(10 << 20) // 10MB

	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf(" retrieving err: %v", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error retrieving the file")
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf(" reading error %v", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error reading the file")
		return
	}
	// openFile := handler.Filename

	fileSize := len(fileBytes)

	fileType := strings.ToLower(filepath.Ext(handler.Filename))
	println("file type", fileType)

	id := uuid.New()

	bucket := os.Getenv("AWS_BUCKET")
	region := os.Getenv("AWS_BUCKET_REGION")

	contentType := mime.TypeByExtension(fileType)
	println("content type", contentType)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	key := id.String() + fileType

	err = aws.Upload(bucket, key, file, contentType)
	if err != nil {
		log.Printf(" upload error %v", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error uploading")
		return
	}

	url := "https://" + bucket + ".s3." + region + ".amazonaws.com/" + key
	response := map[string]interface{}{
		"fileName": handler.Filename,
		"fileType": fileType,
		"fileSize": fileSize,
		"url":      url,
	}
	UserID, err := uuid.Parse(params.UserID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request - invalid event ID"})
		return
	}
	err = dbConfig.DB.UpdateUserProfileImage(r.Context(), database.UpdateUserProfileImageParams{
		ID:              UserID,
		ProfileImageUrl: sql.NullString{String: url, Valid: true},
	})
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "failed to upload image"})
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	dbConfig := config.AppConfig.DBConfig

	decoder := json.NewDecoder(r.Body)
	var params CreateUserParams

	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request"})
		return
	}

	if !isValidEmail(params.Email) {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "invalid email format"})
		return
	}
	params.SignupType = "standard"
	params.SubscriptionType = "free"
	if params.SubscriptionType == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "username and subscription type cannot be empty"})
		return
	}
	if params.Username == "" {
		params.Username = uuid.NewString()
	}

	user, err := dbConfig.DB.CreateUserEmailPassword(r.Context(), database.CreateUserEmailPasswordParams{
		Username:         params.Username,
		Email:            sql.NullString{String: params.Email, Valid: true},
		Password:         sql.NullString{String: params.Password, Valid: params.Password != ""},
		SignupType:       database.NullAccountType{AccountType: database.AccountType(params.SignupType), Valid: params.SignupType != ""},
		SubscriptionType: params.SubscriptionType,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	dbConfig := config.AppConfig.DBConfig

	users, err := dbConfig.DB.GetAllUsers(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	dbConfig := config.AppConfig.DBConfig

	decoder := json.NewDecoder(r.Body)
	var params CreateUserWalletParams
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Bad request"})
		return
	}

	WalletIndex, err := strconv.Atoi(params.WalletIndex)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid wallet index provided"})
		return
	}

	ChainID, err := strconv.Atoi(params.ChainID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid chain ID provided"})
		return
	}

	UserID, err := uuid.Parse(params.UserID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid user ID format"})
		return
	}

	_, err = dbConfig.DB.GetUserByID(r.Context(), UserID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"message": "User not found"})
		return
	}

	if params.WalletAddress == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "no wallet address sent"})
		return
	}

	if params.WalletName == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "no wallet name sent"})
		return
	}

	wallet, err := dbConfig.DB.CreateUserWallet(r.Context(), database.CreateUserWalletParams{
		UserID:        UserID,
		WalletAddress: params.WalletAddress,
		WalletName:    params.WalletName,
		WalletIndex:   int32(WalletIndex),
		ChainID:       int32(ChainID),
	})
	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "failed to create user wallet")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, wallet)
}

func GetUserWallets(w http.ResponseWriter, r *http.Request) {
	dbConfig := config.AppConfig.DBConfig
	decoder := json.NewDecoder(r.Body)
	var params struct {
		UserID string `json:"userID"`
	}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request"})
		return
	}

	if params.UserID == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "userId type cannot be empty"})
		return
	}

	UserID, err := uuid.Parse(params.UserID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request"})
		return
	}

	wallets, err := dbConfig.DB.GetUserWallets(r.Context(), UserID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request"})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, wallets)
}

func CreateUserCryptoPaymentOption(w http.ResponseWriter, r *http.Request) {
	dbConfig := config.AppConfig.DBConfig

	decoder := json.NewDecoder(r.Body)
	var params struct {
		UserID       string `json:"userId"`
		Name         string `json:"name"`
		WalletID     string `json:"walletId"`
		Payment_type string `json:"paymentType"`
	}

	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request - invalid JSON"})
		return
	}

	if params.Name == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request - name is required"})
		return
	}

	UserID, err := uuid.Parse(params.UserID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request - invalid user ID"})
		return
	}

	WalletID, err := strconv.Atoi(params.WalletID)
	if err != nil || WalletID <= 0 {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request - invalid wallet ID"})
		return
	}
	_, err = dbConfig.DB.GetUserByID(r.Context(), UserID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"message": "User not found"})
		return
	}

	paymentOption, err := dbConfig.DB.CreateUserCryptoPaymentOption(r.Context(), database.CreateUserCryptoPaymentOptionParams{
		UserID:      UserID,
		Name:        sql.NullString{String: params.Name, Valid: params.Name != ""},
		WalletID:    sql.NullInt32{Int32: int32(WalletID), Valid: params.WalletID != ""},
		PaymentType: database.PaymentTypeCrypto,
	})
	if err != nil {
		log.Println("Error creating user crypto payment option:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "failed to create user crypto payment option")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, paymentOption)
}

func GetUserPaymentOptions(w http.ResponseWriter, r *http.Request) {
	dbConfig := config.AppConfig.DBConfig

	decoder := json.NewDecoder(r.Body)
	var params struct {
		UserID string `json:"userId"`
	}

	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request - invalid JSON"})
		return
	}

	userID, err := uuid.Parse(params.UserID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "bad request - invalid user ID"})
		return
	}

	_, err = dbConfig.DB.GetUserByID(r.Context(), userID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"message": "User not found"})
		return
	}

	paymentOptions, err := dbConfig.DB.GetUserPaymentOptions(r.Context(), userID)
	if err != nil {
		log.Println("Error retrieving user payment options:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "failed to retrieve user payment options")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, paymentOptions)
}

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	dbConfig := config.AppConfig.DBConfig

	users, err := dbConfig.DB.GetAllUsers(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}
