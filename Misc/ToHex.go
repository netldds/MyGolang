package Misc

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/golang/glog"
)

const hextable = "0123456789abcdef"

func ToHex() {
	strs := "20190101"
	h := sha1.New()
	h.Write([]byte(strs))
	glog.Info(h.Sum(nil))
	glog.Info(hex.EncodeToString(h.Sum(nil)))

	res := make([]byte, len([]byte(h.Sum(nil)))*2)
	for i, v := range h.Sum(nil) {
		res[i*2] = hextable[v>>4]
		res[i*2+1] = hextable[v&0x0f]
	}
	glog.Info(string(res))
}
