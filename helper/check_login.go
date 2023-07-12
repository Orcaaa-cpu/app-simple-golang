package helper

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CheckLogin(c echo.Context) {
	sess, _ := session.Get("session", c)
	if len(sess.Values) == 0 || sess.Values["loggedIn"] != true {
		c.Redirect(http.StatusSeeOther, "/")
	}
}
