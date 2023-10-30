package helpers

import (
	"fmt"
	"time"
)

func FormatToIndonesia(datetime time.Time) string {
	months := [...]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	return fmt.Sprintf("%02d %s %04d %02d:%02d:%02d", datetime.Day(), months[datetime.Month()-1], datetime.Year(), datetime.Hour(), datetime.Minute(), datetime.Second())
}