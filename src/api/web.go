package web

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Profile struct {
	ListenPort      string `json:"listenPort" form:"listenPort" xml:"listenPort"`
	DestinationIp   string `json:"destinationIp" form:"destinationIp" xml:"destinationIp"`
	DestinationPort string `json:"destinationPort" form:"destinatioPort" xml:"destinationPort"`
}

const logfile = "/etc/mhd/gosample/logs/api.log"

var Logfile, LogfileErr = os.Create(logfile)
var logs = log.New(Logfile, "::: API :::", log.Lshortfile)

func addProfile(c *fiber.Ctx) error {
	var profile Profile
	err := c.BodyParser(&profile)
	if err != nil {
		logs.Printf("%s", err)
	}
	logs.Printf("%s,%s,%s", profile.ListenPort, profile.DestinationIp, profile.DestinationPort)
	return nil
}

func handleMain(c *fiber.Ctx) error {
	return c.Render("/etc/mhd/gosample/.public/html/main.html", fiber.Map{}, "html")
}

func Router() {

	router := fiber.New()

	router.Static("/static/css", "/etc/mhd/.public/static/css/style.css")

	router.Add("GET", "/", handleMain)
	router.Add("POST", "/", addProfile)

	router.Listen("localhost:8080")
}
