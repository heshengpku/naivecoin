package controllers

import (
	"encoding/json"
	"strconv"

	"naivecoin/models/block"

	"github.com/astaxie/beego"
)

// Operations about Blockchain
type BlockchainController struct {
	beego.Controller
}

// @Title Get
// @Description find Blockchain by index/hash
// @Param	index	path 	string 	true		"the index/hash you want to get"
// @Success 200 {Block} block.Block
// @Failure 403 :index is empty
// @router /:index [get]
func (o *BlockchainController) Get() {
	index := o.Ctx.Input.Param(":index")
	indexBlock, err := strconv.Atoi(index)
	if err == nil {
		ob := block.GetBlockByIndex(indexBlock)
		o.Data["json"] = ob
	} else {
		ob := block.GetBlockByHash(index)
		o.Data["json"] = ob
	}
	o.ServeJSON()
}

// @Title GetLatest
// @Description find the latest block
// @Success 200 {Block} block.Block
// @Failure 403 blockchain is empty
// @router /latest [get]
func (o *BlockchainController) GetLatest() {
	ob := block.GetLatestBlock()
	o.Data["json"] = ob
	o.ServeJSON()
}

// @Title GetAll
// @Description get all blocks
// @Success 200 {BlockChain} block.BlockChain
// @Failure 403 Blockchain is empty
// @router / [get]
func (o *BlockchainController) GetAll() {
	obs := block.GetAllBlocks()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the latest block
// @Param	body		body 	block.Block	true		"The body"
// @Success 200 {Block} block.Block
// @Failure 403 Invalid block
// @router / [put]
func (o *BlockchainController) Put() {
	var b block.Block
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &b)
	if err != nil {
		o.CustomAbort(403, err.Error())
	} else {
		err := block.AddBlock(&b)
		if err == nil {
			o.Data["json"] = "update success!"
		} else {
			o.CustomAbort(403, err.Error())
		}
	}
	o.ServeJSON()
}
