package entity

type Sekolah struct {
	SekolahID             string  `gorm:"type:uuid;primaryKey" json:"sekolah_id"`
	Nama                  string  `gorm:"type:varchar(255)" json:"nama"`
	NSS                   string  `gorm:"type:varchar(20)" json:"nss"`
	NPSN                  string  `gorm:"type:varchar(20)" json:"npsn"`
	BentukPendidikanID    int     `json:"bentuk_pendidikan_id"`
	BentukPendidikanIDStr string  `gorm:"type:varchar(100)" json:"bentuk_pendidikan_id_str"`
	StatusSekolah         string  `gorm:"type:varchar(50)" json:"status_sekolah"`
	StatusSekolahStr      string  `gorm:"type:varchar(100)" json:"status_sekolah_str"`
	AlamatJalan           string  `gorm:"type:text" json:"alamat_jalan"`
	RT                    string  `gorm:"type:varchar(10)" json:"rt"`
	RW                    string  `gorm:"type:varchar(10)" json:"rw"`
	KodeWilayah           string  `gorm:"type:varchar(20)" json:"kode_wilayah"`
	KodePos               string  `gorm:"type:varchar(10)" json:"kode_pos"`
	NomorTelepon          *string `gorm:"type:varchar(20)" json:"nomor_telepon,omitempty"`
	NomorFax              *string `gorm:"type:varchar(20)" json:"nomor_fax,omitempty"`
	Email                 *string `gorm:"type:varchar(100)" json:"email,omitempty"`
	Website               *string `gorm:"type:varchar(100)" json:"website,omitempty"`
	IsSKS                 bool    `json:"is_sks"`
	Lintang               string  `gorm:"type:varchar(50)" json:"lintang"`
	Bujur                 *string `gorm:"type:varchar(50)" json:"bujur,omitempty"`
	Dusun                 *string `gorm:"type:varchar(100)" json:"dusun,omitempty"`
	DesaKelurahan         string  `gorm:"type:varchar(100)" json:"desa_kelurahan"`
	Kecamatan             string  `gorm:"type:varchar(100)" json:"kecamatan"`
	KabupatenKota         string  `gorm:"type:varchar(100)" json:"kabupaten_kota"`
	Provinsi              string  `gorm:"type:varchar(100)" json:"provinsi"`

	Token string `gorm:"type:varchar(20)" json:"token"`
}

func (Sekolah) TableName() string {
	return "sekolah"
}
