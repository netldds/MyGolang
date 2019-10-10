package Misc

import (
	"fmt"
	"runtime"
)

//使用对象数量,内分配数量
func printUsage() {
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)
	//fmt.Printf("Sys MB %v	", stats.Sys>>20)
	//fmt.Printf("PauseTotalNs %v	", stats.PauseTotalNs)
	//fmt.Printf("Alloc MB %v	", stats.Alloc>>20)
	//fmt.Printf("Lookups %v	", stats.Lookups)
	//fmt.Printf("HeapObjects %v	", stats.HeapObjects)
	//fmt.Printf("live object %v\n", stats.Mallocs-stats.Frees)
	//fmt.Printf("live object %v\n", stats.HeapInuse)

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

	fmt.Printf("%d in use objects (%d in use bytes) | Alloc: %d TotalAlloc: %d\n",
		total.InUseObjects(), total.InUseBytes(), stats.Alloc, stats.TotalAlloc)
}
