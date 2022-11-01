package service

import (
	"log"

	dataaccess "github.com/yonmoyonmo/new_small_talk_helper_server/data_access"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
	"golang.org/x/crypto/bcrypt"
)

func AdminLoginService(inputAdmin model.Admin) bool {
	if inputAdmin.AdminName == "" || inputAdmin.Password == "" {
		log.Println("no name no gain")
		return false
	}

	existingAdmin := dataaccess.GetAdmin(inputAdmin.AdminName)
	if existingAdmin.Id == 0 {
		log.Println("this admin not exists")
		return false
	} else {
		log.Println(existingAdmin.Password)
		hashedPassword := existingAdmin.Password
		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputAdmin.Password))
		if err != nil {
			log.Println("password not matched")
			return false
		} else {
			return true
		}
	}

}

func RegisterAdmin(newAdmin model.Admin) bool {
	if newAdmin.AdminName == "" || newAdmin.Password == "" {
		log.Println("no name no gain")
		return false
	}
	isAdminIsValid := dataaccess.GetAdmin(newAdmin.AdminName)
	if isAdminIsValid.Id != 0 {
		log.Panicln("Existing admin name")
		return false
	} else {
		hashed, err := bcrypt.GenerateFromPassword([]byte(newAdmin.Password), 10)
		if err != nil {
			log.Println(err)
			return false
		}
		newAdmin.Password = string(hashed)
		var result bool = dataaccess.InsertNewAdmin(newAdmin.AdminName, newAdmin.Password)
		return result
	}
}

func CreateNewSugguestions(sugguestion model.Sugguestion) model.SimpleResponseType {
	log.Println("in create new sugguestion service")
	log.Println(sugguestion)

	var newSimpleResponse model.SimpleResponseType
	newSimpleResponse.Message = "good"
	newSimpleResponse.Success = true

	sqlResult := dataaccess.InsertNewSugguestion(sugguestion)
	if !sqlResult {
		newSimpleResponse.Message = "failed"
		newSimpleResponse.Success = false
	}

	return newSimpleResponse
}

func UpdateSugguestion(targetSugg model.Sugguestion) model.SimpleResponseType {
	log.Println("update existing sugguestion")

	var newSimpleResponse model.SimpleResponseType
	newSimpleResponse.Message = "good"
	newSimpleResponse.Success = true

	sqlResult := dataaccess.UpdateSugguestion(targetSugg)
	if !sqlResult {
		newSimpleResponse.Message = "failed"
		newSimpleResponse.Success = false
	}

	return newSimpleResponse
}

func DeleteSugguestion(targetSugg model.Sugguestion) model.SimpleResponseType {
	log.Println("delete existing sugguestion")

	var newSimpleResponse model.SimpleResponseType
	newSimpleResponse.Message = "good"
	newSimpleResponse.Success = true

	sqlResult := dataaccess.DeleteSugguestion(targetSugg)
	if !sqlResult {
		newSimpleResponse.Message = "failed"
		newSimpleResponse.Success = false
	}

	return newSimpleResponse
}
