package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go HandelRequest1(ctx)

	time.Sleep(1 * time.Second)
	fmt.Println("It's time to stop all sub goroutines!")
	cancel()

	//Just for test whether sub goroutines exit or not
	time.Sleep(10 * time.Second)
}

func HandelRequest1(ctx context.Context) {
	go WriteRedis1(ctx)
	go WriteDatabase1(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(5 * time.Second)
		}
	}
}

func WriteRedis1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(5 * time.Second)
		}
	}
}

func WriteDatabase1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(5 * time.Second)
		}
	}
}

