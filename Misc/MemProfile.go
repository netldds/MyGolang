package Misc

import (
	"fmt"
	"runtime"
)

//使用对象数量,内分配数量
func PrintUsage() {
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)
	fmt.Printf("SysMB %v	\t", stats.Sys>>20)
	//fmt.Printf("PauseTotalNs %v	\n", stats.PauseTotalNs)
	fmt.Printf("AllocMB %-10v	\t", stats.Alloc>>20)
	//fmt.Printf("Lookups %-10v	\n", stats.Lookups)
	//fmt.Printf("HeapObjects %-10v	\n", stats.HeapObjects)
	//fmt.Printf("live object %v\n", stats.Mallocs-stats.Frees)
	//fmt.Printf("live object %v\n", stats.HeapInuse)
	fmt.Printf("GC %v \t", stats.NumGC)
	fmt.Printf("HEAPMB %v \t\n", stats.HeapAlloc>>20)
	n, ok := runtime.MemProfile(nil, false)
	var p []runtime.MemProfileRecord
	for {
		p = make([]runtime.MemProfileRecord, n+50)
		n, ok = runtime.MemProfile(p, false)
		if ok {
			p = p[0:n]
			break
		}
	}
	var total runtime.MemProfileRecord
	for i := range p {
		r := &p[i]
		total.AllocBytes += r.AllocBytes
		total.AllocObjects += r.AllocObjects
		total.FreeBytes += r.FreeBytes
		total.FreeObjects += r.FreeObjects
	}

	//fmt.Printf("%d in use objects (%d in use bytes) | Alloc: %d TotalAlloc: %d\n",
	//	total.InUseObjects(), total.InUseBytes(), stats.Alloc, stats.TotalAlloc)
	//fmt.Printf("FreeBytes %v MB AllocBytes %v MB",total.FreeBytes>>20,total.AllocBytes>>20)
}
