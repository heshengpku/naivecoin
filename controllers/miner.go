package controllers

import (
	"naivecoin/models/block"

	"github.com/astaxie/beego"
)

// Operations about Miner
type MinerController struct {
	beego.Controller
}

// @Title Get
// @Description mine new block
// @Param	data	path	string	fasle	"the data you want to write into blockchain"
// @Success 200 {Block} block.Block
// @Failure 403 invalid
// @router /:data [get]
func (o *MinerController) Get() {
	data := o.Ctx.Input.Param(":data")
	ob := block.MineBlock(data)
	o.Data["json"] = ob
	o.ServeJSON()
}
