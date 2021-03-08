package terra

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/pkg/errors"
)

const (
	NonceSize = 12
)

func EncryptAESGCM(key []byte, data []byte) ([]byte, error) {
	aesCypher, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, errors.Wrap(err, "new cipher")
	}

	gcm, err := cipher.NewGCM(aesCypher)
	if err != nil {
		return nil, errors.Wrap(err, "new AEAD")
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, errors.Wrap(err, "fetch random nonce")
	}
	return gcm.Seal(nonce, nonce, data, nil), nil
}

func DecryptAESGCM(key []byte, nonce []byte, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "new cipher")
	}

	aead, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return nil, errors.Wrap(err, "new AEAD")
	}

	decrypted, err := aead.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, errors.Wrap(err, "decrypt data")
	}
	return decrypted, nil
}
