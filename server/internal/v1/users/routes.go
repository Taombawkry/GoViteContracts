package users

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	userRouter := http.NewServeMux()

	userRouter.HandleFunc("POST /sendMail", MailHandler)
	userRouter.HandleFunc("POST /sendHTML", HtmlMailHandler)
	userRouter.HandleFunc("/fileForm", FileForm)
	userRouter.HandleFunc("/upload", Upload)
	userRouter.HandleFunc("/listObj", ListObj)
	userRouter.HandleFunc("/getObj", GetObj)
	userRouter.HandleFunc("/createUser", CreateUser)
	userRouter.HandleFunc("GET /getUsers", GetAllUsers)
	userRouter.HandleFunc("/createWallet", CreateWallet)
	userRouter.HandleFunc("POST /getWallets", GetUserWallets)
	userRouter.HandleFunc("POST /createCryptoPaymentOption", CreateUserCryptoPaymentOption)
	userRouter.HandleFunc("POST /getCryptoPaymentOptions", GetUserPaymentOptions)

	return userRouter
}
