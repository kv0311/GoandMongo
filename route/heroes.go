package route

import (
	"github.com/labstack/echo"
	"gitlab.ghn.vn/vinhnlk/gowithmongodb/handler"
	"gitlab.ghn.vn/vinhnlk/gowithmongodb/middleware"
)

func AddHeroesRouter(e *echo.Echo) {
	middleware.LogInfo(e)
	e.GET("/addHero", handler.AddHeroes)
	e.GET("/removeHero", handler.RemoveHeroes)
	e.GET("/removeHeroFilter", handler.RemoveHeroesFilter)
	e.GET("/addHeroFilter", handler.AddHeroesFilter)
}
