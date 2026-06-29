INSERT INTO users (name, email, password, role)
VALUES
('Admin Takah', 'admin@takah.com', 'password123', 'admin'),
('User Takah', 'user@takah.com', 'password123', 'user');

INSERT INTO master_takah
(code, name, description, sort_order, created_by, updated_by)
VALUES
('SKET', 'Surat Keterangan', 'Master surat keterangan', 1, 1, 1),
('SKK', 'Surat Keterangan Kerja', 'Master surat keterangan kerja', 2, 1, 1),
('SP', 'Surat Peringatan', 'Master surat peringatan', 3, 1, 1),
('SIK', 'Surat Ijin Kegiatan', 'Master surat ijin kegiatan', 4, 1, 1),
('UND', 'Surat Undangan', 'Master surat undangan', 5, 1, 1),
('MEM', 'Surat Memorandum', 'Master surat memorandum', 6, 1, 1),
('ND', 'Nota Dinas', 'Master nota dinas', 7, 1, 1);

INSERT INTO config_nomor_surat
(takah_id, company_code, division_code, reset_type, last_number)
SELECT id, 'CBN', NULL, 'monthly', 0 FROM master_takah;

INSERT INTO template_surat (takah_id, template_name, content)
VALUES
((SELECT id FROM master_takah WHERE code = 'UND'), 'Template Surat Undangan',
'Dengan hormat,

Kami mengundang {{nama_tujuan}} untuk menghadiri kegiatan {{nama_kegiatan}} yang akan dilaksanakan pada {{tanggal_kegiatan}} bertempat di {{tempat_kegiatan}}.

Demikian surat undangan ini kami sampaikan. Atas perhatian dan kehadirannya, kami ucapkan terima kasih.'),

((SELECT id FROM master_takah WHERE code = 'SKET'), 'Template Surat Keterangan',
'Yang bertanda tangan di bawah ini menerangkan bahwa {{nama_pemohon}} benar membutuhkan surat keterangan untuk keperluan {{keperluan}}.

Demikian surat keterangan ini dibuat untuk digunakan sebagaimana mestinya.'),

((SELECT id FROM master_takah WHERE code = 'SKK'), 'Template Surat Keterangan Kerja',
'Yang bertanda tangan di bawah ini menerangkan bahwa {{nama_karyawan}} menjabat sebagai {{jabatan}} dan bekerja pada perusahaan.

Surat keterangan kerja ini dibuat untuk keperluan {{keperluan}}.'),

((SELECT id FROM master_takah WHERE code = 'SP'), 'Template Surat Peringatan',
'Surat peringatan ini diberikan kepada {{nama_karyawan}} atas pelanggaran {{jenis_pelanggaran}} yang dilakukan pada {{tanggal_pelanggaran}}.

Diharapkan yang bersangkutan dapat memperbaiki sikap dan kinerja ke depannya.'),

((SELECT id FROM master_takah WHERE code = 'SIK'), 'Template Surat Ijin Kegiatan',
'Dengan ini kami mengajukan izin kegiatan {{nama_kegiatan}} yang akan dilaksanakan pada {{tanggal_kegiatan}} bertempat di {{tempat_kegiatan}}.

Demikian surat izin kegiatan ini dibuat untuk dipergunakan sebagaimana mestinya.'),

((SELECT id FROM master_takah WHERE code = 'MEM'), 'Template Surat Memorandum',
'Kepada {{nama_tujuan}},

Dengan ini disampaikan memorandum mengenai {{perihal_memo}}.

Demikian memorandum ini dibuat untuk menjadi perhatian.'),

((SELECT id FROM master_takah WHERE code = 'ND'), 'Template Nota Dinas',
'Kepada {{nama_tujuan}},

Nota dinas ini disampaikan terkait {{perihal_nota}}.

Demikian nota dinas ini dibuat untuk ditindaklanjuti sebagaimana mestinya.');

INSERT INTO parameter_surat
(template_id, parameter_name, parameter_key, input_type, is_required)
VALUES
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Undangan'), 'Nama Tujuan', 'nama_tujuan', 'text', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Undangan'), 'Nama Kegiatan', 'nama_kegiatan', 'text', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Undangan'), 'Tanggal Kegiatan', 'tanggal_kegiatan', 'date', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Undangan'), 'Tempat Kegiatan', 'tempat_kegiatan', 'text', TRUE),

((SELECT id FROM template_surat WHERE template_name = 'Template Surat Keterangan'), 'Nama Pemohon', 'nama_pemohon', 'text', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Keterangan'), 'Keperluan', 'keperluan', 'textarea', TRUE),

((SELECT id FROM template_surat WHERE template_name = 'Template Surat Keterangan Kerja'), 'Nama Karyawan', 'nama_karyawan', 'text', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Keterangan Kerja'), 'Jabatan', 'jabatan', 'text', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Keterangan Kerja'), 'Keperluan', 'keperluan', 'textarea', TRUE),

((SELECT id FROM template_surat WHERE template_name = 'Template Surat Peringatan'), 'Nama Karyawan', 'nama_karyawan', 'text', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Peringatan'), 'Jenis Pelanggaran', 'jenis_pelanggaran', 'textarea', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Peringatan'), 'Tanggal Pelanggaran', 'tanggal_pelanggaran', 'date', TRUE),

((SELECT id FROM template_surat WHERE template_name = 'Template Surat Ijin Kegiatan'), 'Nama Kegiatan', 'nama_kegiatan', 'text', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Ijin Kegiatan'), 'Tanggal Kegiatan', 'tanggal_kegiatan', 'date', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Ijin Kegiatan'), 'Tempat Kegiatan', 'tempat_kegiatan', 'text', TRUE),

((SELECT id FROM template_surat WHERE template_name = 'Template Surat Memorandum'), 'Nama Tujuan', 'nama_tujuan', 'text', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Surat Memorandum'), 'Perihal Memo', 'perihal_memo', 'textarea', TRUE),

((SELECT id FROM template_surat WHERE template_name = 'Template Nota Dinas'), 'Nama Tujuan', 'nama_tujuan', 'text', TRUE),
((SELECT id FROM template_surat WHERE template_name = 'Template Nota Dinas'), 'Perihal Nota', 'perihal_nota', 'textarea', TRUE);
