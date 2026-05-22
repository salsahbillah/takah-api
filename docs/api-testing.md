# API Testing - Takah API

Dokumen ini digunakan untuk mencatat proses testing endpoint API pada aplikasi Takah.

Testing dilakukan menggunakan:
- Postman
- Browser
- Localhost environment

---

# Base URL

```text
http://localhost:8080/api/v1
```

---

# Health Check

## Endpoint

```http
GET /health
```

## Example Request

```text
GET http://localhost:8080/api/v1/health
```

## Success Response

```json
{
  "message": "API Takah berjalan"
}
```

## Status

- Success
- Endpoint berjalan normal

---

# Authentication Testing

## Login

### Endpoint

```http
POST /auth/login
```

### Example Request

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

### Status

- Success
- Dummy authentication berjalan normal

---

# Master Takah Testing

## Get All Takah

### Endpoint

```http
GET /takah
```

### Status

- Success

---

## Get Takah By ID

### Endpoint

```http
GET /takah/:id
```

### Example

```text
GET /takah/1
```

### Status

- Success

---

## Create Takah

### Endpoint

```http
POST /takah
```

### Example Request

```json
{
  "code": "UND",
  "name": "Surat Undangan",
  "description": "Jenis surat undangan",
  "order": 1
}
```

### Status

- Success

---

## Update Takah

### Endpoint

```http
PUT /takah/:id
```

### Example

```text
PUT /takah/1
```

### Status

- Success

---

## Delete Takah

### Endpoint

```http
DELETE /takah/:id
```

### Example

```text
DELETE /takah/1
```

### Status

- Success

---

# Surat Testing

## Get All Surat

### Endpoint

```http
GET /surat
```

### Status

- Success

---

## Get Surat By ID

### Endpoint

```http
GET /surat/:id
```

### Status

- Success

---

## Create Surat

### Endpoint

```http
POST /surat
```

### Example Request

```json
{
  "nomor_surat": "001/UND/2026",
  "judul": "Surat Undangan Meeting",
  "pengirim": "HRD",
  "penerima": "Staff",
  "status": "draft"
}
```

### Status

- Success

---

## Update Surat

### Endpoint

```http
PUT /surat/:id
```

### Status

- Success

---

## Delete Surat

### Endpoint

```http
DELETE /surat/:id
```

### Status

- Success

---

# Current Testing Status

| Module | Status | Notes |
| --- | --- | --- |
| Health Check | Done | Endpoint berjalan |
| Login | Done | Dummy authentication |
| Master Takah | Done | CRUD berjalan |
| Surat | Done | CRUD berjalan |
| Database Testing | Todo | Database belum tersedia |
| JWT Testing | Todo | JWT belum tersedia |

---

# Future Testing Plan

Testing yang direncanakan selanjutnya:

- Database integration testing
- JWT authentication testing
- Validation testing
- Error handling testing
- File upload testing
- Generate nomor surat testing
- Monitoring surat testing
- Load testing
- API documentation testing