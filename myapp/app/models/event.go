package models

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	EventName      string `json:"event_name"`
	Date           string `json:"date"`
	Place          string `json:"place"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	Description    string `json:"description"`
	Capacity       string `json:"capacity"`
	IsOfficerOnly  bool   `json:"is_officer_onlly"`
	DepartmentName string `json:"department_name"`
}
