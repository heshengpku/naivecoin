package controllers

import (
	"encoding/json"
	"naivecoin/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about Blockchain
type BlockchainController struct {
	beego.Controller
}

// @Title Get
// @Description find Blockchain by index/hash
// @Param	index	path 	string 	true		"the index/hash you want to get"
// @Success 200 {Block} models.Block
// @Failure 403 :index is empty
// @router /:index [get]
func (o *BlockchainController) Get() {
	index := o.Ctx.Input.Param(":index")
	indexBlock, err := strconv.Atoi(index)
	if err == nil {
		ob := models.GetBlockByIndex(indexBlock)
		o.Data["json"] = ob
	} else {
		ob := models.GetBlockByHash(index)
		o.Data["json"] = ob
	}
	o.ServeJSON()
}

// @Title GetLatest
// @Description find the latest block
// @Success 200 {Block} models.Block
// @Failure 403 blockchain is empty
// @router /latest [get]
func (o *BlockchainController) GetLatest() {
	ob := models.GetLatestBlock()
	o.Data["json"] = ob
	o.ServeJSON()
}

// @Title GetAll
// @Description get all blocks
// @Success 200 {BlockChain} models.BlockChain
// @Failure 403 Blockchain is empty
// @router / [get]
func (o *BlockchainController) GetAll() {
	obs := models.GetAllBlocks()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the latest block
// @Param	body		body 	models.Block	true		"The body"
// @Success 200 {Block} models.Block
// @Failure 403 Invalid block
// @router / [put]
func (o *BlockchainController) Put() {
	var block models.Block
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &block)
	if err != nil {
		o.CustomAbort(403, err.Error())
	} else {
		if models.AddBlock(&block) {
			o.Data["json"] = "update success!"
		} else {
			o.CustomAbort(403, "Invalid block")
		}
	}
	o.ServeJSON()
}
