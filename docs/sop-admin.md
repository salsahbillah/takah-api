# SOP Admin - Takah API

Dokumen ini menjelaskan prosedur penggunaan aplikasi Takah dari sisi admin.

---

# Tujuan

SOP Admin dibuat untuk membantu admin dalam mengelola data Master Takah, Parameter Surat, Template Surat, Config Nomor Surat, Surat Keluar, Surat Masuk, Approval Surat, dan Monitoring Surat pada aplikasi Takah.

---

# Hak Akses Admin

Admin memiliki akses untuk:

* Login ke sistem.
* Mengelola Master Takah.
* Mengelola Parameter Surat.
* Mengelola Template Surat.
* Mengelola Config Nomor Surat.
* Mengelola Surat Keluar.
* Mengelola Surat Masuk.
* Melakukan review surat.
* Melakukan approval atau reject surat.
* Melihat Monitoring Surat.
* Melihat riwayat approval.
* Melihat riwayat surat masuk.

---

# Flow Admin

```text
Admin login
↓
Mengelola Master Takah
↓
Mengelola Parameter Surat
↓
Mengelola Template Surat
↓
Mengelola Config Nomor Surat
↓
Mengelola Surat Keluar
↓
Mengelola Surat Masuk
↓
Melakukan review dan approval
↓
Monitoring seluruh proses surat
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

Setelah berhasil login, admin menggunakan token JWT untuk mengakses endpoint yang membutuhkan autentikasi.

---

# 2. Manage Master Takah

Admin dapat:

* Menambah data Master Takah.
* Mengubah data Master Takah.
* Menghapus data Master Takah.
* Melihat daftar Master Takah.

Master Takah digunakan sebagai referensi pada:

* Parameter Surat
* Template Surat
* Config Nomor Surat
* Surat Keluar

Setiap jenis surat memiliki kode surat yang digunakan sebagai acuan generate nomor surat dan pemilihan template.

---

# 3. Manage Parameter Surat

Admin dapat:

* Menambah Parameter Surat.
* Mengubah Parameter Surat.
* Menghapus Parameter Surat.
* Melihat daftar Parameter Surat.

Parameter Surat digunakan untuk menentukan field yang harus diisi ketika membuat surat berdasarkan Template Surat.

Contoh parameter:

* Nama Tujuan
* Jabatan
* Unit Kerja
* Tanggal Kegiatan
* Tempat Kegiatan
* Keperluan

Setiap Template Surat dapat memiliki parameter yang berbeda sesuai kebutuhan jenis surat.

---

# 4. Manage Template Surat

Admin dapat:

* Menambah Template Surat.
* Mengubah Template Surat.
* Menghapus Template Surat.
* Melihat daftar Template Surat.

Template Surat dibuat berdasarkan jenis surat pada Master Takah dan dapat menggunakan Parameter Surat sebagai placeholder agar isi surat dapat dihasilkan secara dinamis.

---

# 5. Manage Config Nomor Surat

Admin dapat mengelola konfigurasi nomor surat untuk setiap jenis surat.

Contoh format:

```text
001/UND/CBN/062026
```

Admin dapat mengatur:

* Jenis surat.
* Kode perusahaan.
* Kode divisi.
* Aturan reset nomor surat.
* Nomor terakhir yang digunakan.

Config Nomor Surat digunakan sebagai acuan sistem untuk menghasilkan nomor surat secara otomatis saat Surat Keluar dibuat.

---

# 6. Manage Surat Keluar

Admin dapat:

* Melihat daftar Surat Keluar.
* Membuat Surat Keluar.
* Mengubah data Surat Keluar.
* Menghapus Surat Keluar.
* Melihat status surat.

Flow Surat Keluar:

```text
Pilih jenis surat
↓
Sistem membaca Config Nomor Surat
↓
Generate nomor surat otomatis
↓
Sistem membaca Template Surat
↓
Sistem membaca Parameter Surat
↓
Admin mengisi data surat
↓
Surat disimpan sebagai draft
↓
Surat dikirim untuk approval
```

Status Surat Keluar:

* draft
* pending
* approved
* rejected
* completed

---

# 7. Manage Surat Masuk

Admin dapat:

* Menambah Surat Masuk.
* Mengubah Surat Masuk.
* Menghapus Surat Masuk.
* Melihat daftar Surat Masuk.

Flow:

```text
Surat diterima
↓
Input data Surat Masuk
↓
Data disimpan ke database
↓
Status menjadi received
↓
Monitoring Surat diperbarui
```

Status Surat Masuk:

* received
* completed
# 8. Review dan Approval Surat

Admin bertanggung jawab melakukan proses review terhadap Surat Keluar yang diajukan sebelum surat dinyatakan selesai.

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
Memberikan catatan review
↓
Approve atau Reject surat
↓
Status Approval diperbarui
↓
Status Surat Keluar ikut diperbarui
↓
Monitoring Surat diperbarui
```

Admin dapat:

* Melihat daftar surat yang menunggu review.
* Melihat detail surat.
* Memberikan catatan review.
* Menyetujui surat.
* Menolak surat.
* Melihat riwayat approval.

Status Approval:

* pending
* approved
* rejected

Setelah approval diperbarui, sistem akan otomatis memperbarui status Surat Keluar sehingga proses monitoring dapat menampilkan status terbaru.

---

# 9. Monitoring Surat

Admin dapat melakukan monitoring terhadap seluruh proses surat.

Monitoring mengambil data dari:

* Surat Keluar
* Approval Surat
* Surat Masuk

Informasi yang dapat dilihat:

* Nomor surat.
* Status surat.
* Approver terakhir.
* Catatan approval terakhir.
* User terakhir yang memperbarui data.
* Waktu terakhir diperbarui.

Status Surat Keluar:

* draft
* pending
* approved
* rejected
* completed

Status Surat Masuk:

* received
* completed

Monitoring digunakan untuk mempermudah admin mengetahui perkembangan setiap surat tanpa harus membuka masing-masing modul.

---

# Current Status

Status implementasi saat ini:

| Feature                 | Status  |
| ----------------------- | ------- |
| JWT Authentication      | Done    |
| CRUD Master Takah       | Done    |
| CRUD Parameter Surat    | Done    |
| CRUD Config Nomor Surat | Done    |
| CRUD Template Surat     | Done    |
| Generate Nomor Surat    | Done    |
| CRUD Surat Keluar       | Done    |
| CRUD Surat Masuk        | Done    |
| Review Surat            | Done    |
| Approval Surat          | Done    |
| Monitoring Surat        | Done    |
| Database Setup          | Done    |
| Database Integration    | Done    |
| Password Hashing        | Partial |
| Authorization Role      | Partial |

---

# Catatan

Seluruh modul utama pada aplikasi Takah telah terintegrasi dengan database MySQL. Data Master Takah, Parameter Surat, Template Surat, Config Nomor Surat, Surat Keluar, Approval Surat, Surat Masuk, dan Monitoring Surat telah menggunakan penyimpanan database sehingga perubahan data dapat langsung tercermin pada setiap modul yang saling berelasi.

Pengembangan selanjutnya difokuskan pada implementasi password hashing, pengelolaan hak akses berdasarkan role (admin dan user), upload file surat, serta penyempurnaan keamanan dan validasi data.
