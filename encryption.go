package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"strconv"
)

func ECCEncrypt(pt []byte, puk ecies.PublicKey) ([]byte, error) {
	ct, err := ecies.Encrypt(rand.Reader, &puk, pt, nil, nil)
	return ct, err
}

func ECCDecrypt(ct []byte, prk ecies.PrivateKey) ([]byte, error) {
	pt, err := prk.Decrypt(ct, nil, nil)
	return pt, err
}
func getKey() (*ecdsa.PrivateKey, error) {
	prk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return prk, err
	}
	return prk, nil
}

func calculateHashcode(data string) string {
	nonce := 0
	var str string
	var check string
	pass := false
	var dif int = 4
	for nonce = 0; ; nonce++ {
		str = ""
		check = ""
		check = data + strconv.Itoa(nonce)
		h := sha256.New()
		h.Write([]byte(check))
		hashed := h.Sum(nil)
		str = hex.EncodeToString(hashed)
		for i := 0; i < dif; i++ {
			if str[i] != '0' {
				break
			}
			if i == dif-1 {
				pass = true
			}
		}
		if pass == true {
			return str
		}
	}
}
