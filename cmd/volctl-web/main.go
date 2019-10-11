package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ionrock/volctl"
	"github.com/ionrock/volctl/cmd/volctl-web/views"
)

type VolumeEvent struct {
	Volume string `json:"volume"`
}

func index(c echo.Context) error {
	curVol, err := volctl.CurrentVolume()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"CurrentVolume": curVol,
	}
	return c.Render(http.StatusOK, "index", data)
}

func update(c echo.Context) error {
	e := new(VolumeEvent)
	if err := c.Bind(e); err != nil {
		return err
	}

	if err := volctl.UpdateVolume(e.Volume); err != nil {
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("Volume Updated to %s", e.Volume))
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = views.Renderer()
	e.GET("/", index)
	e.POST("/update", update)

	e.Logger.Fatal(e.Start(":1323"))
}
