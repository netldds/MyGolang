package Misc

import (
	"flag"
	"github.com/codeskyblue/go-sh"
	"github.com/golang/glog"
	"gopkg.in/ini.v1"
	"path/filepath"
	"strings"
)

const iFile = "/home/dx/bts.mp4"

//获取视频码率
func Start() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	transcode()
}

func getInfo() {
	var err error
	ns := sh.NewSession()
	res, err := ns.SetDir("/home/dx").Command("ffprobe", "-show_streams", iFile).Output()
	glog.Info(err)
	conf, err := ini.Load([]byte(strings.Split(string(res), "[/STREAM]")[0]))
	ss, err := conf.GetSection("STREAM")
	v, err := ss.GetKey("bit_rate")
	glog.Info(v.Int())
}

func transcode() {
	TrimName := strings.Trim(filepath.Base(iFile), filepath.Ext(iFile))
	glog.Info(TrimName)
	ns := sh.NewSession()
	var err error
	res, err := ns.SetDir("/home/dx").Command("ffmpeg", "-i", iFile, "-y", "-vcodec", "h264", "-b:v", "1000", "-acodec", "copy", TrimName+"_bak.mp4").Output()
	glog.Info(string(res), err)
}
