package main

import (
	"MyGolang/Misc"
	"log"
)

func main() {
	if err := Misc.AudioRun(); err != nil {
		log.Fatal(err)
	}
}
