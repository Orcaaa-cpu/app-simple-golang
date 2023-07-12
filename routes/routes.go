package routes

import (
	autproductcontroller "product/aut-product/aut-product-controller"
	"product/entities"
	itemsproductcontroller "product/items-product/items-product-controller"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Validator = &entities.CustomValidator{Validator: validator.New()}

	e.GET("/", autproductcontroller.ViewLogin)
	e.POST("/login", autproductcontroller.Login)
	e.GET("/logout", autproductcontroller.Logout)
	e.GET("/register", autproductcontroller.ViewRegister)
	e.POST("/register", autproductcontroller.Register)

	e.GET("/google/login", autproductcontroller.GoogleLoginHandler)
	e.GET("/auth/google/callback", autproductcontroller.GoogleCallbackHandler)

	e.GET("/pasien", itemsproductcontroller.GetPasien)
	e.GET("/pasien/add", itemsproductcontroller.ViewCreatePasien)
	e.POST("/pasien/add", itemsproductcontroller.CreatePasien)
	e.GET("/pasien/edit", itemsproductcontroller.ViewEdit)
	e.POST("/pasien/edit", itemsproductcontroller.EditPasien)
	e.GET("/pasien/delete", itemsproductcontroller.DeletePasien)

	return e
}
