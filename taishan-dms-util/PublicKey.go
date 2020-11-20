package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var priKey = `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJwj8jmbx7ZIkUNg
yGSGaLWcHZjYqvGIRHZv3BO0KUCmBWczrg9MOgrP4a/yF26j5ZoIMk+KtLEORplh
/EhcTZW/oIgGF6LDS3qrMxOxaFW98e5nhZhkCfPH2wQosebZl8sH2DjBUL0i2cny
LWLsqGFh1z+o6L6XbjfsFwKhQp43AgMBAAECgYA6oWDWcwGGGC+7zj7RSItPDqUq
fMmL0rBqjMxl9bO729uRih1lDymIX9EOQWi9GfwgX82Mgrgg+AxYkiuqfEaBS4F1
1rMhJKQeDw/Ir3zrbT5rAdu494LPpfOidxKMZwCYt7YRESla5kI3k3nL+46A2FK/
7AMVyFXbTxAChXoH2QJBAMltT7pXQOvdbcTaL1YS4+bkWxGOU1+co44y98jyUiwv
wnq+N8F6NsuzU77rAFF4BNaEj2hwueR1mqv5vzftsdMCQQDGcZ9CmPJ7p5HteYcl
URDn8C9vkdsC6Rfq+lb5jVdZBQtTU8W7mjBAcJ/QcLF3I3U99lyrK1JHLqD3DYdD
YH+NAkB8LMlj5Pp+7ckH/EIGXCrnYovJ7OX1IYmq1jzvQoPp/Z91L+MLgZ5aQbk1
D4bosoa5AIuwJR5UezPZJWP+xKhFAkEAkEU+wc4sTBXxk7KMvGaJYfZOplBl52HL
T7wcy2UkocV3DGeVE+TvO4olxgaIHtOagye/C3p9YN7Xi4U8V5GqaQI/Y7rZGc++
UOQtpo+eWPKxPN3ImHZU3Y3tzn1gzwc3cOeDdjeBBMc922eFTFDnvFMiUlS2OJpW
dqCZu7/8U4Xt
-----END PRIVATE KEY-----`
var pubKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCcI/I5m8e2SJFDYMhkhmi1nB2Y
2KrxiER2b9wTtClApgVnM64PTDoKz+Gv8hduo+WaCDJPirSxDkaZYfxIXE2Vv6CI
Bheiw0t6qzMTsWhVvfHuZ4WYZAnzx9sEKLHm2ZfLB9g4wVC9ItnJ8i1i7KhhYdc/
qOi+l2437BcCoUKeNwIDAQAB
-----END PUBLIC KEY-----`

func main() {

}

//RSA处理
// 加密
func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS8格式的私钥
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pk := priv.(*rsa.PrivateKey)
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, pk, ciphertext)
}
