package dto

type SemesterResponse struct {
	SemesterID  string `json:"id"`
	TahunAjaran string `json:"tahun_ajaran"`
	SemesterKe  int    `json:"semester_ke"`
	Aktif       bool   `json:"aktif"`
}

type UserResponse struct {
	Username   string `json:"username"`
	Nama       string `json:"nama"`
	PeranIDStr string `json:"peran_id_str"`
	Alamat     string `json:"alamat"`
	NoTelepon  string `json:"no_telepon"`
	NoHP       string `json:"no_hp"`
}

type SiswaResponse struct {
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

	BeratBadan          string `json:"berat_badan"`
	SemesterID          string `json:"semester_id"`
	AnggotaRombelID     string `json:"anggota_rombel_id"`
	RombonganBelajarID  string `json:"rombongan_belajar_id"`
	TingkatPendidikanID string `json:"tingkat_pendidikan_id"`
}

type PenggunaResponse struct {
	PenggunaID     string  `json:"pengguna_id"`
	SekolahID      string  `json:"sekolah_id"`
	Username       string  `json:"username"`
	Nama           string  `json:"nama"`
	PeranIDStr     string  `json:"peran_id_str"`
	Password       string  `json:"password"`
	Alamat         *string `json:"alamat"`
	NoTelepon      *string `json:"no_telepon"`
	NoHP           *string `json:"no_hp"`
	PTKID          *string `json:"ptk_id"`
	PesertaDidikID *string `json:"peserta_didik_id"`
	KurikulumID    int     `json:"kurikulum_id"`
	KurikulumIDStr string  `json:"kurikulum_id_str"`
}

type SekolahResponse struct {
	SekolahID             string  `json:"sekolah_id"`
	Nama                  string  `json:"nama"`
	Alamat                string  `json:"alamat"`
	NoTelepon             string  `json:"no_telepon"`
	NoHP                  string  `json:"no_hp"`
	Email                 string  `json:"email"`
	Website               string  `json:"website"`
	NSS                   string  `json:"nss"`
	NPSN                  string  `json:"npsn"`
	KodeRegistrasi        string  `json:"kode_registrasi"`
	ProvinsiID            string  `json:"provinsi_id"`
	ProvinsiStr           string  `json:"provinsi_str"`
	KabupatenID           string  `json:"kabupaten_id"`
	KabupatenStr          string  `json:"kabupaten_str"`
	KecamatanID           string  `json:"kecamatan_id"`
	KecamatanStr          string  `json:"kecamatan_str"`
	KelurahanID           string  `json:"kelurahan_id"`
	KelurahanStr          string  `json:"kelurahan_str"`
	BentukPendidikanID    int     `json:"bentuk_pendidikan_id"`
	BentukPendidikanIDStr string  `json:"bentuk_pendidikan_id_str"`
	StatusSekolah         string  `json:"status_sekolah"`
	StatusSekolahStr      string  `json:"status_sekolah_str"`
	Lintang               string  `json:"lintang"`
	Bujur                 *string `json:"bujur,omitempty"`
	Dusun                 *string `json:"dusun,omitempty"`
	DesaKelurahan         string  `json:"desa_kelurahan"`
	Kecamatan             string  `json:"kecamatan"`
	KabupatenKota         string  `json:"kabupaten_kota"`
	Provinsi              string  `json:"provinsi"`
	IsSKS                 bool    `json:"is_sks"`
	Token                 string  `json:"token"`
}

type RombelResponse struct {
	RombonganBelajarID   string                  `json:"rombongan_belajar_id"`
	Nama                 string                  `json:"nama"`
	TingkatPendidikanID  string                  `json:"tingkat_pendidikan_id"`
	TingkatPendidikanStr string                  `json:"tingkat_pendidikan_id_str"`
	JenisRombel          string                  `json:"jenis_rombel"`
	JenisRombelStr       string                  `json:"jenis_rombel_str"`
	KurikulumID          int                     `json:"kurikulum_id"`
	KurikulumIDStr       string                  `json:"kurikulum_id_str"`
	IDRuang              string                  `json:"id_ruang"`
	IDRuangStr           string                  `json:"id_ruang_str"`
	MovingClass          string                  `json:"moving_class"`
	PTKID                string                  `json:"ptk_id"`
	PTKIDStr             string                  `json:"ptk_id_str"`
	JurusanID            string                  `json:"jurusan_id"`
	JurusanIDStr         string                  `json:"jurusan_id_str"`
	IDKelasEkskul        *string                 `json:"id_kelas_ekskul"`
	IDEkskul             *int                    `json:"id_ekskul"`
	NMEkskul             *string                 `json:"nm_ekskul"`
	SKEkskul             *string                 `json:"sk_ekskul"`
	IDEkskulStr          *string                 `json:"id_ekskul_str"`
	SemesterID           string                  `json:"semester_id"`
	Semester             SemesterResponse        `json:"semester"`
	AnggotaRombel        []AnggotaRombelResponse `json:"anggota_rombel"`
	Pembelajaran         []PembelajaranResponse  `json:"pembelajaran"`
}
type AnggotaRombelResponse struct {
	AnggotaRombelID     string `json:"anggota_rombel_id"`
	RombonganBelajarID  string `json:"rombongan_belajar_id"`
	PesertaDidikID      string `json:"peserta_didik_id"`
	JenisPendaftaranID  string `json:"jenis_pendaftaran_id"`
	JenisPendaftaranStr string `json:"jenis_pendaftaran_id_str"`
}

type PembelajaranResponse struct {
	PembelajaranID       string  `json:"pembelajaran_id"`
	RombonganBelajarID   string  `json:"rombongan_belajar_id"`
	MataPelajaranID      int     `json:"mata_pelajaran_id"`
	MataPelajaranIDStr   string  `json:"mata_pelajaran_id_str"`
	NamaMataPelajaran    string  `json:"nama_mata_pelajaran"`
	PTKTerdaftarID       string  `json:"ptk_terdaftar_id"`
	PTKID                string  `json:"ptk_id"`
	IndukPembelajaranID  *string `json:"induk_pembelajaran_id"`
	JamMengajarPerMinggu string  `json:"jam_mengajar_per_minggu"`
	StatusDiKurikulum    string  `json:"status_di_kurikulum"`
	StatusDiKurikulumStr string  `json:"status_di_kurikulum_str"`
}
