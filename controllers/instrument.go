package controllers

import (
	"github.com/picobank/instruments/models"

	"github.com/astaxie/beego"
)

// InstrumentController has the operations about intruments
type InstrumentController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (controller *InstrumentController) GetAll() {
	controller.Data["json"] = models.GetAllInstruments()
	controller.ServeJSON()
}
