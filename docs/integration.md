# Integration Documentation - Takah API

Dokumen ini menjelaskan integrasi antar modul pada aplikasi Takah.

---

# Tujuan Integrasi

Integrasi dibuat agar setiap modul dalam aplikasi Takah dapat saling terhubung sehingga proses administrasi surat dapat berjalan secara terstruktur mulai dari pembuatan surat hingga proses monitoring.

---

# Main Integration Flow

```text
Master Takah
│
├── Config Nomor Surat
│
├── Template Surat
│   │
│   └── Parameter Surat
│
└── Surat Keluar
      │
      ├── Approval Surat
      │
      └── Monitoring Surat

Surat Masuk
│
└── Monitoring Surat
```

---

# 1. Master Takah Integration

Master Takah digunakan sebagai data utama jenis surat yang tersedia pada sistem.

Contoh data:

* SKET → Surat Keterangan
* SKK → Surat Keterangan Kerja
* SP → Surat Peringatan
* SIK → Surat Izin Kegiatan
* UND → Surat Undangan
* MEM → Memorandum
* ND → Nota Dinas

Master Takah digunakan oleh:

* Config Nomor Surat
* Template Surat
* Surat Keluar

Master Takah menjadi acuan agar setiap jenis surat memiliki konfigurasi nomor, template, dan proses surat yang sesuai.

Status saat ini:

* CRUD Master Takah telah menggunakan MySQL.
* Relasi dengan modul lain telah diterapkan.

---

# 2. Template Surat Integration

Template Surat terhubung langsung dengan Master Takah.

Template digunakan sebagai format dasar isi surat berdasarkan jenis surat yang dipilih.

Flow:

```text
User memilih Master Takah
↓
Sistem membaca Template Surat
↓
Sistem membaca Parameter Surat
↓
User mengisi data surat
↓
Surat dibuat
```

Contoh:

* Surat Undangan menggunakan template UND.
* Surat Keterangan menggunakan template SKET.

Status saat ini:

* CRUD Template Surat telah menggunakan MySQL.
* Template telah berelasi dengan Master Takah.
* Template digunakan saat proses pembuatan Surat Keluar.

---

# 3. Parameter Surat Integration

Parameter Surat digunakan untuk menentukan field input pada setiap Template Surat.

Setiap template dapat memiliki parameter yang berbeda sesuai kebutuhan jenis surat.

Contoh parameter:

* Nama Tujuan
* Nama Kegiatan
* Tanggal Kegiatan
* Tempat Kegiatan
* Keperluan

Flow:

```text
Admin membuat Template Surat
↓
Admin menambahkan Parameter Surat
↓
User memilih Template
↓
Sistem membaca Parameter Surat
↓
User mengisi data sesuai parameter
↓
Data digunakan untuk membuat Surat Keluar
```

Contoh penggunaan pada template:

```text
Dengan hormat,

Kami mengundang {{nama_tujuan}}
untuk menghadiri {{nama_kegiatan}}
pada {{tanggal_kegiatan}}
bertempat di {{tempat_kegiatan}}.
```

Status saat ini:

* CRUD Parameter Surat telah tersedia.
* Data Parameter Surat telah menggunakan MySQL.
* Parameter Surat terhubung dengan Template Surat.

---

# 4. Config Nomor Surat Integration

Config Nomor Surat digunakan untuk menghasilkan nomor surat secara otomatis.

Config Nomor Surat memiliki relasi dengan Master Takah sehingga setiap jenis surat dapat memiliki format nomor yang berbeda.

Flow:

```text
User membuat Surat Keluar
↓
Sistem membaca Master Takah
↓
Sistem membaca Config Nomor Surat
↓
Generate nomor surat otomatis
↓
Nomor surat disimpan pada Surat Keluar
```

Contoh:

```text
001/UND/CBN/062026
001/SKET/CBN/062026
```

Status saat ini:

* CRUD Config Nomor Surat telah menggunakan MySQL.
* Relasi dengan Master Takah telah berjalan.
* Nomor surat dibuat otomatis berdasarkan konfigurasi.
* Nilai `last_number` diperbarui setiap nomor surat berhasil dibuat.

---

# 5. Surat Keluar Integration

Surat Keluar menggunakan beberapa modul yang saling terhubung, yaitu:

* Master Takah
* Config Nomor Surat
* Template Surat
* Parameter Surat

Flow:

```text
User login
↓
Memilih jenis surat
↓
Generate nomor surat
↓
Memilih Template Surat
↓
Sistem membaca Parameter Surat
↓
User mengisi data surat
↓
Surat disimpan sebagai draft
↓
User mengirim surat untuk approval
↓
Status Surat Keluar menjadi pending
```

Status saat ini:

* CRUD Surat Keluar telah menggunakan MySQL.
* Nomor surat dibuat otomatis.
* Status awal surat adalah `draft`.
* Surat dapat dikirim ke proses Approval.
# 6. Approval Surat Integration

Approval Surat digunakan untuk proses review dan persetujuan Surat Keluar sebelum surat dinyatakan selesai.

Approval Surat juga digunakan untuk mencatat riwayat proses persetujuan beserta approver yang melakukan review.

Flow:

