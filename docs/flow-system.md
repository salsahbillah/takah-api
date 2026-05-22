# Flow System - Takah API

Dokumen ini menjelaskan gambaran sistem dan alur utama aplikasi Takah.

---

# Tujuan Sistem

Takah API dibuat untuk membantu proses administrasi surat menyurat secara digital di lingkungan perusahaan atau instansi.

Sistem ini dirancang untuk:
- Mengelola master jenis surat.
- Membantu pembuatan surat keluar.
- Membantu pencatatan surat masuk.
- Membuat nomor surat otomatis.
- Menyimpan arsip surat secara digital.
- Mempermudah monitoring status surat.

---

# Modul Utama

## 1. Master Takah

Master Takah digunakan untuk menyimpan data jenis surat yang digunakan dalam sistem.

Contoh data:
- SKET → Surat Keterangan
- SKK → Surat Keterangan Kerja
- SP → Surat Peringatan
- UND → Surat Undangan
- MEM → Memorandum
- ND → Nota Dinas

Fungsi:
- Digunakan pada pembuatan surat.
- Digunakan untuk generate nomor surat.
- Digunakan untuk template surat.
- Digunakan untuk monitoring dan reporting.

---

## 2. Config Nomor Surat

Config nomor surat digunakan untuk menentukan format penomoran surat otomatis.

Contoh format:

```text
001/UND/CBN/052026
```

Keterangan:
- `001` → nomor urut surat
- `UND` → kode surat dari Master Takah
- `CBN` → kode perusahaan/divisi
- `052026` → bulan dan tahun

Aturan:
- Nomor surat otomatis bertambah.
- Nomor reset setiap bulan dan tahun.

Contoh:

```text
001/UND/CBN/042026
002/UND/CBN/042026
001/UND/CBN/052026
```

---

## 3. Template Surat

Template surat digunakan untuk menyimpan format isi surat berdasarkan jenis surat.

Contoh:
- Template Surat Undangan
- Template Surat Keterangan
- Template Surat Peringatan

Tujuan:
- Mempermudah pembuatan surat.
- Mengurangi penulisan manual.
- Menjaga format surat tetap konsisten.

---

## 4. Surat Keluar

Surat keluar digunakan untuk membuat dan menyimpan surat yang dikirim keluar perusahaan atau instansi.

Data utama:
- Nomor surat
- Jenis surat
- Tujuan surat
- Perihal
- Lampiran
- Tanggal surat
- File surat

Flow:

```text
User login
↓
User memilih jenis surat
↓
User memilih template surat
↓
User mengisi parameter surat
↓
Sistem generate nomor surat otomatis
↓
Surat disimpan
↓
Surat masuk ke monitoring surat
```

---

## 5. Surat Masuk

Surat masuk digunakan untuk mencatat surat yang diterima dari pihak luar.

Data utama:
- Nomor surat
- Pengirim
- Tanggal surat
- Perihal
- File surat

Fungsi:
- Arsip digital surat masuk.
- Monitoring surat masuk.
- Riwayat surat masuk.

---

## 6. Monitoring Surat

Monitoring surat digunakan untuk melihat status dan riwayat surat.

Contoh status:
- draft
- pending
- approved
- rejected
- completed

Flow monitoring:

```text
Surat dibuat
↓
Status draft
↓
Surat diproses
↓
Status berubah
↓
Riwayat surat tersimpan
```

---

# User Role

## Admin

Admin memiliki akses:
- Mengelola Master Takah
- Mengelola template surat
- Mengelola config nomor surat
- Monitoring seluruh surat
- Approve surat

---

## User / Staff

User memiliki akses:
- Membuat surat
- Melihat surat
- Monitoring status surat sendiri

---

# Future Development

Fitur yang dapat dikembangkan selanjutnya:
- Database integration
- JWT authentication
- Upload file PDF
- Digital signature
- Export PDF
- Dashboard reporting
- Email notification
- Disposisi surat