package main

/*
MLB Leverage Index

The values are from: http://www.insidethebook.com/li.shtml

http://thenewstack.io/make-a-restful-json-api-go/
http://www.alexedwards.net/blog/golang-response-snippets#json
*/

import "net/http"

func init() {
	r := Routes()
	http.Handle("/", r)
}

func main() {
	r := Routes()
	http.ListenAndServe(":8080", r)
}
