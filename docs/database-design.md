# Database Design - Takah API

Dokumen ini menjelaskan rancangan database yang digunakan pada aplikasi Takah.

Database digunakan untuk menyimpan data user, master surat, konfigurasi nomor surat, template surat, parameter surat, surat keluar, approval surat, surat masuk, dan monitoring surat.

---

# Database Engine

Database yang digunakan:

* MySQL
* Laragon (local development)

Status database:

* Database `takah_db` sudah dibuat.
* Seluruh tabel utama telah dibuat menggunakan file `schema.sql`.
* Data awal telah ditambahkan menggunakan file `seed.sql`.
* Backend Golang telah berhasil terhubung dengan MySQL.
* Konfigurasi database menggunakan file `.env`.
* Endpoint utama telah menggunakan query MySQL.

---

# Tables

## 1. users

Digunakan untuk menyimpan data user aplikasi.

| Field      | Type      | Notes                 |
| ---------- | --------- | --------------------- |
| id         | bigint    | Primary key           |
| name       | varchar   | Nama user             |
| email      | varchar   | Email login           |
| password   | varchar   | Password user         |
| role       | varchar   | admin / user          |
| created_at | timestamp | Waktu dibuat          |
| updated_at | timestamp | Waktu update terakhir |

Catatan:

* Field `email` digunakan untuk proses login.
* Field `password` saat ini masih menggunakan plaintext.
* Field `role` digunakan untuk membedakan hak akses admin dan user.
* Password hashing akan dikembangkan pada tahap berikutnya.

---

## 2. master_takah

Digunakan untuk menyimpan data master jenis surat.

| Field       | Type      | Notes                |
| ----------- | --------- | -------------------- |
| id          | bigint    | Primary key          |
| code        | varchar   | Kode surat           |
| name        | varchar   | Nama surat           |
| description | text      | Deskripsi surat      |
| sort_order  | int       | Urutan tampilan      |
| created_by  | bigint    | User pembuat data    |
| updated_by  | bigint    | User update terakhir |
| created_at  | timestamp | Waktu dibuat         |
| updated_at  | timestamp | Waktu update         |

Contoh data Master Takah:

| Code | Name                   |
| ---- | ---------------------- |
| SKET | Surat Keterangan       |
| SKK  | Surat Keterangan Kerja |
| SP   | Surat Peringatan       |
| SIK  | Surat Ijin Kegiatan    |
| UND  | Surat Undangan         |
| MEM  | Surat Memorandum       |
| ND   | Nota Dinas             |

Catatan:

* Field `sort_order` digunakan sebagai pengganti `order` agar tidak bentrok dengan keyword SQL.
* Field `created_by` dan `updated_by` berelasi dengan `users.id`.
* Master Takah digunakan sebagai acuan untuk Config Nomor Surat, Template Surat, dan Surat Keluar.

---

## 3. config_nomor_surat

Digunakan untuk konfigurasi format nomor surat otomatis.

| Field         | Type      | Notes                         |
| ------------- | --------- | ----------------------------- |
| id            | bigint    | Primary key                   |
| takah_id      | bigint    | Relasi master_takah           |
| company_code  | varchar   | Kode perusahaan               |
| division_code | varchar   | Kode divisi                   |
| reset_type    | varchar   | monthly / yearly              |
| last_number   | int       | Nomor terakhir yang digunakan |
| created_at    | timestamp | Waktu dibuat                  |
| updated_at    | timestamp | Waktu update                  |

Catatan:

* Field `takah_id` berelasi dengan `master_takah.id`.
* Field `last_number` memiliki nilai awal `0`.
* Field `last_number` akan bertambah otomatis saat Surat Keluar dibuat.
* Field `reset_type` digunakan untuk menentukan aturan reset nomor surat.
* Config Nomor Surat digunakan saat sistem melakukan generate nomor surat otomatis.

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
| updated_at    | timestamp | Waktu update        |

Catatan:

* Field `takah_id` berelasi dengan `master_takah.id`.
* Template surat digunakan sebagai dasar isi surat berdasarkan jenis surat.
* Template surat dapat memiliki parameter yang berbeda sesuai kebutuhan jenis surat.
* Template Surat berelasi dengan Parameter Surat.

