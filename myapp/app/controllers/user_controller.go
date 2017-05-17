package controllers

import (
	"myapp/app/models"

	"github.com/revel/revel"
)

type UserController struct {
	BaseController
}

func (c UserController) Index() revel.Result {
	db := ConnectDB()

	var allUsers []models.User
	db.Find(&allUsers)

	err := models.JsonError{0, ""}
	result := ToJsonFormat("success!", allUsers, err)

	return c.RenderJSON(result)
}

func (c UserController) Create(userName, address, tell, sex string, age int, eventID uint) revel.Result {
	db := ConnectDB()

	user := models.User{UserName: userName, Age: age, Address: address, Tell: tell, Sex: sex, EventID: eventID}

	db.Create(&user)

	return c.RenderJSON(user)
}

func (c UserController) Read(userID uint) revel.Result {
	db := ConnectDB()

	var user models.User

	db.First(&user, userID)

	err := models.JsonError{0, ""}
	result := ToJsonFormat("success!", user, err)

	return c.RenderJSON(result)
}

func (c UserController) Update(userID uint, userName, address, tell, sex string, age int, eventID uint) revel.Result {
	db := ConnectDB()

	var user models.User

	db.First(&user, userID)

	user.UserName = userName
	user.Age = age
	user.Address = address
	user.Tell = tell
	user.Sex = sex
	user.EventID = eventID

	db.Save(&user)

	err := models.JsonError{0, ""}
	result := ToJsonFormat("success!", user, err)

	return c.RenderJSON(result)
}

func (c UserController) Delete(userID uint) revel.Result {
	db := ConnectDB()

	var user models.User

	db.First(&user, userID)

	db.Delete(&user)

	return c.RenderJSON(user)
}
