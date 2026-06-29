# Flow System - Takah API

Dokumen ini menjelaskan gambaran sistem dan alur utama aplikasi Takah.

---

# Tujuan Sistem

Takah API dibuat untuk membantu proses administrasi surat menyurat secara digital di lingkungan perusahaan atau instansi.

Sistem ini dirancang untuk:

* Mengelola Master Takah.
* Mengelola Config Nomor Surat.
* Mengelola Template Surat.
* Mengelola Parameter Surat.
* Membantu pembuatan Surat Keluar.
* Membantu pencatatan Surat Masuk.
* Membuat nomor surat secara otomatis.
* Mengelola proses approval surat.
* Mempermudah monitoring status surat.

---

# Modul Utama

## 1. Master Takah

Master Takah digunakan untuk menyimpan daftar jenis surat yang tersedia pada sistem.

Contoh data:

* SKET → Surat Keterangan
* SKK → Surat Keterangan Kerja
* SP → Surat Peringatan
* SIK → Surat Izin Kegiatan
* UND → Surat Undangan
* MEM → Memorandum
* ND → Nota Dinas

Fungsi:

* Menjadi data master seluruh jenis surat.
* Digunakan pada Config Nomor Surat.
* Digunakan pada Template Surat.
* Digunakan pada Surat Keluar.
* Menjadi acuan Generate Nomor Surat.

---

## 2. Config Nomor Surat

Config Nomor Surat digunakan untuk menentukan format penomoran surat secara otomatis.

Setiap jenis surat memiliki konfigurasi nomor yang berbeda sesuai Master Takah.

Contoh:

Surat Undangan

```text
001/UND/CBN/062026
```

Surat Keterangan

```text
001/SKET/CBN/062026
```

Keterangan:

* 001 → nomor urut
* UND / SKET → kode surat
* CBN → kode perusahaan
* 062026 → bulan dan tahun

Flow:

```text
User memilih jenis surat
↓
Sistem membaca Master Takah
↓
Sistem membaca Config Nomor Surat
↓
Generate nomor surat otomatis
↓
Nomor surat digunakan pada Surat Keluar
```

---

## 3. Template Surat

Template Surat digunakan untuk menyimpan format surat berdasarkan jenis surat.

Setiap Template Surat terhubung dengan Master Takah sehingga setiap jenis surat dapat memiliki template yang berbeda.

Tujuan:

* Mempercepat pembuatan surat.
* Menjaga format surat tetap konsisten.
* Mendukung penggunaan parameter dinamis.

Flow:

```text
User memilih jenis surat
↓
Sistem membaca Template Surat
↓
Sistem membaca Parameter Surat
↓
User mengisi data sesuai parameter
↓
Surat dibuat
```

---

## 4. Parameter Surat

Parameter Surat digunakan untuk menentukan data yang harus diisi pada setiap Template Surat.

Setiap template dapat memiliki parameter yang berbeda.

Contoh parameter:

* Nama Tujuan
* Nama Kegiatan
* Tanggal Kegiatan
* Tempat Kegiatan
* Keperluan

Tujuan:

* Membuat form input menjadi dinamis.
* Menyesuaikan kebutuhan setiap jenis surat.
* Mengurangi input yang tidak diperlukan.
* Mempermudah pengisian data surat.

Flow:

```text
Admin membuat Template Surat
↓
Admin menambahkan Parameter Surat
↓
User memilih Template Surat
↓
Sistem membaca Parameter Surat
↓
User mengisi data
↓
Parameter digunakan pada Surat Keluar
```

---

## 5. Surat Keluar

