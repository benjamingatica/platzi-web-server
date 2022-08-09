package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() MiddleWare {
	return func(f http.HandlerFunc) http.HandlerFunc {
		fmt.Println("first function CheckAuth")
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("second function CheckAuth")
			flag := true
			fmt.Println("Checking authentication")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

/* func CheckAuth2() MiddleWare {
	return func(f http.HandlerFunc) http.HandlerFunc {
		fmt.Println("first function CheckAuth2")
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("second function CheckAuth2")
			flag := true
			fmt.Println("Checking authentication 2")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}
*/

func Logging() MiddleWare {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
