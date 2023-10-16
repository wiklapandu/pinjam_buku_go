package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
type Person struct {
	Name string
}

func HomeHandler(c buffalo.Context) error {
	user := Person{
		Name: "Hello World",
	}

	return c.Render(http.StatusOK, r.JSON(user))
}
