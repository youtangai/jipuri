package controllers

import (
	"myapp/app/models"

	"github.com/revel/revel"
)

type EventController struct {
	BaseController
}

func (c EventController) Index() revel.Result {
	db := ConnectDB()

	var allEvents []models.Event
	db.Find(&allEvents)

	err := models.JsonError{0, ""}
	result := ToJsonFormat("success!", allEvents, err)

	return c.RenderJSON(result)
}

func (c EventController) Create(eventName, date, place, startTime, endTime, description, capacity, departmentName string, isOfficerOnly bool) revel.Result {
	db := ConnectDB()

	event := models.Event{EventName: eventName, Date: date, Place: place, StartTime: startTime, EndTime: endTime, Description: description, Capacity: capacity, DepartmentName: departmentName, IsOfficerOnly: isOfficerOnly}

	db.Create(&event)

	return c.RenderJSON(event)
}

func (c EventController) Read(eventID uint) revel.Result {
	db := ConnectDB()

	var event models.Event

	db.First(&event, eventID)

	err := models.JsonError{0, ""}
	result := ToJsonFormat("success!", event, err)

	return c.RenderJSON(result)
}

func (c EventController) Update(eventID uint, eventName, date, place, startTime, endTime, description, capacity, departmentName string, isOfficerOnly bool) revel.Result {
	db := ConnectDB()

	var event models.Event

	db.First(&event, eventID)

	event.EventName = eventName
	event.Date = date
	event.Place = place
	event.StartTime = startTime
	event.EndTime = endTime
	event.Description = description
	event.Capacity = capacity
	event.DepartmentName = departmentName
	event.IsOfficerOnly = isOfficerOnly

	db.Save(&event)

	err := models.JsonError{0, ""}
	result := ToJsonFormat("success!", event, err)

	return c.RenderJSON(result)
}

func (c EventController) Delete(eventID uint) revel.Result {
	db := ConnectDB()

	var event models.Event

	db.First(&event, eventID)

	db.Delete(&event)

	return c.RenderJSON(event)
}
