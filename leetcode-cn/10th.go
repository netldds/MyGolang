package main

import (
	"github.com/golang/glog"
	"strings"
)

const (
	J      = "aA"
	S      = "aAAbbbb"
	LOVELY = "LOVELY"
)

func numJewelsInStones(J string, S string) int {
	count := 0
	if len(S) > 1 {
		count += numJewelsInStones(J, S[1:])
	}
	for i, _ := range J {
		if J[i] == S[0] {
			count++
		}
	}
	return count
}
func toLowerCase(str string) string {
	var chrs []rune
	for _, v := range str {
		if v <= 90 && v >= 65 {
			chrs = append(chrs, v+32)
		}else {
			chrs=append(chrs,v)
		}
	}
	return string(chrs)
}
var emails =[]string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
func numUniqueEmails(emails []string) int {
	filter:=func(address string)string{
		localName:=strings.Split(address,"@")
		var addr []rune
		for _,v:=range localName[0]{
			if v=='.'{
				continue
			}
			if v=='+'{
				break
			}
			addr=append(addr,v)
		}
		return string(addr)+"@"+string(localName[1])
	}
	exclusiveAddr:=make(map[string]uint8)
	for i:=0;i<len(emails);i++{
		exclusiveAddr[filter(emails[i])]+=1
	}
	var count uint8
	for _,v:=range exclusiveAddr{
		count+=v
	}
	return int(count)
}

var grid =[][]int{[]int{3,0,8,4},[]int{2,4,5,7},[]int{9,2,6,3},[]int{0,3,1,0}}
func maxIncreaseKeepingSkyline(grid [][]int) int {
	width:=len(grid)
	top2bottom:=make([]int,width)
	left2right:=make([]int,width)
	for i:=0;i<width;i++{
		var x,y int
		for j:=0;j<width;j++{
			if grid[i][j]>x{
				x=grid[i][j]
			}
			if grid[j][i]>y{
				y=grid[j][i]
			}
		}
		left2right[i]=x
		top2bottom[i]=y
	}
	count:=0
	for k:=0;k<width;k++{
		for l:=0;l<width;l++{
			pivot:=0
			if left2right[k]<top2bottom[l]{
				pivot=left2right[k]
			}else{
				pivot=top2bottom[l]
			}
			count+=pivot-grid[k][l]
		}
	}
	return count
}
func main() {
	//fmt.Println(numJewelsInStones(J, S))
	//fmt.Println(toLowerCase(LOVELY))
	numUniqueEmails(emails)
	glog.Info(maxIncreaseKeepingSkyline(grid))
}
