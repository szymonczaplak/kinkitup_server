package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

// use provides a cleaner interface for chaining middleware for single routes.
// Middleware functions are simple HTTP handlers (w http.ResponseWriter, r *http.Request)
//
//  r.HandleFunc("/login", use(loginHandler, rateLimit, csrf))
//  r.HandleFunc("/form", use(formHandler, csrf))
//  r.HandleFunc("/about", aboutHandler)
//
// See https://gist.github.com/elithrar/7600878#comment-955958 for how to extend it to suit simple http.Handler's
func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

func myHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Authenticated!"))
	return
}

func getPublicKey(w http.ResponseWriter, r *http.Request) {
	puk2 := keyPair.publicKey
	w.Header().Set("Content-Type", "application/json")
	returnStruct := generatedPublicKey{publicKey: hex.EncodeToString(puk2)}
	response, err := json.Marshal(returnStruct)
	if err != nil{
		fmt.Println("Error while creating json response")
		_, _ = w.Write([]byte("Error in json formatting"))
	}
	fmt.Println(response)
	fmt.Println(err)

	_, _ = w.Write(response)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}