package xy

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

func xy() {
	e := echo.New()

	e.GET("/home", Home)
}

func Home(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/index.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}
	return tmplate.Execute(x.Response(), nil)
}
