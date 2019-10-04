package route

import (
	"github.com/labstack/echo"
	"gitlab.ghn.vn/vinhnlk/gowithmongodb/handler"
)

func AddHeroesRouter(e *echo.Echo) {
	e.GET("/addHero", handler.AddHeroes)
	e.GET("/removeHero", handler.RemoveHeroes)
}
