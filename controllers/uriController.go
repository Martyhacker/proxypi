package controllers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func PostUri(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	uri := data["uri"]

	response, err := http.Get(uri)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    responseString := string(responseData)
    return c.SendString(responseString)
}
