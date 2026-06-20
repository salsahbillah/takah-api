package helper

import (
	"fmt"
	"time"
)

func GenerateNomorSurat(lastNumber int, takahCode string, companyCode string, resetType string) string {
	nextNumber := lastNumber + 1
	now := time.Now()

	sequence := fmt.Sprintf("%03d", nextNumber)

	var dateCode string
	if resetType == "yearly" {
		dateCode = now.Format("2006")
	} else {
		dateCode = now.Format("012006")
	}

	return fmt.Sprintf("%s/%s/%s/%s", sequence, takahCode, companyCode, dateCode)
}