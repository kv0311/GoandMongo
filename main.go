package main

import (
	"fmt"

	"github.com/labstack/echo"
	"gitlab.ghn.vn/vinhnlk/gowithmongodb/route"
)

func main() {
	e := echo.New()
	fmt.Println(e)
	route.AddHeroesRouter(e)
	fmt.Println("Welcome to the webserver")
	e.Start(":8000")
}
