package main

import (
	"context"
	"fmt"
	"time"
)

func timeContext() {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Done task")
	case <-ctxWithTimeout.Done():
		fmt.Println("Timeout task")
	}
}
