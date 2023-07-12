package helper

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	autproductmodel "product/aut-product/aut-product-model"
	"product/entities"
	"product/helper"

	"crypto/tls"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig *oauth2.Config
	httpClient  *http.Client
)

func init() {
	oauthConfig = &oauth2.Config{
		ClientID:     "35285915289-ljfdb39e8t3tlk8ech1b53ja3jg56c4p.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient = &http.Client{
		Transport: transport,
	}
}

func GoogleLoginHandler(c echo.Context) error {
	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallbackHandler(c echo.Context) error {
	code := c.QueryParam("code")
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return err
	}

	client := oauthConfig.Client(context.Background(), token)
	client.Transport = httpClient.Transport
	response, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	userInfo := struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}{}
	err = json.Unmarshal(body, &userInfo)
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

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/pasien")
}

func ViewLogin(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	helper.Template(c, "view/login.html", nil)

	return c.NoContent(http.StatusOK)
}

func Login(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	username := c.FormValue("username")
	password := c.FormValue("password")

	user := entities.Users{}

	err := autproductmodel.Login(&user, username, password)
	if err != nil {
		err = errors.New("Username atau Password salah")
		data := map[string]interface{}{
			"error": err,
		}
		helper.Template(c, "view/login.html", data)
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["loggedIn"] = true
	sess.Values["email"] = user.Email
	sess.Values["username"] = user.Username
	sess.Values["name"] = user.Name

	sess.Save(c.Request(), c.Response())

	c.Redirect(http.StatusSeeOther, "/pasien")

	return c.NoContent(http.StatusOK)
}

func ViewRegister(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	helper.Template(c, "view/register.html", nil)

	return c.NoContent(http.StatusOK)
}

func Register(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	user := entities.Users{
		Name:      c.FormValue("name"),
		Email:     c.FormValue("email"),
		Username:  c.FormValue("username"),
		Password:  c.FormValue("password"),
		Cpassword: c.FormValue("cpassword"),
	}

	errorMessage := make(map[string]interface{})

	if err := c.Validate(user); err != nil {
		arr := helper.ConvertErr(err)

		errorMessage["validation"] = arr
		errorMessage["user"] = user

		helper.Template(c, "view/register.html", errorMessage)
	} else {
		email := autproductmodel.Unic(&user, user.Email, "email")
		username := autproductmodel.Unic(&user, user.Username, "username")

		if email || username {
			unic := make(map[string]interface{})
			if email {
				unic["Email"] = "Email Sudah Di Gunakan"
			} else {
				unic["Username"] = "Username Sudah Di Gunakan"
			}
			errorMessage["validation"] = unic
			errorMessage["user"] = user

			helper.Template(c, "view/register.html", errorMessage)
		} else {

			user.Password, _ = helper.HashPassword(user.Password)
			user.Cpassword = user.Password
			err := autproductmodel.Register(&user)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			errorMessage["pesan"] = "Registrasi Berhasil, Silahkan Login"

			helper.Template(c, "view/register.html", errorMessage)
		}
	}

	return c.NoContent(http.StatusOK)
}

func Logout(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	sess, _ := session.Get("session", c)

	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	c.Redirect(http.StatusSeeOther, "/")

	return c.NoContent(http.StatusOK)
}
