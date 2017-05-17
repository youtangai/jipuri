package controllers

import (
	"fmt"
	"myapp/app/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type BaseController struct {
	*revel.Controller
}

func ToJsonFormat(success string, data interface{}, err models.JsonError) models.JsonResponse {
	return models.JsonResponse{success, data, err}
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", GetConnectionString())

	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}
	return db
}

func GetParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func GetConnectionString() string {
	host := GetParamString("db.host", "")
	port := GetParamString("db.port", "3306")
	user := GetParamString("db.user", "")
	pass := GetParamString("db.password", "")
	dbname := GetParamString("db.name", "")
	protocol := GetParamString("db.protocol", "tcp")
	dbargs := GetParamString("dbargs", " ")
	timezone := GetParamString("db.timezone", "parseTime=true&loc=Asia%2FTokyo")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s?%s", user, pass, protocol, host, port, dbname, dbargs, timezone)
}
