package main

import "github.com/ethereum/go-ethereum/crypto/ecies"

type generatedPublicKey struct {
	publicKey  string
}

type failedToGetPublicKey struct {
	errorMessage string
}

type KeyPair struct{
	publicKey ecies.PublicKey
	privateKey *ecies.PrivateKey
}