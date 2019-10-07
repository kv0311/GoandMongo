package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"gitlab.ghn.vn/vinhnlk/gowithmongodb/model"
	"gitlab.ghn.vn/vinhnlk/gowithmongodb/repo"
	"go.mongodb.org/mongo-driver/bson"
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
func RemoveHeroesFilter(c echo.Context) error {
	hero := bson.M{}
	err := c.Bind(&hero)
	if err != nil {
		log.Printf("%s", err)
	}
	log.Println(hero)
	repo.RemoveOneHeroesFilter(hero)
	return c.String(http.StatusOK, "delete hero by filter success")
}
