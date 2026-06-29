# SOP User - Takah API

Dokumen ini menjelaskan prosedur penggunaan aplikasi Takah dari sisi user atau staff.

---

# Tujuan

SOP User dibuat untuk membantu pengguna dalam membuat Surat Keluar, mengirim surat untuk proses approval, memantau status surat, serta melihat riwayat surat melalui aplikasi Takah.

---

# Hak Akses User

User memiliki akses untuk:

* Login ke sistem.
* Membuat Surat Keluar.
* Memilih jenis surat.
* Menggunakan Template Surat.
* Mengisi Parameter Surat.
* Melihat Surat Keluar yang dibuat.
* Mengirim surat untuk proses approval.
* Melihat status approval.
* Melihat Monitoring Surat.
* Melihat riwayat surat.

---

# Flow User

```text
User login
↓
Memilih jenis surat
↓
Sistem membaca Config Nomor Surat
↓
Generate nomor surat otomatis
↓
Sistem membaca Template Surat
↓
Sistem membaca Parameter Surat
↓
User mengisi data surat
↓
Surat disimpan sebagai draft
↓
User mengirim surat untuk approval
↓
Admin melakukan review
↓
Status surat diperbarui
↓
User melihat status surat
```

---

# 1. Login User

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

Setelah berhasil login, user menggunakan token JWT untuk mengakses endpoint yang memerlukan autentikasi.

---

# 2. Memilih Jenis Surat

User dapat membuat Surat Keluar berdasarkan jenis surat yang tersedia pada Master Takah.

Contoh jenis surat:

* SKET → Surat Keterangan
* SKK → Surat Keterangan Kerja
* SP → Surat Peringatan
* SIK → Surat Izin Kegiatan
* UND → Surat Undangan
* MEM → Memorandum
* ND → Nota Dinas

Jenis surat yang dipilih akan menentukan:

* Template Surat yang digunakan.
* Parameter Surat yang ditampilkan.
* Config Nomor Surat yang digunakan.

---

# 3. Mengisi Parameter Surat

Setelah memilih jenis surat, sistem akan menampilkan Parameter Surat sesuai dengan Template Surat yang dipilih.

Contoh parameter:

* Nama Tujuan
* Jabatan
* Unit Kerja
* Tanggal Kegiatan
* Tempat Kegiatan
* Keperluan

Parameter Surat digunakan untuk melengkapi isi surat secara otomatis berdasarkan template.

---

# 4. Menggunakan Template Surat

Template Surat digunakan sebagai format dasar pembuatan surat.

Tujuan penggunaan Template Surat:

* Mempermudah pembuatan surat.
* Menjaga format surat tetap konsisten.
* Mengurangi penulisan secara manual.
* Menghasilkan isi surat berdasarkan Parameter Surat yang diisi.

---

# 5. Generate Nomor Surat

Nomor surat dibuat secara otomatis berdasarkan Config Nomor Surat.

Contoh format:

```text
001/UND/CBN/062026
```

Keterangan:

* `001` → nomor urut surat.
* `UND` → kode jenis surat.
* `CBN` → kode perusahaan.
* `062026` → bulan dan tahun.

Aturan:

* Nomor surat dibuat otomatis.
* Nomor surat mengikuti Config Nomor Surat.
* Setiap jenis surat memiliki format nomor yang berbeda.
* Nomor terakhir akan bertambah secara otomatis setiap surat baru dibuat.
# 6. Proses Review dan Approval

Setelah Surat Keluar selesai dibuat, user dapat mengirim surat untuk diproses oleh admin.

Flow:

```text
User membuat Surat Keluar
↓
Surat disimpan sebagai draft
↓
User mengirim surat untuk approval
↓
Status Surat Keluar menjadi pending
↓
Admin melakukan review
↓
Admin memberikan catatan
↓
Approve atau Reject surat
↓
Status Surat Keluar diperbarui
↓
Riwayat approval tersimpan
↓
Monitoring Surat diperbarui
```

User dapat melihat:

* Status surat.
* Status approval.
* Catatan review.
* Nama approver.
* Waktu approval.

Status approval yang tersedia:

* pending
* approved
* rejected

---

# 7. Monitoring Surat

User dapat melakukan monitoring terhadap seluruh Surat Keluar yang telah dibuat.

Informasi yang ditampilkan meliputi:

* Nomor surat.
* Jenis surat.
* Status surat.
* Approver terakhir.
* Catatan approval terakhir.
* Waktu terakhir diperbarui.

Status Surat Keluar:

* draft
* pending
* approved
* rejected
* completed

Monitoring membantu user mengetahui perkembangan proses surat tanpa harus menanyakan langsung kepada approver.

---

# 8. Riwayat Surat

Riwayat Surat digunakan untuk melihat seluruh Surat Keluar yang pernah dibuat oleh user.

Informasi yang ditampilkan meliputi:

* Nomor surat.
* Jenis surat.
* Tujuan surat.
* Perihal.
* Tanggal surat.
* Status surat.

Riwayat surat dapat digunakan sebagai arsip dan referensi ketika diperlukan kembali.

---

# Current Status

Status implementasi saat ini:

| Feature              | Status  |
| -------------------- | ------- |
| JWT Authentication   | Done    |
| CRUD Surat Keluar    | Done    |
| CRUD Parameter Surat | Done    |
| CRUD Template Surat  | Done    |
| Generate Nomor Surat | Done    |
| Review Surat         | Done    |
| Approval Surat       | Done    |
| Monitoring Surat     | Done    |
| Database Setup       | Done    |
| Database Integration | Done    |
| Password Hashing     | Partial |
| Authorization Role   | Partial |

---

# Catatan

Seluruh proses pembuatan Surat Keluar telah terintegrasi dengan database MySQL. Mulai dari pemilihan jenis surat, generate nomor surat otomatis, penggunaan Template Surat dan Parameter Surat, proses approval, hingga Monitoring Surat telah menggunakan penyimpanan data pada database sehingga perubahan status dapat langsung ditampilkan kepada pengguna.

Pengembangan berikutnya difokuskan pada implementasi password hashing, pengaturan hak akses berdasarkan role, upload file surat, serta peningkatan keamanan dan validasi data pada setiap proses.
