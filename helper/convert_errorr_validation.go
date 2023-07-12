package helper

import (
	"strings"
)

var (
	user     = "NameEmailUsernamePasswordCpasswordNamaLengkapNIKJenisKelaminTempatLahirTanggalLahirAlamatNoHp"
	validate = "emailgteeqfieldrequired"
	message  = map[string]string{
		"Name":         "Nama Lengkap",
		"Cpassword":    "Konfirmasi Password",
		"NamaLengkap":  "Nama Lengkap",
		"JenisKelamin": "Jenis Kelamin",
		"TempatLahir":  "Tempat Lahir",
		"TanggalLahir": "Tanggal Lahir",
		"NoHp":         "Nomor HP",
	}
	message1 = " Tidak Boleh Kosong"
)

func ConvertErr(e error) interface{} {
	s := ""

	for _, e := range e.Error() {
		s += string(e)
	}

	users := []string{}
	valid := []string{}

	str := ""
	for _, v := range s {
		if string(v) == "'" {
			if len(str) != 0 {
				if strings.Contains(user, str) {
					users = append(users, str)
				} else if strings.Contains(validate, str) {
					valid = append(valid, str)
				}
			}
		} else if v == ' ' {
			str = ""
		} else {
			str += string(v)
		}
	}

	result := MessageConvert(users, valid)

	return result
}

func MessageConvert(users, valid []string) interface{} {
	arr := make(map[string]string)
	for i, v := range users {

		switch {
		case valid[i] == "required":
			if len(message[v]) != 0 {
				arr[v] = message[v] + message1
			} else {
				arr[v] = v + message1
			}
		case valid[i] == "email":
			arr[v] = "Masukan Format Email Dengan Benar"
		case valid[i] == "eqfield":
			arr[v] = "Konfirmasi Password Tidak Sama Dengan Password"
		case valid[i] == "gte":
			if v == "Username" {
				arr[v] = v + " Minimal 3 Character"
			} else {
				arr[v] = v + " Minimal 6 Character"
			}
		}
	}

	return arr
}
