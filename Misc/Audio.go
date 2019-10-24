package Misc

import (
	"fmt"
	"io"
	"os"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"
)

func AudioRun() error {
	f, err := os.Open("./DATA/apple.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer p.Close()

	fmt.Printf("Length: %d[bytes]\n", d.Length())

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}
