package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ConstantVariabler struct {
	Body []byte
}

func (c *ConstantVariabler) ToCsv(csvW *csv.Writer) error {
	//写入:分隔符+index,列名,行
	err := csvW.Write([]string{SPLITE + strconv.FormatInt(int64(CONSTANTINDEX), 10)})
	if err != nil {
		panic(err)
	}
	csvW.Write([]string{"简体", "繁体", "英文"})
	r := bytes.NewReader([]byte(CONSTANT_VARIABLES))
	nr := bufio.NewScanner(r)
	for nr.Scan() {
		s := strings.Split(nr.Text(), ",")
		if len(s) == 1 {
			continue
		}
		csvW.Write(s)
	}
	err = csvW.Write([]string{SPLITE + strconv.FormatInt(int64(CONSTANTINDEX), 10)})
	csvW.Flush()
	return nil
}

//转回
func (c *ConstantVariabler) ToOrigin(reader *csv.Reader) error {
	bw := bytes.Buffer{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			panic(">未完成前,文件读完")
		}
		//结束分隔符
		if strings.Contains(record[0], SPLITE) {
			//保存
			c.Body = bw.Bytes()
			c.OriginSave()
			return nil
		}
		//跳过列
		if record[0] == "变量名" {
			continue
		}
		//csvW.Write([]string{"简体", "繁体", "英文"})
		bw.WriteString(record[0] + "," + record[1] + "," + record[2])
		bw.WriteString("\n")
	}
	return nil
}
func (c *ConstantVariabler) OriginSave() {
	dir, _ := os.Getwd()
	ioutil.WriteFile(filepath.Join(dir, os.TempDir(), "constant_variables.txt"), c.Body, os.ModePerm)
}

//file name
func (c *ConstantVariabler) GetFileNames() []string {
	return []string{"内置变量"}
}
