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

// @router /:instrumentId [get]
func (o *InstrumentController) Get() {
	instrumentId := o.Ctx.Input.Param(":instrumentId")
	if instrumentId != "" {
		ob, err := models.GetOneInstrument(instrumentId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}
