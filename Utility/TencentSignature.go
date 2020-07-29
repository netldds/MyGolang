package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var appID = "appid"
var appKEY = "appkey"

func main() {
	g := gin.Default()
	g.Use(gin.HandlerFunc(SignatureMiddleWare))
	g.GET("/api/v3/project", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	g.POST("/api/v3/project", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	g.Run(":2080")
}

func SignatureMiddleWare(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	//appID := c.Request.Header.Get("Appid")
	signature := authorization[strings.Index(authorization, "Signature")+len("Signature")+1:]
	//ID检查,获取key
	method := c.Request.Method
	urlPath := c.Request.URL.Path
	cthdr := c.Request.Header.Get("Content-Type")
	host := c.Request.Host
	//host检查
	urlQuery := c.Request.URL.RawQuery
	xTimestamp := c.Request.Header.Get("X-Timestamp")
	//时间检查
	i, _ := strconv.ParseInt(xTimestamp, 10, 64)
	tm := time.Unix(i, 0)
	fmt.Printf("time at:%v\n", tm.Local())
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	bodyStr := ""
	if len(bodyBytes) != 0 && err == nil {
		bodyStr = string(bodyBytes)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
		method,
		urlPath,
		fmt.Sprintf("content-type:%s\nhost:%s",
			cthdr,
			host),
		urlQuery,
		HexSHA256(bodyStr),
	)
	fmt.Println(stringToSign)
	secretDate := HexHMACSHA256(xTimestamp, appKEY)
	fmt.Printf("secretDate:%v\n", secretDate)
	sum := HexHMACSHA256(stringToSign, secretDate)
	fmt.Printf("signature client side:%v\n", signature)
	fmt.Printf("signature server side:%v\n", sum)
	if signature == sum {
		return
	} else {
		c.JSON(http.StatusOK, "denied")
		c.Abort()
		return
	}
}
func HexSHA256(value string) string {
	b := sha256.Sum256([]byte(value))
	return hex.EncodeToString(b[:])
}
func HexHMACSHA256(value, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(value))
	return string(hex.EncodeToString(hashed.Sum(nil)))
}
