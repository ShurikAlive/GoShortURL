package main

import (
	"ShortUrl/urls"
	"net/http"
)


func main() {
	http.ListenAndServe(":8181", urls.SetUpURLs())
}
