package controllers

import (
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/picobank/instruments/models"

	pb "protobuf"

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
		id, err := strconv.ParseInt(instrumentId, 10, 32)
		if err != nil {
			o.Data["json"] = err.Error()
		}
		ob, err := models.GetOneInstrument(uint32(id))
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @router /proto/:instrumentId [get]
func (o *InstrumentController) GetProto() {
	instrumentId := o.Ctx.Input.Param(":instrumentId")
	if instrumentId != "" {
		id, err := strconv.ParseInt(instrumentId, 10, 32)
		if err != nil {
			o.Data["json"] = err.Error()
		}
		ob, err := models.GetOneInstrument(uint32(id))
		if err != nil {
			o.Data["json"] = err.Error()
		}
		instrument := &pb.Instrument{
			Id:          ob.ID,
			Name:        ob.Name,
			Description: ob.Description,
			Class: &pb.InstrumentClass{
				Id:   ob.Class.ID,
				Name: ob.Class.Name,
			},
		}
		data, err := proto.Marshal(instrument)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Ctx.Output.Body(data)
			o.Ctx.Output.ContentType("application/octet-stream")
		}
	}
}
