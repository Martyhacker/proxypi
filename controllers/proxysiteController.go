package controllers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ProxySite(c *fiber.Ctx) error {
	site := c.Query("q")

	resp, err := http.Get(site)
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		log.Fatal(err)
	}

	responseString := string(responseData)
	c.Response().Header.Set("Content-Type","text/html")
	return c.SendString(responseString)
}