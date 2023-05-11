package models

type Karyawan struct {
	Id                  int    `json:"id"`
	Nik                 string `json:"nik"`
	Nama                string `json:"nama"`
	Tempat_lahir        string `json:"tempat_lahir"`
	Tangga_lahir        string `json:"tanggal_lahir"`
	Kewarganegaraan     string `json:"kewarganegaraan"`
	Pendidikan_terakhir string `json:"pendidikan_terakhir"`
	Agama               string `json:"agama"`
	Alamat              string `json:"alamat"`
	Telpon              string `json:"telpon"`
	Email               string `json:"email" gorm:"unique"`
	Mulai_bekerja       string `json:"mulai_bekerja"`
	Jabatan             string `json:"jabatan"`
}
