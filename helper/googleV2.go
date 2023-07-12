package helper

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/oauth2/v2"
)

func GoogleV2(s *http.Client, c echo.Context) error {
	service, err := oauth2.New(s)
	if err != nil {
		return err
	}

	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		return err
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["loggedIn"] = true
	sess.Values["email"] = userInfo.Email
	sess.Values["name"] = userInfo.Name

	sess.Save(c.Request(), c.Response())

	return nil
}
