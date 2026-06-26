# API Testing - Takah API

Dokumen ini digunakan untuk mencatat proses testing endpoint API pada aplikasi Takah.

Testing dilakukan menggunakan:

* Postman
* Browser
* Localhost environment

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

* Success
* Endpoint berjalan normal

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
    "token": "jwt-token"
  }
}
```

### Status

* Success
* JWT token berhasil dibuat

---

# Master Takah Testing

## Get All Takah

### Endpoint

```http
GET /takah
```

### Status

* Success

---

## Get Takah By ID

### Endpoint

```http
GET /takah/:id
```

### Status

* Success

---

## Create Takah

### Endpoint

```http
POST /takah
```

### Status

* Success

---

## Update Takah

### Endpoint

```http
PUT /takah/:id
```

### Status

* Success

---

## Delete Takah

### Endpoint

```http
DELETE /takah/:id
```

### Status

* Success

---

# Surat Testing

## Get All Surat

### Endpoint

```http
GET /surat
```

### Status

* Success

---

## Get Surat By ID

### Endpoint

```http
GET /surat/:id
```

### Status

* Success

---

## Create Surat

### Endpoint

```http
POST /surat
```

### Status

* Success

---

## Update Surat

### Endpoint

```http
PUT /surat/:id
```

### Status

* Success

---

## Delete Surat

### Endpoint

```http
DELETE /surat/:id
```

### Status

* Success

---

# Config Nomor Surat Testing

## Get All Config Nomor Surat

### Endpoint

```http
GET /config-nomor
```

### Status

* Success

---

## Get Config Nomor Surat By ID

### Endpoint

```http
GET /config-nomor/:id
```

### Status

* Success

---

## Create Config Nomor Surat

### Endpoint

```http
POST /config-nomor
```

### Status

* Success

---

## Update Config Nomor Surat

### Endpoint

```http
PUT /config-nomor/:id
```

### Status

* Success

---

## Delete Config Nomor Surat

### Endpoint

```http
DELETE /config-nomor/:id
```

### Status

* Success

---

# Template Surat Testing

## Get All Template Surat

### Endpoint

```http
GET /template-surat
```

### Status

* Success

---

## Get Template Surat By ID

### Endpoint

```http
GET /template-surat/:id
```

### Status

* Success

---

## Create Template Surat

### Endpoint

```http
POST /template-surat
```

### Status

* Success

---

## Update Template Surat

### Endpoint

```http
PUT /template-surat/:id
```

### Status

* Success

---

## Delete Template Surat

### Endpoint

```http
DELETE /template-surat/:id
```

### Status

* Success

---

# Surat Keluar Testing

## Get All Surat Keluar

### Endpoint

```http
GET /surat-keluar
```

### Status

* Success

---

## Get Surat Keluar By ID

### Endpoint

```http
GET /surat-keluar/:id
```

### Status

* Success

---

## Create Surat Keluar

### Endpoint

```http
POST /surat-keluar
```

### Status

* Success

### Notes

* Nomor surat berhasil dibuat otomatis berdasarkan Config Nomor Surat.
* Relasi Master Takah berhasil digunakan saat generate nomor surat.

---

## Update Surat Keluar

### Endpoint

```http
PUT /surat-keluar/:id
```

### Status

* Success

---

## Delete Surat Keluar

### Endpoint

```http
DELETE /surat-keluar/:id
```

### Status

* Success

---

# Surat Masuk Testing

## Get All Surat Masuk

### Endpoint

```http
GET /surat-masuk
```

### Status

* Success

---

## Get Surat Masuk By ID

### Endpoint

```http
GET /surat-masuk/:id
```

### Status

* Success

---

## Create Surat Masuk

### Endpoint

```http
POST /surat-masuk
```

### Status

* Success

---

## Update Surat Masuk

### Endpoint

```http
PUT /surat-masuk/:id
```

### Status

* Success

---

## Delete Surat Masuk

### Endpoint

```http
DELETE /surat-masuk/:id
```

### Status

* Success

---

# Approval Surat Testing

## Get All Approval

### Endpoint

```http
GET /approval
```

### Status

* Success

---

## Get Approval By ID

### Endpoint

```http
GET /approval/:id
```

### Status

* Success

---

## Create Approval

### Endpoint

```http
POST /approval
```

### Status

* Success

---

## Update Approval

### Endpoint

```http
PUT /approval/:id
```

### Status

* Success

---

# Monitoring Surat Testing

## Get All Monitoring

### Endpoint

```http
GET /monitoring
```

### Status

* Success

---

## Get Monitoring By ID

### Endpoint

```http
GET /monitoring/:id
```

### Status

* Success

---

## Create Monitoring

### Endpoint

```http
POST /monitoring
```

### Status

* Success

---

## Update Monitoring

### Endpoint

```http
PUT /monitoring/:id
```

### Status

* Success

---

## Delete Monitoring

### Endpoint

```http
DELETE /monitoring/:id
```

### Status

* Success

---

# Current Testing Status

| Module                   | Status | Notes                   |
| ------------------------ | ------ | ----------------------- |
| Health Check             | Done   | Endpoint berjalan       |
| Authentication           | Done   | JWT login berhasil      |
| Master Takah             | Done   | CRUD berjalan           |
| Surat                    | Done   | CRUD berjalan           |
| Config Nomor Surat       | Done   | CRUD berjalan           |
| Template Surat           | Done   | CRUD berjalan           |
| Surat Keluar             | Done   | CRUD berjalan           |
| Generate Nomor Surat     | Done   | Berjalan otomatis       |
| Surat Masuk              | Done   | CRUD berjalan           |
| Approval Surat           | Done   | CRUD berjalan           |
| Monitoring Surat         | Done   | CRUD berjalan           |
| Database Testing         | Todo   | Database belum tersedia |
| Password Hashing Testing | Todo   | Planned                 |

---

# Future Testing Plan

Testing yang direncanakan selanjutnya:

* Database integration testing
* Password hashing testing
* Authorization testing
* Validation testing
* Error handling testing
* File upload testing
* Load testing
* API documentation testing
