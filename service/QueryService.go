package service

import (
	"errors"
	"go-iris-mv/config"
)

// service ini untuk melakukan sebuah query

// fungsi untul get all user
func GetAll(model interface{}) error {
	db := config.GetDatabaseConnection()
	defer db.Close()
	err := db.Find(model).Error
	if err != nil {
		return errors.New("Could not find the user")
	}
	return nil
}

// fungsi untuk get user by id
func GetById(model interface{}, id string) error {
	db := config.GetDatabaseConnection()
	defer db.Close()
	err := db.Find(model, id).Error
	if err != nil {
		return errors.New("Could not find the user")
	}
	return nil
}

// function get where
func GetWithCondition(model interface{}, condition string) error {
	db := config.GetDatabaseConnection()
	defer db.Close()
	err := db.Where(condition).Find(model).Error
	if err != nil {
		return errors.New("Not Found")
	}
	return nil
}

// Update data
func UpdateData(model interface{}, newData interface{}) error {
	db := config.GetDatabaseConnection()
	defer db.Close()
	err := db.Model(model).Updates(newData).Error
	if err != nil {
		return errors.New("Failed to Update data")
	}
	return nil
}
