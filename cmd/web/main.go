package main

import (
	"github.com/gin-gonic/gin"
	_ "online.shop.autmaple.com/internal/configs/db"
)

func main() {
  r := gin.Default()
  RegisterRoutes(r)
  r.Run()
}
