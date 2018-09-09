package main;

import (
	"runtime"
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var (
	state int
	wait sync.WaitGroup
)

func routine(i int) {
	if state < i {
		state = i
	}
	tick := time.Duration(rand.Intn(10000))
	time.Sleep(tick * time.Millisecond);
	wait.Done();
}

func main() {
	runtime.GOMAXPROCS(4);
	for i := 1; i <= 1000000; i++ {
		wait.Add(1)
		go routine(i);
	}
	wait.Wait();
	fmt.Printf("state: %v\n", state);
}
