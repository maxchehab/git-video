package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, r.URL.Query().Get("v"))
	fmt.Print(formatRequest(r))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func formatRequest(r *http.Request) string {
	var request []string

	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)

	request = append(request, fmt.Sprintf("Host: %v", r.Host))

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())

	}
	return strings.Join(request, "\n")

}
