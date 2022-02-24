package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range ticker.C {
			fmt.Println("ticked at", time.Now())
		}
	}()
	ticker2 := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range ticker2.C {
			fmt.Println("ticked22222 at", time.Now())
		}
	}()
	select {}
}
