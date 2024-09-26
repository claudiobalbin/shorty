package services

import (
	"crypto/sha256"
	"errors"
	"shorty/utils"
	"time"

	"github.com/jxskiss/base62"
	"golang.org/x/exp/rand"
)

type ShortenerService struct {
}

func NewShortenerService() *ShortenerService {
	return &ShortenerService{}
}

func (s *ShortenerService) GenerateShortURL(longURL string) (string, error) {
	if !utils.ValidURL(longURL) {
		return "", errors.New("invalid url")
	}

	// Generate a random salt to add entropy
	rand.Seed(uint64(time.Now().UnixNano()))
	salt := make([]byte, 8)
	rand.Read(salt)

	// Combine the long URL and salt
	input := append([]byte(longURL), salt...)

	// Calculate the SHA-256 hash
	hash := sha256.Sum256(input)

	// Encode the hash using base62 for a shorter and more readable URL
	encodedHash := base62.EncodeToString(hash[:])

	// Take the first 8 characters of the encoded hash as the short URL
	return encodedHash[:8], nil
}
