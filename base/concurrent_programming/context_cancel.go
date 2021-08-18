package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go HandelRequest1(ctx)

	time.Sleep(1 * time.Second)
	log.Println("It's time to stop all sub goroutines!")
	cancel()

	//Just for test whether sub goroutines exit or not
	time.Sleep(10 * time.Second)
	log.Println("over")
}

func HandelRequest1(ctx context.Context) {
	go WriteRedis1(ctx)
	go WriteDatabase1(ctx)
	for {
		select {
		case <-ctx.Done():
			log.Println("HandelRequest Done.")
			return
		default:
			log.Println("HandelRequest running")
			time.Sleep(5 * time.Second)
		}
	}
}

func WriteRedis1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("WriteRedis Done.")
			return
		default:
			log.Println("WriteRedis running")
			time.Sleep(5 * time.Second)
		}
	}
}

func WriteDatabase1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("WriteDatabase Done.")
			return
		default:
			log.Println("WriteDatabase running")
			time.Sleep(5 * time.Second)
		}
	}
}

