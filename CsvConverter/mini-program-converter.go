package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

type MiniProgram struct {
	body              MiniProgramBody
	bodyTradition     MiniProgramBody
	bodyEnglish       MiniProgramBody
	FileName          string
	FileNameTradition string
	FileNameEnglish   string
	seek              int
}

func (s *MiniProgram) ToCsv(csvW *csv.Writer) error {
	s.openFiles()
	//写入:分隔符+index,列名,行
	err := csvW.Write([]string{SPLITE + strconv.FormatInt(int64(MINIPROGRAMINDEX), 10)})
	if err != nil {
		panic(err)
	}
	csvW.Write([]string{"简体", "繁体", "英文"})
	bodyValue := reflect.ValueOf(s.body)
	bodyTraditionValue := reflect.ValueOf(s.bodyTradition)
	bodyEnglishValue := reflect.ValueOf(s.bodyEnglish)
	for i := 0; i < bodyValue.NumField(); i++ {
		err := csvW.Write([]string{bodyValue.Field(i).String(), bodyTraditionValue.Field(i).String(), bodyEnglishValue.Field(i).String()})
		if err != nil {
			panic(err)
		}
	}
	err = csvW.Write([]string{SPLITE + strconv.FormatInt(int64(MINIPROGRAMINDEX), 10)})
	csvW.Flush()
	return nil
}

//转回
func (s *MiniProgram) ToOrigin(reader *csv.Reader) error {
	v := reflect.ValueOf(&s.body)
	vt := reflect.ValueOf(&s.bodyTradition)
	ve := reflect.ValueOf(&s.bodyEnglish)
	s.seek = 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			panic(">未完成前,文件读完")
		}
		//结束分隔符
		if strings.Contains(record[0], SPLITE) {
			//保存
			s.OriginSave()
			return nil
		}
		//跳过列
		if record[0] == "简体" {
			continue
		}
		//	csvW.Write([]string{"简体", "繁体", "英文"})
		v.Elem().Field(s.seek).SetString(record[0])
		vt.Elem().Field(s.seek).SetString(record[1])
		ve.Elem().Field(s.seek).SetString(record[2])
		s.seek++
	}
	return nil
}
func (s *MiniProgram) OriginSave() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(s.body)
	if err != nil {
		panic(err)
	}
	bt, err := json.Marshal(s.bodyTradition)
	if err != nil {
		panic(err)
	}
	be, err := json.Marshal(s.bodyEnglish)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filepath.Join(dir, os.TempDir(), s.FileName), b, os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filepath.Join(dir, os.TempDir(), s.FileNameTradition), bt, os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filepath.Join(dir, os.TempDir(), s.FileNameEnglish), be, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

//file name
func (s *MiniProgram) GetFileNames() []string {
	return []string{s.FileName, s.FileNameTradition, s.FileNameEnglish}
}
func (t *MiniProgram) openFiles() {
	dir, _ := os.Getwd()
	b, err := ioutil.ReadFile(filepath.Join(dir, t.FileName))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &t.body)
	if err != nil {
		panic(err)
	}
	b, err = ioutil.ReadFile(filepath.Join(dir, t.FileNameTradition))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &t.bodyTradition)
	if err != nil {
		panic(err)
	}
	b, err = ioutil.ReadFile(filepath.Join(dir, t.FileNameEnglish))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &t.bodyEnglish)
	if err != nil {
		panic(err)
	}

}

