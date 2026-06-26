SOP User - Takah API

Dokumen ini menjelaskan prosedur penggunaan aplikasi Takah dari sisi user/staff.

Tujuan

SOP User dibuat untuk membantu user dalam menggunakan aplikasi Takah untuk pembuatan dan monitoring surat.

Hak Akses User

User memiliki akses untuk:

Login ke sistem
Membuat surat
Melihat surat sendiri
Monitoring status surat
Melihat riwayat surat
Melihat status approval surat
Flow User
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
Surat diproses approver
↓
User monitoring status surat
1. Login User
Endpoint
POST /api/v1/auth/login
Request
{
  "email": "admin@takah.com",
  "password": "password123"
}
Success Response
{
  "message": "Login berhasil",
  "data": {
    "token": "jwt-token"
  }
}
2. Membuat Surat

User dapat membuat surat baru berdasarkan jenis surat yang tersedia pada Master Takah.

Contoh jenis surat:

SKET
SKK
UND
SP
3. Memilih Template Surat

User memilih template surat sesuai kebutuhan.

Contoh:

Template Surat Undangan
Template Surat Keterangan
Template Surat Peringatan

Tujuan:

Mempermudah pembuatan surat
Mengurangi penulisan manual
Menjaga format surat tetap konsisten
4. Generate Nomor Surat

Sistem akan membuat nomor surat otomatis.

Contoh format:

001/UND/CBN/052026

Keterangan:

001 = nomor urut
UND = kode jenis surat
CBN = kode perusahaan/divisi
052026 = bulan dan tahun

Aturan:

Nomor bertambah otomatis
Nomor dibuat berdasarkan Config Nomor Surat
5. Monitoring Surat

User dapat melihat:

Status surat
Riwayat surat
Surat pending
Surat approved
Surat rejected

User juga dapat melihat:

Status approval surat
Riwayat review surat
Catatan approval/reject
6. Riwayat Surat

Riwayat surat digunakan untuk melihat surat yang pernah dibuat user.

Informasi yang ditampilkan:

Nomor surat
Jenis surat
Status surat
Tanggal surat
Current Status

Status implementasi saat ini:

Feature	Status
JWT Authentication	Done
CRUD Surat	Done
Template Surat	Done
Generate Nomor Surat	Done
Surat Keluar	Done
Surat Masuk	Done
Monitoring Surat	Done
Approval Tracking	Done
Database Integration	Todo
Future Development

Pengembangan user module selanjutnya:

Upload file surat
MySQL integration
Password hashing
Authorization role