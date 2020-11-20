package ModuleExample

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func MutexSample() {
	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	var state = make(map[int]int)

	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for j := 0; j < 10; j++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				state[key] = value
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)

	fmt.Printf("read:%d\n", readOpsFinal)
	fmt.Printf("write:%d\n", writeOpsFinal)

	mutex.Lock()
	fmt.Println(state)
	mutex.Unlock()
}
