package main

import (
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/gorilla/mux"
	"net/http"
)

var keyPair KeyPair


func main() {
	//var mt = "20181111"
	//var pn = "18811881188"
	//var ln = "001"
	//var mn = "importantmeeting"
	//var rn = "216"
	//data := "This IS HIDDEN MESSAGE"
	//hdata := calculateHashcode(data)
	//fmt.Println("string:", data)
	//fmt.Println("sha256 encrypted:", hdata)
	//bdata := []byte(hdata)

	//Generate keys
	prk, err := getKey()

	if err != nil{
		panic(err)
	}

	prk2 := ecies.ImportECDSA(prk)
	keyPair = KeyPair{prk2.PublicKey, prk2}

	//endata, err := ECCEncrypt([]byte(bdata), puk2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("ecc public key encrypted:", hex.EncodeToString(endata))
	//dedata, err := ECCDecrypt(endata, *prk2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Private Key Decryption:", string(dedata))
	//fmt.Println("Decrypted:", string(dedata))

	r := mux.NewRouter()
	r.HandleFunc("/getPublicKey", getPublicKey)

	r.HandleFunc("/form", use(myHandler, basicAuth))
	r.HandleFunc("/hello", hello)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

