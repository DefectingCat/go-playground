package main

import (
	"fmt"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}

	_, err = fmt.Fprintf(w, "Hello world!")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		return
	}
}
