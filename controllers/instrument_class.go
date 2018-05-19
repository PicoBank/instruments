package controllers

import (
	"github.com/picobank/instruments/models"

	"github.com/astaxie/beego"
)

// InstrumentClassController has the operations about intruments
type InstrumentClassesController struct {
	beego.Controller
}

// @router / [get]
func (controller *InstrumentClassesController) GetAll() {
	controller.Data["json"] = models.GetAllInstrumentClasses()
	controller.ServeJSON()
}
