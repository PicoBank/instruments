package controllers

import (
	"strconv"

	"github.com/picobank/instruments/models"

	"github.com/astaxie/beego"
)

// InstitutionController has the operations about intruments
type InstitutionController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (controller *InstitutionController) GetAll() {
	controller.Data["json"] = models.GetAllInstitutions()
	controller.ServeJSON()
}

// @router /:institutionId [get]
func (o *InstitutionController) Get() {
	institutionId := o.Ctx.Input.Param(":institutionId")
	if institutionId != "" {
		id, err := strconv.ParseInt(institutionId, 10, 32)
		if err != nil {
			o.Data["json"] = err.Error()
		}
		ob, err := models.GetOneInstitution(uint32(id))
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}
