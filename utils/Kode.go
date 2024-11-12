package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomCode(length int) string {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[r.Intn(len(charset))]
	}
	return string(code)
}

// Fungsi untuk mengecek duplikasi
func CheckDuplicates(kodeVotes []string) bool {
	// Membuat map untuk melacak item yang sudah ditemukan
	seen := make(map[string]bool)

	// Memeriksa setiap kode dalam slice
	for _, kode := range kodeVotes {
		if seen[kode] {
			// Jika kode sudah ada di map, berarti ada duplikasi
			return true
		}
		seen[kode] = true
	}
	// Jika tidak ada duplikasi
	return false
}
