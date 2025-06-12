package dto

import (
	"school-backend/internal/entity"
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
