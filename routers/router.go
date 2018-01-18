// @APIVersion 1.0.0
// @Title naivecoin Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact heshengpku@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"naivecoin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/blockchain",
			beego.NSInclude(
				&controllers.BlockchainController{},
			),
		), beego.NSNamespace("/miner",
			beego.NSInclude(
				&controllers.MinerController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
