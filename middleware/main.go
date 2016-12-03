package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page!")
	fmt.Println("about")
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome Page!")
	fmt.Println("welcome")
}

func midLogger(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("log")
		next.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), end.Sub(start))
		fmt.Println("log done")
	}
}

func midAuth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authenticating")
		next.ServeHTTP(w, r)
		panic("auth panicked")
		fmt.Println("authenticating done")
	}
}

func midRecover(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered", r)
			} else {
				fmt.Println("not recovered")
			}
		}()

		next.ServeHTTP(w, r)
	}
}

func main() {

	http.HandleFunc("/about", midLogger(aboutHandler))
	http.HandleFunc("/", midRecover(midAuth(midLogger(handler))))
	http.ListenAndServe(":8080", nil)
}
