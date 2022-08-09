package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("init HandleHome")
	fmt.Fprintf(w, "Welcome at home!")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprint(w, "error: ", err)
		return
	}

	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Body %T\n", r.Body)
	fmt.Printf("r.Body %v\n", r.Body)
	decoder := json.NewDecoder(r.Body)
	fmt.Printf("decoder %T\n", decoder)
	fmt.Printf("decoder %v\n", decoder)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprint(w, "error: ", err)
		return
	}

	fmt.Printf("user %T\n", user)
	fmt.Printf("user %v\n", user)

	userJ, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("userJ %T\n", userJ)
	fmt.Printf("userJ %v\n", userJ)

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJ)
}
