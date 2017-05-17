package controllers

import (
	"myapp/app/models"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) MigrateDB() revel.Result {
	db := ConnectDB()
	db.AutoMigrate(models.Event{})
	db.AutoMigrate(models.Information{})
	db.AutoMigrate(models.User{})
	db.Model(models.User{}).AddForeignKey("event_id", "events(id)", "RESTRICT", "RESTRICT")
	return c.RenderJSON("migrate complete")
}
