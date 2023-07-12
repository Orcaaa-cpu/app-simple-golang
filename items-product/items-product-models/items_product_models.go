package itemsproductmodels

import (
	"errors"
	"product/config"
	"product/entities"
)

func GetPasien() ([]entities.Pasiens, error) {
	con := config.CreateCon()

	var pasiens []entities.Pasiens
	err := con.Model(&pasiens).Select()
	if err != nil {
		return nil, err
	}

	return pasiens, nil
}

func CreatePasien(pasien *entities.Pasiens) error {
	con := config.CreateCon()

	_, err := con.Model(pasien).Insert()
	if err != nil {
		return err
	}

	return nil
}

func FindId(id int64, pasien *entities.Pasiens) error {
	con := config.CreateCon()

	err := con.Model(pasien).
		Where("id = ?", id).
		Select()
	if err != nil {
		return err
	}

	return nil
}

func EditPasien(pasien entities.Pasiens) error {
	con := config.CreateCon()

	_, err := con.Model(&pasien).
		WherePK().
		Update()
	if err != nil {
		return err
	}

	return nil
}

func DeletePasien(id int64) error {
	con := config.CreateCon()

	res, err := con.Model((*entities.Pasiens)(nil)).
		Where("id = ?", id).
		Delete()
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return errors.New("Pasien tidak ditemukan")
	}

	return nil
}
