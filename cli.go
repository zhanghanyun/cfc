package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"net/http"
	"sync"
	"time"
)

type CFIp struct {
	Ip    string
	Delay time.Duration
}

var (
	wg    sync.WaitGroup
	tasks []CFIp
	host  string
)

func run() {
	go func() {
		http.Get("https://hitscounter.dev/api/hit?url=https%3A%2F%2Fgithub.com%2Fzhanghanyun%2Fcfc")
	}()
	ips := getIps()
	if len(ips) == 0 {
		log.Fatal("no ips found")
	}

	log.Infof("Get ips: %v", len(ips))

	for _, ip := range ips {
		wg.Add(1)
		go func() {
			defer wg.Done()
			delay := httping(ip)
			tasks = append(tasks, CFIp{ip, delay})
		}()
	}
	wg.Wait()
	setHost(tasks[0].Ip, host)
	fmt.Println("按下 回车键 或 Ctrl+C 退出。")
	fmt.Scanln()
}
