package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ErrorTabler struct {
	FileObj map[string]ErrorTable
}
type ErrorTable map[string]ErrorObject

func (t *ErrorTabler) ToOrigin(reader *csv.Reader) error {
	for {
		record, err := reader.Read()
		if err == io.EOF {
			panic(">未完成前,文件读完")
		}
		//结束分隔符
		if strings.Contains(record[0], SPLITE) {
			//保存
			t.OriginSave()
			return nil
		}
		//跳过列
		if record[0] == "文件名" {
			continue
		}
		//csvW.Write([]string{"文件名","对象名称", "KEY", "简体", "繁体", "英文"})
		var obj ErrorTable
		obj = make(map[string]ErrorObject)
		obj[record[1]] = ErrorObject{
			Key:       record[2],
			English:   record[5],
			Tradition: record[4],
			Msg:       record[3],
		}
		for k, v := range t.FileObj[record[0]] {
			obj[k] = v
		}
		t.FileObj[record[0]] = obj
	}
	return nil
}
func (t *ErrorTabler) OriginSave() {
	for name, v := range t.FileObj {
		dir, _ := os.Getwd()
		b, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(filepath.Join(dir, os.TempDir(), name), b, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
func (t *ErrorTabler) ToCsv(csvW *csv.Writer) error {
	t.openFiles()
	//写入:分隔符+index,列名,行
	err := csvW.Write([]string{SPLITE + strconv.FormatInt(int64(ERRORSTABLE), 10)})
	if err != nil {
		panic(err)
	}
	csvW.Write([]string{"文件名", "对象名称", "KEY", "简体", "繁体", "英文"})
	for fileName, v := range t.FileObj {
		for name, obj := range v {
			csvW.Write([]string{fileName, name, obj.Key, obj.Msg, obj.Tradition, obj.English})
		}
	}
	err = csvW.Write([]string{SPLITE + strconv.FormatInt(int64(ERRORSTABLE), 10)})
	csvW.Flush()
	return nil
}
func (t *ErrorTabler) GetFileNames() []string {
	names := []string{}
	for name, _ := range t.FileObj {
		names = append(names, name)
	}
	return names
}
func (t *ErrorTabler) openFiles() {
	for name, _ := range t.FileObj {
		dir, _ := os.Getwd()
		b, err := ioutil.ReadFile(filepath.Join(dir, name))
		if err != nil {
			panic(err)
		}
		v := make(map[string]ErrorObject)
		err = json.Unmarshal(b, &v)
		if err != nil {
			panic(err)
		}
		t.FileObj[name] = v
	}
}