---

## 5. parameter_surat

Digunakan untuk menyimpan daftar parameter yang digunakan pada setiap template surat.

Parameter surat memungkinkan setiap template memiliki field input yang berbeda sesuai kebutuhan surat.

| Field          | Type      | Notes                             |
| -------------- | --------- | --------------------------------- |
| id             | bigint    | Primary key                       |
| template_id    | bigint    | Relasi template_surat             |
| parameter_name | varchar   | Nama parameter                    |
| parameter_key  | varchar   | Key parameter pada template surat |
| input_type     | varchar   | text / textarea / date / number   |
| is_required    | boolean   | Parameter wajib diisi atau tidak  |
| created_at     | timestamp | Waktu dibuat                      |
| updated_at     | timestamp | Waktu update                      |

Catatan:

* Field `template_id` berelasi dengan `template_surat.id`.
* Parameter surat digunakan sebagai field input saat user membuat surat.
* Setiap template surat dapat memiliki lebih dari satu parameter.
* Parameter digunakan untuk menggantikan placeholder pada template surat.
* CRUD Parameter Surat telah terhubung dengan database MySQL.

Contoh parameter:

| Template         | Parameter Name   | Parameter Key    | Input Type |
| ---------------- | ---------------- | ---------------- | ---------- |
| Surat Undangan   | Nama Tujuan      | nama_tujuan      | text       |
| Surat Undangan   | Nama Kegiatan    | nama_kegiatan    | text       |
| Surat Undangan   | Tanggal Kegiatan | tanggal_kegiatan | date       |
| Surat Undangan   | Tempat Kegiatan  | tempat_kegiatan  | text       |
| Surat Keterangan | Nama Pemohon     | nama_pemohon     | text       |
| Surat Keterangan | Keperluan        | keperluan        | textarea   |

## 6. surat_keluar

Digunakan untuk menyimpan data surat keluar.

| Field | Type | Notes |
|------|------|------|
| id | bigint | Primary key |
| nomor_surat | varchar | Nomor surat otomatis |
| takah_id | bigint | Relasi Master Takah |
| tujuan_surat | varchar | Tujuan surat |
| perihal | varchar | Perihal surat |
| lampiran | varchar | Lampiran surat |
| tanggal_surat | date | Tanggal surat |
| file_surat | varchar | File surat |
| status | varchar | draft / pending / approved / rejected / completed |
| created_by | bigint | User pembuat |
| created_at | timestamp | Waktu dibuat |
| updated_at | timestamp | Waktu diperbarui |

Catatan:

* Field `takah_id` berelasi dengan `master_takah.id`.
* Nomor surat dibuat otomatis berdasarkan Config Nomor Surat.
* Status awal surat adalah `draft`.
* Setelah dikirim untuk review status berubah menjadi `pending`.
* Setelah proses approval selesai status akan berubah menjadi `approved`, `rejected`, atau `completed`.

---

## 7. approval_surat

Digunakan untuk menyimpan proses review dan persetujuan surat keluar.

| Field | Type | Notes |
|------|------|------|
| id | bigint | Primary key |
| surat_keluar_id | bigint | Relasi surat keluar |
| approver_id | bigint | User approver |
| approver_name | varchar | Nama approver |
| approval_status | varchar | pending / approved / rejected |
| notes | text | Catatan approval |
| approved_at | timestamp | Waktu approval |
| created_at | timestamp | Waktu dibuat |
| updated_at | timestamp | Waktu diperbarui |

Catatan:

* Field `surat_keluar_id` berelasi dengan `surat_keluar.id`.
* Field `approver_id` berelasi dengan `users.id`.
* Riwayat approval akan disimpan setiap proses review dilakukan.
* Status approval digunakan untuk memperbarui status surat keluar.

---

## 8. surat_masuk

Digunakan untuk menyimpan surat masuk dari pihak luar.

