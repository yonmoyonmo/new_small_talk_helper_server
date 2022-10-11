package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
	"github.com/yonmoyonmo/new_small_talk_helper_server/service"
)

func RegisterNewDonatorHandler(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var newDonator model.Donator = model.Donator{}
		json.NewDecoder(req.Body).Decode(&newDonator)
		log.Printf("new donator : %v", newDonator.DonatorName)
		result := service.RegisterNewDonator(newDonator)
		if result.Success == true {
			resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(resWriter).Encode(result)
		} else {
			resWriter.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func CheckDonatorHandler(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var targetDonator model.Donator = model.Donator{}
		json.NewDecoder(req.Body).Decode(&targetDonator)
		log.Printf("check donator : %v", targetDonator.DonatorName)
		result := service.IsDonatorIsExist(targetDonator)
		if result == true {
			resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(resWriter).Encode(result)
		} else {
			resWriter.WriteHeader(http.StatusInternalServerError)
		}

	}
}
