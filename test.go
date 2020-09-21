package main

import (
	"proxy-collect/scheduler"
	"time"
)

func main() {
	s := scheduler.CheckIp{}
	s.Run()
	time.Sleep(50 * time.Second)
}