| Field | Type | Notes |
|------|------|------|
| id | bigint | Primary key |
| nomor_surat | varchar | Nomor surat |
| pengirim | varchar | Pengirim surat |
| penerima | varchar | Penerima surat |
| perihal | varchar | Perihal surat |
| file_surat | varchar | File surat |
| tanggal_surat | date | Tanggal surat |
| keterangan | text | Keterangan |
| status | varchar | received / completed |
| created_by | bigint | User input |
| created_at | timestamp | Waktu dibuat |
| updated_at | timestamp | Waktu diperbarui |

Catatan:

* Field `created_by` berelasi dengan `users.id`.
* Status awal surat masuk adalah `received`.
* Data surat masuk telah menggunakan CRUD berbasis MySQL.

---

## 9. monitoring_surat

Digunakan untuk memantau perkembangan seluruh proses surat.

| Field | Type | Notes |
|------|------|------|
| id | bigint | Primary key |
| surat_keluar_id | bigint | Relasi surat keluar |
| surat_masuk_id | bigint | Relasi surat masuk |
| nomor_surat | varchar | Nomor surat |
| status | varchar | Status terbaru |
| last_approver | varchar | Approver terakhir |
| last_notes | text | Catatan terakhir |
| updated_by | bigint | User terakhir mengubah |
| created_at | timestamp | Waktu dibuat |
| updated_at | timestamp | Waktu diperbarui |

Catatan:

* Monitoring dapat digunakan untuk surat keluar maupun surat masuk.
* Salah satu field `surat_keluar_id` atau `surat_masuk_id` dapat bernilai `NULL`.
* Monitoring menampilkan status terakhir dari proses surat.
* Data monitoring diperbarui setelah proses approval selesai.

---

# Table Relationship

```text
users
│
├── master_takah
├── surat_keluar
├── surat_masuk
├── approval_surat
└── monitoring_surat

master_takah
│
├── config_nomor_surat
├── template_surat
│
└── surat_keluar
    │
    ├── approval_surat
    └── monitoring_surat

template_surat
│
└── parameter_surat

surat_masuk
│
└── monitoring_surat
```

Relationship:

* `master_takah.created_by` → `users.id`
* `master_takah.updated_by` → `users.id`
* `config_nomor_surat.takah_id` → `master_takah.id`
* `template_surat.takah_id` → `master_takah.id`
* `parameter_surat.template_id` → `template_surat.id`
* `surat_keluar.takah_id` → `master_takah.id`
* `surat_keluar.created_by` → `users.id`
* `approval_surat.surat_keluar_id` → `surat_keluar.id`
* `approval_surat.approver_id` → `users.id`
* `surat_masuk.created_by` → `users.id`
* `monitoring_surat.surat_keluar_id` → `surat_keluar.id`
* `monitoring_surat.surat_masuk_id` → `surat_masuk.id`
* `monitoring_surat.updated_by` → `users.id`

---

# Generate Nomor Surat

Nomor surat dibuat secara otomatis berdasarkan konfigurasi pada tabel `config_nomor_surat`.

Format nomor surat:

```text
001/UND/CBN/062026
```

Komponen nomor surat:

* Nomor urut
* Kode surat
* Kode perusahaan
* Bulan dan tahun

Status implementasi:

* Generate nomor surat berjalan otomatis.
* Nomor terakhir disimpan pada field `last_number`.
* Nomor bertambah setiap Surat Keluar berhasil dibuat.
* Mendukung pengembangan reset nomor bulanan maupun tahunan.

---

# Current Status

Status database saat ini:

* Database `takah_db` telah dibuat.
* Seluruh tabel utama telah tersedia.
* Seluruh endpoint CRUD utama telah menggunakan MySQL.
* Relasi antar tabel telah diterapkan sesuai kebutuhan aplikasi.
* Query menggunakan package `database/sql`.
* Database dikembangkan sebagai dasar implementasi REST API Takah.

---

# Future Development

Pengembangan selanjutnya:

* Implementasi foreign key constraint secara penuh.
* Password hashing menggunakan bcrypt.
* Audit log aktivitas pengguna.
* Soft delete pada data tertentu.
* Upload file surat.
* Role based authorization.
* Optimasi query dan indexing database.