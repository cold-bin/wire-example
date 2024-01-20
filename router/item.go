package router

import (
	"github.com/gin-gonic/gin"
	"wire-example/handler"
)

type ItemRouter struct {
	Ih *handler.ItemHandler
}

func (ir *ItemRouter) ItemRoute(app *gin.Engine) {
	item := app.Group("/item")
	{
		item.POST("update", ir.Ih.Update)
	}
}
