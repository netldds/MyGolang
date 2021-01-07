package Algorithm

import (
	"fmt"
	"os"
	"os/signal"
)

func SignalExample() {

	sigs := make(chan os.Signal)
	done := make(chan bool)

	signal.Notify(sigs, os.Interrupt)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting.")
}
