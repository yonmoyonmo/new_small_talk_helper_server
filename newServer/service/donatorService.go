package service

import (
	"log"

	dataaccess "github.com/yonmoyonmo/new_small_talk_helper_server/data_access"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
	"golang.org/x/crypto/bcrypt"
)

func RegisterNewDonator(newDonator model.Donator) model.DonationMessage {
	if newDonator.DonatorName == "" || newDonator.Password == "" {
		return model.DonationMessage{
			Success: false,
			Message: "no name, no password",
		}
	} else {
		checkExistingDonator := dataaccess.GetDonator(newDonator.DonatorName)
		//check duplication of name
		if checkExistingDonator.Id != 0 {
			return model.DonationMessage{
				Success: false,
				Message: "duplicated name",
			}
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(newDonator.Password), 10)
		if err != nil {
			log.Println(err)
			return model.DonationMessage{
				Success: false,
				Message: "password hash error",
			}
		}
		newDonator.Password = string(hashed)
		var result bool = dataaccess.InsertNewDonator(newDonator.DonatorName, newDonator.Password)
		if result {
			return model.DonationMessage{
				Success: true,
				Message: "success",
			}
		} else {
			return model.DonationMessage{
				Success: false,
				Message: "failed",
			}
		}
	}
}

func IsDonatorIsExist(targetDonator model.Donator) bool {
	if targetDonator.DonatorName == "" || targetDonator.Password == "" {
		return false
	} else {
		checkExistingDonator := dataaccess.GetDonator(targetDonator.DonatorName)
		log.Println("checkExistingDonator")
		log.Println(checkExistingDonator)
		if checkExistingDonator.Id == 0 {
			return false
		}
		if err := bcrypt.CompareHashAndPassword([]byte(checkExistingDonator.Password), []byte(targetDonator.Password)); err != nil {
			log.Printf("incorrect password, %v", err)
			return false
		} else {
			return true
		}
	}
}
