package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
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
	ErrorsTable.FileObj["core-errors.json"] = nil
	ErrorsTable.FileObj["container.json"] = nil
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
	g := gin.Default()
	g.LoadHTMLGlob("templates/*")
	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uploader.html", nil)
		return
	})
	g.POST("/upload_file", func(c *gin.Context) {
		m, _ := c.MultipartForm()
		f := m.File["file"][0]
		if f == nil {
			c.String(http.StatusBadRequest, "")
			return
		}
		fn, _ := f.Open()
		dir, _ := os.Getwd()
		body, _ := ioutil.ReadAll(fn)
		csvName := "result.csv"
		os.Remove(filepath.Join(dir, csvName))
		err := ioutil.WriteFile(filepath.Join(dir, csvName), body, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		Reverse()
		f1, _ := os.OpenFile(filepath.Join(dir, os.TempDir(), "小程序zh-CN.json"), os.O_RDONLY, os.ModePerm)
		f2, _ := os.OpenFile(filepath.Join(dir, os.TempDir(), "小程序zh-TW.json"), os.O_RDONLY, os.ModePerm)
		f3, _ := os.OpenFile(filepath.Join(dir, os.TempDir(), "小程序en-US.json"), os.O_RDONLY, os.ModePerm)
		f4, _ := os.OpenFile(filepath.Join(dir, os.TempDir(), "zh-CN.json"), os.O_RDONLY, os.ModePerm)
		f5, _ := os.OpenFile(filepath.Join(dir, os.TempDir(), "zh-TW.json"), os.O_RDONLY, os.ModePerm)
		f6, _ := os.OpenFile(filepath.Join(dir, os.TempDir(), "en-US.json"), os.O_RDONLY, os.ModePerm)
		fsn := []*os.File{f1, f2, f3, f4, f5, f6}
		archiveName := "archive.zip"
		os.Remove(filepath.Join(dir, os.TempDir(), archiveName))
		Compress(fsn, filepath.Join(dir, os.TempDir(), archiveName))
		filePath := path.Join(dir, os.TempDir(), archiveName)
		c.Header("Content-Disposition", "attachment; filename="+archiveName)
		c.File(filePath)
	})
	g.Run(":6060")
}
func Reverse() {
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
		if record == nil {
			fmt.Println("record 为空")
			break
		}
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
func main2() {
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
