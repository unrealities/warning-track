package main

/*
MLB Leverage Index

The values are from: http://www.insidethebook.com/li.shtml

http://thenewstack.io/make-a-restful-json-api-go/
*/

import "net/http"

func init() {
	r := Routes()
	http.Handle("/", r)
	http.HandleFunc("/tv", redirectHandler("/tv.html"))
}
