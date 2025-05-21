package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	port = 8080
)

func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s/\n", r.Method)
			fmt.Printf("server: query id %s\n", r.URL.Query().Get("id"))
			fmt.Printf("server: content-type %s\n", r.Header.Get("Content-Type"))
			fmt.Printf("server: headers:\n")
			for headerName, headerValue := range r.Header {
				fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
			}

			reqBody, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("server could not read request body: %s\n", err)
			}
			fmt.Printf("server: request body: %s\n", reqBody)

			fmt.Fprintf(w, `{"message": "hello!"}`)
			//time.Sleep(11 * time.Second)
		})

		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
			fmt.Println("listening on port:", port)
		}
	}()

	time.Sleep(time.Second) //sleep to routine start

	jsonBody := []byte(`{client_message": "hello!"}`)
	bodyReader := bytes.NewReader(jsonBody)

	// /POST
	reqURL := fmt.Sprintf("http://localhost:%d?id=1234", port)
	req, err := http.NewRequest(http.MethodPost, reqURL, bodyReader)
	if err != nil {
		fmt.Printf("error creating http request %s\n", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: time.Second * 10,
	}

	// /GET
	//req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	//if err != nil {
	//	fmt.Printf("error creating http request %s\n", err)
	//	os.Exit(1)
	//}

	resp, err := client.Do(req)
	//resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error running http request %s\n", err)
		os.Exit(1)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading response body %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: status code: %d\n", resp.StatusCode)
	fmt.Printf("client: response body: %s\n", respBody)
}
