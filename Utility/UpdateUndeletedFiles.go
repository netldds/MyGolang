package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

/*
删除文件夹后，更新未正确标记的文件和文件夹
*/

func main() {
	flag.Set("logtostderr", "true")
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
		return
	}
	tx := db.Begin()
	Update(tx)
	tx.Commit()
}
func Update(db *gorm.DB) {
	//从回收站找到所有，被删除（0），彻底删除（-2）的文件夹记录
	type RecycleRecord struct {
		Id       string `gorm:"column:id"`
		ObjectId string `gorm:"column:object_id"`
	}
	var recycleRecords []RecycleRecord
	//文件夹
	err := db.Table("recycle_bins").Where("object_type = 3 AND status <> -1 ").Scan(&recycleRecords).Error
	if err != nil {
		glog.Errorln(err)
	}
	for _, v := range recycleRecords {
		err := RecursiveFolder("/"+v.ObjectId, v.Id, db)
		if err != nil {
			glog.Errorln(err)
		}
	}
	//目录
	recycleRecords = nil
	err = db.Table("recycle_bins").Where("object_type = 1 AND status <> -1 ").Scan(&recycleRecords).Error
	for _, c := range recycleRecords {
		var foldersId []string
		err = db.Table("project_folders").Where("project_id = ? and  (recycle_id ='' or recycle_id is null ) and deleted_at is null and parent_path = '/' ", c.ObjectId).Pluck("id", &foldersId).Error
		if err != nil {
			glog.Errorln(err)
			return
		}

		for _, v := range foldersId {
			err := RecursiveFolder("/"+v, c.Id, db)
			if err != nil {
				glog.Errorln(err)
			}
		}

	}
}
func RecursiveFolder(folerPath, recycleId string, db *gorm.DB) (err error) {
	//递归找子文件夹
	var subPaths []string
	err = db.Table("project_folders").Where("parent_path = ? and  (recycle_id ='' or recycle_id is null ) ", folerPath).Pluck("path", &subPaths).Error
	if err != nil {
		glog.Errorln(err)
		return
	}
	//glog.Infoln(subPaths)
	for _, v := range subPaths {
		err = RecursiveFolder(v, recycleId, db)
		if err != nil {
			glog.Errorln(err)
			return
		}
	}
	//更新当前文件夹
	err = db.Table("project_folders").Where("path = ?", folerPath).Update(map[string]interface{}{"recycle_id": recycleId, "status": -1, "deleted_at": time.Now()}).Error
	if err != nil {
		glog.Info(err)
		return
	}
	glog.Infof("folder %v has been updated\n",folerPath)
	//找当前文件夹下的文件，更新状态
	var filesId []string
	err = db.Table("files").Where("parent_path = ? and  (recycle_id ='' or recycle_id is null ) and deleted_at is null  ", folerPath).Pluck("id", &filesId).Update(map[string]interface{}{"recycle_id": recycleId, "status": -1, "deleted_at": time.Now()}).Error
	if err != nil {
		glog.Info(err)
		return
	}
	err = db.Table("file_versions").Where("file_id in (?)", filesId).Update(map[string]interface{}{"recycle_id": recycleId, "status": -1, "deleted_at": time.Now()}).Error
	if err != nil {
		glog.Info(err)
		return
	}
	err = db.Table("raw_files").Where("file_id in (?)", filesId).Update(map[string]interface{}{"recycle_id": recycleId, "status": -1, "deleted_at": time.Now()}).Error
	if err != nil {
		glog.Info(err)
		return
	}
	glog.Infof("files id %v has been updated\n",filesId)
	return nil
}
