package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Templater struct {
	TemplateFileName           string
	TemplateFileName_EN        string
	TemplateFileName_Tradition string
	body                       TemplateBodyArray
	body_tradition             TemplateBodyArray
	body_en                    TemplateBodyArray
	csvBody                    [][]string
}
type TemplateBodyArray []TemplateBody
type TemplateBody struct {
	Id            int    `json:"id"`
	Text          string `json:"text"`
	EmailContent  string `json:"email_content"`
	EmailTitle    string `json:"email_title"`
	WechatContent string `json:"wechat_content"`
	EmailType     int    `json:"email_type"`
	WebUrl        string `json:"web_url,omitempty"`
}

func (t *Templater) ToCsv(csvW *csv.Writer) error {
	err := t.OpenFiles()
	if err != nil {
		return err
	}
	//写入:分隔符+index,列名,行
	err = csvW.Write([]string{SPLITE + strconv.FormatInt(int64(TEMPLATE_INDEX), 10)})
	if err != nil {
		panic(err)
	}
	//csvW.Write([]string{t.TemplateFileName, t.TemplateFileName_Tradition, t.TemplateFileName_EN})
	csvW.Write([]string{"ID", "web_url", "text简体", "text繁体", "text英文", "email_title简体", "email_title繁体", "email_title英文"})
	for i := 0; i < len(t.body); i++ {
		ID := t.body[i].Id
		webUrl := t.body[i].WebUrl
		chn := t.body[i].Text
		chn_tradition := ""
		en := ""
		etitle := t.body[i].EmailTitle
		etitle_tradition := ""
		etitle_en := ""
		for _, v := range t.body_tradition {
			if v.Id == ID {
				chn_tradition = v.Text
				etitle_tradition = v.EmailTitle
			}
		}
		for _, v := range t.body_en {
			if v.Id == ID {
				en = v.Text
				etitle_en = v.EmailTitle
			}
		}
		csvW.Write([]string{
			strconv.FormatInt(int64(ID), 10), webUrl,
			chn, chn_tradition, en,
			etitle, etitle_tradition, etitle_en,
		})
	}
	err = csvW.Write([]string{SPLITE + strconv.FormatInt(int64(TEMPLATE_INDEX), 10)})
	csvW.Flush()
	return nil
}
func (t *Templater) ToOrigin(reader *csv.Reader) error {
	//写入:分隔符+index,列名,行
	//读到空格要返回
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
		if record[0] == "ID" {
			continue
		}
		//csvW.Write([]string{"ID", "web_url", "text简体", "text繁体", "text英文", "email_title简体", "email_title繁体", "email_title英文"})
		id, err := strconv.ParseInt(record[0], 10, 32)
		if err != nil {
			fmt.Println(err)
			continue
		}
		t.body = append(t.body, TemplateBody{
			Id:            int(id),
			Text:          record[2],
			EmailContent:  record[2],
			EmailTitle:    record[5],
			WechatContent: record[2],
			EmailType:     1,
			WebUrl:        record[1],
		},
		)
		t.body_tradition = append(t.body_tradition, TemplateBody{
			Id:            int(id),
			Text:          record[3],
			EmailContent:  record[3],
			EmailTitle:    record[6],
			WechatContent: record[3],
			EmailType:     1,
			WebUrl:        record[1],
		})
		t.body_en = append(t.body_en, TemplateBody{
			Id:            int(id),
			Text:          record[4],
			EmailContent:  record[4],
			EmailTitle:    record[7],
			WechatContent: record[4],
			EmailType:     1,
			WebUrl:        record[1],
		})
	}
	return nil
}
func (t *Templater) GetFileNames() []string {
	return []string{t.TemplateFileName, t.TemplateFileName_Tradition, t.TemplateFileName_EN}
}
func (t *Templater) OriginSave() {
	dir, _ := os.Getwd()
	dir = filepath.Join(dir, os.TempDir())
	os.Mkdir(dir, os.ModePerm)
	f, err := os.OpenFile(filepath.Join(dir, t.TemplateFileName), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(t.body)
	f.Write(b)
	f.Close()

	f, err = os.OpenFile(filepath.Join(dir, t.TemplateFileName_EN), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	b, err = json.Marshal(t.body_en)
	f.Write(b)
	f.Close()

	f, err = os.OpenFile(filepath.Join(dir, t.TemplateFileName_Tradition), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	b, err = json.Marshal(t.body_tradition)
	f.Write(b)
	f.Close()
}
func (t *Templater) OpenFiles() error {
	dir, _ := os.Getwd()
	b, err := ioutil.ReadFile(filepath.Join(dir, t.TemplateFileName))
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &t.body)
	if err != nil {
		return err
	}
	if len(t.body) == 0 {
		panic("len(t.body) == 0")
	}
	b, err = ioutil.ReadFile(filepath.Join(dir, t.TemplateFileName_EN))
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &t.body_en)
	if err != nil {
		return err
	}
	if len(t.body_en) == 0 {
		panic("len(t.body) == 0")
	}
	b, err = ioutil.ReadFile(filepath.Join(dir, t.TemplateFileName_Tradition))
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &t.body_tradition)
	if err != nil {
		return err
	}
	if len(t.body_tradition) == 0 {
		panic("len(t.body) == 0")
	}
	return nil
}
