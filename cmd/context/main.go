package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct{}

func main() {
	c := context.WithValue(context.Background(), paramKey{}, "val")
	c, cancelFunc := context.WithTimeout(c, 20*time.Second)
	defer cancelFunc()
	go mainTask(c)
	var cmd string
	for {
		fmt.Scan(&cmd)
		if cmd == "c" {
			cancelFunc()
			time.Sleep(1 * time.Second)
			break
		}
	}
}

func mainTask(c context.Context) {
	fmt.Printf("main task started with param %q\n", c.Value(paramKey{}))
	go func() {
		c1, cancel := context.WithTimeout(c, 10*time.Second)
		defer cancel()
		smallTask(c1, "task1", 8*time.Second)
	}()
	smallTask(c, "task2", 8*time.Second)
}

func smallTask(c context.Context, name string, d time.Duration) {
	fmt.Printf("%s start with param %q\n", name, c.Value(paramKey{}))
	select {
	case <-time.After(d):
		fmt.Printf("%s done\n", name)
	case <-c.Done():
		fmt.Printf("%s cancelled\n", name)
	}
}
