# Flow System - Takah API

Dokumen ini menjelaskan gambaran sistem dan alur utama aplikasi Takah.

---

# Tujuan Sistem

Takah API dibuat untuk membantu proses administrasi surat menyurat secara digital di lingkungan perusahaan atau instansi.

Sistem ini dirancang untuk:

* Mengelola master jenis surat.
* Membantu pembuatan surat keluar.
* Membantu pencatatan surat masuk.
* Membuat nomor surat otomatis.
* Menyimpan arsip surat secara digital.
* Mempermudah monitoring status surat.
* Membantu proses approval dan review surat.

---

# Modul Utama

## 1. Master Takah

Master Takah digunakan untuk menyimpan data jenis surat yang digunakan dalam sistem.

Contoh data:

* SKET → Surat Keterangan
* SKK → Surat Keterangan Kerja
* SP → Surat Peringatan
* UND → Surat Undangan
* MEM → Memorandum
* ND → Nota Dinas

Fungsi:

* Digunakan pada pembuatan surat.
* Digunakan untuk generate nomor surat.
* Digunakan untuk template surat.
* Digunakan untuk monitoring dan reporting.

---

## 2. Config Nomor Surat

Config nomor surat digunakan untuk menentukan format penomoran surat otomatis.

Config nomor surat memiliki relasi dengan Master Takah.

Setiap jenis surat pada Master Takah dapat memiliki format nomor surat yang berbeda sesuai kebutuhan perusahaan atau instansi.

Contoh:

* Surat Undangan (UND)

```text
001/UND/CBN/052026
```

* Surat Keterangan (SKET)

```text
001/SKET/CBN/052026
```

Keterangan:

* `001` → nomor urut surat
* `UND` / `SKET` → kode surat dari Master Takah
* `CBN` → kode perusahaan/divisi
* `052026` → bulan dan tahun

Aturan:

* Nomor surat otomatis bertambah.
* Nomor reset setiap bulan dan tahun.
* Format nomor surat dapat berbeda untuk setiap jenis surat.

Contoh:

```text
001/UND/CBN/042026
002/UND/CBN/042026
001/SKET/CBN/052026
```
---

## 3. Template Surat

Template surat digunakan untuk menyimpan format isi surat berdasarkan jenis surat.

Contoh:

* Template Surat Undangan
* Template Surat Keterangan
* Template Surat Peringatan

Tujuan:

* Mempermudah pembuatan surat.
* Mengurangi penulisan manual.
* Menjaga format surat tetap konsisten.

---

## 4. Surat Keluar

Surat keluar digunakan untuk membuat dan menyimpan surat yang dikirim keluar perusahaan atau instansi.

Data utama:

* Nomor surat
* Jenis surat
* Tujuan surat
* Perihal
* Lampiran
* Tanggal surat
* File surat

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
Surat disimpan sebagai draft
↓
User mengirim surat untuk approval
↓
Admin / approver melakukan review surat
↓
Approve / reject surat
↓
Status surat berubah
↓
Riwayat approval tersimpan
↓
Surat masuk ke monitoring surat
```

---

## 5. Surat Masuk

Surat masuk digunakan untuk mencatat surat yang diterima dari pihak luar.

Data utama:

* Nomor surat
* Pengirim
* Tanggal surat
* Perihal
* File surat

Fungsi:

* Arsip digital surat masuk.
* Monitoring surat masuk.
* Riwayat surat masuk.

---

## 6. Monitoring Surat

Monitoring surat digunakan untuk melihat status dan riwayat surat.

Contoh status:

* draft
* pending
* approved
* rejected
* completed

Monitoring surat juga digunakan untuk melihat proses approval surat dan riwayat review surat.

Informasi monitoring:

* Status surat
* Riwayat approval
* User approver
* Waktu approval
* Catatan approval

Flow monitoring:

```text
Surat dibuat
↓
Status draft
↓
Surat dikirim untuk approval
↓
Approver melakukan review
↓
Status approval berubah
↓
Riwayat approval tersimpan
↓
Monitoring surat diperbarui
```

---

# User Role

## Admin

Admin memiliki akses:

* Mengelola Master Takah
* Mengelola template surat
* Mengelola config nomor surat
* Monitoring seluruh surat
* Approve surat
* Review surat
* Melihat riwayat approval surat

---

## User / Staff

User memiliki akses:

* Membuat surat
* Melihat surat
* Monitoring status surat sendiri
* Melihat status approval surat

---
