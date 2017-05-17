package controllers

import (
	"myapp/app/models"

	"github.com/revel/revel"
)

type InfoController struct {
	BaseController
}

func (c InfoController) Index() revel.Result {
	db := ConnectDB()

	var allInfos []models.Information
	db.Find(&allInfos)

	err := models.JsonError{0, ""}
	result := ToJsonFormat("success!", allInfos, err)

	return c.RenderJSON(result)
}

func (c InfoController) Create(content, title, departmentName string, isOfficerOnly bool) revel.Result {
	db := ConnectDB()

	info := models.Information{Content: content, Title: title, DepartmentName: departmentName, IsOfficerOnly: isOfficerOnly}

	db.Create(&info)

	return c.RenderJSON(info)
}

func (c InfoController) Read(infoID uint) revel.Result {
	db := ConnectDB()

	var info models.Information

	db.First(&info, infoID)

	err := models.JsonError{0, ""}
	result := ToJsonFormat("success!", info, err)

	return c.RenderJSON(result)
}

func (c InfoController) Update(infoID uint, content, title, departmentName string, isOfficerOnly bool) revel.Result {
	db := ConnectDB()

	var info models.Information

	db.First(&info, infoID)

	info.Content = content
	info.Title = title
	info.DepartmentName = departmentName
	info.IsOfficerOnly = isOfficerOnly

	db.Save(&info)

	err := models.JsonError{0, ""}
	result := ToJsonFormat("success!", info, err)

	return c.RenderJSON(result)
}

func (c InfoController) Delete(infoID uint) revel.Result {
	db := ConnectDB()

	var info models.Information

	db.First(&info, infoID)

	db.Delete(&info)

	return c.RenderJSON(info)
}
