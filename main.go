package main

import (
	"context"
	"net/http"
	"net/url"
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
		logf(LogLevelError, "Usage: %s <url>", os.Args[0])
		os.Exit(1)
	}

	r, err := http.NewRequest(METHOD, os.Args[1], nil)
	if err != nil {
		logf(LogLevelFatal, "Creating request: %+v", err)
	}
	logf(LogLevelDebug, "Starting %s requests to %s", METHOD, r.URL)
	go doRequest(r)

	log(LogLevelDebug, "Press Ctrl+C to exit\n")
	<-ctx.Done()
	log(LogLevelInfo, "\n Stopping!\n")
	os.Exit(0)
}

func doRequest(r *http.Request) {
	var count uint64
	for {
		count++
		resp, err := http.DefaultClient.Do(r)
		if err != nil {
			logf(LogLevelError, "#%06X;%s;%s;;request err: %+v", count, METHOD, r.URL.Host, err.(*url.Error))
		} else if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			// Status 2xx
			logf(LogLevelInfo, "#%06X;%s;%s;%s", count, METHOD, r.URL.Host, resp.Status)
		} else {
			// Status 1xx, 3xx, 4xx, 5xx
			logf(LogLevelError, "#%06X;%s;%s;%s;bad status: %+v", count, METHOD, r.URL.Host, resp.Status, resp.Header)
		}
		time.Sleep(SLEEP_SEC * time.Second)
	}
}
