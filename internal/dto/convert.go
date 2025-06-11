package dto

import (
	"school-backend/internal/entity"
	"time"
)

func (d *SiswaDapodikFull) ToEntity() (*entity.Siswa, error) {
	tgl, err := time.Parse("2006-01-02", d.TanggalLahir)
	if err != nil {
		return nil, err
	}
	return &entity.Siswa{
		RegistrasiID:        d.RegistrasiID,
		PesertaDidikID:      d.PesertaDidikID,
		NIPD:                d.NIPD,
		Nama:                d.Nama,
		NISN:                d.NISN,
		JenisKelamin:        d.JenisKelamin,
		NIK:                 d.NIK,
		TanggalLahir:        tgl,
		TempatLahir:         d.TempatLahir,
		NomorTeleponSeluler: d.NomorTeleponSeluler,
		SemesterID:          d.SemesterID,
		RombonganBelajarID:  d.RombonganBelajarID,
		NamaRombel:          d.NamaRombel,
		AnggotaRombelID:     d.AnggotaRombelID,
	}, nil
}
