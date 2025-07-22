package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func computeHash(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(content)
	return hex.EncodeToString(sum[:]), nil
}
