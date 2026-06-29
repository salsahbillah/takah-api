# Technical Backend - Takah API

Dokumen ini menjelaskan arsitektur backend, stack teknologi, endpoint API, struktur project, dan status pengembangan backend aplikasi Takah.

---

## Backend Stack

Backend Takah API menggunakan teknologi berikut:

* Golang
* Gin Framework
* REST API
* JWT Authentication
* MySQL
* database/sql
* Laragon sebagai local database server

---

## Project Structure

```text
takah-api/
├── cmd/
│   └── web/
│
├── database/
│
├── docs/
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── handler/
│   ├── helper/
│   ├── middleware/
│   ├── model/
│   └── route/
│
├── .env
├── go.mod
├── go.sum
└── README.md
```

| Folder | Kegunaan |
|---------|----------|
| `cmd` | Entry point aplikasi |
| `database` | SQL Schema database |
| `docs` | Dokumentasi sistem |
| `config` | Konfigurasi aplikasi |
| `internal/database` | Koneksi MySQL |
| `handler` | CRUD endpoint API |
| `helper` | Fungsi pendukung |
| `middleware` | JWT Authentication |
| `model` | Request & Response struct |
| `route` | Routing endpoint |

---

## Backend Architecture

Backend menggunakan arsitektur sederhana berbasis REST API dengan pemisahan layer agar proses pengembangan lebih terstruktur.

Layer yang digunakan:

* Config
* Database
* Route
* Handler
* Model
* Helper
* Middleware

Setiap request akan diterima oleh Route, diteruskan ke Handler, diproses menggunakan Helper maupun Database, kemudian menghasilkan response JSON kepada client.

---

## Config Layer

Folder

```text
internal/config
```

Layer ini digunakan untuk menyimpan konfigurasi aplikasi.

Fungsi:

* Membaca konfigurasi dari file `.env`
* Menyimpan konfigurasi database
* Menyimpan konfigurasi JWT Secret
* Menyediakan konfigurasi host, port, username, password dan nama database

---

## Database Layer

Folder

```text
internal/database
```

Layer database bertugas menghubungkan backend dengan database MySQL.

Fungsi:

* Membuka koneksi database
* Mengecek status koneksi
* Menyediakan object database untuk seluruh handler
* Menjalankan query CRUD

Database yang digunakan:

```text
takah_db
```

Status

* Database berhasil dibuat.
* Koneksi MySQL berhasil diimplementasikan.
* Seluruh endpoint utama telah menggunakan database MySQL.

---

## Route Layer

Folder

```text
internal/route
```

Base URL

```text
/api/v1
```

Endpoint yang tersedia

| Module | Endpoint |
|---------|----------|
| Health | GET `/health` |
| Authentication | POST `/auth/login` |
| Master Takah | CRUD `/takah` |
| Config Nomor Surat | CRUD `/config-nomor` |
| Template Surat | CRUD `/template-surat` |
| Surat Keluar | CRUD `/surat-keluar` |
| Approval Surat | CRUD `/approval` |
| Monitoring Surat | CRUD `/monitoring` |
| Surat Masuk | CRUD `/surat-masuk` |
| Surat | CRUD `/surat` |

Semua endpoint selain login menggunakan JWT Authentication melalui middleware.

---

## Middleware Layer

Folder

```text
internal/middleware
```

Middleware digunakan untuk mengamankan endpoint API menggunakan JWT.

Fungsi:

* Membaca Authorization Header
* Memvalidasi JWT Token
* Mengambil informasi user dari token
* Membatasi akses endpoint yang dilindungi

Seluruh endpoint selain Login berada pada protected route sehingga hanya dapat diakses setelah user berhasil melakukan autentikasi.

---

## Helper Layer

Folder

```text
internal/helper
```

Helper digunakan sebagai fungsi pendukung yang dapat dipanggil oleh beberapa module.

Fungsi utama:

* Generate nomor surat otomatis
* Formatting nomor surat
* Generate kode berdasarkan konfigurasi
* Utility function lainnya

Contoh hasil generate nomor surat

```text
001/SKET/CBN/062026
```

Generate nomor surat dilakukan secara otomatis ketika data Surat Keluar dibuat berdasarkan konfigurasi pada Config Nomor Surat.

