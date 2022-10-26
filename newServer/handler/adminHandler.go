package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
	"github.com/yonmoyonmo/new_small_talk_helper_server/service"
)

func generateJWT(secretKey []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("error in token gen")
		log.Println(secretKey)
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}

func AdminLoginHandler(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var inputAdmin model.Admin
		err := json.NewDecoder(req.Body).Decode(&inputAdmin)
		if err != nil {
			log.Println(err)
			return
		}
		loginResult := service.AdminLoginService(inputAdmin)
		log.Println(loginResult)
		if loginResult == false {
			resWriter.WriteHeader(http.StatusBadRequest)
		} else {
			resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
			var tokenResponse model.AdminToken = model.AdminToken{}

			//jwt
			secret := []byte(os.Getenv("JWTKEY"))
			tokenString, err := generateJWT(secret)
			if err != nil {
				tokenResponse.Message = "failed"
				tokenResponse.Token = ""
				json.NewEncoder(resWriter).Encode(tokenResponse)
				return
			}

			tokenResponse.Message = "success"
			tokenResponse.Token = tokenString
			json.NewEncoder(resWriter).Encode(tokenResponse)
		}
	}
}

func AdminRegisterHandler(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var newAdmin model.Admin
		json.NewDecoder(req.Body).Decode(&newAdmin)
		log.Printf("new Admin : %v", newAdmin.AdminName)

		wonmoKey := req.URL.Query().Get("key")

		if wonmoKey != os.Getenv("WONMOKEY") {
			resWriter.WriteHeader(http.StatusBadRequest)
		}

		result := service.RegisterAdmin(newAdmin)
		if result == true {
			resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(resWriter).Encode(result)
		} else {
			resWriter.WriteHeader(http.StatusInternalServerError)
		}
	}
}
