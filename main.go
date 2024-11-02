package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	METHOD    = "HEAD"
	SLEEP_SEC = 5
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	if len(os.Args) < 2 {
		fmt.Printf("[ERROR] Usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	r, err := http.NewRequest(METHOD, os.Args[1], nil)
	if err != nil {
		fmt.Printf("[ERROR] Creating request: %+v\n", err)
		panic(err)
	}
	fmt.Printf("[INFO] Starting %s requests to %s\n", METHOD, r.URL)
	go doRequest(r)

	fmt.Println("[INFO] Press Ctrl+C to exit")
	fmt.Println("")
	<-ctx.Done()
	fmt.Println("\n[INFO] Stopping!")
	fmt.Println("")
	os.Exit(0)
}

func doRequest(r *http.Request) {
	for {
		resp, err := http.DefaultClient.Do(r)
		if err != nil {
			fmt.Printf("[ERROR] Doing request: %+v\n", err)
			panic(err)
		}

		fmt.Printf("[INFO] %s;%s;%s\n", METHOD, r.URL.Host, resp.Status)
		time.Sleep(SLEEP_SEC * time.Second)
	}
}
