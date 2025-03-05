package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type Organisation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var orgMap = make(map[string]Organisation)

func newApp() *echo.Echo {
	e := echo.New()
	e.GET("/orgs", getAllOrganisation)
	e.POST("/orgs", createOrganisation)
	return e
}

func main() {

	err := newApp().Start("localhost:8080")
	if err != nil {
		panic(err)
	}

}

func getAllOrganisation(c echo.Context) error {

	return c.JSON(http.StatusOK, orgMap)
}

func createOrganisation(c echo.Context) error {
	org := new(Organisation)
	if err := c.Bind(org); err != nil {
		return err
	}
	uuid := uuid.New()
	org.Id = uuid.String()
	orgMap[org.Id] = *org
	return c.JSON(http.StatusCreated, org)

}
