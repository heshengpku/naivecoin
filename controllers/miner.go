package controllers

import (
	"naivecoin/models"

	"github.com/astaxie/beego"
)

// Operations about Miner
type MinerController struct {
	beego.Controller
}

// @Title Get
// @Description mine new block
// @Param	data	path	string	true	"the data you want to write into blockchain"
// @Success 200 {Block} models.Block
// @Failure 403 invalid
// @router /:data [get]
func (o *MinerController) Get() {
	data := o.Ctx.Input.Param(":data")
	ob := models.MineBlock(data)
	o.Data["json"] = ob
	o.ServeJSON()
}
