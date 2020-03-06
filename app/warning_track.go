package main

import (
	"log"
	"net/http"
	"os"

	"github.com/unrealities/warning-track/routers"
	"google.golang.org/appengine"
)

func main() {
	r := routers.Routes()
	http.Handle("/", r)
	http.HandleFunc("/tv", routers.RedirectHandler("/tv.html"))
	http.HandleFunc("/faq", routers.RedirectHandler("/faq.html"))

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]

	appengine.Main()
}
