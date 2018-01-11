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
// @Param	data	body	string	true	"the data you want to write into blockchain"
// @Success 200 {Block} models.Block
// @Failure 403 invalid
// @router [get]
func (o *MinerController) Get() {
	ob := models.MineBlock(string(o.Ctx.Input.RequestBody))
	o.Data["json"] = ob
	o.ServeJSON()
}
