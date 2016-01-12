package routers

import (
	"github.com/ob3/dapetaje/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
