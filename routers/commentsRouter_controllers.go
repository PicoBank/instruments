package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:InstrumentClassesController"] = append(beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:InstrumentClassesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:InstrumentController"] = append(beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:InstrumentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:InstrumentController"] = append(beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:InstrumentController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:instrumentId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/picobank/instruments/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