Surat Keluar digunakan untuk membuat surat yang akan dikirim kepada pihak lain.

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
Generate nomor surat otomatis
↓
User memilih Template Surat
↓
Sistem membaca Parameter Surat
↓
User mengisi data surat
↓
Surat disimpan dengan status draft
↓
User mengirim surat untuk approval
↓
Status berubah menjadi pending
```

## 6. Approval Surat

Approval Surat digunakan untuk melakukan proses review dan persetujuan terhadap Surat Keluar sebelum surat dinyatakan selesai.

Approval juga digunakan untuk mencatat riwayat review setiap surat.

Data yang dicatat:

* Surat Keluar
* Nama approver
* Status approval
* Catatan review
* Waktu approval

Status approval:

* pending
* approved
* rejected

Flow:

```text
User mengirim Surat Keluar
↓
Status surat menjadi pending
↓
Approver melakukan review
↓
Sistem menyimpan nama approver
↓
Sistem menyimpan catatan review
↓
Approver memilih approve atau reject
↓
Status approval diperbarui
↓
Status Surat Keluar diperbarui
↓
Monitoring Surat diperbarui
```

---

## 7. Surat Masuk

Surat Masuk digunakan untuk mencatat surat yang diterima dari pihak luar perusahaan atau instansi.

Data utama:

* Nomor surat
* Pengirim
* Penerima
* Perihal
* File surat
* Tanggal surat
* Keterangan

Fungsi:

* Menyimpan arsip surat masuk.
* Memudahkan pencarian surat.
* Menjadi bagian dari monitoring surat.
* Menyimpan riwayat surat masuk.

Flow:

```text
Admin menerima surat
↓
Input data Surat Masuk
↓
Upload file surat
↓
Data disimpan
↓
Status menjadi received
↓
Monitoring Surat diperbarui
```

---

## 8. Monitoring Surat

Monitoring Surat digunakan untuk menampilkan perkembangan seluruh proses surat.

Monitoring mengambil informasi dari:

* Surat Keluar
* Approval Surat
* Surat Masuk

Status Surat Keluar:

* draft
* pending
* approved
* rejected
* completed

Status Surat Masuk:

* received
* completed

Informasi yang ditampilkan:

* Nomor surat
* Status surat
* Approver terakhir
* Catatan approval terakhir
* Waktu terakhir diperbarui

Flow Monitoring Surat Keluar:

```text
Surat dibuat
↓
Status draft
↓
Surat dikirim untuk approval
↓
Approver melakukan review
↓
Status Surat Keluar diperbarui
↓
Monitoring Surat diperbarui
↓
Status terbaru ditampilkan
```

Flow Monitoring Surat Masuk:

```text
Surat diterima
↓
Data Surat Masuk dicatat
↓
Status received
↓
Monitoring Surat diperbarui
```

---

# User Role

## Admin

Admin memiliki hak akses untuk:

* Login ke sistem.
* Mengelola Master Takah.
* Mengelola Config Nomor Surat.
* Mengelola Template Surat.
* Mengelola Parameter Surat.
* Mengelola Surat Keluar.
* Mengelola Surat Masuk.
* Melakukan review dan approval surat.
* Melihat Monitoring Surat.
* Melihat riwayat approval.

---

## User / Staff

User memiliki hak akses untuk:

* Login ke sistem.
* Membuat Surat Keluar.
* Memilih Template Surat.
* Mengisi Parameter Surat.
* Mengirim surat untuk approval.
* Melihat status Surat Keluar.
* Melihat hasil approval.
* Mencatat Surat Masuk.
* Melihat Monitoring Surat miliknya.

---

# Ringkasan Alur Sistem

Alur utama aplikasi Takah:

```text
Master Takah
      │
      ▼
Config Nomor Surat
      │
      ▼
Generate Nomor Surat
      │
      ▼
Template Surat
      │
      ▼
Parameter Surat
      │
      ▼
Surat Keluar
      │
      ▼
Approval Surat
      │
      ▼
Monitoring Surat

Surat Masuk
      │
      ▼
Monitoring Surat
```

---

# Status Pengembangan

Status implementasi saat ini:

* CRUD Master Takah telah menggunakan MySQL.
* CRUD Config Nomor Surat telah menggunakan MySQL.
* CRUD Template Surat telah menggunakan MySQL.
* CRUD Parameter Surat telah menggunakan MySQL.
* CRUD Surat Keluar telah menggunakan MySQL.
* CRUD Approval Surat telah menggunakan MySQL.
* CRUD Surat Masuk telah menggunakan MySQL.
* CRUD Monitoring Surat telah menggunakan MySQL.
* Generate nomor surat berjalan otomatis berdasarkan Config Nomor Surat.
* Monitoring menampilkan status terbaru berdasarkan proses Approval maupun Surat Masuk.
