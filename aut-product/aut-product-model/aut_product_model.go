package autproductmodel

import (
	"errors"
	"product/config"
	"product/entities"
	"product/helper"

	"github.com/go-pg/pg/v10"
)

func Login(user *entities.Users, username, password string) error {
	con := config.CreateCon()

	err := con.Model(user).
		Where("username = ?", username).
		Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return errors.New("Username tidak ditemukan")
		}
		return err
	}

	match, err := helper.CheckPasswordHash(password, user.Password)
	if !match {
		return err
	}

	return nil
}

func Register(user *entities.Users) error {
	con := config.CreateCon()

	_, _ = con.Model(user).Insert()

	return nil
}

func Unic(user *entities.Users, value, param string) bool {
	con := config.CreateCon()

	err := con.Model(user).
		Where(param+" = ?", value).
		Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return false
		}
		return false
	}

	return true
}
