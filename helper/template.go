package helper

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

func Template(c echo.Context, url string, data interface{}) {
	temp, err := template.ParseFiles(url)
	PanicError(err)

	temp.Execute(c.Response(), data)
}
