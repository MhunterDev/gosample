package web

import (
	"log"
	"net/http"
	"os"

	"github.com/MhunterDev/gosample/src/db"
	"github.com/gin-gonic/gin"
)

type Profile struct {
	ListenPort      string `json:"listenPort" form:"listenPort" xml:"listenPort"`
	DestinationIp   string `json:"destinationIp" form:"destinationIp" xml:"destinationIp"`
	DestinationPort string `json:"destinationPort" form:"destinatioPort" xml:"destinationPort"`
}

const logfile = "/etc/mhd/gosample/logs/replicate.log"

var Logfile, LogfileErr = os.Create(logfile)
var logs = log.New(Logfile, "::: API :::", log.Lshortfile)

func addProfile(c *gin.Context) {
	var p Profile
	c.Bind(&p)

	var response struct {
		message string
	}

	err := db.AddProfile(p.ListenPort, p.DestinationIp, p.DestinationPort)
	if err != nil {
		response.message = "Error adding profile"
		c.JSON(http.StatusBadGateway, response)
		return
	}
	response.message = "Profile added"
	c.JSON(200, response)
}

func handleMain(c *gin.Context) {
	var response struct {
		message string
	}
	response.message = "Api service is running on port 5000"
	c.JSON(202, nil)
}

func Router() {

	router := gin.Default()
	router.Handle("GET", "/api", handleMain)
	router.Handle("POST", "/api/add/profiles", addProfile)

	router.Run("localhost:5000")
}
