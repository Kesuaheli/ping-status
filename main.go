package main

import (
	"context"
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
		logf(LogLevelError, "Usage: %s <url>", os.Args[0])
		os.Exit(1)
	}

	r, err := http.NewRequest(METHOD, os.Args[1], nil)
	if err != nil {
		logf(LogLevelFatal, "Creating request: %+v", err)
	}
	logf(LogLevelDebug, "Starting %s requests to %s", METHOD, r.URL)
	go doRequest(r)

	log(LogLevelInfo, "Press Ctrl+C to exit\n")
	<-ctx.Done()
	log(LogLevelInfo, "\n Stopping!\n")
	os.Exit(0)
}

func doRequest(r *http.Request) {
	for {
		resp, err := http.DefaultClient.Do(r)
		if err != nil {
			logf(LogLevelFatal, "Doing request: %+v", err)
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			// Status 2xx
			logf(LogLevelInfo, "%s;%s;%s", METHOD, r.URL.Host, resp.Status)
		} else {
			logf(LogLevelError, "%s;%s;%s;%+v", METHOD, r.URL.Host, resp.Status, resp.Header)
		}
		time.Sleep(SLEEP_SEC * time.Second)
	}
}
