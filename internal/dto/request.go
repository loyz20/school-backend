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
