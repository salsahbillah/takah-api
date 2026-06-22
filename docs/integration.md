# Integration Documentation - Takah API

Dokumen ini menjelaskan integrasi antar module pada aplikasi Takah.

---

# Tujuan Integrasi

Integrasi dibuat agar setiap module dalam aplikasi Takah dapat saling terhubung dan digunakan dalam satu alur sistem administrasi surat.

---

# Main Integration Flow

```text
Master Takah
↓
Template Surat
↓
Config Nomor Surat
↓
Surat Keluar
↓
Approval Surat
↓
Monitoring Surat

Surat Masuk
↓
Monitoring Surat
```

---

# 1. Master Takah Integration

Master Takah digunakan sebagai data master jenis surat.

Contoh:

* SKET
* SKK
* UND
* SP

Master Takah digunakan oleh:

* Template Surat
* Generate Nomor Surat
* Surat Keluar
* Monitoring Surat
* Approval Surat

---

# 2. Template Surat Integration

Template surat terhubung dengan Master Takah.

Flow:

```text
Master Takah dipilih
↓
Template surat digunakan
↓
User mengisi parameter surat
↓
Surat dibuat
```

Contoh:

* Surat Undangan menggunakan template UND
* Surat Keterangan menggunakan template SKET

Status saat ini:

* Belum diimplementasikan
* Masih tahap perencanaan

---

# 3. Config Nomor Surat Integration

Config nomor surat digunakan untuk generate nomor otomatis.

Config nomor surat memiliki relasi dengan Master Takah.

Setiap jenis surat pada Master Takah dapat memiliki format nomor surat yang berbeda sesuai kebutuhan perusahaan atau instansi.

Config nomor surat juga digunakan untuk menentukan aturan reset nomor surat.

Flow:

```text
User membuat surat
↓
Sistem membaca jenis surat dari Master Takah
↓
Sistem membaca config nomor surat berdasarkan jenis surat
↓
Sistem membaca aturan reset nomor surat
↓
Sistem generate nomor otomatis
↓
Nomor surat disimpan
```

Contoh:

```text
001/UND/CBN/052026
001/SKET/CBN/052026
```

Aturan reset:

* monthly → reset setiap bulan
* yearly → reset setiap tahun

Status saat ini:

* CRUD Config Nomor Surat tersedia
* Relasi dengan Master Takah tersedia
* Generate nomor surat sudah terintegrasi dengan Surat Keluar

---

# 4. Surat Keluar Integration

Surat keluar menggunakan:

* Master Takah
* Template Surat
* Config Nomor Surat

Flow:

```text
User login
↓
Pilih jenis surat
↓
Pilih template surat
↓
Generate nomor surat
↓
Simpan surat sebagai draft
↓
Kirim surat untuk approval
↓
Approver melakukan review
↓
Approve / reject surat
↓
Monitoring surat
```

Status saat ini:

* CRUD Surat Keluar tersedia
* Nomor surat dibuat otomatis berdasarkan Config Nomor Surat
* Status awal surat adalah draft

---

# 5. Approval Surat Integration

Approval surat digunakan untuk proses review dan persetujuan surat sebelum surat dinyatakan selesai.

Flow:

```text
User membuat surat
↓
Surat masuk ke approval
↓
Approver melakukan review
↓
Approve / reject surat
↓
Status surat diperbarui
↓
Riwayat approval tersimpan
```

Tujuan:

* Mengetahui siapa yang melakukan review surat
* Mengetahui status approval surat
* Menyimpan riwayat approval surat
* Monitoring proses surat

Status saat ini:

* CRUD Approval Surat tersedia
* Data approver dan catatan approval dapat disimpan

---

# 6. Surat Masuk Integration

Surat masuk digunakan untuk mencatat surat yang diterima dari pihak luar perusahaan atau instansi.

Flow:

```text
Surat diterima dari pihak luar
↓
User / admin input data surat masuk
↓
Upload file surat
↓
Data surat masuk disimpan
↓
Monitoring surat masuk
```

Tujuan:

* Menyimpan arsip surat masuk secara digital
* Mempermudah pencarian surat masuk
* Menyimpan riwayat surat yang diterima
* Monitoring surat masuk

Status saat ini:

* Belum diimplementasikan
* Masih tahap perencanaan

---

# 7. Monitoring Surat Integration

Monitoring surat digunakan untuk tracking status surat keluar dan surat masuk.

Status surat keluar:

* draft
* pending
* approved
* rejected
* completed

Status surat masuk:

* received
* completed

Flow surat keluar:

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

Flow surat masuk:

```text
Surat diterima
↓
Data surat dicatat
↓
File surat disimpan
↓
Status received
↓
Monitoring surat diperbarui
```

Status saat ini:

* CRUD Monitoring Surat tersedia
* Monitoring menyimpan status surat
* Monitoring menyimpan approver terakhir
* Monitoring menyimpan catatan approval terakhir

---

# Authentication Integration

Authentication digunakan untuk membatasi akses user.

Saat ini:

* Authentication menggunakan JWT
* Login menghasilkan token JWT
* Middleware digunakan untuk memvalidasi token pada endpoint yang dilindungi

Role plan:

* admin
* user

Admin:

* Manage master data
* Manage surat masuk
* Monitoring semua surat
* Approval surat
* Review surat

User:

* Membuat surat
* Melihat surat sendiri
* Input surat masuk
* Melihat riwayat surat masuk
* Monitoring status approval surat

---

# Database Integration Plan

Database yang direncanakan:

| Table              | Purpose                 |
| ------------------ | ----------------------- |
| users              | Data user               |
| master_takah       | Jenis surat             |
| template_surat     | Template surat          |
| config_nomor_surat | Config penomoran        |
| surat_keluar       | Data surat keluar       |
| approval_surat     | Data approval surat     |
| surat_masuk        | Data surat masuk        |
| monitoring_surat   | Monitoring status surat |

---

# Current Integration Status

| Module               | Status | Notes                                |
| -------------------- | ------ | ------------------------------------ |
| Master Takah         | Done   | CRUD tersedia                        |
| Surat                | Done   | CRUD dummy tersedia                  |
| Authentication       | Done   | JWT authentication tersedia          |
| Config Nomor Surat   | Done   | Relasi dengan Master Takah tersedia  |
| Generate Nomor Surat | Done   | Helper generate nomor surat tersedia |
| Surat Keluar         | Done   | Generate nomor otomatis berjalan     |
| Approval Surat       | Done   | CRUD approval tersedia               |
| Monitoring Surat     | Done   | CRUD monitoring tersedia             |
| Database             | Todo   | Belum terintegrasi MySQL             |
| Template Surat       | Todo   | Belum diimplementasikan              |
| Surat Masuk          | Todo   | Belum diimplementasikan              |

---

# Future Integration Plan

Pengembangan integrasi selanjutnya:

* MySQL integration
* Upload file surat
* Template surat
* Surat masuk
* Approval multi level
* Notification system
* Audit log
* Dashboard monitoring
