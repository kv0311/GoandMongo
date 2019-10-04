package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"gitlab.ghn.vn/vinhnlk/gowithmongodb/model"
	"gitlab.ghn.vn/vinhnlk/gowithmongodb/repo"
)

func AddHeroes(c echo.Context) error {
	hero := model.Hero{}
	err := c.Bind(&hero)
	if err != nil {
		log.Printf("%s", err)
	}
	log.Println(hero)
	repo.InsertOneHeroes(hero)
	return c.String(http.StatusOK, "add hero success")
}
func RemoveHeroes(c echo.Context) error {
	hero := model.Hero{}
	err := c.Bind(&hero)
	if err != nil {
		log.Printf("%s", err)
	}
	log.Println(hero)
	repo.RemoveOneHeroes(hero)
	return c.String(http.StatusOK, "delete hero success")
}
