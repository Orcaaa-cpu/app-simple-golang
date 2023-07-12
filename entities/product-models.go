package entities

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type (
	Users struct {
		Id        int64
		Name      string `json:"nama" validate:"required"`
		Email     string `json:"email" validate:"required,email"`
		Username  string `json:"username" validate:"required,gte=3"`
		Password  string `validate:"required,gte=6"`
		Cpassword string `validate:"required,eqfield=Password"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}

	Pasiens struct {
		Id           int64
		NamaLengkap  string `validate:"required"`
		NIK          string `validate:"required"`
		JenisKelamin string `validate:"required"`
		TempatLahir  string `validate:"required"`
		TanggalLahir string `validate:"required"`
		Alamat       string `validate:"required"`
		NoHp         string `validate:"required"`
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
