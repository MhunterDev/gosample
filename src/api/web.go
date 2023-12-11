package web

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Profile struct {
	ListenPort      string `json:"listenPort" form:"listenPort" xml:"listenPort"`
	DestinationIp   string `json:"destinationIp" form:"destinationIp" xml:"destinationIp"`
	DestinationPort string `json:"destinationPort" form:"destinatioPort" xml:"destinationPort"`
}

const logfile = "/etc/mhd/gosample/logs/api.log"

var Logfile, LogfileErr = os.Create(logfile)
var logs = log.New(Logfile, "::: API :::", log.Lshortfile)

func addProfile(c *gin.Context) {
	var profile Profile
	err := c.Bind(&profile)
	if err != nil {
		logs.Println(err)
	}
	logs.Printf("%s,%s,%s", profile.ListenPort, profile.DestinationIp, profile.DestinationPort)
}

func handleMain(c *gin.Context) {
	c.HTML(200, "main.html", nil)
}

func Router() {

	router := gin.Default()
	router.LoadHTMLFiles("/etc/mhd/gosample/.public/html/main.html")
	router.GET("/", handleMain)
	router.POST("/", addProfile)

	router.Run("localhost:8080")
}
