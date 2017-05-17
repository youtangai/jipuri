package models

import (
	"github.com/jinzhu/gorm"
)

type Information struct {
	gorm.Model
	Content        string `json:"content"`
	Title          string `json:"title"`
	IsOfficerOnly  bool   `json:"is_officer_onlly"`
	DepartmentName string `json:"department_name"`
}
