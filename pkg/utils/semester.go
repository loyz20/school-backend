package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// GenerateSemesterID membentuk ID seperti: "20241" dari "2024/2025" dan semester ke-1
func GenerateSemesterID(tahunAjaran string, semesterKe int) (string, error) {
	parts := strings.Split(tahunAjaran, "/")
	if len(parts) != 2 || len(parts[0]) != 4 {
		return "", fmt.Errorf("format tahun ajaran tidak valid, contoh: 2024/2025")
	}
	if semesterKe != 1 && semesterKe != 2 {
		return "", fmt.Errorf("semester ke harus 1 atau 2")
	}
	return fmt.Sprintf("%s%d", parts[0], semesterKe), nil
}

// ParseSemesterID membagi semester ID seperti "20241" menjadi tahun ajaran dan semester ke
func ParseSemesterID(semesterID string) (tahunAjaran string, semesterKe int, err error) {
	if len(semesterID) != 5 {
		return "", 0, fmt.Errorf("semester ID harus terdiri dari 5 digit, contoh: 20241")
	}

	tahunStr := semesterID[:4]
	semesterKeStr := semesterID[4:]

	semesterKe, err = strconv.Atoi(semesterKeStr)
	if err != nil || (semesterKe != 1 && semesterKe != 2) {
		return "", 0, fmt.Errorf("digit semester tidak valid, harus 1 atau 2")
	}

	tahun, err := strconv.Atoi(tahunStr)
	if err != nil {
		return "", 0, fmt.Errorf("tahun ajaran tidak valid")
	}

	tahunAjaran = fmt.Sprintf("%d/%d", tahun, tahun+1)
	return tahunAjaran, semesterKe, nil
}
