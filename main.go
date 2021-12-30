package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	router := http.NewServeMux()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		rw.WriteHeader(200)
		rw.Write([]byte(`{status:200}`))
	})

	http.ListenAndServe(":"+port, router)
}
