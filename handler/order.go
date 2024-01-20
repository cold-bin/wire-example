package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wire-example/entity"
	"wire-example/service"
)

type ItemHandler struct {
	Itemsvc *service.ItemSvc
}

func (oh *ItemHandler) Update(c *gin.Context) {
	item := &entity.Item{}
	if err := oh.Itemsvc.Update(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "server is bad",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
