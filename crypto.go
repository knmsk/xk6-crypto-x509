package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"hash"
	"regexp"
	"strings"

	"log"

	"go.k6.io/k6/js/modules"
)

var mapOfHashes map[string]hash.Hash = make(map[string]hash.Hash)

func init() {
	mapOfHashes["sha512"] = sha512.New()
	mapOfHashes["sha256"] = sha256.New()
	modules.Register("k6/x/crypto-x509", new(CRYPTO))
}

type CRYPTO struct{}

func parseRSAPublicKey(pemDecodedBlock []byte) *rsa.PublicKey {
	key, err := x509.ParsePKCS1PublicKey(pemDecodedBlock)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

func parsePublicKeyToRsa(pemDecodedBlock []byte) *rsa.PublicKey {
	keyInstance, err := x509.ParsePKIXPublicKey(pemDecodedBlock)
	if err != nil {
		log.Fatal(err)
	}
	key, ok := keyInstance.(*rsa.PublicKey)
	if !ok {
		log.Fatal("failed to key rsa")
	}
	return key
}

func loadPublicKey(publicKey string) *rsa.PublicKey {
	isRSAPublicKey := strings.Contains(publicKey, "RSA")
	block, _ := pem.Decode([]byte(publicKey))
	if isRSAPublicKey {
		return parseRSAPublicKey(block.Bytes)
	} else {
		return parsePublicKeyToRsa(block.Bytes)
	}
}

func encryptData(data []byte, pubKey *rsa.PublicKey, hash hash.Hash) []byte {
	encrypted, err := rsa.EncryptOAEP(hash, rand.Reader, pubKey, data, nil)
	if err != nil {
		log.Fatal("failed to encrypt the data")
	}
	return encrypted
}

func removeExtraSpaces(key string) string {
	whitespacesBlock := regexp.MustCompile(`(\s{2,})`)
	allFormatted := whitespacesBlock.ReplaceAllString(key, "\n")

	return allFormatted
}

func (*CRYPTO) EncryptRsa(publicKey string, data string, algorithm string) string {
	if algorithm == "" {
		algorithm = "sha512"
	}
	hash := mapOfHashes[algorithm]

	publicKeyFormatted := removeExtraSpaces(publicKey)
	pubKey := loadPublicKey(publicKeyFormatted)
	encryptedData := encryptData([]byte(data), pubKey, hash)
	return string(encryptedData)
}
