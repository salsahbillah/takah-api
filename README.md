![Go](https://img.shields.io/badge/Go-1.21-blue)
![Gin](https://img.shields.io/badge/Framework-Gin-green)
![Status](https://img.shields.io/badge/status-development-yellow)

# 🚀 Takah API

Backend REST API sederhana untuk master data Takah menggunakan Golang dan Gin Framework.

Project ini dibuat untuk pembelajaran dan pengembangan backend API dengan implementasi basic CRUD, routing REST API, dan dokumentasi teknis awal.

---

## 📦 Project Structure

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

## ⚙️ Configuration

Saat ini project masih menggunakan konfigurasi default:

- Port: `8080`
- Mode: `debug`
- Data: dummy data
- Database: belum terintegrasi

---

## 🚀 Run Application

```bash
go run cmd/web/main.go
```

Server akan berjalan di:

```bash
http://localhost:8080
```

---

## 📌 Available Endpoints

### Health
- GET `/api/v1/health`

### Auth
- POST `/api/v1/auth/login`

### Master Takah
- GET `/api/v1/takah`
- GET `/api/v1/takah/:id`
- POST `/api/v1/takah`
- PUT `/api/v1/takah/:id`
- DELETE `/api/v1/takah/:id`

### Surat
- GET `/api/v1/surat`
- GET `/api/v1/surat/:id`
- POST `/api/v1/surat`
- PUT `/api/v1/surat/:id`
- DELETE `/api/v1/surat/:id`

---

## 📡 API Example

### Login

**POST** `/api/v1/auth/login`

Request:

```json
{
  "email": "admin@takah.com",
  "password": "password123"
}
```

Response:

```json
{
  "message": "Login berhasil",
  "data": {
    "token": "dummy-token"
  }
}
```

---

### Create Master Takah

**POST** `/api/v1/takah`

Request:

```json
{
  "code": "UND",
  "name": "Surat Undangan",
  "description": "Jenis surat undangan",
  "order": 1
}
```

Response:

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

## 📚 Documentation

Dokumentasi teknis tersedia pada folder `docs/`:

- [Technical Backend](./docs/backend.md)
- [Flow System](./docs/flow-system.md)
- [Database Design](./docs/database-design.md)
- [Integration Documentation](./docs/integration.md)
- [API Testing](./docs/api-testing.md)
- [SOP Admin](./docs/sop-admin.md)
- [SOP User](./docs/sop-user.md)
- [Task List & Roadmap](./docs/task-list.md)

---

## 🧪 API Testing

Gunakan Postman atau browser:

```bash
http://localhost:8080/api/v1/takah
```

atau:

```bash
http://localhost:8080/api/v1/surat
```

---

## 🚧 Development Status

Project masih dalam tahap development awal:

- ✅ Basic authentication API
- ✅ CRUD Master Takah API
- ✅ CRUD Surat API
- ✅ REST API routing dengan Gin
- ✅ Technical documentation
- ❌ Belum menggunakan database
- ❌ Belum menggunakan JWT authentication
