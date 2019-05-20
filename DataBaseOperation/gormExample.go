package DataBaseOperation

import (
	"dx/taishan/core/db"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type RawFile struct {
	db.Model
	ContainerId string `json:"container_id" db:"container_id"`
	FileId      string `json:"file_id" db:"file_id"`
	UploaderId  string `json:"uploader_id" db:"uploader_id"`
	VersionId   string `json:"version_id" db:"version_id"`
	Type        int    `json:"type" db:"type"`
	Size        int64  `json:"size" db:"size"`
	Name        string `json:"name" db:"name"`
	Thumbnail   string `json:"thumbnail" db:"thumbnail"`
	RecycleId   string `json:"recycle_id" db:"recycle_id"`
	Status      int    `json:"status" db:"status"`
}

func RungOrm() {

	//var CMdata []ContainerMember
	//var RAWfiles []RawFile
	//userId := "2fd8584564ad47798ac4d23cf4a03ea1"
	db, err := gorm.Open("mysql", "root:123@/taishan_dev2?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	Webapp(db)
	//fmt.Println(time.Now())
}
func Webapp(db *gorm.DB) {
	//return
	//timestampv := time.Now().Local().Format("2006-01-02 15:04:05")
	//timestampv = time.Now().Local().String()
	//bigTimeStampv,_:=time.Parse("2006-01-02 15:04:05","2020-12-13 13:14:15")
	//fmt.Println(bigTimeStampv)
	//if db.Table("user_web_apps").Where("due_time >= ? and user_id = ? and app_id = ? ", time.Now().Local().String(), "e28ff0cdab2a482b939ac3dc154df1ad", db.Table("web_apps").Select("id").Where("app_key = ?","Watermark").QueryExpr()).First(&userWebApp).RecordNotFound() {
	//	fmt.Println("0")
	//}

	//db.Table("user_web_apps").Where("due_time >= ? and user_id = ? and app_id = ? ", time.Now().Local().String(), "e28ff0cdab2a482b939ac3dc154df1ad",db.Table("web_apps").Select("id").Where("app_key = ?","Watermark").SubQuery()).First(&userWebApp)
	//db.Table("projects").Select("creator_id,owner_id").Where("creator_id = ?", "0d732263d5ad43ca9430ccaf9031ca94").Update(map[string]interface{}{"creator_id": "abc", "owner_id": "abc"})

	//fmt.Println(webApp)
	//fmt.Println(userWebApp)
	//err := db.Where("due_time >= ? and user_id = ? and app_id in (?) ", time.Now().Local().String(), "e28ff0cdab2a482b939ac3dc154df1ad", db.Table("web_apps").Select("id").Where("app_key =?", "Watermark").QueryExpr()).First(&userWebApp).Error

	//fmt.Println(err)

	//row := db.Table("raw_files").Select("id,file_state,thumbnail_state").Where("id = ?", "90cd1b66dffb43d99e636d56dc057aef").Row()
	//
	//var id string
	//var file_state int8
	//var thumbnail_state int8
	//
	//row.Scan(&id, &file_state, &thumbnail_state)
	//fmt.Printf("%v\n%v\n%v\n", id, file_state, thumbnail_state)

	//projectMember.Id="id1"
	//projectMember.Status = 1
	//projectMember.ProjectId = "projectID"
	//projectMember.OwnerId = "OwnerID"
	//projectMember.MemberType = 1
	//fmt.Println(db.NewRecord(&projectMember))
	//dbSearch1 = dbSearch1.Joins("UNION ?", dbSearch3.SubQuery())
	//dbSearch1=db.Table("containers").Select("*")

	//rows,err:=dbSearch1.Rows()
	//fmt.Println(err)
	//for rows.Next(){
	//	var lvl ,name string
	//	err=rows.Scan(&lvl)
	//	if err !=nil{
	//		fmt.Println(err)
	//	}
	//	fmt.Println(lvl," ",name)
	//}
	//dbSearch1.Scan(&containers)

	//db.Raw("? UNION ? UNION ?",dbSearch1.SubQuery(),dbSearch2.SubQuery(),dbSearch3.SubQuery()).Find(&containers)
	//db.Raw("select *,'1' lvl from containers where containers.name LIKE '%女一号%' ",).Scan(&containers)

	//db.Raw("(select *,'1' lvl  from containers where containers.name = ?) UNION  (select *,'2' lvl  from containers where containers.name LIKE ?) UNION  (select *,'3' lvl  from containers where containers.name LIKE ?) ORDER BY lvl", "女一号", "女一号_%","%_女一号_%").Scan(&containers)

	//var num int
	//db.Table("oauths").Where("user_id = ? and delete_at not null ", "1fcf019ec96c4ba1ac88bfc408febd97").Count(&num)

	//body := []struct {
	//	Date time.Time
	//	Count int
	//}{}
	//err:= db.Table("files").Limit(-1).Select("DATE(created_at)as date ,count(*) count").Group("DATE(created_at)").Where("created_at > ? AND created_at < ? AND parent_path like ?","2019-01-10","2019-05-16","%"+""+"%").Scan(&body).Error
	//fmt.Println(err)
	//fmt.Println(body)

	var fileSizes int64
	var count int
	db.Table("raw_files").Select("sum(size),count(*)").Where("container_id = ?", "8637c3f6270c48a2a7b8289d6ca4fe7d").Order("size DESC").Row().Scan(&fileSizes,&count)
	fmt.Println(fileSizes)

}
func FuncUser(db *gorm.DB) {
	//db.Table("container_members").Where("container_id =? and member_type =?", "4fc637a53d2242fdbfed3e3195906175", 2).Not("deleted_at", nil).Find(&containerMembers)
	//fmt.Println(CMS)
	//db.Where("owner_id = ? AND member_type=? ", userId, 1).Find(&CMdata)
	//fmt.Println(CMdata)

	//v, _ := db.Table("container_members").Select("id").Where("owner_id = ? AND member_type=? ", userId, 1).Rows()

	//var ids []struc
	//var ids []string
	//v, _ := db.Table("container_members").Select("id").Where("owner_id = ? AND member_type=? ", userId, 1).Rows()
	//v, err := db.Table("container_members").Select("id").Where("owner_id = ? AND member_type=? ", userId, 1).Rows()
	//if v.Next() {
	//	var id string
	//	v.Scan(&id)
	//	ids = append(ids, id)
	//	fmt.Println(id)
	//}

	//ids=[]string{"c4898f0f42d844579e395762d0099c0b","f24445cba3d64bf4bb70f16a20628835"}

	//v,_=db.Table("raw_files").Select("id").Where("uploader_id=?", ids).Rows()
	//if v.Next(){
	//
	//}
	//fmt.Println(ids)
	//db.Table("raw_files").Where("uploader_id = ?", "1dc6e02381114f26aba01fc89868200a").Find(&RAWfiles)

	//db.Table("raw_files").Where("uploader_id=?")
	//var files []FileStruct
	//db.Table("files").Joins("JOIN raw_files on raw_files.id = files.raw_file_id").Where("raw_files.uploader_id in (?)",ids).Find(&files)
	//db.Table("files").Where("raw_file_id =? ",).Find(&fss)

	//fmt.Println(files)

	//db.Table("raw_files").Where("uploader_id = ?", ids).Find(&RAWfiles)
	//fmt.Println(RAWfiles)

	//var profiles []UserProfile

	//v, err := db.Table("users").Select("id").Where("id in (?)", db.Table("container_members").Select("owner_id").Where("container_id in (?) and member_type = ?", db.Table("containers").Select("id").Where("creator_id = ?", userId).QueryExpr(), 2).QueryExpr()).Rows()
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//var ids []string
	//if v.Next() {
	//	var id string
	//	v.Scan(&id)
	//	ids = append(ids, id)
	//}
	//fmt.Println(ids)

	var result []UserForProfiles
	//var result2 []UserProfile
	//var result3 []User
	//err = db.Table("users").Select("*").Joins("JOIN user_profiles on user_profiles.user_id=users.id").Where("users.id in (?)", db.Table("container_members").Select("owner_id").Where("container_id in (?) and member_type = ?", db.Table("containers").Select("id").Where("creator_id = ?", userId).QueryExpr(), 2).QueryExpr()).Find(&result).Error

	//err = db.Table("users").Select("*").Joins("left join user_profiles on user_profiles.user_id=users.id").Where("users.id in (?)", "dbd09cf3d711455d8c69421225d5eb2b").Find(&result).Error
	//err = db.Table("user_profiles").Where("user_id in (?)", "dbd09cf3d711455d8c69421225d5eb2b").Find(&result2).Error
	fmt.Println(result)

	//p := fmt.Println
	//p(result[0].Name)
	//p(result[0].Id)
	//p(result[0].Avatar)
	//p(result[0].Email)
	//p(result[0].CompanyId)
	//p(result[0].Status)
}
