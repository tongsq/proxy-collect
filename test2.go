package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range ticker.C {
			fmt.Println("ticked at %v", time.Now())
		}
	}()
	ticker2 := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range ticker2.C {
			fmt.Println("ticked22222 at %v", time.Now())
		}
	}()
	select {}
}
