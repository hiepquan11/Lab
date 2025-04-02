package main

import (
	"document-management/core/kernel"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	kernel.InitKernel()
	kernel.LoadPlugins(r)
	r.Run(":8080")
}
