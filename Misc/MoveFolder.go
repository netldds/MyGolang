package Misc

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

//复制文件到目标目录,oldStr替换名字到NewStr
func MoveFolder(srcDir, destDir string, oldStr string, NewStr string, filters []string) {
	filtersMap := make(map[string]bool)
	for _, v := range filters {
		filtersMap[v] = true
	}
	filepath.Walk(srcDir, func(fpath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		_, ok := filtersMap[info.Name()]
		if ok {
			return nil
		}
		destFile := strings.Replace(fpath, srcDir, destDir, -1)
		destFile = strings.Replace(destFile, oldStr, NewStr, -1)
		err = os.MkdirAll(filepath.Dir(destFile), 0775)
		if err != nil {
			// glog2.Infoln(err)
			return err
		}
		err = Copy(fpath, destFile)
		return err
	})
}
func Copy(srcFile, destFile string) error {
	// glog2.Infof("->\n%v\n%v", srcFile, destFile)
	out, err := os.Create(destFile)
	defer out.Close()
	if err != nil {
		return err
	}

	in, err := os.Open(srcFile)
	defer in.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
