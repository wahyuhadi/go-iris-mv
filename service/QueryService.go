package service

import (
	"fmt"
	"github.com/mochadwi/go-iris-mv/config"
	"github.com/mochadwi/go-iris-mv/model"
)

// ini untuk proses query
// ini get all harusnya return slice of interface, right?
func GetAll(obj interface{}) (interface{}, error) {
	db := config.GetDatabaseConnection()
	defer db.Close()
	db.Find(obj.(model.User))
	return obj, fmt.Errorf("exception throws")
}
