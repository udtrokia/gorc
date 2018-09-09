package main;

import (
	"runtime"
	"fmt"
	"sync"
	"time"
	"math/rand"
)

type State struct{
	mu    sync.Mutex
	state int
}

var state = State{state: 0}

func routine() {
	for i := 0; i <= 1000000; i++ {	
		go func() {
			if state.state < i { state.state = i }
			state.mu.Lock()
			tick := time.Duration(rand.Intn(100000))
			time.Sleep(tick * time.Millisecond);
			state.mu.Unlock();
			fmt.Printf("\r %d\n", i);
		}()
	}
	fmt.Printf("state: %v\n", state.state);		
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU());
	routine();
}
