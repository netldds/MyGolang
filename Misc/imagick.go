package Misc

import (
	"encoding/json"
	"fmt"
	"gopkg.in/gographics/imagick.v2/imagick"
	"os"
	"path/filepath"
)

func do() {
	var UnJson map[string]interface{}
	json.Unmarshal([]byte(jsonData), &UnJson)
	var JsonStrValue string
	if result, ok := UnJson["result"].(map[string]interface{}); ok {
		if OcrDataList, ok := result["ocr_data_list"].([]interface{}); ok {
			for _, v := range OcrDataList {
				if value, ok := v.(map[string]interface{}); ok {
					JsonStrValue = fmt.Sprintf("%s\n%s", JsonStrValue, value["value"])
				}
			}
		}
	}
	fmt.Println(JsonStrValue)
}

func Start() {

	imagick.Initialize()
	defer imagick.Terminate()

	mv := imagick.NewMagickWand()
	defer mv.Destroy()

	err := mv.SetResolution(200, 200)

	err = mv.ReadImage(pdfPATH)
	//fmt.Println(	mv.GetImageWidth())
	err = mv.SetCompression(100)
	index := mv.GetIteratorIndex()
	err = mv.SetFormat("jpg")
	pt, _ := os.Getwd()
	pt = filepath.Join(pt, "a")
	err = os.MkdirAll(pt, os.ModePerm)
	for index > 0 {
		mv.SetIteratorIndex(int(index - 1))
		targetPath := fmt.Sprintf("%s/%s-%d.jpg", pt, "a", index-1)
		err = mv.WriteImage(targetPath)
		index--
	}
	err = mv.WriteImage(pdfPATH + ".jpg")
	fmt.Println(err)
}

var jsonData string = `{"cost_time":354,"result":{"rotated_image_width":662,"rotated_image_height":188,"image_angle":0,"ocr_data_list":[{"value":"一-------------------ー一","position":"79,44,617,42,617,54,79,54","key":"text","description":"文本"},{"value":"重1,有事主体的经骨范岡由幸程确定.垒营范川中国于法津,迭凰脱定腕古经批准的项目,取爵赤可事盘文件后方","position":"41,55,622,55,622,71,41,71","key":"text","description":"文本"},{"value":"河开展相关经营话动","position":"73,72,185,73,185,85,73,83","key":"text","description":"文本"},{"value":"2,商事东体经背范川和许可审盘项目等有关事通及年报信息和其他信用信总,请登杀深期市市场和质験意督管理i","position":"73,90,622,88,622,101,73,102","key":"text","description":"文本"},{"value":"委小会商事主体信用信配公不平台(阿址htp:\/\/www.szcredit.com.cn)成村横执黑的堆科专询","position":"73,103,542,103,542,115,73,115","key":"text","description":"文本"},{"value":"3,商事主体期于好年1月1日-6月30日向商事俊记机关提安上一年度的年度假告,商事未体成当核照《食业信总会;","position":"73,119,622,118,622,131,73,131","key":"text","description":"文本"},{"value":"示永新行条例)等總建同社会金本商事業体師ーーーーーーーーーーー","position":"47,131,622,143,622,158,47,144","key":"text","description":"文本"},{"value":"---","position":"63,146,302,146,302,156,63,155","key":"text","description":"文本"}]},"code":200,"message":"success"}`
