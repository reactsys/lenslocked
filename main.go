package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "Welcome to my awesome site!")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an emial to <a href=\"mailto:support@lenlocked.com\"> support@lenslocked.com</a>")
	} else {
		fmt.Fprint(w, "Sorry, but an error occured")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
