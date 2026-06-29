# API Testing - Takah API

Dokumen ini digunakan untuk mencatat proses pengujian endpoint API pada aplikasi Takah.

Pengujian dilakukan menggunakan:

* Postman
* Browser
* Localhost Environment

---

# Base URL

```text
http://localhost:8080/api/v1
```

---

# Health Check Testing

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
* Endpoint berjalan normal.

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
* JWT token berhasil dibuat.
* Middleware berhasil memvalidasi token.
* Endpoint yang diproteksi berhasil diakses menggunakan JWT.

---

# Dashboard API Testing

## Dashboard Summary

### Endpoint

```http
GET /surat
```

### Status

* Success

### Notes

* Endpoint digunakan sebagai dashboard ringkasan aplikasi.
* Menampilkan informasi jumlah data pada sistem.

---

# Master Takah Testing

## Get All Master Takah

### Endpoint

```http
GET /takah
```

### Status

* Success

---

## Get Master Takah By ID

### Endpoint

```http
GET /takah/:id
```

### Status

* Success

---

## Create Master Takah

### Endpoint

```http
POST /takah
```

### Status

* Success

### Notes

* Data berhasil disimpan ke database MySQL.

---

## Update Master Takah

### Endpoint

```http
PUT /takah/:id
```

### Status

* Success

* Data berhasil diperbarui pada database.

---

## Delete Master Takah

### Endpoint

```http
DELETE /takah/:id
```

### Status

* Success

* Data berhasil dihapus dari database.

---

# Parameter Surat Testing

## Get All Parameter Surat

### Endpoint

```http
GET /parameter-surat
```

### Status

* Success

---

## Get Parameter Surat By ID

### Endpoint

```http
GET /parameter-surat/:id
```

### Status

* Success

---

## Create Parameter Surat

### Endpoint

```http
POST /parameter-surat
```

### Status

* Success

### Notes

* Parameter surat berhasil disimpan berdasarkan Master Takah.
* Parameter surat berhasil digunakan sebagai acuan pembuatan template surat.

---

## Update Parameter Surat

### Endpoint

```http
PUT /parameter-surat/:id
```

### Status

* Success

---

## Delete Parameter Surat

### Endpoint

