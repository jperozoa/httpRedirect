package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

var debug = false

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("NEW REQUEST: %v - %v %v\n", time.Now(), r.Method, r.URL.RequestURI())

	if debug {
		b, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(b))
	}

	newRequest, err := http.NewRequest(r.Method, redirectURL+r.URL.RequestURI(), r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	newRequest.Header = r.Header

	client := &http.Client{}

	newResponse, err := client.Do(newRequest)
	if err != nil {
		log.Fatalln(err)
	}
	defer newResponse.Body.Close()

	b, err := io.ReadAll(newResponse.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(b)
}

var redirectURL = "http://localhost:8070"

func main() {
	redirectURL, _ = os.LookupEnv("REDIRECT_BASE_URL")
	if len(os.Args) >= 2 {
		redirectURL = os.Args[1]
	}
	fmt.Printf("Redirecting to: %v\n", redirectURL)

	mux := http.NewServeMux()

	mux.HandleFunc("/", redirect)

	http.ListenAndServe(":6000", mux)
}
