package handler

import (
	"net/http"

	"takah-api/internal/database"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	userID := c.GetInt("user_id")
	role := c.GetString("user_role")

	isAdmin := role == "admin"

	var totalSurat int
	var suratKeluar int
	var suratMasuk int
	var approval int
	var pending int

	if isAdmin {
		database.DB.QueryRow(`SELECT COUNT(*) FROM surat_keluar`).Scan(&suratKeluar)
		database.DB.QueryRow(`SELECT COUNT(*) FROM surat_masuk`).Scan(&suratMasuk)
		database.DB.QueryRow(`SELECT COUNT(*) FROM approval_surat`).Scan(&approval)
		database.DB.QueryRow(`
			SELECT COUNT(*) 
			FROM monitoring_surat 
			WHERE status = 'pending'
		`).Scan(&pending)
	} else {
		database.DB.QueryRow(`
			SELECT COUNT(*) 
			FROM surat_keluar 
			WHERE created_by = ?
		`, userID).Scan(&suratKeluar)

		database.DB.QueryRow(`
			SELECT COUNT(*) 
			FROM surat_masuk 
			WHERE created_by = ?
		`, userID).Scan(&suratMasuk)

		database.DB.QueryRow(`
			SELECT COUNT(*) 
			FROM approval_surat a
			JOIN surat_keluar sk ON sk.id = a.surat_keluar_id
			WHERE sk.created_by = ?
		`, userID).Scan(&approval)

		database.DB.QueryRow(`
			SELECT COUNT(*) 
			FROM monitoring_surat m
			LEFT JOIN surat_keluar sk ON sk.id = m.surat_keluar_id
			LEFT JOIN surat_masuk sm ON sm.id = m.surat_masuk_id
			WHERE m.status = 'pending'
			AND (sk.created_by = ? OR sm.created_by = ?)
		`, userID, userID).Scan(&pending)
	}

	totalSurat = suratKeluar + suratMasuk

	chart := getDashboardChart(userID, isAdmin)
	activities := getDashboardActivities(userID, isAdmin)
	latestLetters := getLatestLetters(userID, isAdmin)

	c.JSON(http.StatusOK, gin.H{
		"message": "Data dashboard berhasil diambil",
		"data": gin.H{
			"summary": gin.H{
				"total_surat":  totalSurat,
				"surat_keluar": suratKeluar,
				"surat_masuk":  suratMasuk,
				"approval":     approval,
				"pending":      pending,
			},
			"chart":          chart,
			"activities":     activities,
			"latest_letters": latestLetters,
		},
	})
}

func getDashboardChart(userID int, isAdmin bool) []gin.H {
	months := []string{
		"Jan", "Feb", "Mar", "Apr", "Mei", "Jun",
		"Jul", "Agu", "Sep", "Okt", "Nov", "Des",
	}

	result := make([]gin.H, 0)

	for index, month := range months {
		monthNumber := index + 1

		var keluar int
		var masuk int

		if isAdmin {
			database.DB.QueryRow(`
				SELECT COUNT(*) 
				FROM surat_keluar 
				WHERE MONTH(tanggal_surat) = ?
			`, monthNumber).Scan(&keluar)

			database.DB.QueryRow(`
				SELECT COUNT(*) 
				FROM surat_masuk 
				WHERE MONTH(tanggal_surat) = ?
			`, monthNumber).Scan(&masuk)
		} else {
			database.DB.QueryRow(`
				SELECT COUNT(*) 
				FROM surat_keluar 
				WHERE MONTH(tanggal_surat) = ?
				AND created_by = ?
			`, monthNumber, userID).Scan(&keluar)

			database.DB.QueryRow(`
				SELECT COUNT(*) 
				FROM surat_masuk 
				WHERE MONTH(tanggal_surat) = ?
				AND created_by = ?
			`, monthNumber, userID).Scan(&masuk)
		}

		result = append(result, gin.H{
			"month":  month,
			"keluar": keluar,
			"masuk":  masuk,
		})
	}

	return result
}

func getDashboardActivities(userID int, isAdmin bool) []gin.H {
	query := `
		SELECT nomor_surat, status, updated_at
		FROM monitoring_surat
		ORDER BY updated_at DESC
		LIMIT 3
	`

	args := []interface{}{}

	if !isAdmin {
		query = `
			SELECT m.nomor_surat, m.status, m.updated_at
			FROM monitoring_surat m
			LEFT JOIN surat_keluar sk ON sk.id = m.surat_keluar_id
			LEFT JOIN surat_masuk sm ON sm.id = m.surat_masuk_id
			WHERE sk.created_by = ? OR sm.created_by = ?
			ORDER BY m.updated_at DESC
			LIMIT 3
		`
		args = append(args, userID, userID)
	}

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return []gin.H{}
	}
	defer rows.Close()

	activities := []gin.H{}

	for rows.Next() {
		var nomorSurat string
		var status string
		var updatedAt string

		rows.Scan(&nomorSurat, &status, &updatedAt)

		activities = append(activities, gin.H{
			"title":    "Surat " + nomorSurat + " " + status,
			"subtitle": "Status surat diperbarui",
			"time":     updatedAt,
			"status":   status,
		})
	}

	return activities
}

func getLatestLetters(userID int, isAdmin bool) []gin.H {
	query := `
		SELECT 
			sk.nomor_surat,
			sk.perihal,
			mt.name,
			COALESCE(u.name, 'Admin Takah') AS pengirim,
			DATE_FORMAT(sk.tanggal_surat, '%d %M %Y') AS tanggal_surat,
			sk.status
		FROM surat_keluar sk
		JOIN master_takah mt ON mt.id = sk.takah_id
		LEFT JOIN users u ON u.id = sk.created_by
		ORDER BY sk.created_at DESC
		LIMIT 3
	`

	args := []interface{}{}

	if !isAdmin {
		query = `
			SELECT 
				sk.nomor_surat,
				sk.perihal,
				mt.name,
				COALESCE(u.name, 'User Takah') AS pengirim,
				DATE_FORMAT(sk.tanggal_surat, '%d %M %Y') AS tanggal_surat,
				sk.status
			FROM surat_keluar sk
			JOIN master_takah mt ON mt.id = sk.takah_id
			LEFT JOIN users u ON u.id = sk.created_by
			WHERE sk.created_by = ?
			ORDER BY sk.created_at DESC
			LIMIT 3
		`
		args = append(args, userID)
	}

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return []gin.H{}
	}
	defer rows.Close()

	letters := []gin.H{}

	for rows.Next() {
		var nomorSurat string
		var perihal string
		var jenisSurat string
		var pengirim string
		var tanggalSurat string
		var status string

		rows.Scan(
			&nomorSurat,
			&perihal,
			&jenisSurat,
			&pengirim,
			&tanggalSurat,
			&status,
		)

		letters = append(letters, gin.H{
			"nomor_surat":   nomorSurat,
			"perihal":       perihal,
			"jenis_surat":   jenisSurat,
			"pengirim":      pengirim,
			"tanggal_surat": tanggalSurat,
			"status":        status,
		})
	}

	return letters
}