```http
DELETE /parameter-surat/:id
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

### Notes

* Relasi dengan Master Takah berhasil digunakan.
* Config nomor surat berhasil disimpan ke database.
* Config nomor surat dapat disimpan berdasarkan jenis surat.

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

### Notes

* Template berhasil terhubung dengan Master Takah.
* Template berhasil disimpan ke database berdasarkan jenis surat.

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

```
```
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

* Data surat keluar berhasil disimpan ke database MySQL.
* Nomor surat berhasil dibuat secara otomatis berdasarkan Config Nomor Surat.
* Nomor surat mengikuti jenis surat yang dipilih pada Master Takah.
* Nomor terakhir pada Config Nomor Surat berhasil diperbarui secara otomatis.
* Status awal surat adalah **draft**.

---

## Update Surat Keluar

### Endpoint

```http
PUT /surat-keluar/:id
```

### Status

* Success

### Notes

* Data surat keluar berhasil diperbarui pada database.

---

## Delete Surat Keluar

### Endpoint

```http
DELETE /surat-keluar/:id
```

### Status

* Success

### Notes

* Data surat keluar berhasil dihapus dari database.

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

### Notes

* Data approval berhasil disimpan ke database.
* Data approver berhasil direlasikan dengan surat keluar.
* Status surat keluar otomatis berubah menjadi **pending**.
* Riwayat approval berhasil tercatat.

---

## Update Approval

### Endpoint

```http
PUT /approval/:id
```

### Status

* Success

### Notes

* Status approval berhasil diperbarui menjadi **approved** atau **rejected**.
* Status pada Surat Keluar otomatis ikut diperbarui.
* Waktu approval berhasil tercatat pada database.

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

### Notes

* Data monitoring berhasil disimpan ke database.
* Status surat berhasil ditampilkan.
* Approver terakhir berhasil ditampilkan.
* Catatan approval terakhir berhasil ditampilkan.
* Monitoring berhasil menyimpan riwayat proses surat.

---

## Update Monitoring

### Endpoint

```http
PUT /monitoring/:id
```

### Status

* Success

### Notes

* Status monitoring berhasil diperbarui.
* Data approver terakhir berhasil diperbarui.

---

## Delete Monitoring

### Endpoint

```http
DELETE /monitoring/:id
```

### Status

* Success

### Notes

* Data monitoring berhasil dihapus dari database.
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

### Notes

* Data surat masuk berhasil disimpan ke database MySQL.
* Status awal surat masuk adalah **received**.
* Data pengirim, penerima, nomor surat, dan tanggal surat berhasil tersimpan.

---

## Update Surat Masuk

### Endpoint

```http
PUT /surat-masuk/:id
```

### Status

* Success

### Notes

* Data surat masuk berhasil diperbarui pada database.

---

## Delete Surat Masuk

### Endpoint

```http
DELETE /surat-masuk/:id
```

### Status

* Success

### Notes

* Data surat masuk berhasil dihapus dari database.

---

# Database Connection Testing

## Database

```text
takah_db
```

## Configuration

```text
.env
```

### Status

* Database MySQL berhasil dibuat.
* Konfigurasi database menggunakan file `.env`.
* Aplikasi berhasil terhubung ke database MySQL.
* Seluruh endpoint CRUD berhasil menggunakan database MySQL.
* Koneksi database berhasil diuji menggunakan package `database/sql`.

---

# Authentication & Security Testing

## JWT Authentication

### Status

* Success

### Notes

* JWT berhasil dibuat saat proses login.
* Token berhasil digunakan untuk mengakses endpoint yang diproteksi.
* Middleware berhasil memvalidasi token.
* Endpoint tanpa token mengembalikan status **401 Unauthorized**.

---

# Current Testing Status

| Module               | Status | Notes                                       |
| -------------------- | ------ | ------------------------------------------- |
| Health Check         | Done   | Endpoint berjalan normal                    |
| Authentication       | Done   | JWT Authentication berhasil                 |
| Dashboard API        | Done   | Endpoint dashboard berjalan normal          |
| Master Takah         | Done   | CRUD menggunakan MySQL                      |
| Parameter Surat      | Done   | CRUD menggunakan MySQL                      |
| Config Nomor Surat   | Done   | CRUD menggunakan MySQL                      |
| Template Surat       | Done   | CRUD menggunakan MySQL                      |
| Generate Nomor Surat | Done   | Nomor surat berhasil dibuat otomatis        |
| Surat Keluar         | Done   | CRUD menggunakan MySQL                      |
| Approval Surat       | Done   | Approval berhasil mengubah status surat     |
| Monitoring Surat     | Done   | CRUD menggunakan MySQL                      |
| Surat Masuk          | Done   | CRUD menggunakan MySQL                      |
| Database Setup       | Done   | Database `takah_db` berhasil dibuat         |
| Database Connection  | Done   | MySQL berhasil terhubung menggunakan `.env` |
| Database CRUD        | Done   | Seluruh endpoint berhasil menggunakan MySQL |

---

# Testing Conclusion

Berdasarkan hasil pengujian menggunakan Postman, Browser, dan Localhost Environment, seluruh endpoint pada Takah API berhasil dijalankan dengan baik. Seluruh proses CRUD pada setiap modul telah menggunakan database MySQL sebagai media penyimpanan data. Proses autentikasi menggunakan JSON Web Token (JWT) juga berhasil diterapkan sehingga endpoint yang diproteksi hanya dapat diakses oleh pengguna yang telah melakukan login.

Implementasi generate nomor surat otomatis berhasil berjalan berdasarkan konfigurasi pada Config Nomor Surat. Selain itu, proses Approval Surat telah terintegrasi dengan Surat Keluar sehingga perubahan status approval secara otomatis memperbarui status surat. Modul Monitoring Surat juga berhasil mencatat perkembangan status surat berdasarkan proses approval yang dilakukan.

Secara keseluruhan, backend Takah API telah berhasil diimplementasikan dan seluruh endpoint utama telah melewati proses pengujian dengan hasil yang sesuai dengan kebutuhan sistem.
