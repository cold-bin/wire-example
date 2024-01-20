package main

import (
	"github.com/gin-gonic/gin"
	"wire-example/cmd/di"
)

func main() {
	engine := gin.Default()

	di.BuildItemRouter().ItemRoute(engine)

	engine.Run()
}
