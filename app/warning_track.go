package main

import (
	"net/http"

	"github.com/unrealities/warning-track/routers"
)

func main() {
	r := routers.Routes()
	http.Handle("/", r)
	http.HandleFunc("/tv", routers.RedirectHandler("/tv.html"))
	http.HandleFunc("/faq", routers.RedirectHandler("/faq.html"))
}
