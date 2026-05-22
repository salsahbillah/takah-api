# Flow Aplikasi Takah

## Tujuan Sistem

Aplikasi Takah dibuat untuk membantu proses administrasi surat menyurat secara digital, mulai dari pengelolaan master jenis surat, template surat, penomoran surat otomatis, hingga monitoring surat.

## Modul Utama

1. Master Takah
2. Config Nomor Surat
3. Template Surat
4. Surat Keluar
5. Surat Masuk
6. Monitoring Surat

## Flow Master Takah

Master Takah digunakan untuk menyimpan jenis-jenis surat yang digunakan dalam perusahaan.

Contoh data:
- SKET - Surat Keterangan
- SKK - Surat Keterangan Kerja
- SP - Surat Peringatan
- SIK - Surat Ijin Kegiatan
- UND - Surat Undangan
- MEM - Surat Memorandum
- ND - Nota Dinas

Alur:
```text
Admin membuka menu Master Takah
↓
Admin melihat daftar jenis surat
↓
Admin dapat menambah, mengubah, atau menghapus data Takah
↓
Data Master Takah digunakan untuk template dan penomoran surat

Flow Surat Keluar

User login
↓
User memilih jenis surat dari Master Takah
↓
User memilih template surat
↓
User mengisi parameter surat
↓
Sistem generate nomor surat otomatis
↓
Surat disimpan
↓
Surat masuk ke monitoring

Flow Penomoran Surat
Nomor surat dibuat otomatis berdasarkan konfigurasi.
Contoh format:
001/UND/CBN/052026
Keterangan:

001 = nomor urut surat
UND = kode jenis surat dari Master Takah
CBN = kode perusahaan/divisi
052026 = bulan dan tahun

Nomor urut akan reset setiap bulan dan tahun.

Contoh:

001/UND/CBN/042026
002/UND/CBN/042026
001/UND/CBN/052026
Flow Monitoring Surat
Surat dibuat
↓
Status awal: draft
↓
Surat diproses / diajukan
↓
Status berubah menjadi pending / approved / rejected
↓
Riwayat surat dapat dilihat pada menu monitoring

Setelah itu save dulu. Baru nanti kita lanjut bikin **`backend.md`** dan **rancangan SQL**.