package faspaydisbursementapi

import (
	"crypto/md5"
	"crypto/sha256"
	"github.com/spacemonkeygo/openssl"
)

// Crypter ...
type Crypter struct {
	key    []byte
	iv     []byte
	cipher *openssl.Cipher
}

// NewCrypter ...
func NewCrypter(key []byte, iv []byte) (*Crypter, error) {
	cipher, err := openssl.GetCipherByName("aes-256-cbc")
	if err != nil {
		return nil, err
	}

	return &Crypter{key, iv, cipher}, nil
}

// Encrypt ...
func (c *Crypter) Encrypt(input []byte) ([]byte, error) {
	ctx, err := openssl.NewEncryptionCipherCtx(c.cipher, nil, c.key, c.iv)
	if err != nil {
		return nil, err
	}

	cipherbytes, err := ctx.EncryptUpdate(input)
	if err != nil {
		return nil, err
	}

	finalbytes, err := ctx.EncryptFinal()
	if err != nil {
		return nil, err
	}

	cipherbytes = append(cipherbytes, finalbytes...)
	return cipherbytes, nil
}

// GenerateIv ...
func GenerateIv(iv string) []byte {
	convertMd5 := md5.Sum([]byte(iv))
	return convertMd5[len(convertMd5)-16:]
}

// GeneratePassword ...
func GeneratePassword(faspaysecret string) []byte {
	hashing := sha256.Sum256([]byte(faspaysecret))
	return hashing[0:32]
}
