package main

import (
	"fmt"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	s1 := 1
	s2 := 2
	fmt.Printf("%v %v", s1, s2)
	s1, s2 = s2, s1
	fmt.Printf("%v %v", s1, s2)
}

func FA() {
	termui.Init()
	defer termui.Close()

	//p := widgets.NewParagraph()
	//p.Border = true
	//p.Text = "Hello World!"
	//p.SetRect(0, 0, 25, 5)
	//termui.Render(p)

	l := widgets.NewBarChart()
	l.Border = true
	l.Title = "list"
	l.SetRect(100, 100, 0, 0)
	termui.Render(l)

	for e := range termui.PollEvents() {
		if e.Type == termui.KeyboardEvent {
			break
		}
	}
}
