CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE semester (
    semester_id TEXT PRIMARY KEY,
    tahun_ajaran TEXT NOT NULL,
    semester_ke INT NOT NULL,
    aktif BOOLEAN DEFAULT FALSE
);

CREATE TABLE sekolah (
    sekolah_id UUID PRIMARY KEY,
    nama TEXT NOT NULL,
    nss VARCHAR(20),
    npsn VARCHAR(20),
    bentuk_pendidikan_id INT,
    bentuk_pendidikan_id_str VARCHAR(100),
    status_sekolah VARCHAR(50),
    status_sekolah_str VARCHAR(100),
    alamat_jalan TEXT,
    rt VARCHAR(10),
    rw VARCHAR(10),
    kode_wilayah VARCHAR(20),
    kode_pos VARCHAR(10),
    nomor_telepon VARCHAR(20),
    nomor_fax VARCHAR(20),
    email VARCHAR(100),
    website VARCHAR(100),
    is_sks BOOLEAN DEFAULT FALSE,
    lintang VARCHAR(50),
    bujur VARCHAR(50),
    dusun VARCHAR(100),
    desa_kelurahan VARCHAR(100),
    kecamatan VARCHAR(100),
    kabupaten_kota VARCHAR(100),
    provinsi VARCHAR(100)
);

CREATE TABLE pengguna (
    pengguna_id UUID PRIMARY KEY,
    sekolah_id UUID,
    username TEXT UNIQUE NOT NULL,
    nama TEXT,
    peran_id_str TEXT,
    password TEXT NOT NULL,
    alamat TEXT,
    no_telepon TEXT,
    no_hp TEXT,
    ptk_id TEXT,
    peserta_didik_id TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pengguna_id UUID NOT NULL,
    token TEXT NOT NULL,
    ip_address TEXT,
    user_agent TEXT,
    expired_at TIMESTAMPTZ NOT NULL,
    is_revoked BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rombongan_belajar (
    rombongan_belajar_id UUID PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    tingkat_pendidikan_id VARCHAR(10),
    tingkat_pendidikan_str VARCHAR(100),
    jenis_rombel VARCHAR(5),
    jenis_rombel_str VARCHAR(100),
    kurikulum_id INTEGER,
    kurikulum_id_str VARCHAR(100),
    id_ruang UUID,
    id_ruang_str VARCHAR(100),
    moving_class VARCHAR(10),
    ptk_id UUID,
    ptk_id_str VARCHAR(100),
    jurusan_id VARCHAR(20),
    jurusan_id_str VARCHAR(100),
    id_kelas_ekskul UUID,
    id_ekskul INTEGER,
    nm_ekskul VARCHAR(100),
    sk_ekskul VARCHAR(100),
    id_ekskul_str VARCHAR(100),
    semester_id TEXT NOT NULL REFERENCES semester(semester_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE anggota_rombel (
    anggota_rombel_id UUID PRIMARY KEY,
    rombongan_belajar_id UUID NOT NULL REFERENCES rombongan_belajar(rombongan_belajar_id),
    peserta_didik_id UUID,
    jenis_pendaftaran_id VARCHAR(5),
    jenis_pendaftaran_str VARCHAR(100)
);

CREATE TABLE pembelajaran (
    pembelajaran_id UUID PRIMARY KEY,
    rombongan_belajar_id UUID NOT NULL REFERENCES rombongan_belajar(rombongan_belajar_id),
    mata_pelajaran_id INTEGER,
    matapelajaran_id_str VARCHAR(100),
    nama_mata_pelajaran VARCHAR(100),
    ptk_terdaftar_id UUID,
    ptk_id UUID,
    induk_pembelajaran_id UUID,
    jam_mengajar_per_minggu VARCHAR(5),
    status_di_kurikulum VARCHAR(5),
    status_kurikulum_str VARCHAR(100)
);

CREATE TABLE siswa (
    registrasi_id UUID PRIMARY KEY,
    jenis_pendaftaran_id TEXT,
    jenis_pendaftaran_id_str TEXT,
    n_ip_d TEXT,
    tanggal_masuk_sekolah TEXT,
    sekolah_asal TEXT,
    peserta_didik_id TEXT,
    nama TEXT,
    nisn TEXT,
    jenis_kelamin TEXT,
    nik TEXT,
    tempat_lahir TEXT,
    tanggal_lahir TEXT,
    agama_id BIGINT,
    agama_id_str TEXT,
    alamat_jalan TEXT,
    nomor_telepon_rumah TEXT,
    nomor_telepon_seluler TEXT,
    nama_ayah TEXT,
    pekerjaan_ayah_id BIGINT,
    pekerjaan_ayah_id_str TEXT,
    nama_ibu TEXT,
    pekerjaan_ibu_id BIGINT,
    pekerjaan_ibu_id_str TEXT,
    nama_wali TEXT,
    pekerjaan_wali_id BIGINT,
    pekerjaan_wali_id_str TEXT,
    tinggi_badan TEXT,
    berat_badan TEXT,
    anggota_rombel_id TEXT,
    rombongan_belajar_id TEXT,
    tingkat_pendidikan_id TEXT,
    nama_rombel TEXT,
    kurikulum_id BIGINT,
    kurikulum_id_str TEXT,
    semester_id TEXT REFERENCES semester(semester_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
