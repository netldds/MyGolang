package main

import (
	"fmt"
	"time"
)

// 生成2, 3, 4, ... 到 channel 'ch'中.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// 从管道复制值 'in' 到 channel 'out',
// 移除可整除的数 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // 接收值 'in'.
		if i%prime != 0 {
			out <- i // 传入 'i' 到 'out'.
		}
	}
}

//costing: 20.6380005
func main() {
	ch := make(chan int) // Create a newchannel.
	go Generate(ch)      // Launch Generate goroutine.
	start := time.Now()
	for i := 0; i < 20000; i++ {
		prime := <-ch
		//print(prime, " ")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	fmt.Println(time.Now().Sub(start).Seconds())
}

/*
102:
4.565076461
Architecture:          x86_64
CPU op-mode(s):        32-bit, 64-bit
Byte Order:            Little Endian
CPU(s):                56
On-line CPU(s) list:   0-55
Thread(s) per core:    2
Core(s) per socket:    14
Socket(s):             2
NUMA node(s):          2
Vendor ID:             GenuineIntel
CPU family:            6
Model:                 63
Model name:            Genuine Intel(R) CPU @ 2.20GHz
Stepping:              1
CPU MHz:               1232.171
CPU max MHz:           2800.0000
CPU min MHz:           1200.0000
BogoMIPS:              4402.03
Virtualization:        VT-x
L1d cache:             32K
L1i cache:             32K
L2 cache:              256K
L3 cache:              35840K
NUMA node0 CPU(s):     0-13,28-41
NUMA node1 CPU(s):     14-27,42-55
Flags:                 fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid dca sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer xsave avx f16c rdrand lahf_lm abm epb invpcid_single kaiser tpr_shadow vnmi flexpriority ept vpid fsgsbase tsc_adjust bmi1 hle avx2 smep bmi2 erms invpcid rtm cqm xsaveopt cqm_llc cqm_occup_llc dtherm ida arat pln pts


101:
3.912391375
Architecture:          x86_64
CPU 运行模式：    32-bit, 64-bit
Byte Order:            Little Endian
CPU(s):                72
On-line CPU(s) list:   0-71
每个核的线程数：2
每个座的核数：  18
Socket(s):             2
NUMA 节点：         2
厂商 ID：           GenuineIntel
CPU 系列：          6
型号：              79
Model name:            Intel(R) Xeon(R) CPU E5-2686 v4 @ 2.30GHz
步进：              1
CPU MHz：             1226.457
CPU max MHz:           3000.0000
CPU min MHz:           1200.0000
BogoMIPS:              4601.73
虚拟化：           VT-x
L1d 缓存：          32K
L1i 缓存：          32K
L2 缓存：           256K
L3 缓存：           46080K
NUMA node0 CPU(s):     0-17,36-53
NUMA node1 CPU(s):     18-35,54-71
Flags:                 fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid dca sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm 3dnowprefetch epb invpcid_single intel_pt ssbd ibrs ibpb stibp kaiser tpr_shadow vnmi flexpriority ept vpid fsgsbase tsc_adjust bmi1 hle avx2 smep bmi2 erms invpcid rtm cqm rdseed adx smap xsaveopt cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local dtherm ida arat pln pts md_clear flush_l1d

*/
