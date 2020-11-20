package signature_sdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
生成签名
*/
func GenerateSignature(r *http.Request, host, appId, appKey string) {
	//数据类型
	CT := "application/json"
	if r.Method == "GET" {
		CT = "application/x-www-form-urlencoded"
	}
	r.Header.Set("Content-Type", CT)

	//时间戳
	Timestamp := time.Now().UTC().Unix()
	r.Header.Set("X-Timestamp", strconv.FormatInt(Timestamp, 10))

	//app id
	r.Header.Set("Appid", appId)

	//host
	r.Host = host

	reqMethod := r.Method
	reqCT := r.Header.Get("Content-Type")
	reqHost := r.Host
	reqQuery := r.URL.RawQuery
	reqTimestamp := r.Header.Get("X-Timestamp")

	//服务端用/backend路径转发,需过滤/backend
	reqPath := r.URL.Path
	reqPath = strings.Replace(reqPath, "/backend", "", 1)

	var bs []byte
	if r.GetBody != nil {
		rc, _ := r.GetBody()
		if rc != nil {
			bs, _ = ioutil.ReadAll(rc)
		}
	}

	//组装签名字段
	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
		reqMethod,
		reqPath,
		fmt.Sprintf("content-type:%s\nhost:%s",
			reqCT,
			reqHost),
		reqQuery,
		HexSHA256(string(bs)),
	)

	//时间密钥
	secretDate := HexHMACSHA256(reqTimestamp, appKey)
	//加密
	sum := HexHMACSHA256(stringToSign, secretDate)

	//Authorization
	r.Header.Set("Authorization", "Signature="+sum)
}

/*
SHA256加密
*/
func HexSHA256(value string) string {
	b := sha256.Sum256([]byte(value))
	return hex.EncodeToString(b[:])
}

/*
带密钥的SHA256加密
*/
func HexHMACSHA256(value, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(value))
	return string(hex.EncodeToString(hashed.Sum(nil)))
}
