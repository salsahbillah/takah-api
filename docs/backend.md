# Technical Backend - Takah API

Dokumen ini menjelaskan arsitektur backend, stack teknologi, endpoint API, struktur project, dan status pengembangan backend aplikasi Takah.

---

## Backend Stack

Backend Takah API menggunakan teknologi berikut:

* Golang
* Gin Framework
* REST API
* JWT Authentication
* MySQL (planned)
* Laragon (planned for local database development)

---

## Project Structure

```bash
takah-api/
├── cmd/
│   └── web/
│       └── main.go
│
├── docs/
│   ├── api-testing.md
│   ├── backend.md
│   ├── database-design.md
│   ├── flow-system.md
│   ├── integration.md
│   ├── sop-admin.md
│   ├── sop-user.md
│   └── task-list.md
│
├── internal/
│   ├── handler/
│   │   ├── auth_handler.go
│   │   ├── takah_handler.go
│   │   ├── surat_handler.go
│   │   ├── surat_keluar_handler.go
│   │   ├── config_nomor_handler.go
│   │   ├── approval_handler.go
│   │   └── monitoring_handler.go
│   │
│   ├── helper/
│   │   └── nomor_surat.go
│   │
│   ├── middleware/
│   │   └── auth_middleware.go
│   │
│   ├── model/
│   │   ├── auth_model.go
│   │   ├── takah_model.go
│   │   ├── surat_model.go
│   │   ├── surat_keluar_model.go
│   │   ├── config_nomor_model.go
│   │   ├── approval_model.go
│   │   └── monitoring_model.go
│   │
│   └── route/
│       └── route.go
│
├── go.mod
├── go.sum
└── README.md
```

---

## Backend Architecture

Backend saat ini menggunakan struktur sederhana berbasis:

* route
* handler
* model
* helper
* middleware

Struktur ini digunakan untuk tahap awal development agar endpoint dan alur request-response mudah dipahami.

---

## Route Layer

Folder:

```text
internal/route
```

Route layer digunakan untuk mendefinisikan endpoint API.

File utama:

```text
internal/route/route.go
```

Base route:

```text
/api/v1
```

Route yang tersedia saat ini:

| Module             | Method | Endpoint               |
| ------------------ | ------ | ---------------------- |
| Health             | GET    | `/api/v1/health`       |
| Auth               | POST   | `/api/v1/auth/login`   |
| Master Takah       | CRUD   | `/api/v1/takah`        |
| Surat Dummy        | CRUD   | `/api/v1/surat`        |
| Config Nomor Surat | CRUD   | `/api/v1/config-nomor` |
| Surat Keluar       | CRUD   | `/api/v1/surat-keluar` |
| Approval Surat     | CRUD   | `/api/v1/approval`     |
| Monitoring Surat   | CRUD   | `/api/v1/monitoring`   |

---

## Middleware Layer

Folder:

```text
internal/middleware
```

Middleware digunakan untuk memvalidasi token JWT sebelum user mengakses endpoint yang dilindungi.

File:

```text
auth_middleware.go
```

Fungsi:

* Validasi JWT token
* Membatasi akses endpoint
* Menyimpan data user dari token ke context request

---

## Helper Layer

Folder:

```text
internal/helper
```

Helper digunakan untuk fungsi pendukung yang dapat digunakan oleh beberapa module.

File:

```text
nomor_surat.go
```

Fungsi:

* Generate nomor surat otomatis
* Membuat format nomor surat berdasarkan konfigurasi
* Digunakan pada proses pembuatan surat keluar

Contoh:

```text
001/SKET/CBN/062026
```

---

## Handler Layer

Folder:

```text
internal/handler
```

Handler digunakan untuk menerima request dari client, memproses data sementara, dan mengembalikan response JSON.

File handler saat ini:

| File                    | Description             |
| ----------------------- | ----------------------- |
| auth_handler.go         | Login authentication    |
| takah_handler.go        | CRUD Master Takah       |
| surat_handler.go        | CRUD Surat dummy        |
| config_nomor_handler.go | CRUD Config Nomor Surat |
| surat_keluar_handler.go | CRUD Surat Keluar       |
| approval_handler.go     | CRUD Approval Surat     |
| monitoring_handler.go   | CRUD Monitoring Surat   |

Catatan:

* Data masih menggunakan dummy data.
* Belum menggunakan database.
* Belum menggunakan service layer.
* Belum menggunakan repository layer.

---

## Model Layer

Folder:

```text
internal/model
```

Model digunakan untuk mendefinisikan request dan response struct.

Model yang tersedia:

