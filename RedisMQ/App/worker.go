package App

import "fmt"

type Worker struct {
}

func NewWorker() *Worker {
	return &Worker{}
}

func (this *Worker) doProcess() {
	fmt.Println("doProcess")
}

func (this *Worker) OriginProcess() {
	fmt.Println("OriginProcess")
}
