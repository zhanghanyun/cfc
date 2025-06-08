package main

import (
	"encoding/json"
	"github.com/charmbracelet/log"
	"io"
	"net/http"
)

func getIps() []string {
	response, err := http.Get("https://vps789.com/openApi/cfIpApi")
	if err != nil {
		log.Fatal("Get ips error")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := map[string]interface{}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	ips := []string{}
	for _, v := range data["data"].(map[string]interface{}) {
		for _, ip := range v.([]interface{}) {
			for k, vv := range ip.(map[string]interface{}) {
				if k == "ip" {
					ips = append(ips, vv.(string))
				}
			}
		}
	}
	return ips
}
