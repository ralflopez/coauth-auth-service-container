package controllers

import (
	"coauth/pkg/exceptions"
	"coauth/pkg/models"
	"coauth/pkg/services"
	"coauth/pkg/utils"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	exceptions.ThrowNotFoundException(w, "Not implemented")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := services.GetUsers()
	utils.StructToJSON(users, w)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Validation
	var user models.User
	utils.JSONToStuct(&user, r.Body)

	err := utils.ValidateStruct(user)
	if err != nil {
		exceptions.ThrowBadRequestException(w, err.Error())
		return
	}

	// Persistence
	services.CreateUser(&user)

	// Return
	// omit password on return
	userReturn := user
	userReturn.Password = ""
	utils.StructToJSON(userReturn, w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	exceptions.ThrowNotFoundException(w, "Not implemented")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	exceptions.ThrowNotFoundException(w, "Not implemented")
}