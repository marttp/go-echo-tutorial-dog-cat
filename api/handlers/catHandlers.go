package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func GetCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s and his type is %s", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "you need to lets us know if you want json or string data",
	})
}

func AddCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed unmarshaling in addCat: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your cat: %#v", cat)
	return c.String(http.StatusOK, "We got your cat")
}
