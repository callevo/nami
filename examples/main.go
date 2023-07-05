package main

import "github/callevo/nami"

var TLogger = nami.TLogger

func main() {
	cl := nami.NewClient()
	err := cl.Connect(
		nami.Options{
			Username: "admin",
			Password: "",
			Port:     5038,
			Events:   "on",
			Url:      " nats://admin:q57jDLKGbPzz3PNs9yDwsz@172.30.0.200:4222",
		})
	if err != nil {
		return
	}
}
