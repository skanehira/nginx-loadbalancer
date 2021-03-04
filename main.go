package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		apiname := os.Getenv("API_NAME")
		fmt.Fprintln(w, apiname)
		b, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		fmt.Fprintln(w, string(b))
	})
	log.Println("start http server :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
