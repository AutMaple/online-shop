package main

import (
	"github.com/gin-gonic/gin"
	_ "online.shop.autmaple.com/internal/configs/db"
	"online.shop.autmaple.com/internal/configs/log"
)

func main() {
	r := gin.New()
	r.Use(gin.LoggerWithWriter(log.Logger()))
	r.Use(gin.Recovery())
	r.Static("/static", "./ui/static")
	r.LoadHTMLGlob("ui/templates/*")
	RegisterRoutes(r)
	RegisterValidator()
	r.Run(":8080")
}
