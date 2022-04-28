package main

import (
	"log"
	"resource_monitor/routers"
)

func main() {
	r := routers.SetupRouter()
	err := r.Run()
	if err != nil {
		log.Printf("run server failed, err:%v\n", err)
		return
	}
}