type MiniProgramBody struct {
	BIM                   string `json:"BIM"`
	POWERED               string `json:"POWERED"`
	GETINTO               string `json:"GETINTO"`
	CLICKLOGIN            string `json:"CLICKLOGIN"`
	SPACE                 string `json:"SPACE"`
	SPACEUSED             string `json:"SPACEUSED"`
	UNUSEDSPACE           string `json:"UNUSEDSPACE"`
	TOTALSPACE            string `json:"TOTALSPACE"`
	SPACESIZE             string `json:"SPACESIZE"`
	ALLAPPS               string `json:"ALLAPPS"`
	NOTLOGINED            string `json:"NOTLOGINED"`
	BEFORELOGIN           string `json:"BEFORELOGIN"`
	NOLOGIN               string `json:"NOLOGIN"`
	LOGIN                 string `json:"LOGIN"`
	SEARCH                string `json:"SEARCH"`
	PSPACE                string `json:"PSPACE"`
	WSPACE                string `json:"WSPACE"`
	NODATA                string `json:"NODATA"`
	NONET                 string `json:"NONET"`
	CATALOG               string `json:"CATALOG"`
	NEWDIR                string `json:"NEWDIR"`
	DIRNAME               string `json:"DIRNAME"`
	DIRNAME15             string `json:"DIRNAME15"`
	ITEMDESC              string `json:"ITEMDESC"`
	ITEMDESC80            string `json:"ITEMDESC80"`
	CATALOGCOVER          string `json:"CATALOGCOVER"`
	ADDMEMBERS            string `json:"ADDMEMBERS"`
	RECYCLEBIN            string `json:"RECYCLEBIN"`
	ASCENDTIME            string `json:"ASCENDTIME"`
	DESCENDTIME           string `json:"DESCENDTIME"`
	ASCENDNAME            string `json:"ASCENDNAME"`
	DESCENDNAME           string `json:"DESCENDNAME"`
	UNABLE                string `json:"UNABLE"`
	SELECTFILE            string `json:"SELECTFILE"`
	CONFIRM               string `json:"CONFIRM"`
	CANCEL                string `json:"CANCEL"`
	LOADING               string `json:"LOADING"`
	DIRNOEMPTY            string `json:"DIRNOEMPTY"`
	ONLYPIC               string `json:"ONLYPIC"`
	EXCEEDMAXUPLOAD       string `json:"EXCEEDMAXUPLOAD"`
	PHONEONLYNUMBER       string `json:"PHONEONLYNUMBER"`
	INVALIDPHONE          string `json:"INVALIDPHONE"`
	NOTSAMEMEMBER         string `json:"NOTSAMEMEMBER"`
	DIRDELETED            string `json:"DIRDELETED"`
	CATALOGDESC           string `json:"CATALOGDESC"`
	SETUP                 string `json:"SETUP"`
	DELETEDIR             string `json:"DELETEDIR"`
	SUREDELETE            string `json:"SUREDELETE"`
	FAILEDDELETE          string `json:"FAILEDDELETE"`
	ITEMEXISTS            string `json:"ITEMEXISTS"`
	FAILEDINVITE          string `json:"FAILEDINVITE"`
	FULLMEMBERS           string `json:"FULLMEMBERS"`
	FAILEDEDIT            string `json:"FAILEDEDIT"`
	ITEMNAME              string `json:"ITEMNAME"`
	ITEMCOVER             string `json:"ITEMCOVER"`
	ITEMMEMBERS           string `json:"ITEMMEMBERS"`
	ITEMLABEL             string `json:"ITEMLABEL"`
	ITEMSERVICES          string `json:"ITEMSERVICES"`
	DELETEITEM            string `json:"DELETEITEM"`
	SUREDELETEITEM        string `json:"SUREDELETEITEM"`
	FILEDELETED           string `json:"FILEDELETED"`
	UNABLESHARE           string `json:"UNABLESHARE"`
	FOLDERNOT             string `json:"FOLDERNOT"`
	FOLDEREXISTS          string `json:"FOLDEREXISTS"`
	FOLDERCREATED         string `json:"FOLDERCREATED"`
	FILEEDITED            string `json:"FILEEDITED"`
	FILEFAILED            string `json:"FILEFAILED"`
	FOLDERDELETED         string `json:"FOLDERDELETED"`
	FILENOTFORMAT         string `json:"FILENOTFORMAT"`
	FILESIZENOT0          string `json:"FILESIZENOT0"`
	FILESIZENOT200        string `json:"FILESIZENOT200"`
	FILEUPLOADED          string `json:"FILEUPLOADED"`
	FILENAMEEXISTS        string `json:"FILENAMEEXISTS"`
	FILENAMEENOT          string `json:"FILENAMEENOT"`
	DESELECTALL           string `json:"DESELECTALL"`
	SELECTALL             string `json:"SELECTALL"`
	SHARE                 string `json:"SHARE"`
	RENAME                string `json:"RENAME"`
	DELETE                string `json:"DELETE"`
	DELETEFILE            string `json:"DELETEFILE"`
	RENAMEFILE            string `json:"RENAMEFILE"`
	REMOVEFOLDERS         string `json:"REMOVEFOLDERS"`
	AFTERFOLDERDELETED    string `json:"AFTERFOLDERDELETED"`
	NEWFOLDER             string `json:"NEWFOLDER"`
	RENAMEFOLDER          string `json:"RENAMEFOLDER"`
	FOLDERNAME15          string `json:"FOLDERNAME15"`
	PERMANENT             string `json:"PERMANENT"`
	DAYS7                 string `json:"DAYS7"`
	DAY1                  string `json:"DAY1"`
	FORWARDED             string `json:"FORWARDED"`
	FORWARDFAILED         string `json:"FORWARDFAILED"`
	NOTSHAREYOUR          string `json:"NOTSHAREYOUR"`
	NAMENOTEMPTY          string `json:"NAMENOTEMPTY"`
	SELECTTOSHARE         string `json:"SELECTTOSHARE"`
	SHARESUCCESS          string `json:"SHARESUCCESS"`
	FILESHARE             string `json:"FILESHARE"`
	SHAREFILE             string `json:"SHAREFILE"`
	TERMVALIDITY          string `json:"TERMVALIDITY"`
	PHONENUM              string `json:"PHONENUM"`
	PETNAME               string `json:"PETNAME"`
	ADDPHONE              string `json:"ADDPHONE"`
	SHAREWECHAT           string `json:"SHAREWECHAT"`
	SHAREPHONE            string `json:"SHAREPHONE"`
	SELECTFILEF           string `json:"SELECTFILEF"`
	FILERESTORED          string `json:"FILERESTORED"`
	DELETION              string `json:"DELETION"`
	RESTOREDIR            string `json:"RESTOREDIR"`
	RESTOREFILES          string `json:"RESTOREFILES"`
	FILE                  string `json:"FILE"`
	REPLACE               string `json:"REPLACE"`
	SKIP                  string `json:"SKIP"`
	KEEP                  string `json:"KEEP"`
	RETAIN7               string `json:"RETAIN7"`
	GOBACK                string `json:"GOBACK"`
	DELETECOM             string `json:"DELETECOM"`
	EMPTYTRASH            string `json:"EMPTYTRASH"`
	RESTORE               string `json:"RESTORE"`
	SUREEMPTYRECYCLE      string `json:"SUREEMPTYRECYCLE"`
	SURERESTORE           string `json:"SURERESTORE"`
	SUREDELETEALL         string `json:"SUREDELETEALL"`
	NOTNEWPROJECT         string `json:"NOTNEWPROJECT"`
	NOTEMPTYSTART         string `json:"NOTEMPTYSTART"`
	LEASTTAG              string `json:"LEASTTAG"`
	TAGEXIST              string `json:"TAGEXIST"`
	MODIFICATION          string `json:"MODIFICATION"`
	ADDED                 string `json:"ADDED"`
	ENTERTAG              string `json:"ENTERTAG"`
	ADDTAGS               string `json:"ADDTAGS"`
	RENDERSERVICE         string `json:"RENDERSERVICE"`
	PROJECTNUM            string `json:"PROJECTNUM"`
	PROJECTSPACE          string `json:"PROJECTSPACE"`
	VALIDUNTIL            string `json:"VALIDUNTIL"`
	MODELLIMIT            string `json:"MODELLIMIT"`
	UPLOADSIZE            string `json:"UPLOADSIZE"`
	CONVERTSIZE           string `json:"CONVERTSIZE"`
	ASSEMBLYSIZE          string `json:"ASSEMBLYSIZE"`
	PEOPLE                string `json:"PEOPLE"`
	NOTHING               string `json:"NOTHING"`
	ONLYUSED              string `json:"ONLYUSED"`
	CREATETASK            string `json:"CREATETASK"`
	DOCWANTDELETE         string `json:"DOCWANTDELETE"`
	VIEWPOINT             string `json:"VIEWPOINT"`
	TASK                  string `json:"TASK"`
	MEETING               string `json:"MEETING"`
	VIEWPOINTDELETED      string `json:"VIEWPOINTDELETED"`
	VIEWPOINTNOTSHARE     string `json:"VIEWPOINTNOTSHARE"`
	DELETEVIEWPOINT       string `json:"DELETEVIEWPOINT"`
	VIEWPOINTSUREDELETE   string `json:"VIEWPOINTSUREDELETE"`
	DELETETASK            string `json:"DELETETASK"`
	SUREDELETETASK        string `json:"SUREDELETETASK"`
	NOTRECOVERED          string `json:"NOTRECOVERED"`
	CREATEDBY             string `json:"CREATEDBY"`
	EXECUTOR              string `json:"EXECUTOR"`
	CREATIONTIME          string `json:"CREATIONTIME"`
	DEADLINE              string `json:"DEADLINE"`
	INEXECUTION           string `json:"INEXECUTION"`
	TOREVIEW              string `json:"TOREVIEW"`
	PASSAUDIT             string `json:"PASSAUDIT"`
	COMPLETED             string `json:"COMPLETED"`
	FAILEDAUDIT           string `json:"FAILEDAUDIT"`
	CANCELTASK            string `json:"CANCELTASK"`
	REOPEN                string `json:"REOPEN"`
	REMARKS               string `json:"REMARKS"`
	SORTNAME              string `json:"SORTNAME"`
	SORTSTATUS            string `json:"SORTSTATUS"`
	SORTCREATOR           string `json:"SORTCREATOR"`
	SORTCREATIONTIME      string `json:"SORTCREATIONTIME"`
	EDIT                  string `json:"EDIT"`
	COPYINVITATION        string `json:"COPYINVITATION"`
	ENTERMEET             string `json:"ENTERMEET"`
	SPONSOR               string `json:"SPONSOR"`
	CONFERENCEPASSWORD    string `json:"CONFERENCEPASSWORD"`
	CONFERENCEID          string `json:"CONFERENCEID"`
	MEETTIME              string `json:"MEETTIME"`
	MOMENT                string `json:"MOMENT"`
	ALLUPLOADED           string `json:"ALLUPLOADED"`
	UPLOADED              string `json:"UPLOADED"`
	FAILEDUPLOADALL       string `json:"FAILEDUPLOADALL"`
	ENCLOSURE             string `json:"ENCLOSURE"`
	LOCALUPLOAD           string `json:"LOCALUPLOAD"`
	SPACEUPLOAD           string `json:"SPACEUPLOAD"`
	ADDCATALOGATT         string `json:"ADDCATALOGATT"`
	EXECUTIVE             string `json:"EXECUTIVE"`
	AUDITOR               string `json:"AUDITOR"`
	NEWTASK               string `json:"NEWTASK"`
	CREATE                string `json:"CREATE"`
	COMPLETECONTENT       string `json:"COMPLETECONTENT"`
	TASKNAMENOTSPACES     string `json:"TASKNAMENOTSPACES"`
	TASKNAMESTARTNOT      string `json:"TASKNAMESTARTNOT"`
	TASKDESCNOTEMPTY      string `json:"TASKDESCNOTEMPTY"`
	TASKTIME              string `json:"TASKTIME"`
	TASKCREATED           string `json:"TASKCREATED"`
	EDITTASK              string `json:"EDITTASK"`
	DETAILS               string `json:"DETAILS"`
	COMMENT               string `json:"COMMENT"`
	PARTICIPANTS          string `json:"PARTICIPANTS"`
	TASKTAG               string `json:"TASKTAG"`
	TASKDESC              string `json:"TASKDESC"`
	TASKNAME              string `json:"TASKNAME"`
	ENTERTASKNAME         string `json:"ENTERTASKNAME"`
	TASKDESCIPTION        string `json:"TASKDESCIPTION"`
	CREATORNAME           string `json:"CREATORNAME"`
	PRESERVATION          string `json:"PRESERVATION"`
	TASKTAGNOTEMPTY       string `json:"TASKTAGNOTEMPTY"`
	EXECUTORNOTEMPTY      string `json:"EXECUTORNOTEMPTY"`
	REVIEWERNOTEMPTY      string `json:"REVIEWERNOTEMPTY"`
	TASKEDITED            string `json:"TASKEDITED"`
	TASKNAMEENTER         string `json:"TASKNAMEENTER"`
	TASKRELEASED          string `json:"TASKRELEASED"`
	APPROVED              string `json:"APPROVED"`
	AUDITFAILED           string `json:"AUDITFAILED"`
	CANCELLED             string `json:"CANCELLED"`
	TASKRECORD            string `json:"TASKRECORD"`
	CHANGETASKSTATUS      string `json:"CHANGETASKSTATUS"`
	ADDVIEWPOINT          string `json:"ADDVIEWPOINT"`
	CREATOR               string `json:"CREATOR"`
	NOACCESSVIEWPOINT     string `json:"NOACCESSVIEWPOINT"`
	TASKPERSPECTIVE       string `json:"TASKPERSPECTIVE"`
	TASKATT               string `json:"TASKATT"`
	UNABLEPREVIEW         string `json:"UNABLEPREVIEW"`
	FILEPROCESSED         string `json:"FILEPROCESSED"`
	FILELIMIT             string `json:"FILELIMIT"`
	NOTACCESSFILE         string `json:"NOTACCESSFILE"`
	ASSOCIATEDFILES       string `json:"ASSOCIATEDFILES"`
	GETCAPTCHA            string `json:"GETCAPTCHA"`
	PHONENOTEMPTY         string `json:"PHONENOTEMPTY"`
	WRONGPHONE            string `json:"WRONGPHONE"`
	PHONEAPPLET           string `json:"PHONEAPPLET"`
	CODEWRONG             string `json:"CODEWRONG"`
	CODENOTEMPTY          string `json:"CODENOTEMPTY"`
	CODEBINDPHONE         string `json:"CODEBINDPHONE"`
	BINDPHONE             string `json:"BINDPHONE"`
	INPUTPHONE            string `json:"INPUTPHONE"`
	INPUTCODE             string `json:"INPUTCODE"`
	BINDNOW               string `json:"BINDNOW"`
	ENTERPHONE            string `json:"ENTERPHONE"`
	ENTERCORRECTPHONE     string `json:"ENTERCORRECTPHONE"`
	INPUTPASSWORD         string `json:"INPUTPASSWORD"`
	INCORRECTPASSWORD     string `json:"INCORRECTPASSWORD"`
	ENTERCODE             string `json:"ENTERCODE"`
	ENTERCORRECTCODE      string `json:"ENTERCORRECTCODE"`
	VERIFICATIONCODE      string `json:"VERIFICATIONCODE"`
	ENTERPASSW            string `json:"ENTERPASSW"`
	FORGETPASSW           string `json:"FORGETPASSW"`
	ACCOUNTPASSWLOGIN     string `json:"ACCOUNTPASSWLOGIN"`
	SMSLOGIN              string `json:"SMSLOGIN"`
	NOTACCOUNT            string `json:"NOTACCOUNT"`
	REGISTER              string `json:"REGISTER"`
	EDITMEET              string `json:"EDITMEET"`
	ENTERMEETSUBJECT      string `json:"ENTERMEETSUBJECT"`
	MEETBEGIN             string `json:"MEETBEGIN"`
	MEETEND               string `json:"MEETEND"`
	ENTERCONFERENCEPASSW  string `json:"ENTERCONFERENCEPASSW"`
	ENTERMEETID           string `json:"ENTERMEETID"`
	SUREINFORMATION       string `json:"SUREINFORMATION"`
	COMPLETE              string `json:"COMPLETE"`
	COMCONTENTS           string `json:"COMCONTENTS"`
	MEETNOTSPACE          string `json:"MEETNOTSPACE"`
	MEETTHEMECHARS        string `json:"MEETTHEMECHARS"`
	CONFERENCEPASSWDIGITS string `json:"CONFERENCEPASSWDIGITS"`
	CONFERENCEIDDIGITS    string `json:"CONFERENCEIDDIGITS"`
	EDITSUCCEEDED         string `json:"EDITSUCCEEDED"`
	SHARESITE             string `json:"SHARESITE"`
	MULTCHOICE            string `json:"MULTCHOICE"`
	SELECTMEMBER          string `json:"SELECTMEMBER"`
	MEETSHARE             string `json:"MEETSHARE"`
	HELPCENTER            string `json:"HELPCENTER"`
	PIECHART              string `json:"PIECHART"`
	PROJECTNOTEMPTY       string `json:"PROJECTNOTEMPTY"`
	PROJECTDESCNOTEMPTY   string `json:"PROJECTDESCNOTEMPTY"`
	PROJECTLABELNOTEMPTY  string `json:"PROJECTLABELNOTEMPTY"`
	PROJECTLABELS20       string `json:"PROJECTLABELS20"`
	NEWPROJECT            string `json:"NEWPROJECT"`
	PROJECTNAME20         string `json:"PROJECTNAME20"`
	PROJECTDESC200        string `json:"PROJECTDESC200"`
	LABELSPROJECTS        string `json:"LABELSPROJECTS"`
	ESTABLISH             string `json:"ESTABLISH"`
	WECHATTHIRDLOGIN      string `json:"WECHATTHIRDLOGIN"`
	WECHATSELFLOGIN       string `json:"WECHATSELFLOGIN"`
	PUBLISHER             string `json:"PUBLISHER"`
	SUBSCRIBED            string `json:"SUBSCRIBED"`
	SUBSCRIBEAGAIN        string `json:"SUBSCRIBEAGAIN"`
	SUBSCRIBENOW          string `json:"SUBSCRIBENOW"`
	EDITION               string `json:"EDITION"`
	RELEASETIME           string `json:"RELEASETIME"`
	FUNCTIONINTR          string `json:"FUNCTIONINTR"`
	ADDSERVICES           string `json:"ADDSERVICES"`
	USERAGR               string `json:"USERAGR"`
	FAILEDUPLOADIMG       string `json:"FAILEDUPLOADIMG"`
	IMGUPLOADED           string `json:"IMGUPLOADED"`
	SENDOUT               string `json:"SENDOUT"`
	YOURPROBLEM           string `json:"YOURPROBLEM"`
	PHOTOGRAPH            string `json:"PHOTOGRAPH"`
	SELECTALBUM           string `json:"SELECTALBUM"`
	SURETODELETE          string `json:"SURETODELETE"`
	USERREG               string `json:"USERREG"`
	NEWPASSWFORMAT        string `json:"NEWPASSWFORMAT"`
	NEWPASSWLEN           string `json:"NEWPASSWLEN"`
	INPUTCORRECTPHONE     string `json:"INPUTCORRECTPHONE"`
	ENTERCORRECTPASSW     string `json:"ENTERCORRECTPASSW"`
	RETRIEVEPASSW         string `json:"RETRIEVEPASSW"`
	REGACCOUNT            string `json:"REGACCOUNT"`
	BINDLOGIN             string `json:"BINDLOGIN"`
	CONFIRMLOGIN          string `json:"CONFIRMLOGIN"`
	PASSWORD              string `json:"PASSWORD"`
	NEWPASSWORD           string `json:"NEWPASSWORD"`
	BINDACCEPT            string `json:"BINDACCEPT"`
	HAVEACC               string `json:"HAVEACC"`
	ENTER                 string `json:"ENTER"`
	DIGITSLETTER16        string `json:"DIGITSLETTER16"`
	UPLOADADDRESS         string `json:"UPLOADADDRESS"`
	UPLOADFILE            string `json:"UPLOADFILE"`
	PREVIEWFILELIMIT      string `json:"PREVIEWFILELIMIT"`
	NOTHAVEACCESS         string `json:"NOTHAVEACCESS"`
	VIEWMODELDELEDTED     string `json:"VIEWMODELDELEDTED"`
	TODELETED             string `json:"TODELETED"`
	SIGNIN                string `json:"SIGNIN"`
}
