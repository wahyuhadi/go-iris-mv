package service

import (
	"errors"
	"go-iris-mv/config"
)

// ini untuk proses query
// ini get all harusnya return slice of interface, right?
// func GetAll(obj interface{}) (interface{}, error) {
// 	db := config.GetDatabaseConnection()
// 	defer db.Close()
// 	db.Find(obj.(model.User))
// 	return obj, fmt.Errorf("exception throws")
// }
func GetAll(obj interface{}) error {
	db := config.GetDatabaseConnection()
	defer db.Close()
	if err := db.Debug().Find(obj).Error; err != nil {
		return errors.New("Could not find the user")
	}
	return nil
}
