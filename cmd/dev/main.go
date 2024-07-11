package main

import (
	"thelastking/kingseafood/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("templates/*")
	router.KingRouters(r)
	r.Run(":3250")
}
