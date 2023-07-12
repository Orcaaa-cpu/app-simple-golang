package itemsproductcontroller

import (
	"net/http"
	"product/entities"
	"product/helper"
	itemsproductmodels "product/items-product/items-product-models"
	"strconv"

	"github.com/labstack/echo/v4"
)

var data = make(map[string]interface{})

func GetPasien(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	helper.CheckLogin(c)

	result, err := itemsproductmodels.GetPasien()
	helper.PanicError(err)

	data["pasien"] = result

	helper.Template(c, "view/index.html", data)

	return c.NoContent(http.StatusOK)
}

func ViewCreatePasien(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	helper.CheckLogin(c)
	helper.Template(c, "view/add.html", nil)

	return c.NoContent(http.StatusOK)
}

func CreatePasien(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	pasien := entities.Pasiens{
		NamaLengkap:  c.FormValue("nama_lengkap"),
		NIK:          c.FormValue("nik"),
		JenisKelamin: c.FormValue("jenis_kelamin"),
		TempatLahir:  c.FormValue("tempat_lahir"),
		TanggalLahir: c.FormValue("tanggal_lahir"),
		Alamat:       c.FormValue("alamat"),
		NoHp:         c.FormValue("no_hp"),
	}

	errorMessage := make(map[string]interface{})

	if err := c.Validate(pasien); err != nil {
		arr := helper.ConvertErr(err)

		errorMessage["validation"] = arr
		errorMessage["pasien"] = pasien

		helper.Template(c, "view/add.html", errorMessage)
	} else {
		err := itemsproductmodels.CreatePasien(&pasien)
		helper.PanicError(err)

		errorMessage["pesan"] = "Berhasil Menambahkan Data"

		helper.Template(c, "view/add.html", errorMessage)
	}

	return c.NoContent(http.StatusOK)
}

func ViewEdit(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	queryString := c.Request().URL.Query()
	id, _ := strconv.Atoi(queryString.Get("id"))

	var pasien entities.Pasiens
	itemsproductmodels.FindId(int64(id), &pasien)

	data := map[string]interface{}{
		"pasien": pasien,
	}

	helper.CheckLogin(c)
	helper.Template(c, "view/edit.html", data)

	return c.NoContent(http.StatusOK)
}

func EditPasien(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	c.Request().ParseForm()

	var pasien entities.Pasiens
	pasien.Id, _ = strconv.ParseInt(c.Request().Form.Get("id"), 10, 64)
	pasien.NamaLengkap = c.Request().Form.Get("nama_lengkap")
	pasien.NIK = c.Request().Form.Get("nik")
	pasien.JenisKelamin = c.Request().Form.Get("jenis_kelamin")
	pasien.TempatLahir = c.Request().Form.Get("tempat_lahir")
	pasien.TanggalLahir = c.Request().Form.Get("tanggal_lahir")
	pasien.Alamat = c.Request().Form.Get("alamat")
	pasien.NoHp = c.Request().Form.Get("no_hp")

	var data = make(map[string]interface{})

	if err := c.Validate(pasien); err != nil {
		arr := helper.ConvertErr(err)
		data["validation"] = arr
		data["pasien"] = pasien
	} else {
		data["pesan"] = "Data pasien berhasil diperbarui"
		itemsproductmodels.EditPasien(pasien)
	}

	helper.Template(c, "view/edit.html", data)

	return c.NoContent(http.StatusOK)
}

func DeletePasien(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	id, _ := strconv.Atoi(c.FormValue("id"))
	itemsproductmodels.DeletePasien(int64(id))

	c.Redirect(http.StatusSeeOther, "/pasien")

	return c.NoContent(http.StatusOK)
}
