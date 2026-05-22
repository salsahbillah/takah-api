# Technical Backend - Takah API

Dokumen ini menjelaskan arsitektur backend, stack teknologi, endpoint API, struktur project, dan status pengembangan backend aplikasi Takah.

---

## Backend Stack

Backend Takah API menggunakan teknologi berikut:

- Golang
- Gin Framework
- REST API
- MySQL (planned)
- Laragon (planned for local database development)

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
│   │   ├── surat_handler.go
│   │   └── takah_handler.go
│   │
│   ├── model/
│   │   ├── auth_model.go
│   │   ├── surat_model.go
│   │   └── takah_model.go
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

- `route`
- `handler`
- `model`

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

| Module | Method | Endpoint | Handler |
| --- | --- | --- | --- |
| Health | GET | `/api/v1/health` | inline handler |
| Auth | POST | `/api/v1/auth/login` | `Login` |
| Master Takah | GET | `/api/v1/takah` | `GetAllTakah` |
| Master Takah | POST | `/api/v1/takah` | `CreateTakah` |
| Master Takah | GET | `/api/v1/takah/:id` | `GetTakahByID` |
| Master Takah | PUT | `/api/v1/takah/:id` | `UpdateTakah` |
| Master Takah | DELETE | `/api/v1/takah/:id` | `DeleteTakah` |
| Surat | GET | `/api/v1/surat` | `GetAllSurat` |
| Surat | POST | `/api/v1/surat` | `CreateSurat` |
| Surat | GET | `/api/v1/surat/:id` | `GetSuratByID` |
| Surat | PUT | `/api/v1/surat/:id` | `UpdateSurat` |
| Surat | DELETE | `/api/v1/surat/:id` | `DeleteSurat` |

---

## Handler Layer

Folder:

```text
internal/handler
```

Handler digunakan untuk menerima request dari client, memproses data sementara, dan mengembalikan response JSON.

File handler saat ini:

| File | Description |
| --- | --- |
| `auth_handler.go` | Handler login dummy |
| `takah_handler.go` | Handler CRUD Master Takah |
| `surat_handler.go` | Handler CRUD Surat dummy |

Catatan:

- Data Master Takah saat ini masih disimpan pada variable `takahData`.
- Data Surat saat ini masih menggunakan dummy data di dalam handler.
- Belum ada service layer dan repository layer.
- Belum terhubung ke database.

---

## Model Layer

Folder:

```text
internal/model
```

Model digunakan untuk mendefinisikan request dan response struct.

File model saat ini:

| File | Struct |
| --- | --- |
| `auth_model.go` | `LoginRequest`, `LoginResponse` |
| `takah_model.go` | `TakahRequest`, `TakahResponse` |
| `surat_model.go` | `SuratRequest`, `SuratResponse` |

Catatan:

- `LoginRequest` digunakan pada handler login.
- `LoginResponse` sudah tersedia, tetapi response login saat ini masih menggunakan `gin.H`.
- `TakahRequest` dan `TakahResponse` digunakan untuk CRUD Master Takah.
- `SuratRequest` dan `SuratResponse` digunakan untuk CRUD Surat.

---

## Authentication

Authentication saat ini masih menggunakan dummy login.

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
    "token": "dummy-token"
  }
}
```

### Failed Response

```json
{
  "message": "Email atau password salah"
}
```

Catatan:

- Belum menggunakan database user.
- Belum menggunakan password hashing.
- Belum menggunakan JWT.
- Token masih dummy.

---

## Master Takah

Master Takah adalah data master untuk jenis surat.

Contoh data:

| Code | Name |
| --- | --- |
| SKET | Surat Keterangan |
| SKK | Surat Keterangan Kerja |

Master Takah digunakan sebagai dasar untuk:

- Jenis surat
- Template surat
- Generate nomor surat
- Monitoring surat
- Filtering laporan surat

### Request Create / Update

```json
{
  "code": "UND",
  "name": "Surat Undangan",
  "description": "Jenis surat undangan",
  "order": 1
}
```

### Response

```json
{
  "message": "Data takah berhasil dibuat",
  "data": {
    "id": 3,
    "code": "UND",
    "name": "Surat Undangan",
    "description": "Jenis surat undangan",
    "order": 1,
    "created_by": "Admin",
    "created_time": "2026-05-04 15:00",
    "updated_by": "Admin",
    "updated_time": "2026-05-04 15:00"
  }
}
```

---

## Surat

Modul Surat saat ini masih berupa basic CRUD dummy yang dibuat pada tahap awal development.

Field utama:

- `nomor_surat`
- `judul`
- `pengirim`
- `penerima`
- `status`

Catatan:

- Modul ini masih dummy.
- Modul ini dapat dikembangkan menjadi Surat Masuk dan Surat Keluar.
- Fokus pengembangan selanjutnya diarahkan ke Master Takah, Config Nomor Surat, Template Surat, Surat Masuk, Surat Keluar, dan Monitoring Surat.

---

## Database Plan

Database belum terintegrasi.

Database yang direncanakan:

- `users`
- `master_takah`
- `config_nomor_surat`
- `template_surat`
- `surat_keluar`
- `surat_masuk`
- `monitoring_surat`

Database engine yang direncanakan:

- MySQL
- Laragon for local development

---

## Generate Nomor Surat Plan

Format nomor surat yang direncanakan:

```text
001/UND/CBN/052026
```

Keterangan:

- `001` = nomor urut
- `UND` = kode jenis surat dari Master Takah
- `CBN` = kode perusahaan/divisi
- `052026` = bulan dan tahun

Aturan:

- Nomor surat bertambah otomatis.
- Nomor surat reset setiap bulan dan tahun.

---

## Current Development Status

| Feature | Status | Notes |
| --- | --- | --- |
| Setup Golang + Gin | Done | Project sudah berjalan lokal |
| Route API v1 | Done | Base route `/api/v1` |
| Health Check | Done | `GET /api/v1/health` |
| Login dummy | Done | Belum JWT |
| CRUD Master Takah | Done | Masih dummy data |
| CRUD Surat | Done | Masih dummy data |
| Database integration | Todo | Planned MySQL |
| JWT Authentication | Todo | Planned |
| Service layer | Todo | Planned |
| Repository layer | Todo | Planned |

---

## Future Development

Pengembangan backend selanjutnya:

- Integrasi MySQL
- Membuat SQL schema
- Membuat repository layer
- Membuat service layer
- JWT authentication
- Request validation lebih lengkap
- Helper generate nomor surat otomatis
- Upload file surat
- Template surat
- Monitoring surat
- Export PDF
- API documentation Swagger/OpenAPI
- Logging dan error handling