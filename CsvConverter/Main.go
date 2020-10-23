package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	//分割
	SPLITE           = "==SPLITE=="
	TEMPLATE_INDEX   = 0
	ERROBJS_INDEX    = 1
	ERRORSTABLE      = 2
	CONSTANTINDEX    = 3
	FRONTENDINDEX    = 4
	MINIPROGRAMINDEX = 5
)

//每个文件的转换发方法
type FileHandler interface {
	//转到Csv
	ToCsv(writer *csv.Writer) error
	//转回
	ToOrigin(reader *csv.Reader) error
	//file name
	GetFileNames() []string
}

//map
var FileConverts map[int]FileHandler

func init() {
	FileConverts = make(map[int]FileHandler)

	templateJson := Templater{
		TemplateFileName:           "template.json",
		TemplateFileName_EN:        "template_en.json",
		TemplateFileName_Tradition: "template_tradition.json",
	}
	FileConverts[TEMPLATE_INDEX] = &templateJson

	ErrObjsConvert := ErrObjConverter{
		FileName: "err-objs.go",
	}
	ErrObjsConvert.ObjFormat = `	
%v = ErrorObject{
		Key:       "%v",
		English:   "%v.",
		Tradition: "%v",
		Msg:       "%v",
	}
`
	ErrObjsConvert.Objs = make(map[string]ErrorObject)
	FileConverts[ERROBJS_INDEX] = &ErrObjsConvert

	ErrorsTable := ErrorTabler{}
	ErrorsTable.FileObj = make(map[string]ErrorTable)
	ErrorsTable.FileObj["admin-errors.json"] = nil
	ErrorsTable.FileObj["alert-errors.json"] = nil
	ErrorsTable.FileObj["alert-errors.json"] = nil
	ErrorsTable.FileObj["core-errors.json"] = nil
	ErrorsTable.FileObj["entity_errors.json"] = nil
	ErrorsTable.FileObj["entity-state-errors.json"] = nil
	ErrorsTable.FileObj["file-errors.json"] = nil
	ErrorsTable.FileObj["folder-errors.json"] = nil
	ErrorsTable.FileObj["freight-errors.json"] = nil
	ErrorsTable.FileObj["notification-errors.json"] = nil
	ErrorsTable.FileObj["oauth-errors.json"] = nil
	ErrorsTable.FileObj["plan-entity-errors.json"] = nil
	ErrorsTable.FileObj["plus-plan-errors.json"] = nil
	ErrorsTable.FileObj["project-errors.json"] = nil
	ErrorsTable.FileObj["project-file-errors.json"] = nil
	ErrorsTable.FileObj["task-errors.json"] = nil
	ErrorsTable.FileObj["user-errors.json"] = nil
	ErrorsTable.FileObj["viewpoint-errors.json"] = nil
	ErrorsTable.FileObj["webapp-erros.json"] = nil
	ErrorsTable.FileObj["yanhuagis-errors.json"] = nil
	FileConverts[ERRORSTABLE] = &ErrorsTable

	c := ConstantVariabler{}
	FileConverts[CONSTANTINDEX] = &c

	frontEnder := FrontEnder{
		FileName:          "zh-CN.json",
		FileNameTradition: "zh-TW.json",
		FileNameEnglish:   "en-US.json",
	}
	FileConverts[FRONTENDINDEX] = &frontEnder

	miniProgram := MiniProgram{
		FileName:          "小程序zh-CN.json",
		FileNameTradition: "小程序zh-TW.json",
		FileNameEnglish:   "小程序en-US.json",
	}
	FileConverts[MINIPROGRAMINDEX] = &miniProgram

}
func main() {
	//确认模式
	fmt.Println(">csv文件名:result.csv")
	fmt.Println(">转换参数默认空,恢复参数加入-r")
	if len(os.Args) == 1 {
		//新建
		csvW := NewCsvFileWriter()
		//转换
		fmt.Printf(">CSV模式\n")
		fmt.Printf(">共%v个文件\n", len(FileConverts))
		for _, v := range FileConverts {
			fmt.Printf(">处理:%v\n", v.GetFileNames())
			err := v.ToCsv(csvW)
			if err != nil {
				panic(err)
			}
		}
	}
	if len(os.Args) == 2 && os.Args[1] == "-r" {
		//恢复
		csvR := NewCsvFileReader()
		fmt.Printf(">恢复模式\n")
		fmt.Printf(">共%v个文件\n", len(FileConverts))
		//选择器
		for {
			record, err := csvR.Read()
			if err == io.EOF {
				fmt.Println(">文件读完")
				break
			}
			//if len(record) != 1 {
			//	fmt.Println("没找到分隔符")
			//	break
			//}
			recordOne := record[0]
			if !strings.Contains(recordOne, SPLITE) {
				fmt.Println("没找到分隔符")
				break
			}
			indexStr := recordOne[len(recordOne)-1]
			index, err := strconv.ParseInt(string(indexStr), 10, 32)
			if err != nil {
				panic(err)
			}
			handler := FileConverts[int(index)]
			if handler == nil {
				panic("nil FileConvert Handler:" + string(indexStr))
			}
			handler.ToOrigin(csvR)
		}
	}
}
func NewCsvFileWriter() *csv.Writer {
	Name := "result.csv"
	dir, _ := os.Getwd()
	os.Mkdir(path.Join(dir, os.TempDir()), os.ModePerm)
	file, err := os.OpenFile(filepath.Join(dir, os.TempDir(), Name), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	file.WriteString("\xEF\xBB\xBF")
	csvW := csv.NewWriter(file)
	return csvW
}
func NewCsvFileReader() *csv.Reader {
	Name := "result.csv"
	dir, _ := os.Getwd()
	file, err := os.OpenFile(filepath.Join(dir, Name), os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	csvR := csv.NewReader(file)
	return csvR
}
