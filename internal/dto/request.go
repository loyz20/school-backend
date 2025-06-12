package dto

type CreateSemesterRequest struct {
	TahunAjaran string `json:"tahun_ajaran" binding:"required"`
	SemesterKe  int    `json:"semester_ke" binding:"required,oneof=1 2"`
	Aktif       bool   `json:"aktif"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type UserUpdateRequest struct {
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	NoTelepon string `json:"no_telepon"`
	NoHP      string `json:"no_hp"`
}

type SiswaDapodikFull struct {
	RegistrasiID          string  `json:"registrasi_id"`
	JenisPendaftaranID    string  `json:"jenis_pendaftaran_id"`
	JenisPendaftaranIDStr string  `json:"jenis_pendaftaran_id_str"`
	NIPD                  string  `json:"nipd"`
	TanggalMasukSekolah   string  `json:"tanggal_masuk_sekolah"`
	SekolahAsal           *string `json:"sekolah_asal"`
	PesertaDidikID        string  `json:"peserta_didik_id"`
	Nama                  string  `json:"nama"`
	NISN                  *string `json:"nisn"`
	JenisKelamin          string  `json:"jenis_kelamin"`
	NIK                   string  `json:"nik"`
	TempatLahir           string  `json:"tempat_lahir"`
	TanggalLahir          string  `json:"tanggal_lahir"`
	AgamaID               int     `json:"agama_id"`
	AgamaIDStr            string  `json:"agama_id_str"`
	AlamatJalan           *string `json:"alamat_jalan"`
	NomorTeleponRumah     *string `json:"nomor_telepon_rumah"`
	NomorTeleponSeluler   *string `json:"nomor_telepon_seluler"`
	NamaAyah              string  `json:"nama_ayah"`
	PekerjaanAyahID       *int    `json:"pekerjaan_ayah_id"`
	PekerjaanAyahIDStr    string  `json:"pekerjaan_ayah_id_str"`
	NamaIbu               string  `json:"nama_ibu"`
	PekerjaanIbuID        *int    `json:"pekerjaan_ibu_id"`
	PekerjaanIbuIDStr     string  `json:"pekerjaan_ibu_id_str"`
	NamaWali              *string `json:"nama_wali"`
	PekerjaanWaliID       int     `json:"pekerjaan_wali_id"`
	PekerjaanWaliIDStr    string  `json:"pekerjaan_wali_id_str"`
	TinggiBadan           string  `json:"tinggi_badan"`
	BeratBadan            string  `json:"berat_badan"`
	SemesterID            string  `json:"semester_id"`
	AnggotaRombelID       string  `json:"anggota_rombel_id"`
	RombonganBelajarID    string  `json:"rombongan_belajar_id"`
	TingkatPendidikanID   string  `json:"tingkat_pendidikan_id"`
	NamaRombel            string  `json:"nama_rombel"`
	KurikulumID           int     `json:"kurikulum_id"`
	KurikulumIDStr        string  `json:"kurikulum_id_str"`
}
