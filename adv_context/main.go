package main

import (
	"context"
	"fmt"
	"time"
)

func tick(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Microsecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-ctx.Done():
			fmt.Println("cancel")
			return
		}

	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go tick(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
