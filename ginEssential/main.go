package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//db := common.GetDB()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
