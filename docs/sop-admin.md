# SOP Admin - Takah API

Dokumen ini menjelaskan prosedur penggunaan aplikasi Takah dari sisi admin.

---

# Tujuan

SOP Admin dibuat untuk membantu admin dalam mengelola data master, surat, dan monitoring surat pada aplikasi Takah.

---

# Hak Akses Admin

Admin memiliki akses untuk:

- Login ke sistem
- Mengelola Master Takah
- Mengelola template surat
- Mengelola config nomor surat
- Monitoring seluruh surat
- Approve surat
- Melihat riwayat surat

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
Admin monitoring surat
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
    "token": "dummy-token"
  }
}
```

---

# 2. Manage Master Takah

Admin dapat:
- Menambah jenis surat
- Mengubah jenis surat
- Menghapus jenis surat
- Melihat daftar jenis surat

Contoh jenis surat:
- SKET
- SKK
- UND
- SP

---

# 3. Manage Template Surat

Admin dapat:
- Membuat template surat
- Mengubah template surat
- Menghapus template surat

Template surat digunakan untuk:
- Surat Undangan
- Surat Keterangan
- Surat Peringatan

---

# 4. Manage Config Nomor Surat

Admin dapat mengatur format nomor surat.

Contoh format:

```text
001/UND/CBN/052026
```

Aturan:
- Nomor otomatis bertambah
- Reset setiap bulan

---

# 5. Monitoring Surat

Admin dapat melihat:
- Status surat
- Riwayat surat
- Surat pending
- Surat approved
- Surat rejected

Contoh status:
- draft
- pending
- approved
- rejected

---

# 6. Approval Surat

Flow approval:

```text
User membuat surat
↓
Surat masuk monitoring
↓
Admin melakukan pengecekan
↓
Admin approve / reject surat
↓
Status surat berubah
```

---

# Current Status

Status implementasi saat ini:

| Feature | Status |
| --- | --- |
| Login dummy | Done |
| CRUD Master Takah | Done |
| CRUD Surat | Done |
| Approval surat | Todo |
| Monitoring surat | Todo |
| Database integration | Todo |

---

# Future Development

Pengembangan admin selanjutnya:
- JWT authentication
- Dashboard admin
- Notification system
- Approval multi level
- Export PDF
- Audit log