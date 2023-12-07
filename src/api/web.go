package web

import (
	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()

	router.Run(":8080")
}
