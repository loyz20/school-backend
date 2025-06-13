package dto

import (
	"school-backend/internal/entity"
	"school-backend/pkg/auth"
)

func (d *SiswaDapodikFull) ToEntity() (*entity.Siswa, error) {
	return &entity.Siswa{
		RegistrasiID:          d.RegistrasiID,
		JenisPendaftaranID:    d.JenisPendaftaranID,
		JenisPendaftaranIDStr: d.JenisPendaftaranIDStr,
		TanggalMasukSekolah:   d.TanggalMasukSekolah,
		SekolahAsal:           d.SekolahAsal,
		PesertaDidikID:        d.PesertaDidikID,
		NIPD:                  d.NIPD,
		Nama:                  d.Nama,
		NISN:                  d.NISN,
		JenisKelamin:          d.JenisKelamin,
		NIK:                   d.NIK,
		TanggalLahir:          d.TanggalLahir,
		TempatLahir:           d.TempatLahir,
		NomorTeleponSeluler:   d.NomorTeleponSeluler,
		SemesterID:            d.SemesterID,
		RombonganBelajarID:    d.RombonganBelajarID,
		NamaRombel:            d.NamaRombel,
		AnggotaRombelID:       d.AnggotaRombelID,
		AgamaID:               d.AgamaID,
		AgamaIDStr:            d.AgamaIDStr,
		AlamatJalan:           d.AlamatJalan,
		NomorTeleponRumah:     d.NomorTeleponRumah,
		NamaAyah:              d.NamaAyah,
		PekerjaanAyahID:       d.PekerjaanAyahID,
		PekerjaanAyahIDStr:    d.PekerjaanAyahIDStr,
		NamaIbu:               d.NamaIbu,
		PekerjaanIbuID:        d.PekerjaanIbuID,
		PekerjaanIbuIDStr:     d.PekerjaanIbuIDStr,
		NamaWali:              d.NamaWali,
		PekerjaanWaliID:       d.PekerjaanWaliID,
		PekerjaanWaliIDStr:    d.PekerjaanWaliIDStr,
		TinggiBadan:           d.TinggiBadan,
		BeratBadan:            d.BeratBadan,
		TingkatPendidikanID:   d.TingkatPendidikanID,
		KurikulumID:           d.KurikulumID,
		KurikulumIDStr:        d.KurikulumIDStr,
	}, nil
}

func (d *PenggunaDapodikFull) ToEntity() (*entity.Pengguna, error) {
	passHass, err := auth.HashPassword("guru1234")
	if err != nil {
		return nil, err
	}
	return &entity.Pengguna{
		PenggunaID:     d.PenggunaID,
		SekolahID:      d.SekolahID,
		Username:       d.Username,
		Nama:           d.Nama,
		PeranIDStr:     d.PeranIDStr,
		Password:       passHass,
		Alamat:         d.Alamat,
		NoTelepon:      d.NoTelepon,
		NoHP:           d.NoHP,
		PTKID:          d.PTKID,
		PesertaDidikID: d.PesertaDidikID,
	}, nil
}

func (d *SekolahDapodikFull) ToEntity(token string) (*entity.Sekolah, error) {
	return &entity.Sekolah{
		SekolahID:             d.SekolahID,
		Nama:                  d.Nama,
		NPSN:                  d.NPSN,
		Website:               d.Website,
		NSS:                   d.NSS,
		BentukPendidikanID:    d.BentukPendidikanID,
		BentukPendidikanIDStr: d.BentukPendidikanIDStr,
		StatusSekolah:         d.StatusSekolah,
		StatusSekolahStr:      d.StatusSekolahStr,
		RT:                    d.RT,
		RW:                    d.RW,
		KodeWilayah:           d.KodeWilayah,
		KodePos:               d.KodePos,
		NomorFax:              d.NomorFax,
		Lintang:               d.Lintang,
		Bujur:                 d.Bujur,
		Dusun:                 d.Dusun,
		DesaKelurahan:         d.DesaKelurahan,
		Kecamatan:             d.Kecamatan,
		KabupatenKota:         d.KabupatenKota,
		Provinsi:              d.Provinsi,
		IsSKS:                 d.IsSKS,
		Token:                 token,
	}, nil
}

func (d *RombelDapodikFull) ToEntity() (*entity.RombonganBelajar, error) {
	return &entity.RombonganBelajar{
		RombonganBelajarID:   d.RombonganBelajarID,
		Nama:                 d.Nama,
		TingkatPendidikanID:  d.TingkatPendidikanID,
		TingkatPendidikanStr: d.TingkatPendidikanStr,
		JenisRombel:          d.JenisRombel,
		JenisRombelStr:       d.JenisRombelStr,
		KurikulumID:          d.KurikulumID,
		KurikulumIDStr:       d.KurikulumIDStr,
		IDRuang:              d.IDRuang,
		IDRuangStr:           d.IDRuangStr,
		MovingClass:          d.MovingClass,
		PTKID:                d.PTKID,
		PTKIDStr:             d.PTKIDStr,
		JurusanID:            d.JurusanID,
		JurusanIDStr:         d.JurusanIDStr,
		IDKelasEkskul:        d.IDKelasEkskul,
		IDEkskul:             d.IDEkskul,
		NMEkskul:             d.NMEkskul,
		SKEkskul:             d.SKEkskul,
		IDEkskulStr:          d.IDEkskulStr,
		SemesterID:           d.SemesterID,
	}, nil
}

func (d *AnggotaRombelDapodikFull) ToEntity(id string) (*entity.AnggotaRombel, error) {
	return &entity.AnggotaRombel{
		AnggotaRombelID:     d.AnggotaRombelID,
		RombonganBelajarID:  id,
		PesertaDidikID:      d.PesertaDidikID,
		JenisPendaftaranID:  d.JenisPendaftaranID,
		JenisPendaftaranStr: d.JenisPendaftaranStr,
	}, nil
}

func (d *PembelajaranDapodikFull) ToEntity(id string) (*entity.Pembelajaran, error) {

	return &entity.Pembelajaran{
		PembelajaranID:       d.PembelajaranID,
		RombonganBelajarID:   id,
		MataPelajaranID:      d.MataPelajaranID,
		MataPelajaranIDStr:   d.MataPelajaranIDStr,
		NamaMataPelajaran:    d.NamaMataPelajaran,
		PTKTerdaftarID:       d.PTKTerdaftarID,
		PTKID:                d.PTKID,
		IndukPembelajaranID:  d.IndukPembelajaranID,
		JamMengajarPerMinggu: d.JamMengajarPerMinggu,
		StatusDiKurikulum:    d.StatusDiKurikulum,
		StatusDiKurikulumStr: d.StatusDiKurikulumStr,
	}, nil
}
