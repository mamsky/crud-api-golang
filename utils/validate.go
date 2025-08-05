package utils

import "regexp"

func IsValidEmail(email string) bool {
	// Regex sederhana untuk validasi email
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return regex.MatchString(email)
}

// IsValidPhone memvalidasi format nomor telepon
func IsValidPhone(phone string) bool {
	// Contoh validasi sederhana: hanya angka, panjang minimal 10 dan maksimal 15
	regex := regexp.MustCompile(`^\d{10,15}$`)
	return regex.MatchString(phone)
}