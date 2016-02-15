package areena

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// DecryptURL decrypts URL according to http://developer.yle.fi/tutorial-playing-a-program/
func DecryptURL(secret string, encStr string) (url string, err error) {
	encrypted := []byte(encStr)
	decoded := make([]byte, len(encrypted))
	n, err := base64.StdEncoding.Decode(decoded, encrypted)
	if err != nil {
		return
	}
	ciph, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return
	}
	iv := decoded[:16]
	msg := decoded[16:n]
	dec := cipher.NewCBCDecrypter(ciph, iv)
	dst := make([]byte, len(msg))
	dec.CryptBlocks(dst, msg)
	return stripInvalidChars(string(dst)), nil
}

func stripInvalidChars(str string) string {
	b := make([]byte, len(str))
	var bl int
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c >= 32 && c < 127 {
			b[bl] = c
			bl++
		}
	}
	return string(b[:bl])
}
