package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
	"github.com/yonmoyonmo/new_small_talk_helper_server/service"
)

func AdminRegisterHandler(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var newAdmin model.Admin = model.Admin{}
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
