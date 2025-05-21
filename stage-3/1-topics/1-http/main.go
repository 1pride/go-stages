package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

type contextKey string

const (
	port          = ":8080"
	keyServerAddr = contextKey("serverAddr")
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")

	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s\\n", ctx.Value(keyServerAddr), hasFirst, first,
		hasSecond, second)

	_, err := io.WriteString(w, "website room\n")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot) //routing

	ctx, cancelCtx := context.WithCancel(context.Background())

	server := &http.Server{
		Addr:    port,
		Handler: mux,
		BaseContext: func(listener net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, listener.Addr().String())
			return ctx
		},
	}

	go func() {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("server 1 closed")
		} else if err != nil {
			fmt.Println("listening on port:", port)
		}
		cancelCtx()
	}()

	<-ctx.Done()
}
