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
			log.Println("decoding req.body failed")
			log.Println(err)
			resWriter.WriteHeader(http.StatusInternalServerError)
			return
		}

		loginResult := service.AdminLoginService(inputAdmin)
		log.Println(loginResult)
		if loginResult == false {
			resWriter.WriteHeader(http.StatusBadRequest)
			return
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
		err := json.NewDecoder(req.Body).Decode(&newAdmin)
		if err != nil {
			log.Println("decoding req.body failed")
			log.Println(err)
			resWriter.WriteHeader(http.StatusInternalServerError)
			return
		}
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

func verifyJWT(req *http.Request) bool {
	secret := []byte(os.Getenv("JWTKEY"))
	token, err := jwt.Parse(req.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return false
	}
	if token.Valid {
		return true
	} else {
		return false
	}
}

func RegisterSugguestionsHandler(resWriter http.ResponseWriter, req *http.Request) {
	//verify token in req header.authentication
	//if verified -> create new sugguestions
	switch req.Method {
	case http.MethodPost:
		if req.Header["Authorization"] != nil {
			if verifyJWT(req) {
				var newSuggestion model.Sugguestion
				err := json.NewDecoder(req.Body).Decode(&newSuggestion)
				if err != nil {
					log.Println("decoding req.body failed")
					log.Println(err)
					resWriter.WriteHeader(http.StatusInternalServerError)
					return
				}
				//create sugguestions
				var result model.SimpleResponseType = service.CreateNewSugguestions(newSuggestion)

				resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
				json.NewEncoder(resWriter).Encode(result)

			} else {
				log.Println("invalid token")
				resWriter.WriteHeader(http.StatusBadRequest)
			}
		} else {
			log.Println("no token in header")
			resWriter.WriteHeader(http.StatusBadRequest)
		}
	}
}
