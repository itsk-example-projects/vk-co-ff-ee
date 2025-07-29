package crypto

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"vk-co-ff-ee/internal/config"
)

// pkcs7Pad applies PKCS#7 padding to data
func pkcs7Pad(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

// pkcs7Unpad removes PKCS#7 padding from data
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, fmt.Errorf("invalid padding")
	}
	padLen := int(data[len(data)-1])
	if padLen == 0 || padLen > blockSize {
		return nil, fmt.Errorf("invalid padding size")
	}
	for i := 0; i < padLen; i++ {
		if data[len(data)-1-i] != byte(padLen) {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	return data[:len(data)-padLen], nil
}

// deriveKey produces a 16‑character key used to encrypt messages when a custom password is provided
func deriveKey(custom string) string {
	plaintext := []byte(custom + viper.GetString(config.KDFSalt)) // Prepare the plaintext for the derivation: the user key followed by the constant suffix
	block, err := aes.NewCipher([]byte(viper.GetString(config.BaseKey)))
	if err != nil {
		// An error here indicates a misconfigured environment.  Return
		// the baseKey as a fallback.
		return viper.GetString(config.BaseKey)
	}
	padded := pkcs7Pad(plaintext, block.BlockSize())
	encrypted := make([]byte, len(padded))
	for i := 0; i < len(padded); i += block.BlockSize() {
		block.Encrypt(encrypted[i:i+block.BlockSize()], padded[i:i+block.BlockSize()])
	}
	b64 := base64.StdEncoding.EncodeToString(encrypted) // Convert the ciphertext to base64 and take the first 16 characters
	if len(b64) < 16 {
		return b64
	}
	return b64[:16]
}

// EncryptVKCoffee applies the VK CO FF EE cipher to the provided plaintext
func EncryptVKCoffee(plain, customKey string) string {
	var prefix string
	var key string
	if customKey != "" {
		prefix = "VK C0 FF EE "
		key = deriveKey(customKey)
	} else {
		prefix = "VK CO FF EE "
		key = viper.GetString(config.BaseKey)
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		// In the unlikely event of an error creating the cipher return
		// the input unmodified.
		return plain
	}
	data := []byte(plain)
	data = pkcs7Pad(data, block.BlockSize())
	encrypted := make([]byte, len(data))
	for i := 0; i < len(data); i += block.BlockSize() {
		block.Encrypt(encrypted[i:i+block.BlockSize()], data[i:i+block.BlockSize()])
	}
	b64 := base64.StdEncoding.EncodeToString(encrypted)
	var sb strings.Builder
	for i := 0; i < len(b64); i++ {
		sb.WriteString(fmt.Sprintf("%02X", b64[i]))
		sb.WriteString(" ")
	}
	hexEncoded := sb.String()
	hexEncoded = strings.TrimRight(hexEncoded, " ")
	suffix := strings.TrimSpace(prefix) // Append the trimmed prefix (without its trailing space) to close the encrypted message (vkcoffee + ciphertext + vkcoffee.trim())
	return prefix + hexEncoded + " " + suffix
}

// DecryptVKCoffee reverses the VK Coffee cipher
func DecryptVKCoffee(cipherText, customKey string) (string, error) {
	s := strings.TrimSpace(cipherText) // Trim outer whitespace and normalise the case for case‑insensitive replacements
	upper := strings.ToUpper(s)
	replacer := strings.NewReplacer( // Remove known prefixes/suffixes and noise tokens
		"VK C0 FF EE", "",
		"VK CO FF EE", "",
		"PP", "",
		"AP ID OG", "",
		"AP ID 0G", "",
		"II", "",
		" ", "",
	)
	cleaned := replacer.Replace(upper)
	raw, err := hex.DecodeString(cleaned) // Convert the hex representation back to bytes, two hex characters become one byte
	if err != nil {
		return "", fmt.Errorf("invalid cipher text: %w", err)
	}
	b64Str := string(raw)                                       // raw now contains ASCII codes representing a base64 string
	cipherBytes, err := base64.StdEncoding.DecodeString(b64Str) // Decode the base64 into the original encrypted bytes
	if err != nil {
		return "", fmt.Errorf("base64 decode failed: %w", err)
	}
	var key string // Determine which key to use.  If the caller supplies a custom key derive a new AES key, otherwise use the fixed one
	if customKey != "" {
		key = deriveKey(customKey)
	} else {
		key = viper.GetString(config.BaseKey)
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("invalid key: %w", err)
	}
	if len(cipherBytes)%block.BlockSize() != 0 {
		return "", fmt.Errorf("cipher length not multiple of block size")
	}
	decrypted := make([]byte, len(cipherBytes))
	for i := 0; i < len(cipherBytes); i += block.BlockSize() {
		block.Decrypt(decrypted[i:i+block.BlockSize()], cipherBytes[i:i+block.BlockSize()])
	}
	unpadded, err := pkcs7Unpad(decrypted, block.BlockSize())
	if err != nil {
		return "", err
	}
	return string(unpadded), nil
}
