package main

import (
	"fmt"
	youtu "github.com/Tencent-YouTu/Go_sdk"
	"io/ioutil"
	"os"
)

const imgFIle = "3b.jpg"

func main() {

	//Register your app on http://open.youtu.qq.com
	//Get the following details
	appID := uint32(1256126591)
	secretID := "AKIDWXUS0TX282QehaFxOglleOzWr4kBDmy2"
	secretKey := "1ymQt9E792kGLBytGOODtAFGrMgy7OWS"
	userID := "100004155798"

	as, err := youtu.NewAppSign(appID, secretID, secretKey, userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewAppSign() failed: %s\n", err)
		return
	}
	imgData, err := ioutil.ReadFile(imgFIle)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ReadFile() failed: %s\n", err)
		return
	}
	yt := youtu.Init(as, youtu.DefaultHost)
	//yt := youtu.Init(as, "http://api.youtu.qq.com")
	//yt.SetDebug(true)
	df, err := yt.GeneralOcr(imgData, 0, "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "DetectFace() failed: %s", err)
		return
	}
	fmt.Printf("df: %#v\n", df.ErrorMsg)
}