---

## Handler Layer

Folder

```text
internal/handler
```

Handler bertugas menerima request dari client, menjalankan proses bisnis, melakukan query database, kemudian mengembalikan response JSON.

File handler

| File | Fungsi |
|------|--------|
| auth_handler.go | Login Authentication |
| takah_handler.go | CRUD Master Takah |
| config_nomor_handler.go | CRUD Config Nomor Surat |
| template_surat_handler.go | CRUD Template Surat |
| surat_keluar_handler.go | CRUD Surat Keluar |
| approval_handler.go | CRUD Approval Surat |
| monitoring_handler.go | CRUD Monitoring Surat |
| surat_masuk_handler.go | CRUD Surat Masuk |
| surat_handler.go | CRUD Surat Dummy |

Status

* Seluruh handler utama telah menggunakan MySQL.
* Seluruh endpoint menghasilkan response JSON.
* Query menggunakan package `database/sql`.
* Seluruh proses CRUD telah berjalan dengan baik.

---

## Model Layer

Folder

```text
internal/model
```

Model digunakan untuk mendefinisikan struktur Request dan Response pada setiap endpoint.

Model yang tersedia

| File |
|------|
| auth_model.go |
| takah_model.go |
| config_nomor_model.go |
| template_surat_model.go |
| surat_keluar_model.go |
| approval_model.go |
| monitoring_model.go |
| surat_masuk_model.go |


Model digunakan sebagai validasi request sekaligus struktur response yang dikirimkan kepada frontend.

## Authentication

Authentication menggunakan JWT (JSON Web Token) untuk mengamankan endpoint API.

### Endpoint

```http
POST /api/v1/auth/login
```

### Request

```json
{
  "email": "admin@takah.com",
  "password": "password123"
}
```

### Success Response

```json
{
  "message": "Login berhasil",
  "data": {
    "token": "jwt-token"
  }
}
```

Fitur yang telah diimplementasikan:

* Login menggunakan email dan password.
* JWT Token berhasil dibuat setelah proses login.
* Middleware melakukan validasi token pada setiap request.
* Informasi role disimpan di dalam JWT.
* Endpoint protected hanya dapat diakses setelah login.

Catatan:

* Password masih menggunakan plaintext.
* Password hashing (bcrypt) belum diimplementasikan.
* Data user masih menggunakan tabel sederhana.
* Pengelolaan user akan dikembangkan pada tahap berikutnya.

---

## Master Takah

Master Takah merupakan data utama yang menyimpan jenis surat yang digunakan pada sistem.

Contoh data:

| Code | Nama Surat |
|------|---------------------------|
| SKET | Surat Keterangan |
| SKK | Surat Keterangan Kerja |
| SP | Surat Peringatan |
| UND | Surat Undangan |
| MEM | Memorandum |
| ND | Nota Dinas |

Master Takah digunakan oleh beberapa module:

* Config Nomor Surat
* Template Surat
* Surat Keluar

Fitur:

* CRUD Master Takah
* Relasi dengan Config Nomor Surat
* Relasi dengan Template Surat
* Relasi dengan Surat Keluar

---

## Config Nomor Surat

Config Nomor Surat digunakan sebagai acuan pembuatan nomor surat otomatis.

Contoh format:

```text
001/SKET/CBN/062026
```

Komponen nomor surat:

* Nomor urut
* Kode surat
* Kode perusahaan
* Bulan
* Tahun

Fitur:

* CRUD Config Nomor Surat
* Relasi dengan Master Takah
* Menyimpan Company Code
* Menyimpan Division Code
* Menyimpan tipe reset nomor surat
* Menyimpan nomor terakhir (Last Number)

Generate nomor surat menggunakan data pada Config Nomor Surat sehingga setiap jenis surat memiliki format nomor yang berbeda.

---

## Template Surat

Template Surat digunakan untuk menyimpan format isi surat berdasarkan jenis surat.

Relasi:

```text
Master Takah
      │
      ▼
Template Surat
```

Data utama:

* Jenis Surat
* Nama Template
* Isi Template

Fitur:

* CRUD Template Surat
* Relasi dengan Master Takah
* Menyimpan isi template
* Digunakan saat pembuatan Surat Keluar

