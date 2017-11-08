package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var userAgent = r.Header.Get("User-Agent")
	if strings.Contains(userAgent, "github-camo") {
		http.Redirect(w, r, "https://img.youtube.com/vi/"+r.URL.Query().Get("v")+"/0.jpg", 301)
	} else {
		http.Redirect(w, r, "https://youtube.com/watch?v="+r.URL.Query().Get("v"), 301)
	}
}

func main() {
	fmt.Print("> Listening on http://localhost:8080\n")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
