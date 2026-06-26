# SOP Admin - Takah API

Dokumen ini menjelaskan prosedur penggunaan aplikasi Takah dari sisi admin.

---

# Tujuan

SOP Admin dibuat untuk membantu admin dalam mengelola data master, surat, approval surat, dan monitoring surat pada aplikasi Takah.

---

# Hak Akses Admin

Admin memiliki akses untuk:

* Login ke sistem
* Mengelola Master Takah
* Mengelola template surat
* Mengelola config nomor surat
* Mengelola surat masuk
* Monitoring seluruh surat
* Approve surat
* Review surat
* Melihat riwayat approval surat
* Melihat riwayat surat masuk

---

# Flow Admin

```text
Admin login
↓
Admin mengelola master data
↓
Admin mengelola template surat
↓
Admin mengelola config nomor surat
↓
Admin mengelola surat masuk
↓
Admin monitoring surat
↓
Admin melakukan review surat
↓
Admin melakukan approval surat
```

---

# 1. Login Admin

## Endpoint

```http
POST /api/v1/auth/login
```

## Request

```json
{
  "email": "admin@takah.com",
  "password": "password123"
}
```

## Success Response

```json
{
  "message": "Login berhasil",
  "data": {
    "token": "jwt-token"
  }
}
```

---

# 2. Manage Master Takah

Admin dapat:

* Menambah jenis surat
* Mengubah jenis surat
* Menghapus jenis surat
* Melihat daftar jenis surat

---

# 3. Manage Template Surat

Admin dapat:

* Membuat template surat
* Mengubah template surat
* Menghapus template surat
* Melihat daftar template surat

---

# 4. Manage Config Nomor Surat

Admin dapat mengatur format nomor surat.

Contoh format:

```text
001/UND/CBN/052026
```

Aturan:

* Nomor otomatis bertambah
* Reset mengikuti config yang digunakan
* Config nomor surat terhubung dengan Master Takah

---

# 5. Manage Surat Keluar

Admin dapat:

* Melihat data surat keluar
* Membuat surat keluar
* Mengubah data surat keluar
* Menghapus data surat keluar
* Melihat status surat keluar

---

# 6. Manage Surat Masuk

Admin dapat:

* Menambah data surat masuk
* Mengubah data surat masuk
* Menghapus data surat masuk
* Melihat riwayat surat masuk

Flow:

```text
Surat diterima dari pihak luar
↓
Admin input data surat masuk
↓
Admin upload file surat
↓
Data surat masuk disimpan
↓
Surat masuk tercatat pada sistem
```

---

# 7. Monitoring Surat

Admin dapat melihat:

* Status surat
* Riwayat surat
* Surat pending
* Surat approved
* Surat rejected
* Riwayat approval surat
* Riwayat surat masuk

Contoh status:

* draft
* pending
* approved
* rejected
* received

---

# 8. Approval Surat

Flow approval:

```text
User membuat surat
↓
Surat masuk monitoring approval
↓
Admin melakukan review surat
↓
Admin approve / reject surat
↓
Status surat berubah
↓
Riwayat approval tersimpan
```

Admin dapat:

* Melihat surat pending approval
* Melakukan review surat
* Memberikan catatan approval
* Approve surat
* Reject surat
* Melihat riwayat approval surat

---

# Current Status

Status implementasi saat ini:

| Feature              | Status |
| -------------------- | ------ |
| JWT Authentication   | Done   |
| CRUD Master Takah    | Done   |
| CRUD Surat           | Done   |
| Config Nomor Surat   | Done   |
| Template Surat       | Done   |
| Generate Nomor Surat | Done   |
| Surat Keluar         | Done   |
| Surat Masuk          | Done   |
| Approval Surat       | Done   |
| Monitoring Surat     | Done   |
| Database Integration | Todo   |

---

# Future Development

Pengembangan admin selanjutnya:

* MySQL integration
* Upload file surat
* Password hashing
* Authorization role

---