---

## Generate Nomor Surat

Nomor surat dibuat secara otomatis ketika Surat Keluar berhasil dibuat.

Contoh:

```text
001/UND/CBN/062026
```

Proses generate:

```text
User membuat Surat Keluar
        │
        ▼
Sistem membaca Config Nomor Surat
        │
        ▼
Mengambil Last Number
        │
        ▼
Generate Nomor Surat
        │
        ▼
Update Last Number
        │
        ▼
Data Surat Keluar disimpan
```

Fitur:

* Generate otomatis.
* Increment nomor surat.
* Mendukung reset bulanan.
* Mendukung reset tahunan.
* Menggunakan transaction database agar nomor surat tidak duplikat.

---

## Surat Keluar

Surat Keluar digunakan untuk membuat surat yang akan dikirim kepada pihak lain.

Data utama:

* Nomor Surat
* Jenis Surat
* Tujuan Surat
* Perihal
* Lampiran
* File Surat
* Tanggal Surat
* Status

Status Surat Keluar:

```text
draft
pending
approved
rejected
completed
```

Flow Surat Keluar

```text
Pilih Jenis Surat
        │
        ▼
Generate Nomor Surat
        │
        ▼
Simpan Surat Keluar
        │
        ▼
Status Draft
        │
        ▼
Kirim Approval
        │
        ▼
Status Pending
        │
        ▼
Approve / Reject
        │
        ▼
Status Surat Berubah
        │
        ▼
Monitoring Surat
```

Fitur:

* CRUD Surat Keluar
* Generate nomor surat otomatis
* Menggunakan transaksi database
* Mengupdate nomor terakhir pada Config Nomor Surat
* Status surat berubah mengikuti proses approval
## Approval Surat

Approval Surat digunakan untuk melakukan proses review dan persetujuan terhadap Surat Keluar sebelum surat dinyatakan selesai diproses.

Data utama:

* Surat Keluar
* Nomor Surat
* Approver
* Status Approval
* Catatan Review
* Waktu Approval

Status Approval

```text
pending
approved
rejected
```

Flow Approval

```text
Surat Keluar
      │
      ▼
Create Approval
      │
      ▼
Status Pending
      │
      ▼
Approver Review
      │
      ▼
Approve / Reject
      │
      ▼
Update Status Surat Keluar
      │
      ▼
Update Monitoring
```

Fitur:

* CRUD Approval Surat
* Menyimpan data approver
* Menyimpan catatan review
* Menyimpan waktu approval
* Mengubah status Surat Keluar secara otomatis
* Menjadi sumber informasi untuk Monitoring Surat

---

## Surat Masuk

Surat Masuk digunakan untuk mencatat seluruh surat yang diterima dari pihak luar perusahaan atau instansi.

Data utama:

* Nomor Surat
* Pengirim
* Penerima
* Perihal
* File Surat
* Tanggal Surat
* Keterangan
* Status

Status Surat Masuk

```text
received
completed
```

Fitur:

* CRUD Surat Masuk
* Menyimpan data surat yang diterima
* Menyimpan informasi pengirim dan penerima
* Menyimpan file surat
* Menyimpan keterangan surat
* Digunakan sebagai salah satu sumber data Monitoring Surat

---

## Monitoring Surat

Monitoring Surat digunakan untuk memantau perkembangan proses surat secara keseluruhan.

Data Monitoring berasal dari:

* Surat Keluar
* Approval Surat
* Surat Masuk

Informasi yang ditampilkan:

* Nomor Surat
* Status Surat
* Approver Terakhir
* Catatan Approval Terakhir
* Waktu Update

Status Monitoring

Untuk Surat Keluar

```text
draft
pending
approved
rejected
completed
```

Untuk Surat Masuk

```text
received
completed
```

Flow Monitoring

```text
Surat Keluar
      │
      ├──────────────┐
      ▼              │
Approval Surat       │
      │              │
      └──────┐       │
             ▼       ▼
        Monitoring Surat
```

Fitur:

* CRUD Monitoring
* Menampilkan status terbaru surat
* Menampilkan approver terakhir
* Menampilkan catatan approval terakhir
* Menampilkan waktu update terakhir
* Membaca data dari Surat Keluar, Approval Surat, dan Surat Masuk

