package Misc

import (
	"fmt"
	"os"
)

const (
	Cpath   = "D:/MyGO/temp"
	tempTxt = "temp_txt.txt"
	tempDoc = "temp_doc.doc"
)

func checkPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string, debugSymbol bool) bool {
	exist := true
	if info, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	} else {
		fmt.Printf("File Name:%15s \nFile Size:%13.2fMB\n", info.Name(), float64(info.Size())/float64(1<<20))
	}
	return exist
}

func readFile(filename string) {
	fullFilename := Cpath + "/" + tempTxt
	if checkFileIsExist(fullFilename, true) {
		f, err := os.Open(fullFilename)
		checkPanic(err)
		defer f.Close()

		dat := make([]byte, 5)

		for {
			if _, err := f.Read(dat); err != nil {
				break
			} else {
				fmt.Print(string(dat))

			}
		}

	}
}
