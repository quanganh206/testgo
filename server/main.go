package main

import (
	"fmt"
	"github.com/quanganh206/testgo/server/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
)

const port string = ":3333"

func main() {
	fmt.Printf("Running at %v\n", port)

	e := echo.New()

	routes.Init(e)

	e.Run(fasthttp.New(port))
}
