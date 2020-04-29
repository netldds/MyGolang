package Misc

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

/*
生成制定大小文件
author:liu
*/
const (
	B  = 1
	KB = 1 << 10
	MB = 1 << 20
	GB = 1 << 30
	//TB = 1 << 40
)

type Size struct {
	Size   int64
	Metric string
}

var ch = make(chan int)
var wg = sync.WaitGroup{}

func Generate() {
	fName := flag.String("n", "chunk", "-n filename")
	sizeStr := flag.String("s", "100M", "-s 3M 1G 2K")
	flag.Parse()
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	sizeArr := parseSize(*sizeStr)
	for _, v := range sizeArr {
		wg.Add(1)
		nameMetric := *fName + v.Metric
		fullName := path.Join(wd, nameMetric)
		go Fill(fullName, v.Size)
	}
	wg.Wait()
}
func Fill(fName string, size int64) {
	defer wg.Done()
	if fName == "" || size == 0 {
		return
	}
	if CheckFreeSpace(path.Dir(fName)) <= size {
		return
	}
	f, err := os.Create(fName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_, err = f.Seek(size-1, 0)
	if err != nil {
		return
	}
	_, err = f.Write([]byte{0})
	if err != nil {
		return
	}

}
func parseSize(sizeStr string) []Size {
	arr := strings.Split(sizeStr, ",")
	var sizeArr []Size
	for _, v := range arr {
		if v == "" {
			continue
		}
		num := v[:len(v)-1]
		n, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			continue
		}
		suffixChar := string(v[len(v)-1])
		switch suffixChar {
		case "B":
			m := strconv.FormatInt(n, 10) + "Byte"
			sizeArr = append(sizeArr, Size{
				Size:   n * B,
				Metric: m,
			})
		case "K":
			m := strconv.FormatInt(n, 10) + "KiB"
			sizeArr = append(sizeArr, Size{
				Size:   n * KB,
				Metric: m,
			})
		case "M":
			m := strconv.FormatInt(n, 10) + "MiB"
			sizeArr = append(sizeArr, Size{
				Size:   n * MB,
				Metric: m,
			})
		case "G":
			m := strconv.FormatInt(n, 10) + "GiB"
			sizeArr = append(sizeArr, Size{
				Size:   n * GB,
				Metric: m,
			})
		}
	}
	return sizeArr
}
func CheckFreeSpace(wd string) int64 {
	var stat syscall.Statfs_t
	syscall.Statfs(wd, &stat)
	size := stat.Bavail * uint64((stat.Bsize))
	fmt.Println(strconv.FormatInt(int64(size/1024/1024), 10) + "MB")
	return int64(size)
}
