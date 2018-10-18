//---------------------------------------------------
// ini bisa digunakan untuk modif query sesuai dengan
// kebutuhan yang akan digunakan di UserController
// @wahyuhadi 18/10/2018
//---------------------------------------------------

package UserQuery

import (
	"errors"
	"go-iris-mv/config"
)

func GetAllAssociation(model interface{}, association string) error {
	db := config.GetDatabaseConnection()
	defer db.Close()
	err := db.Preload(association).Find(model).Error
	if err != nil {
		return errors.New("Could not find the user")
	}
	return nil
}
