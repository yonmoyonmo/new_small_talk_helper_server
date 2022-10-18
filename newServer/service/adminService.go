package service

import (
	"log"

	dataaccess "github.com/yonmoyonmo/new_small_talk_helper_server/data_access"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
	"golang.org/x/crypto/bcrypt"
)

func RegisterAdmin(newAdmin model.Admin) bool {
	if newAdmin.AdminName == "" || newAdmin.Password == "" {
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
