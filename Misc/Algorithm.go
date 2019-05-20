package Misc

import (
	"fmt"
	"math/rand"
)

var a chan string

func algorithmStart(name string) {

	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, rand.Intn(100))
	}
	switch name {
	case "QuickSort":
		Misc.QuickSort(a, 0, len(a)-1)
		fmt.Printf("Sorted array is :%v\n", a)

	case "MergeSort":
		Misc.MergeSort(a, 0, len(a)-1)
		fmt.Printf("Sorted array is :%v\n", a)

	default:
		fmt.Println("Correct name was expected.")
	}
}
func MergeSort(a []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		MergeSort(a, low, mid)
		MergeSort(a, mid+1, high)
		merge(a, low, mid, high)
	}
}
func merge(a []int, low int, mid int, high int) {
	N := high - low + 1
	var b = make([]int, N)
	left := low
	right := mid + 1
	bIndex := 0
	for left <= mid && right <= high {
		if a[left] <= a[right] {
			b[bIndex] = a[left]
			bIndex++
			left++
		} else {
			b[bIndex] = a[right]
			bIndex++
			right++
		}
	}
	for left <= mid {
		b[bIndex] = a[left]
		bIndex++
		left++
	}
	for right <= high {
		b[bIndex] = a[right]
		bIndex++
		right++
	}
	for i := 0; i < N; i++ {
		a[i+low] = b[i]
	}
}

func QuickSort(a []int, low, high int) {
	if low < high {
		pivot := partition(a, low, high)
		QuickSort(a, low, pivot)
		QuickSort(a, pivot+1, high)
	}

}

func partition(a []int, low int, high int) int {
	m := low
	p := a[low]
	for k := low + 1; k <= high; k++ {
		if a[k] < p {
			m++
			a[k], a[m] = a[m], a[k]
		}
	}
	a[low], a[m] = a[m], a[low]
	return m
}

type Body struct {
	Value int
	Name  string
}
type Bodys struct {
	bodys []Body
	Flag  bool
}

func (b *Bodys) Sort() {
	b.quickSort(0, len(b.bodys))

}
func (b *Bodys) quickSort(low, height int) {
	if low < height {
		pivot := b.pivot(low, height)
		b.quickSort(low, pivot)
		b.quickSort(pivot+1, height)
	}

}
func (b *Bodys) pivot(low, high int) int {
	pivotValue := b.bodys[low].Value
	pivot := low
	for i := low + 1; i < high; i++ {
		if b.Flag {
			if b.bodys[i].Value < pivotValue {
				pivot++
				b.bodys[i], b.bodys[pivot] = b.bodys[pivot], b.bodys[i]
			}
		} else {
			if b.bodys[i].Value > pivotValue {
				pivot++
				b.bodys[i], b.bodys[pivot] = b.bodys[pivot], b.bodys[i]
			}
		}
	}
	b.bodys[low], b.bodys[pivot] = b.bodys[pivot], b.bodys[low]
	return pivot

}
