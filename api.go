package example

import (
	"net/http"

	"github.com/labstack/echo"
)

type api struct {
	server Server
}

func NewAPI(server Server) error {
	a := &api{
		server: server,
	}

	e := echo.New()

	e.POST("/users", a.createUser)

	return e.Start(":8080")
}

func (a api) createUser(c echo.Context) error {
	user := &User{}
	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := a.server.CreateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}
