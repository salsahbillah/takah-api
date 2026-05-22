# Database Design - Takah API

Dokumen ini menjelaskan rancangan database awal untuk aplikasi Takah.

Database digunakan untuk menyimpan data master surat, template surat, surat masuk, surat keluar, dan monitoring surat.

---

# Database Engine

Database yang direncanakan:

- MySQL
- Laragon (local development)

Catatan:
- Database saat ini belum diimplementasikan.
- Data masih menggunakan dummy data pada handler.
- SQL schema masih tahap perancangan.

---

# Planned Tables

## 1. users

Digunakan untuk menyimpan data user aplikasi.

| Field | Type | Notes |
| --- | --- | --- |
| id | bigint | Primary key |
| name | varchar | Nama user |
| email | varchar | Email login |
| password | varchar | Password hash |
| role | varchar | admin / user |
| created_at | timestamp | Waktu dibuat |

---

## 2. master_takah

Digunakan untuk menyimpan data master jenis surat.

Contoh:
- SKET
- SKK
- UND
- SP

| Field | Type | Notes |
| --- | --- | --- |
| id | bigint | Primary key |
| code | varchar | Kode surat |
| name | varchar | Nama surat |
| description | text | Deskripsi surat |
| order | int | Urutan tampilan |
| created_by | varchar | Pembuat data |
| created_time | varchar | Waktu dibuat |
| updated_by | varchar | User update terakhir |
| updated_time | varchar | Waktu update |

---

## 3. config_nomor_surat

Digunakan untuk konfigurasi format nomor surat otomatis.

Contoh format:

```text
001/UND/CBN/052026
```

| Field | Type | Notes |
| --- | --- | --- |
| id | bigint | Primary key |
| company_code | varchar | Kode perusahaan |
| division_code | varchar | Kode divisi |
| reset_type | varchar | monthly / yearly |
| created_at | timestamp | Waktu dibuat |

---

## 4. template_surat

Digunakan untuk menyimpan template surat berdasarkan jenis surat.

| Field | Type | Notes |
| --- | --- | --- |
| id | bigint | Primary key |
| takah_id | bigint | Relasi master_takah |
| template_name | varchar | Nama template |
| content | text | Isi template surat |
| created_at | timestamp | Waktu dibuat |

---

## 5. surat_keluar

Digunakan untuk menyimpan surat keluar.

| Field | Type | Notes |
| --- | --- | --- |
| id | bigint | Primary key |
| nomor_surat | varchar | Nomor surat |
| takah_id | bigint | Relasi master_takah |
| tujuan_surat | varchar | Tujuan surat |
| perihal | varchar | Perihal surat |
| lampiran | varchar | File lampiran |
| status | varchar | draft / pending / approved |
| created_by | bigint | User pembuat |
| created_at | timestamp | Waktu dibuat |

---

## 6. surat_masuk

Digunakan untuk menyimpan surat masuk dari pihak luar.

| Field | Type | Notes |
| --- | --- | --- |
| id | bigint | Primary key |
| nomor_surat | varchar | Nomor surat |
| pengirim | varchar | Pengirim surat |
| perihal | varchar | Perihal surat |
| file_surat | varchar | File scan surat |
| tanggal_surat | date | Tanggal surat |
| created_at | timestamp | Waktu dibuat |

---

## 7. monitoring_surat

Digunakan untuk monitoring status surat.

| Field | Type | Notes |
| --- | --- | --- |
| id | bigint | Primary key |
| surat_keluar_id | bigint | Relasi surat keluar |
| status | varchar | draft / pending / approved / rejected |
| notes | text | Catatan monitoring |
| updated_by | bigint | User update |
| updated_at | timestamp | Waktu update |

---

# Table Relationship

```text
master_takah
│
├── template_surat
│
└── surat_keluar
    │
    └── monitoring_surat
```

Relationship:
- `template_surat.takah_id` → `master_takah.id`
- `surat_keluar.takah_id` → `master_takah.id`
- `monitoring_surat.surat_keluar_id` → `surat_keluar.id`

---

# Generate Nomor Surat

Nomor surat dibuat otomatis berdasarkan:
- nomor urut
- kode surat
- kode perusahaan/divisi
- bulan dan tahun

Contoh:

```text
001/UND/CBN/052026
002/UND/CBN/052026
001/UND/CBN/062026
```

Keterangan:
- `001` → nomor urut surat
- `UND` → kode jenis surat
- `CBN` → kode perusahaan/divisi
- `052026` → bulan dan tahun

Aturan:
- Nomor surat bertambah otomatis.
- Nomor surat reset setiap bulan.

---

# Current Status

Current database status:

- Database belum diimplementasikan
- Data masih menggunakan dummy data
- SQL schema masih tahap perancangan
- Integrasi MySQL masih planned

---

# Future Development

Pengembangan database selanjutnya:

- Approval table
- Audit log
- Disposisi surat
- Notification table
- Attachment table
- Digital signature
- Multi company support