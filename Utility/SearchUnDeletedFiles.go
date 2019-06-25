package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"path/filepath"
)

//查询dam后端删除后，扔在磁盘的文件

const containerPath = "/home/ubuntu/storage/dam-data/uptaishan/files"
const modelPath = "/home/ubuntu/storage/taishan"

func main() {
	flag.Set("-alsologtostderr", "true")
	host := flag.String("h", "127.0.0.1", "-h 127.0.0.1")
	passwd := flag.String("p", "123", "-p 123")
	dbName := flag.String("db", "taishan_dev2", "-db taishan")
	flag.Parse()
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		*passwd,
		*host,
		3306,
		*dbName)
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		glog.Warning(err)
	}
	fmt.Println(db)

	//扎到被永久删除的目录
	var folderPaths []string
	db.Table("project_folders").Where("deleted_at is not null and project_folders.id in (?)", db.Table("recycle_bins").Select("object_id").Where(" recycle_bins.status = -2 and recycle_bins.object_type = 3").QueryExpr()).Pluck("path", &folderPaths)

	//fmt.Println(folderPaths)
	LogFilePath := func(containerId, fileId string, fileType int) {
		var fpath string
		if fileType == 0 || fileType == 1 || fileType == 4 || fileType == 8 {
			fpath = filepath.Join(modelPath, containerId, fileId)
		} else {
			fpath = filepath.Join(containerPath, containerId, fileId)
		}
		if _, err := os.Stat(fpath); err == nil {
			fmt.Println(fpath)
		}
	}
	for _, v := range folderPaths {
		walk(db, v, LogFilePath)
	}

}

func walk(tx *gorm.DB, ParentPath string, wlkFun func(containerId, fileId string, fileType int)) {
	//找到当前文件夹下文件
	type Result struct {
		Id          string
		ContainerId string
		Type        int
	}
	var res []Result
	err := tx.Table("files").Select("id,container_id,type").Where("parent_path = ?", ParentPath).Scan(&res).Error
	if err != nil {
		glog.Errorln(err)
	}
	for _, v := range res {
		wlkFun(v.ContainerId, v.Id, v.Type)
	}
	//找到子文件夹
	var folderPath []string
	//.Update("deleted_at", time.Now())
	err = tx.Table("project_folders").Where("parent_path = ? ", ParentPath).Pluck("path", &folderPath).Error
	if err != nil {
		glog.Errorln(err)
	}
	for _, v := range folderPath {
		walk(tx, v, wlkFun)
	}
}
