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
	RegistrasiID        string `json:"registrasi_id"`
	PesertaDidikID      string `json:"peserta_didik_id"`
	NIPD                string `json:"nipd"`
	Nama                string `json:"nama"`
	NISN                string `json:"nisn"`
	JenisKelamin        string `json:"jenis_kelamin"`
	NIK                 string `json:"nik"`
	TempatLahir         string `json:"tempat_lahir"`
	TanggalLahir        string `json:"tanggal_lahir"` // YYYY-MM-DD
	AgamaIDStr          string `json:"agama_id"`
	NomorTeleponSeluler string `json:"nomor_telepon_seluler"`
	RombonganBelajarID  string `json:"rombongan_belajar_id"`
	NamaRombel          string `json:"nama_rombel"`
	SemesterID          string `json:"semester_id"`
	AnggotaRombelID     string `json:"anggota_rombel_id"`
}
