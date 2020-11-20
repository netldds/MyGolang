package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const GeneralOCR string = "https://ocr-api.ccint.com/ocr_service"
const AppKey = "e7273a10a954a2609bb31972e419bcba"
const AppSecret = "04c389658ab472e64b9b1c1252e56374"

const ImageFileDir = "OCRFiles"

func main() {

	dir, _ := os.Getwd()

	fs, _ := ioutil.ReadDir(filepath.Join(dir, ImageFileDir))
	for _, v := range fs {
		Start(filepath.Join(ImageFileDir, v.Name()))
	}

}

type PostBody struct {
	ImageData string `json:"image_data"`
	AppSecret string `json:"app_secret"`
}

func Start(filePath string) {
	imgData, err := ioutil.ReadFile(filePath)
	fmt.Println(filePath)
	if err != nil {
		fmt.Println(err)
	}
	var ccBody PostBody
	ccBody.AppSecret = "04c389658ab472e64b9b1c1252e56374"
	imgStr := base64.StdEncoding.EncodeToString(imgData)
	ccBody.ImageData = imgStr

	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
		//Transport: tr,
	}
	bodyData, _ := json.Marshal(ccBody)
	httpreq, err := http.NewRequest("POST", GeneralOCR, strings.NewReader(string(bodyData)))
	if err != nil {
		return
	}
	query := url.Values{}
	query.Add("app_key", AppKey)
	httpreq.URL.RawQuery = query.Encode()
	httpreq.Header.Add("Content-Type", "text/json")
	httpreq.Header.Add("Accept", "*/*")
	//httpreq.Close = true
	fmt.Println(httpreq.URL)
	resp, err := client.Do(httpreq)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		fmt.Sprintf("httperrorcode: %d \n", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	rsp, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(rsp))
}
