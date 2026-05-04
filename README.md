![Go](https://img.shields.io/badge/Go-1.21-blue)
![Gin](https://img.shields.io/badge/Framework-Gin-green)
![Status](https://img.shields.io/badge/status-development-yellow)
# 🚀 Takah API

Backend API sederhana untuk sistem manajemen surat menggunakan Golang (Gin Framework).

---

## 📦 Project Structure

```
takah-api/
├── cmd/web/main.go        # Entry point aplikasi
├── internal/
│   ├── handler/           # Handler (logic API)
│   ├── model/             # Struct request & response
│   └── route/             # Routing API
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Configuration

Saat ini project masih menggunakan konfigurasi default:

* Port: `8080`
* Mode: `debug`
* Data: masih dummy (belum database)

---

## 🚀 Run Application

### ▶️ Run server

```bash
go run cmd/web/main.go
```

Server akan berjalan di:

```
http://localhost:8080
```

---

## 📌 Available Endpoints

---

- POST /api/v1/auth/login
- GET /api/v1/surat
- GET /api/v1/surat/:id
- POST /api/v1/surat
  
---

## 📡 API Spec
---

### 🔐 Auth

#### Login

* Method: `POST`
* Endpoint: `/api/v1/auth/login`

**Request:**

```json
{
  "email": "admin@takah.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "message": "Login berhasil",
  "token": "dummy-token"
}
```

---

### 📄 Surat

#### Get All Surat

* Method: `GET`
* Endpoint: `/api/v1/surat`

**Response:**

```json
{
  "message": "success",
  "data": []
}
```

---

#### Get Surat by ID

* Method: `GET`
* Endpoint: `/api/v1/surat/:id`

**Example:**

```
/api/v1/surat/1
```

---

#### Create Surat

* Method: `POST`
* Endpoint: `/api/v1/surat`

**Request:**

```json
{
  "nomor_surat": "003/ADM/2026",
  "judul": "Surat Permohonan Data",
  "pengirim": "Administrasi",
  "penerima": "HRD",
  "status": "draft"
}
```

**Response:**

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

## 🧪 Testing API

Gunakan Postman atau browser:

```
http://localhost:8080/api/v1/surat
```

---

## 🚧 Status

Project masih dalam tahap development awal:
- Basic API sudah dibuat (Auth & Surat)
- Data masih menggunakan dummy
- Database belum terintegrasi

---
