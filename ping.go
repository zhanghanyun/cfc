package main

import (
	"context"
	"github.com/charmbracelet/log"
	"net"
	"net/http"
	"time"
)

func httping(ip string) time.Duration {
	hc := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return (&net.Dialer{Timeout: 5 * time.Second}).DialContext(ctx, network, ip+":443")
			},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	success := 0
	var delay time.Duration
	for i := 0; i < 1; i++ {
		req, err := http.NewRequest(http.MethodGet, "https://ptchdbits.co/", nil)
		if err != nil {
			return time.Hour
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
		startTime := time.Now()
		resp, err := hc.Do(req)
		if err != nil || resp.StatusCode >= 400 {
			return time.Hour
		}
		resp.Body.Close()
		success++
		duration := time.Since(startTime)
		delay += duration
	}
	return delay
}

func tcping(ip string) time.Duration {
	startTime := time.Now()
	conn, err := net.DialTimeout("tcp", ip+":443", 5*time.Second)
	if err != nil {
		log.Error(err)
		return time.Hour
	}
	defer conn.Close()
	duration := time.Since(startTime)
	return duration
}
