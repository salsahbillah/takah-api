# Database Design - Takah API

Dokumen ini menjelaskan rancangan database awal untuk aplikasi Takah.

Database digunakan untuk menyimpan data master surat, template surat, surat masuk, surat keluar, approval surat, dan monitoring surat.

---

# Database Engine

Database yang direncanakan:

* MySQL
* Laragon (local development)

Catatan:

* Database saat ini belum diimplementasikan.
* Data masih menggunakan dummy data pada handler.
* SQL schema masih tahap perancangan.

---

# Planned Tables

## 1. users

Digunakan untuk menyimpan data user aplikasi.

| Field      | Type      | Notes         |
| ---------- | --------- | ------------- |
| id         | bigint    | Primary key   |
| name       | varchar   | Nama user     |
| email      | varchar   | Email login   |
| password   | varchar   | Password hash |
| role       | varchar   | admin / user  |
| created_at | timestamp | Waktu dibuat  |

---

## 2. master_takah

Digunakan untuk menyimpan data master jenis surat.

Contoh:

* SKET
* SKK
* UND
* SP

| Field        | Type    | Notes                |
| ------------ | ------- | -------------------- |
| id           | bigint  | Primary key          |
| code         | varchar | Kode surat           |
| name         | varchar | Nama surat           |
| description  | text    | Deskripsi surat      |
| order        | int     | Urutan tampilan      |
| created_by   | varchar | Pembuat data         |
| created_time | varchar | Waktu dibuat         |
| updated_by   | varchar | User update terakhir |
| updated_time | varchar | Waktu update         |

---

## 3. config_nomor_surat

Digunakan untuk konfigurasi format nomor surat otomatis.

Setiap konfigurasi nomor surat memiliki relasi dengan Master Takah sehingga setiap jenis surat dapat menggunakan format nomor yang berbeda.

Config nomor surat juga digunakan untuk menentukan aturan reset nomor surat.

Contoh format:

```text
001/UND/CBN/052026
```

| Field         | Type      | Notes               |
| ------------- | --------- | ------------------- |
| id            | bigint    | Primary key         |
| takah_id      | bigint    | Relasi master_takah |
| company_code  | varchar   | Kode perusahaan     |
| division_code | varchar   | Kode divisi         |
| reset_type    | varchar   | monthly / yearly    |
| created_at    | timestamp | Waktu dibuat        |

Keterangan:

* `monthly` → nomor surat reset setiap bulan
* `yearly` → nomor surat reset setiap tahun

---

## 4. template_surat

Digunakan untuk menyimpan template surat berdasarkan jenis surat.

| Field         | Type      | Notes               |
| ------------- | --------- | ------------------- |
| id            | bigint    | Primary key         |
| takah_id      | bigint    | Relasi master_takah |
| template_name | varchar   | Nama template       |
| content       | text      | Isi template surat  |
| created_at    | timestamp | Waktu dibuat        |

---

## 5. surat_keluar

Digunakan untuk menyimpan surat keluar.

| Field        | Type      | Notes                                 |
| ------------ | --------- | ------------------------------------- |
| id           | bigint    | Primary key                           |
| nomor_surat  | varchar   | Nomor surat                           |
| takah_id     | bigint    | Relasi master_takah                   |
| tujuan_surat | varchar   | Tujuan surat                          |
| perihal      | varchar   | Perihal surat                         |
| lampiran     | varchar   | File lampiran                         |
| status       | varchar   | draft / pending / approved / rejected |
| created_by   | bigint    | User pembuat                          |
| created_at   | timestamp | Waktu dibuat                          |

---

## 6. approval_surat

Digunakan untuk menyimpan proses approval surat keluar.

| Field           | Type      | Notes                         |
| --------------- | --------- | ----------------------------- |
| id              | bigint    | Primary key                   |
| surat_keluar_id | bigint    | Relasi surat keluar           |
| approver_id     | bigint    | User approver                 |
| approval_status | varchar   | pending / approved / rejected |
| notes           | text      | Catatan approval              |
| approved_at     | timestamp | Waktu approval                |

Flow approval:

* Surat dibuat user
* Surat dikirim ke approver
* Approver melakukan review
* Status approval tersimpan

---

## 7. surat_masuk

Digunakan untuk menyimpan surat masuk dari pihak luar perusahaan atau instansi.

| Field         | Type      | Notes            |
| ------------- | --------- | ---------------- |
| id            | bigint    | Primary key      |
| nomor_surat   | varchar   | Nomor surat      |
| pengirim      | varchar   | Pengirim surat   |
| penerima      | varchar   | Penerima surat   |
| perihal       | varchar   | Perihal surat    |
| file_surat    | varchar   | File scan surat  |
| tanggal_surat | date      | Tanggal surat    |
| keterangan    | text      | Keterangan surat |
| created_by    | bigint    | User input       |
| created_at    | timestamp | Waktu dibuat     |

Flow surat masuk:

* Surat diterima dari pihak luar
* Data surat dicatat ke sistem
* File surat diupload
* Surat masuk tersimpan

---

## 8. monitoring_surat

Digunakan untuk monitoring status surat.

| Field           | Type      | Notes                                 |
| --------------- | --------- | ------------------------------------- |
| id              | bigint    | Primary key                           |
| surat_keluar_id | bigint    | Relasi surat keluar                   |
| status          | varchar   | draft / pending / approved / rejected |
| notes           | text      | Catatan monitoring                    |
| updated_by      | bigint    | User update                           |
| updated_at      | timestamp | Waktu update                          |

---

# Table Relationship

```text
master_takah
│
├── template_surat
│
├── config_nomor_surat
│
└── surat_keluar
    │
    ├── approval_surat
    │
    └── monitoring_surat

surat_masuk
│
└── monitoring_surat
```

Relationship:

* `template_surat.takah_id` → `master_takah.id`
* `config_nomor_surat.takah_id` → `master_takah.id`
* `surat_keluar.takah_id` → `master_takah.id`
* `approval_surat.surat_keluar_id` → `surat_keluar.id`
* `monitoring_surat.surat_keluar_id` → `surat_keluar.id`
---


# Generate Nomor Surat

Nomor surat dibuat otomatis berdasarkan konfigurasi nomor surat yang terhubung dengan Master Takah.

Nomor surat dibuat berdasarkan:

* nomor urut
* kode surat
* kode perusahaan/divisi
* bulan dan tahun

Contoh reset bulanan:

```text
001/UND/CBN/052026
002/UND/CBN/052026
001/UND/CBN/062026
```

Contoh reset tahunan:

```text
001/SKET/CBN/2026
002/SKET/CBN/2026
001/SKET/CBN/2027
```

Aturan:

* Nomor surat bertambah otomatis.
* Reset nomor surat mengikuti config yang digunakan.
* Reset dapat dilakukan per bulan atau per tahun.

---

# Current Status

Current database status:

* Database belum diimplementasikan
* Data masih menggunakan dummy data
* SQL schema masih tahap perancangan
* Integrasi MySQL masih planned

---

# Future Development

Pengembangan database selanjutnya:

* Audit log
* Disposisi surat
* Notification table
* Attachment table
* Digital signature
* Multi company support
* Approval multi level
