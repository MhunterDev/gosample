package web

import (
	"github.com/gin-gonic/gin"
)

type Profile struct {
	Listen-port `json:"listen-port" form:"listen-port" xml:"listen-port" binding:"required`
	Destination-ip `json:"destination-ip" form:"destination-ip" xml:"destination-ip" binding:"required"`
	Destination-port `json:"destination-port" form:"destination-port" xml:"destination-port" binding:"required"`
}

func addProfile(c *gin.Context){
	
}

func Router() {

	router := gin.Default()

	router.POST("/", addProfile)

	router.Run(":8080")
}