| File                  |
| --------------------- |
| auth_model.go         |
| takah_model.go        |
| surat_model.go        |
| config_nomor_model.go |
| surat_keluar_model.go |
| approval_model.go     |
| monitoring_model.go   |

---

## Authentication

Authentication saat ini menggunakan JWT token.

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

Catatan:

* Menggunakan JWT token.
* Middleware authentication sudah tersedia.
* Password hashing belum diimplementasikan.
* Data user masih menggunakan dummy data.
* Database user belum tersedia.

---

## Master Takah

Master Takah digunakan sebagai data master jenis surat.

Contoh:

| Code | Name                   |
| ---- | ---------------------- |
| SKET | Surat Keterangan       |
| SKK  | Surat Keterangan Kerja |
| UND  | Surat Undangan         |
| SP   | Surat Peringatan       |

Master Takah digunakan oleh:

* Config Nomor Surat
* Generate Nomor Surat
* Surat Keluar
* Approval Surat
* Monitoring Surat

---

## Config Nomor Surat

Config nomor surat digunakan untuk menentukan format nomor surat otomatis.

Relasi:

```text
Config Nomor Surat
↓
Master Takah
```

Contoh:

```text
001/SKET/CBN/062026
```

Fitur:

* CRUD Config Nomor Surat
* Relasi dengan Master Takah
* Menyimpan kode perusahaan
* Menyimpan tipe reset nomor surat
* Menyimpan nomor terakhir

---

## Generate Nomor Surat

Generate nomor surat dilakukan secara otomatis saat surat keluar dibuat.

Format:

```text
001/SKET/CBN/062026
```

Komponen:

* Nomor urut
* Kode surat
* Kode perusahaan
* Bulan dan tahun

Aturan:

* Nomor bertambah otomatis.
* Berdasarkan konfigurasi nomor surat.
* Mendukung reset bulanan dan tahunan (planned improvement).

---

## Surat Keluar

Surat keluar digunakan untuk membuat surat yang dikirim keluar perusahaan.

Data utama:

* Nomor surat
* Jenis surat
* Tujuan surat
* Perihal
* Lampiran
* Tanggal surat
* Status

Status:

```text
draft
pending
approved
rejected
completed
```

Flow:

```text
Pilih jenis surat
↓
Generate nomor surat
↓
Simpan draft
↓
Approval
↓
Monitoring
```

---

## Approval Surat

Approval surat digunakan untuk proses review dan persetujuan surat.

Data utama:

* Surat keluar
* Approver
* Status approval
* Catatan approval

Status:

```text
pending
approved
rejected
```

Tujuan:

* Mengetahui siapa yang melakukan review
* Menyimpan catatan approval
* Menyimpan riwayat approval

---

## Monitoring Surat

Monitoring surat digunakan untuk melihat status dan riwayat surat.

Informasi:

* Nomor surat
* Status surat
* Approver terakhir
* Catatan approval terakhir
* Waktu update

Tujuan:

* Monitoring status surat
* Tracking proses approval
* Melihat riwayat review

---

## Database Plan

Database belum terintegrasi.

Database yang direncanakan:

* users
* master_takah
* config_nomor_surat
* template_surat
* surat_keluar
* approval_surat
* surat_masuk
* monitoring_surat

Database engine:

* MySQL
* Laragon

---

## Current Development Status

| Feature                 | Status | Notes                   |
| ----------------------- | ------ | ----------------------- |
| Setup Golang + Gin      | Done   | Project berjalan normal |
| Route API v1            | Done   | Base route tersedia     |
| Health Check            | Done   | Berjalan normal         |
| JWT Authentication      | Done   | Middleware tersedia     |
| CRUD Master Takah       | Done   | Dummy data              |
| CRUD Surat              | Done   | Dummy data              |
| CRUD Config Nomor Surat | Done   | Relasi Master Takah     |
| Generate Nomor Surat    | Done   | Helper tersedia         |
| CRUD Surat Keluar       | Done   | Berjalan normal         |
| CRUD Approval Surat     | Done   | Berjalan normal         |
| CRUD Monitoring Surat   | Done   | Berjalan normal         |
| Database Integration    | Todo   | Planned MySQL           |
| Template Surat          | Todo   | Planned                 |
| Surat Masuk             | Todo   | Planned                 |
| Service Layer           | Todo   | Planned                 |
| Repository Layer        | Todo   | Planned                 |

---

## Future Development

Pengembangan backend selanjutnya:

* Integrasi MySQL
* SQL schema
* Repository layer
* Service layer
* Password hashing
* Authorization role
* Template surat
* Surat masuk
* Upload file surat
