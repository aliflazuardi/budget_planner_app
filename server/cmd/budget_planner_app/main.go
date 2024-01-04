package main

import "net/http"

// main app for backend web server and api
func main() {
	// start http server
	http.ListenAndServe(":8080", nil)
}