```text
User membuat Surat Keluar
↓
Surat dikirim untuk approval
↓
Status Surat Keluar menjadi pending
↓
Approver melakukan review
↓
Approver memberikan catatan
↓
Approve / Reject
↓
Status Approval diperbarui
↓
Status Surat Keluar ikut diperbarui
↓
Monitoring Surat diperbarui
```

Data approval yang disimpan:

* ID Surat Keluar
* Nomor Surat
* ID Approver
* Nama Approver
* Status Approval
* Catatan Approval
* Waktu Approval

Tujuan:

* Mengetahui siapa yang melakukan review.
* Menyimpan catatan approval atau reject.
* Memperbarui status Surat Keluar.
* Menjadi sumber data Monitoring Surat.

Status saat ini:

* CRUD Approval Surat telah menggunakan MySQL.
* Status approval dapat berubah menjadi `pending`, `approved`, atau `rejected`.
* Status Surat Keluar diperbarui secara otomatis setelah approval diubah.

---

# 7. Surat Masuk Integration

Surat Masuk digunakan untuk mencatat surat yang diterima dari pihak luar perusahaan atau instansi.

Flow:

```text
Surat diterima
↓
Admin menginput data Surat Masuk
↓
Data disimpan ke database
↓
Status menjadi received
↓
Monitoring Surat diperbarui
```

Tujuan:

* Menyimpan arsip surat masuk.
* Mempermudah pencarian surat.
* Mendukung monitoring surat.

Status saat ini:

* CRUD Surat Masuk telah menggunakan MySQL.
* Status awal surat masuk adalah `received`.
* Upload file surat masih menggunakan placeholder dan akan dikembangkan pada tahap berikutnya.

---

# 8. Monitoring Surat Integration

Monitoring Surat digunakan untuk menampilkan perkembangan seluruh proses surat.

Monitoring memperoleh data dari:

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

Flow Surat Keluar:

```text
Surat dibuat
↓
Status draft
↓
Approval dikirim
↓
Review dilakukan
↓
Status Approval berubah
↓
Status Surat Keluar diperbarui
↓
Monitoring diperbarui
```

Flow Surat Masuk:

```text
Surat diterima
↓
Data Surat Masuk disimpan
↓
Status received
↓
Monitoring diperbarui
```

Data Monitoring yang disimpan:

* ID Surat Keluar atau Surat Masuk
* Nomor Surat
* Status Surat
* Approver terakhir
* Catatan approval terakhir
* User terakhir yang memperbarui data
* Waktu update

Status saat ini:

* CRUD Monitoring Surat telah menggunakan MySQL.
* Monitoring menampilkan status surat terbaru.
* Monitoring membaca data Approval Surat dan Surat Masuk.

---

# Authentication Integration

Authentication digunakan untuk membatasi akses pengguna ke endpoint API.

Saat ini:

* Authentication menggunakan JWT.
* Login menghasilkan token JWT.
* Middleware melakukan validasi token pada endpoint yang dilindungi.

Role yang direncanakan:

* admin
* user

Admin memiliki akses:

* Mengelola Master Takah
* Mengelola Config Nomor Surat
* Mengelola Template Surat
* Mengelola Parameter Surat
* Mengelola Surat Masuk
* Melakukan Approval Surat
* Melihat Monitoring Surat

User memiliki akses:

* Membuat Surat Keluar
* Melihat status surat
* Melihat hasil approval
* Menginput Surat Masuk sesuai hak akses

---

# Database Integration

Database yang digunakan:

| Table              | Purpose                 |
| ------------------ | ----------------------- |
| users              | Data user               |
| master_takah       | Data jenis surat        |
| config_nomor_surat | Konfigurasi nomor surat |
| template_surat     | Template surat          |
| parameter_surat    | Parameter template      |
| surat_keluar       | Data surat keluar       |
| approval_surat     | Data approval           |
| surat_masuk        | Data surat masuk        |
| monitoring_surat   | Monitoring surat        |

Relasi utama:

```text
users
│
├── master_takah
├── surat_keluar
├── surat_masuk
├── approval_surat
└── monitoring_surat

master_takah
│
├── config_nomor_surat
├── template_surat
│   │
│   └── parameter_surat
│
└── surat_keluar
      │
      ├── approval_surat
      └── monitoring_surat

surat_masuk
│
└── monitoring_surat
```

---

# Current Integration Status

| Module               | Status  | Notes                                            |
| -------------------- | ------- | ------------------------------------------------ |
| Master Takah         | Done    | CRUD menggunakan MySQL                           |
| Authentication       | Done    | JWT Authentication                               |
| Config Nomor Surat   | Done    | Relasi dengan Master Takah                       |
| Template Surat       | Done    | CRUD menggunakan MySQL                           |
| Parameter Surat      | Done    | CRUD dan relasi dengan Template Surat            |
| Generate Nomor Surat | Done    | Nomor otomatis berdasarkan konfigurasi           |
| Surat Keluar         | Done    | CRUD menggunakan MySQL                           |
| Approval Surat       | Done    | Approval memperbarui status Surat Keluar         |
| Surat Masuk          | Done    | CRUD menggunakan MySQL                           |
| Monitoring Surat     | Done    | CRUD menggunakan MySQL dan membaca data approval |
| Database             | Done    | Seluruh modul telah terintegrasi dengan MySQL    |
| Upload File Surat    | Partial | Tahap pengembangan selanjutnya                   |
| Role Authorization   | Partial | Pengembangan hak akses admin dan user            |
