package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	// log.Fatal(http.ListenAndServe(":1999", http.FileServer(http.Dir("./"))))
	// demo
	log.Fatal(http.ListenAndServeTLS(":1999", "/etc/fullchain.pem", "/etc/letsencrypt/privkey.pem", http.FileServer(http.Dir("./"))))

}
