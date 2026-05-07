![Go](https://img.shields.io/badge/Go-1.21-blue)
![Gin](https://img.shields.io/badge/Framework-Gin-green)
![Status](https://img.shields.io/badge/status-development-yellow)

# 🚀 Takah API

Backend REST API sederhana untuk master data Takah (jenis surat) menggunakan Golang dan Gin Framework.

---

# 📦 Project Structure

```bash
takah-api/
├── cmd/
│   └── web/
│       └── main.go
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

# ⚙️ Configuration

Saat ini project masih menggunakan konfigurasi default:

- Port: `8080`
- Mode: `debug`
- Data: dummy data (belum database)

---

# 🚀 Run Application

## ▶️ Run server

```bash
go run cmd/web/main.go
```

Server akan berjalan di:

```bash
http://localhost:8080
```

---

# 📌 Available Endpoints

## Health
- GET `/api/v1/health`

## Auth
- POST `/api/v1/auth/login`

## Master Takah
- GET `/api/v1/takah`
- GET `/api/v1/takah/:id`
- POST `/api/v1/takah`
- PUT `/api/v1/takah/:id`
- DELETE `/api/v1/takah/:id`

## Surat
- GET `/api/v1/surat`
- GET `/api/v1/surat/:id`
- POST `/api/v1/surat`
- PUT `/api/v1/surat/:id`
- DELETE `/api/v1/surat/:id`

---

# 📡 API Specification

# 🔐 Auth

## Login

- Method: `POST`
- Endpoint: `/api/v1/auth/login`

### Request

```json
{
  "email": "admin@takah.com",
  "password": "password123"
}
```

### Response

```json
{
  "message": "Login berhasil",
  "token": "dummy-token"
}
```

---

# 📂 Master Takah API

## 📄 Get All Takah

- Method: `GET`
- Endpoint: `/api/v1/takah`

### Response

```json
{
  "message": "Data takah berhasil diambil",
  "data": []
}
```

---

## 🔍 Get Takah By ID

- Method: `GET`
- Endpoint: `/api/v1/takah/:id`

### Example

```bash
/api/v1/takah/1
```

---

## ➕ Create Takah

- Method: `POST`
- Endpoint: `/api/v1/takah`

### Request

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
    "id": 1,
    "code": "UND",
    "name": "Surat Undangan",
    "description": "Jenis surat undangan",
    "order": 1
  }
}
```

---

## ✏️ Update Takah

- Method: `PUT`
- Endpoint: `/api/v1/takah/:id`

### Request

```json
{
  "code": "SKET",
  "name": "Surat Keterangan",
  "description": "Update data takah",
  "order": 2
}
```

### Response

```json
{
  "message": "Data takah berhasil diupdate"
}
```

---

## ❌ Delete Takah

- Method: `DELETE`
- Endpoint: `/api/v1/takah/:id`

### Response

```json
{
  "message": "Data takah berhasil dihapus"
}
```

---

# 📄 Surat API

## 📄 Get All Surat

- Method: `GET`
- Endpoint: `/api/v1/surat`

### Response

```json
{
  "message": "success",
  "data": []
}
```

---

## ➕ Create Surat

- Method: `POST`
- Endpoint: `/api/v1/surat`

### Request

```json
{
  "nomor_surat": "003/ADM/2026",
  "judul": "Surat Permohonan Data",
  "pengirim": "Administrasi",
  "penerima": "HRD",
  "status": "draft"
}
```

### Response

```json
{
  "message": "Data surat berhasil dibuat",
  "data": {
    "id": 3,
    "nomor_surat": "003/ADM/2026",
    "judul": "Surat Permohonan Data",
    "pengirim": "Administrasi",
    "penerima": "HRD",
    "status": "draft"
  }
}
```

---

# 🧪 API Testing

Gunakan Postman atau browser:

```bash
http://localhost:8080/api/v1/takah
```

atau

```bash
http://localhost:8080/api/v1/surat
```

---

# 🚧 Development Status

Project masih dalam tahap development awal:

- ✅ Basic authentication API
- ✅ CRUD Master Takah API
- ✅ CRUD Surat API
- ✅ REST API routing dengan Gin
- ❌ Belum menggunakan database
