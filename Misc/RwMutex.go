package Misc

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex

var l *sync.WaitGroup

func RwMutexExample() {

	m = new(sync.RWMutex)
	l = new(sync.WaitGroup)
	l.Add(2)
	go readAsset(1)
	go writeAsset(2)
	l.Wait()
}

func readAsset(i int) {
	fmt.Println(i, "read start")
	m.RLock()
	fmt.Println(i, "reading")
	time.Sleep(2 * time.Second)
	fmt.Println(i, "read over")

	defer func() {
		//time.Sleep(1 * time.Second)

		m.RUnlock()
		l.Done()
	}()
}

func writeAsset(i int) {
	fmt.Println(i, "write start")
	time.Sleep(1 * time.Second)
	m.Lock()
	fmt.Println(i, "writing")
	//time.Sleep(1 * time.Second)
	fmt.Println(i, "write over")
	defer func() {
		m.Unlock()
		l.Done()
	}()
}

func Add(x, y int) int {
	return x + y
}
