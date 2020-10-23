package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ErrObjConverter struct {
	FileName  string
	Objs      map[string]ErrorObject
	ObjFormat string
}
type ErrorObject struct {
	Key       string `json:"key"`
	English   string `json:"en"`
	Tradition string `json:"tradition"`
	Msg       string `json:"msg"`
}

func (t *ErrObjConverter) ToCsv(csvW *csv.Writer) error {
	t.OpenFile()
	//写入:分隔符+index,列名,行
	err := csvW.Write([]string{SPLITE + strconv.FormatInt(int64(ERROBJS_INDEX), 10)})
	if err != nil {
		panic(err)
	}
	//csvW.Write([]string{t.TemplateFileName, t.TemplateFileName_Tradition, t.TemplateFileName_EN})
	csvW.Write([]string{"变量名", "KEY", "简体", "繁体", "英文"})
	for key, v := range t.Objs {
		csvW.Write([]string{key, v.Key, v.Msg, v.Tradition, v.English})
	}
	err = csvW.Write([]string{SPLITE + strconv.FormatInt(int64(ERROBJS_INDEX), 10)})
	csvW.Flush()
	return nil
}
func (t *ErrObjConverter) OriginSave() {
	dir, _ := os.Getwd()
	dir = filepath.Join(dir, os.TempDir())
	os.Mkdir(dir, os.ModePerm)
	f, err := os.OpenFile(filepath.Join(dir, t.FileName), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	bf := bufio.NewWriter(f)
	bf.WriteString(`package main
var (`)
	for key, v := range t.Objs {
		/*
			%v = ErrorObject{
				Key:       "%v",
				English:   "%v.",
				Tradition: "%v",
				Msg:       "%v",
			}*/
		bf.WriteString(fmt.Sprintf(t.ObjFormat, key, v.Key, v.English, v.Tradition, v.Msg))
	}
	bf.WriteString(`)`)
	bf.Flush()
	f.Close()
}
func (t *ErrObjConverter) ToOrigin(reader *csv.Reader) error {
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
		if record[0] == "变量名" {
			continue
		}
		//csvW.Write([]string{"变量名","KEY", "简体", "繁体", "英文"})
		t.Objs[record[0]] = ErrorObject{
			Key:       record[1],
			English:   record[4],
			Tradition: record[3],
			Msg:       record[2],
		}
	}
	return nil
}
func (t *ErrObjConverter) GetFileNames() []string {
	return []string{t.FileName}
}
func (t *ErrObjConverter) OpenFile() {
	objs := make(map[string]ErrorObject)
	objs["ErrIllegalParameters"] = ErrIllegalParameters
	objs["ErrIDConvert"] = ErrIDConvert
	objs["ErrPackageNotFound"] = ErrPackageNotFound
	objs["ERRMarshalJSON"] = ERRMarshalJSON
	objs["ErrExpireIncorrect"] = ErrExpireIncorrect
	objs["ErrCloudServiceNotFound"] = ErrCloudServiceNotFound
	objs["ErrSignaturePatterns"] = ErrSignaturePatterns
	objs["ErrServiceUnavailable"] = ErrServiceUnavailable
	objs["ErrServiceUnAvailableMember"] = ErrServiceUnAvailableMember
	objs["ErrExceedMaxStorage"] = ErrExceedMaxStorage
	objs["ErrExceedMaxStorageeMember"] = ErrExceedMaxStorageeMember
	objs["ErrExceedOperatorAmount"] = ErrExceedOperatorAmount
	objs["ErrExceedOperatorAmountMember"] = ErrExceedOperatorAmountMember
	objs["ErrExceedMaxUploadSize"] = ErrExceedMaxUploadSize
	objs["ErrExceedMaxUploadSizeMember"] = ErrExceedMaxUploadSizeMember
	objs["ErrExceedMaxAssembleSize"] = ErrExceedMaxAssembleSize
	objs["ErrExceedMaxAssembleSizeMember"] = ErrExceedMaxAssembleSizeMember
	objs["ErrExceed6GB"] = ErrExceed6GB
	objs["ErrExceed6GBMember"] = ErrExceed6GBMember
	objs["ErrExceedFreeProjectAmount"] = ErrExceedFreeProjectAmount
	objs["ErrFreeProjectUploadModellimitation"] = ErrFreeProjectUploadModellimitation
	objs["ErrFreeProjectUploadModellimitationMember"] = ErrFreeProjectUploadModellimitationMember
	objs["ErrSuperVisorOnly"] = ErrSuperVisorOnly
	objs["ErrContainerMaximumLimitation"] = ErrContainerMaximumLimitation
	objs["ErrContainerUploadMaximumLimitation"] = ErrContainerUploadMaximumLimitation
	objs["ErrFileExtensionLimited"] = ErrFileExtensionLimited
	objs["ErrReturnToRoot"] = ErrReturnToRoot
	objs["ErrFileNotExisted"] = ErrFileNotExisted
	objs["ErrCloudServiceStatus"] = ErrCloudServiceStatus
	objs["ErrOverCoverFileLimit"] = ErrOverCoverFileLimit
	objs["ErrMeetingTimeFormat"] = ErrMeetingTimeFormat
	objs["ErrMeetingCreate"] = ErrMeetingCreate
	objs["ErrMeetingNotExist"] = ErrMeetingNotExist
	objs["ErrMeetingUpdateFail"] = ErrMeetingUpdateFail
	objs["ErrMeetingParseFail"] = ErrMeetingParseFail
	objs["ErrMeetingDeleteFail"] = ErrMeetingDeleteFail
	objs["ErrMeetingPermission"] = ErrMeetingPermission
	objs["ErrMeetingNO"] = ErrMeetingNO
	objs["ErrMeetingSharePermission"] = ErrMeetingSharePermission
	objs["ErrMeetingShareFail"] = ErrMeetingShareFail
	objs["ErrMeetingExists"] = ErrMeetingExists
	objs["ErrWeakPassword"] = ErrWeakPassword
	objs["ErrCaptcha"] = ErrCaptcha
	t.Objs = objs
}
