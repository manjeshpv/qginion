package subscriptionroutes

import (
	"github.com/manjeshpv/qginion/api/subscription/controller"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	router.GET("/api/subscriptions", subscriptioncontroller.GetAll)
//	router.POST("/api/subscriptions", todocontroller.NewTodo)
//	router.DELETE("/api/subscriptions/:id", todocontroller.RemoveTodo)
}
