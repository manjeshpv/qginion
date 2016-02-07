package routes

import (
	"github.com/manjeshpv/qginion/api/subscription/routes"
//	"github.com/manjeshpv/qginion/common/static"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	subscriptionroutes.Init(router)
//	static.Init(router)
}