---

## Database

Database yang digunakan

```text
MySQL
```

Database Server

```text
Laragon
```

Nama Database

```text
takah_db
```

Tabel yang digunakan

* users
* master_takah
* config_nomor_surat
* template_surat
* surat_keluar
* approval_surat
* surat_masuk
* monitoring_surat

Relasi utama

```text
users
│
├── master_takah
├── surat_keluar
├── approval_surat
├── surat_masuk
└── monitoring_surat

master_takah
│
├── config_nomor_surat
├── template_surat
└── surat_keluar
      │
      ├── approval_surat
      └── monitoring_surat

surat_masuk
│
└── monitoring_surat
```

Status Database

* Database MySQL berhasil dibuat.
* Koneksi database berhasil diimplementasikan.
* Seluruh endpoint utama telah menggunakan MySQL.
* Seluruh proses CRUD berjalan menggunakan query SQL.
* Transaction digunakan pada proses Generate Nomor Surat.

---

## Current Development Status

| Feature | Status | Notes |
|---------|--------|--------------------------------|
| Setup Golang + Gin | Done | Project berjalan normal |
| Route API v1 | Done | Seluruh endpoint tersedia |
| Health Check | Done | Endpoint berjalan normal |
| JWT Authentication | Done | Login dan Middleware berhasil |
| Master Takah | Done | CRUD MySQL |
| Config Nomor Surat | Done | CRUD MySQL |
| Template Surat | Done | CRUD MySQL |
| Generate Nomor Surat | Done | Otomatis menggunakan Helper |
| Surat Keluar | Done | CRUD + Generate Nomor Surat |
| Approval Surat | Done | CRUD + Update Status Surat |
| Surat Masuk | Done | CRUD MySQL |
| Monitoring Surat | Done | CRUD MySQL |
| Database MySQL | Done | Terhubung |
| SQL Schema | Done | Seluruh tabel telah dibuat |
| Database Integration | Done | Seluruh modul telah menggunakan MySQL |
| Service Layer | Planned | Belum diimplementasikan |
| Repository Layer | Planned | Belum diimplementasikan |

---

## Current Backend Status

Status backend saat ini:

* Struktur project telah selesai disusun.
* Backend menggunakan Golang dan Gin Framework.
* Seluruh endpoint utama telah tersedia.
* JWT Authentication telah berhasil diimplementasikan.
* Middleware berhasil mengamankan endpoint.
* Seluruh modul CRUD telah menggunakan database MySQL.
* Generate nomor surat otomatis telah berjalan.
* Approval Surat telah terintegrasi dengan Surat Keluar.
* Monitoring Surat telah terintegrasi dengan Surat Keluar, Approval Surat, dan Surat Masuk.
* Konfigurasi database menggunakan file `.env`.
* Backend siap digunakan sebagai REST API untuk frontend.

---

## Development Notes

Backend Takah API dikembangkan menggunakan arsitektur REST API sederhana dengan pemisahan layer Config, Database, Route, Handler, Helper, Middleware, dan Model agar struktur project lebih terorganisir dan mudah dikembangkan.

Seluruh modul utama seperti Master Takah, Config Nomor Surat, Template Surat, Surat Keluar, Approval Surat, Monitoring Surat, dan Surat Masuk telah berhasil diimplementasikan menggunakan database MySQL. Setiap endpoint telah mendukung operasi CRUD serta menghasilkan response dalam format JSON yang siap digunakan oleh frontend.

Implementasi Generate Nomor Surat telah memanfaatkan helper dan transaction database sehingga nomor surat dapat dibuat secara otomatis berdasarkan konfigurasi yang tersimpan pada Config Nomor Surat. Selain itu, proses Approval Surat telah terintegrasi dengan Surat Keluar sehingga perubahan status approval akan memperbarui status surat secara otomatis dan dapat dimonitor melalui modul Monitoring Surat.

Dokumentasi ini menjadi acuan pengembangan backend selanjutnya, seperti implementasi Service Layer, Repository Layer, upload file surat, manajemen user, password hashing, serta pengembangan fitur lainnya untuk mendukung kebutuhan aplikasi Takah secara menyeluruh.