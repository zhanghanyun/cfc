package main

import (
	"github.com/charmbracelet/log"
	"github.com/goodhosts/hostsfile"
)

func setHost(ip, host string) {
	//log.Info("Hello World")
	hosts, err := hostsfile.NewHosts()
	if err != nil {
		log.Fatal(err)
	}
	if err = hosts.RemoveByHostname(host); err != nil {
		log.Fatal(err)
	}
	if err = hosts.Add(ip, host); err != nil {
		log.Fatal(err)
	}
	if err = hosts.Flush(); err != nil {
		log.Fatal(err)
	}
	log.Infof("Add Host: %v %v ✅", ip, host)
}
