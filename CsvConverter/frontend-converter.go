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

type FrontEnder struct {
	body              FrontEndBody
	bodyTradition     FrontEndBody
	bodyEnglish       FrontEndBody
	FileName          string
	FileNameTradition string
	FileNameEnglish   string
	seek              int
}

func (s *FrontEnder) ToCsv(csvW *csv.Writer) error {
	s.openFiles()
	//写入:分隔符+index,列名,行
	err := csvW.Write([]string{SPLITE + strconv.FormatInt(int64(FRONTENDINDEX), 10)})
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
	err = csvW.Write([]string{SPLITE + strconv.FormatInt(int64(FRONTENDINDEX), 10)})
	csvW.Flush()
	return nil
}

//转回
func (s *FrontEnder) ToOrigin(reader *csv.Reader) error {
	v := reflect.ValueOf(&s.body)
	vt := reflect.ValueOf(&s.bodyTradition)
	ve := reflect.ValueOf(&s.bodyEnglish)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			panic(">未完成前,文件读完")
		}
		//结束分隔符
		if strings.Contains(record[0], SPLITE) {
			//保存
			s.OriginSave()
			s.seek = 0
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
func (s *FrontEnder) OriginSave() {
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
func (s *FrontEnder) GetFileNames() []string {
	return []string{s.FileName, s.FileNameTradition, s.FileNameEnglish}
}
func (t *FrontEnder) openFiles() {
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

type FrontEndBody struct {
	Num403                       string `json:"403"`
	Num404                       string `json:"404"`
	Num500                       string `json:"500"`
	CHARGECHOOSENAME             string `json:"CHARGE_CHOOSE_NAME"`
	CHARGEACCOUNTSPAYABLE        string `json:"CHARGE_ACCOUNTS_PAYABLE"`
	CHARGEREALPAYMENT            string `json:"CHARGE_REAL_PAYMENT"`
	CHARGEPROJECTCHANGEPACKAGE   string `json:"CHARGE_PROJECT_CHANGE_PACKAGE"`
	CHARGEPROJECTSTORAGE         string `json:"CHARGE_PROJECT_STORAGE"`
	CHARGEPROJECTPEOPLE          string `json:"CHARGE_PROJECT_PEOPLE"`
	CHARGEPROJECTPACKAGE         string `json:"CHARGE_PROJECT_PACKAGE"`
	CHARGECONFIRMCONFIG          string `json:"CHARGE_CONFIRM_CONFIG"`
	CHARGEPROJECTNAME            string `json:"CHARGE_PROJECT_NAME"`
	CHARGEPROJECTID              string `json:"CHARGE_PROJECT_ID"`
	CHARGEBUY                    string `json:"CHARGE_BUY"`
	CHARGERENDERER               string `json:"CHARGE_RENDERER"`
	CHARGEBUYTIME                string `json:"CHARGE_BUY_TIME"`
	CHARGEMOMEY                  string `json:"CHARGE_MOMEY"`
	CHARGEACCOUNTNAME            string `json:"CHARGE_ACCOUNT_NAME"`
	CHARGECHOOSEPACKAGE          string `json:"CHARGE_CHOOSE_PACKAGE"`
	CHARGEMORESPACE              string `json:"CHARGE_MORE_SPACE"`
	CHARGEPURCHASEAPPLICATION    string `json:"CHARGE_PURCHASE_APPLICATION"`
	UNITPEOPLE                   string `json:"UNIT_PEOPLE"`
	UNITYEAR                     string `json:"UNIT_YEAR"`
	REFRESH                      string `json:"REFRESH"`
	CHOOSEPROJECT                string `json:"CHOOSE_PROJECT"`
	CHARGEUPGRADETIP             string `json:"CHARGE_UPGRADE_TIP"`
	UPLOAD                       string `json:"UPLOAD"`
	UPLOADLIST                   string `json:"UPLOAD_LIST"`
	PROGRESSLIST                 string `json:"PROGRESS_LIST"`
	COMPLETELIST                 string `json:"COMPLETE_LIST"`
	UPLOADCOMPLETE               string `json:"UPLOAD_COMPLETE"`
	CLOSEUPLOAD                  string `json:"CLOSE_UPLOAD"`
	CLOSETIP                     string `json:"CLOSE_TIP"`
	CLICKDRAWUPLOAD              string `json:"CLICK_DRAW_UPLOAD"`
	UPLOADSUPPORTFORMAT          string `json:"UPLOAD_SUPPORT_FORMAT"`
	STARTUPLOAD                  string `json:"START_UPLOAD"`
	BASICINFOTITLE               string `json:"BASICINFO_TITLE"`
	USERNAME                     string `json:"USER_NAME"`
	NAME                         string `json:"NAME"`
	USERSEX                      string `json:"USER_SEX"`
	USERBIRTHDAY                 string `json:"USER_BIRTHDAY"`
	USERCOMPANY                  string `json:"USER_COMPANY"`
	USERAREA                     string `json:"USER_AREA"`
	CHANGEPHONE                  string `json:"CHANGE_PHONE"`
	NEWPHONE                     string `json:"NEW_PHONE"`
	REVERTIP                     string `json:"REVER_TIP"`
	BOY                          string `json:"BOY"`
	GIRL                         string `json:"GIRL"`
	LINK                         string `json:"LINK"`
	FETCHCODE                    string `json:"FETCH_CODE"`
	SELECTDATE                   string `json:"SELECT_DATE"`
	SELECTPROVINCE               string `json:"SELECT_PROVINCE"`
	SELECTCITY                   string `json:"SELECT_CITY"`
	SELECTAREA                   string `json:"SELECT_AREA"`
	SELECTCOUNTY                 string `json:"SELECT_COUNTY"`
	ENTERPOSITION                string `json:"ENTER_POSITION"`
	ENTERSTARTIME                string `json:"ENTER_STARTIME"`
	ENTERENDTIME                 string `json:"ENTER_ENDTIME"`
	ADDRESUME                    string `json:"ADDRESUME"`
	ADDRESUMETIP                 string `json:"ADDRESUME_TIP"`
	CHANGEWEIXIN                 string `json:"CHANGE_WEIXIN"`
	BINDWEIXIN                   string `json:"BIND_WEIXIN"`
	CHANGEQQ                     string `json:"CHANGE_QQ"`
	BINDQQ                       string `json:"BIND_QQ"`
	BINDSUCCESS                  string `json:"BIND_SUCCESS"`
	UNBOUNDSUC                   string `json:"UNBOUND_SUC"`
	SCANCODE                     string `json:"SCANCODE"`
	SCANCODEQQ                   string `json:"SCANCODE_QQ"`
	DIFFPASSWORD                 string `json:"DIFF_PASSWORD"`
	ORIGINALPASSWORDEMPTY        string `json:"ORIGINAL_PASSWORD_EMPTY"`
	NEWPASSWORDEMPTY             string `json:"NEW_PASSWORD_EMPTY"`
	NEWPASSWORDLENGTH            string `json:"NEW_PASSWORD_LENGTH"`
	COMFIRMPASSWORDEMPTY         string `json:"COMFIRM_PASSWORD_EMPTY"`
	COMFIRMPASSWORDLENGTH        string `json:"COMFIRM_PASSWORD_LENGTH"`
	MODIFYSUCCESS                string `json:"MODIFY_SUCCESS"`
	SETTINGSUCCESS               string `json:"SETTING_SUCCESS"`
	EMAILREPEAT                  string `json:"EMAIL_REPEAT"`
	EMAILINVALID                 string `json:"EMAIL_INVALID"`
	EMAILEMPTY                   string `json:"EMAIL_EMPTY"`
	EMAILREGTIP                  string `json:"EMAIL_REG_TIP"`
	BINDEMAILERROR               string `json:"BIND_EMAIL_ERROR"`
	MODIFYEMAILERROR             string `json:"MODIFY_EMAIL_ERROR"`
	BINDSUCCESS2                 string `json:"BIND_SUCCESS2"`
	REMOVEBIND                   string `json:"REMOVE_BIND"`
	REMOVEBINDCONFIM             string `json:"REMOVE_BIND_CONFIM"`
	PHONEREPEAT                  string `json:"PHONE_REPEAT"`
	PHONEINVALID                 string `json:"PHONE_INVALID"`
	PHONEEMPTY                   string `json:"PHONE_EMPTY"`
	PHONEREGTIP                  string `json:"PHONE_REG_TIP"`
	PERSONALSPACETOTAL           string `json:"PERSONAL_SPACE_TOTAL"`
	USEDSPACETOTAL               string `json:"USED_SPACE_TOTAL"`
	UNUSEDSPACETOTAL             string `json:"UNUSED_SPACE_TOTAL"`
	TOTALSPACE                   string `json:"TOTAL_SPACE"`
	MUNICIPALDISTRICTS           string `json:"MUNICIPAL_DISTRICTS"`
	USERNAMEEMPTY                string `json:"USER_NAME_EMPTY"`
	INPUTUSERNAME                string `json:"INPUT_USER_NAME"`
	MYSHARE                      string `json:"MY_SHARE"`
	MYFRIENDS                    string `json:"MY_FRIENDS"`
	MYRECOVERY                   string `json:"MY_RECOVERY"`
	RECOVERTIP                   string `json:"RECOVER_TIP"`
	FILENAME                     string `json:"FILENAME"`
	SHAREDATE                    string `json:"SHAREDATE"`
	VALIDDATE                    string `json:"VALIDDATE"`
	VIEWCOUNT                    string `json:"VIEWCOUNT"`
	SHARELINKFAIL                string `json:"SHARE_LINK_FAIL"`
	SHARELINKNOT                 string `json:"SHARE_LINK_NOT"`
	SHARELINKVIEWED              string `json:"SHARE_LINK_VIEWED"`
	DELETESHARELINK              string `json:"DELETE_SHARE_LINK"`
	DELETESHARELINKTIP           string `json:"DELETE_SHARE_LINK_TIP"`
	HASFAILURE                   string `json:"HASFAILURE"`
	SHAREAGAIN                   string `json:"SHARE_AGAIN"`
	DELETESHARE                  string `json:"DELETE_SHARE"`
	ORIGINALLOCATION             string `json:"ORIGINAL_LOCATION"`
	DELETEDATE                   string `json:"DELETE_DATE"`
	UPLOADER                     string `json:"UPLOADER"`
	REMAININGTIME                string `json:"REMAINING_TIME"`
	FILESIZE                     string `json:"FILE_SIZE"`
	FILE                         string `json:"FILE"`
	ALLFILE                      string `json:"ALL_FILE"`
	TODELETE                     string `json:"TO_DELETE"`
	REPLACE                      string `json:"REPLACE"`
	APPLYTOALL                   string `json:"APPLY_TO_ALL"`
	KEEPBOTH                     string `json:"KEEP_BOTH"`
	SKIP                         string `json:"SKIP"`
	RESTOREDIRECTORY             string `json:"RESTORE_DIRECTORY"`
	RESTOREFILE                  string `json:"RESTORE_FILE"`
	DELETING                     string `json:"DELETING"`
	DELETESUCCESS                string `json:"DELETE_SUCCESS"`
	DELETEDEEP                   string `json:"DELETE_DEEP"`
	DELETINGTIP                  string `json:"DELETING_TIP"`
	CHOOSEFILE                   string `json:"CHOOSE_FILE"`
	REDUCTION                    string `json:"REDUCTION"`
	REDUCTIONTIP                 string `json:"REDUCTION_TIP"`
	EMPTYRECYCLE                 string `json:"EMPTY_RECYCLE"`
	PUTBACK                      string `json:"PUT_BACK"`
	PUTBACKSUCCESS               string `json:"PUT_BACK_SUCCESS"`
	EMPTYRECYCLECONFIRM          string `json:"EMPTY_RECYCLE_CONFIRM"`
	MESSAGE                      string `json:"MESSAGE"`
	MESSAGECENTER                string `json:"MESSAGE_CENTER"`
	MESSAGECONTENT               string `json:"MESSAGE_CONTENT"`
	DELETEMESSAGE                string `json:"DELETE_MESSAGE"`
	DELETEMESSAGECONFIRM         string `json:"DELETE_MESSAGE_CONFIRM"`
	DELETEMESSAGECHOOSE          string `json:"DELETE_MESSAGE_CHOOSE"`
	ALL                          string `json:"ALL"`
	UNREAD                       string `json:"UNREAD"`
	READ                         string `json:"READ"`
	DELETERECORDS                string `json:"DELETE_RECORDS"`
	SETREAD                      string `json:"SET_READ"`
	SETMESSAGE                   string `json:"SET_MESSAGE"`
	SETMESSAGESUCCESS            string `json:"SET_MESSAGE_SUCCESS"`
	STATIONPUSH                  string `json:"STATION_PUSH"`
	USINGTUTORIAL                string `json:"USING_TUTORIAL"`
	ABOUTPLATFORM                string `json:"ABOUT_PLATFORM"`
	USERMANUAL                   string `json:"USER_MANUAL"`
	TECHNICALSUPPORT             string `json:"TECHNICAL_SUPPORT"`
	WECHARCODE                   string `json:"WECHAR_CODE"`
	CURRENTVERSION               string `json:"CURRENT_VERSION"`
	NEWVERSION                   string `json:"NEW_VERSION"`
	NEWVERSIONDESCRIPTION        string `json:"NEW_VERSION_DESCRIPTION"`
	RESERVED                     string `json:"RESERVED"`
	PACKAGETITLE1                string `json:"PACKAGE_TITLE1"`
	PACKAGETITLE2                string `json:"PACKAGE_TITLE2"`
	PACKAGETITLE3                string `json:"PACKAGE_TITLE3"`
	PACKAGETITLE4                string `json:"PACKAGE_TITLE4"`
	PACKAGETITLE5                string `json:"PACKAGE_TITLE5"`
	PACKAGETITLE6                string `json:"PACKAGE_TITLE6"`
	PACKAGERSTYPE1               string `json:"PACKAGE_RSTYPE1"`
	PACKAGERSTYPE2               string `json:"PACKAGE_RSTYPE2"`
	PACKAGERSTYPE3               string `json:"PACKAGE_RSTYPE3"`
	PACKAGERSTYPE4               string `json:"PACKAGE_RSTYPE4"`
	PACKAGERSTYPE5               string `json:"PACKAGE_RSTYPE5"`
	PACKAGERSTYPE6               string `json:"PACKAGE_RSTYPE6"`
	PROJECTOPENINGSIZE           string `json:"PROJECT_OPENING_SIZE"`
	NOTACCESSOPENFILE1           string `json:"NOT_ACCESS_OPEN_FILE1"`
	NOTACCESSOPENFILE2           string `json:"NOT_ACCESS_OPEN_FILE2"`
	NOTACCESSOPENFILE3           string `json:"NOT_ACCESS_OPEN_FILE3"`
	NOTACCESSOPENFILE4           string `json:"NOT_ACCESS_OPEN_FILE4"`
	MYSPACENOTSUPPORT2           string `json:"MYSPACE_NOT_SUPPORT2"`
	NOTACCESSOPENASSEMBLE        string `json:"NOT_ACCESS_OPEN_ASSEMBLE"`
	NOTACCESSOPENASSEMBLE2       string `json:"NOT_ACCESS_OPEN_ASSEMBLE2"`
	DELETEFILESSUCCESS           string `json:"DELETE_FILES_SUCCESS"`
	FILESIZEZERO                 string `json:"FILE_SIZE_ZERO"`
	UPLOADASSEMBLESERVERTIP      string `json:"UPLOAD_ASSEMBLE_SERVER_TIP"`
	NOTSUPPORTOPEN               string `json:"NOT_SUPPORT_OPEN"`
	PROJECTUPLOADTIP1            string `json:"PROJECT_UPLOAD_TIP1"`
	PROJECTUPLOADTIP2            string `json:"PROJECT_UPLOAD_TIP2"`
	PROJECTUPLOADTIP3            string `json:"PROJECT_UPLOAD_TIP3"`
	PROJECTUPLOADTIP4            string `json:"PROJECT_UPLOAD_TIP4"`
	PROJECTUPLOADTIP5            string `json:"PROJECT_UPLOAD_TIP5"`
	SPACEINSUFFICIENT1           string `json:"SPACE_INSUFFICIENT1"`
	SPACEINSUFFICIENT2           string `json:"SPACE_INSUFFICIENT2"`
	UPDATAFILEFORMAT             string `json:"UPDATA_FILE_FORMAT"`
	DESCRIPTIONLIMITATIONS       string `json:"DESCRIPTION_LIMITATIONS"`
	INPUTVERSION                 string `json:"INPUT_VERSION"`
	CHOOSEFILES                  string `json:"CHOOSE_FILES"`
	NOTSUPPORTFORMAT             string `json:"NOT_SUPPORT_FORMAT"`
	STOP                         string `json:"STOP"`
	MEETINGCREATED               string `json:"MEETING_CREATED"`
	MEETINFOREDITED              string `json:"MEET_INFOR_EDITED"`
	MEETSHAREING                 string `json:"MEET_SHAREING"`
	NOTHINGCLIPBOARD             string `json:"NOTHING_CLIPBOARD"`
	BROWSERNOTPASTE              string `json:"BROWSER_NOT_PASTE"`
	CONFERENCETHEME              string `json:"CONFERENCE_THEME"`
	MEETINGTIME                  string `json:"MEETING_TIME"`
	MEETINGID                    string `json:"MEETING_ID"`
	CONFERENCEPASSW              string `json:"CONFERENCE_PASSW"`
	OTHERWISE                    string `json:"OTHERWISE"`
	MEETINGNOTEMPTY              string `json:"MEETING_NOT_EMPTY"`
	MEETINGTHEME                 string `json:"MEETING_THEME"`
	MEETINGTIMENOTEMPTY          string `json:"MEETING_TIME_NOT_EMPTY"`
	MEETINGPASSWFORMAT           string `json:"MEETING_PASSW_FORMAT"`
	MEETINGIDNOTEMPTY            string `json:"MEETING_ID_NOT_EMPTY"`
	MEETINGIDFORMAT              string `json:"MEETING_ID_FORMAT"`
	COMPLETEBTN                  string `json:"COMPLETE_BTN"`
	EDITBTN                      string `json:"EDIT_BTN"`
	EDITMEETING                  string `json:"EDIT_MEETING"`
	CREATEMEETING                string `json:"CREATE_MEETING"`
	PASTEMEETING                 string `json:"PASTE_MEETING"`
	SPONSOR                      string `json:"SPONSOR"`
	OPERATION                    string `json:"OPERATION"`
	ENTERMEETING                 string `json:"ENTER_MEETING"`
	SHARESITE                    string `json:"SHARE_SITE"`
	SHAREMEETING                 string `json:"SHARE_MEETING"`
	ONLYCREATORDELETE            string `json:"ONLY_CREATOR_DELETE"`
	DELETEMEETING                string `json:"DELETE_MEETING"`
	DELETEEDMEETINGRESTORED      string `json:"DELETEED_MEETING_RESTORED"`
	MEETINGNUM                   string `json:"MEETING_NUM"`
	COPYBTN                      string `json:"COPY_BTN"`
	ENTERSEARCH                  string `json:"ENTER_SEARCH"`
	NEWMEETING                   string `json:"NEW_MEETING"`
	COPYINVIT                    string `json:"COPY_INVIT"`
	SELECTSPACE                  string `json:"SELECT_SPACE"`
	INVOICE                      string `json:"INVOICE"`
	INVOICEINFOR                 string `json:"INVOICE_INFOR"`
	VATINVOICE                   string `json:"VAT_INVOICE"`
	COMPNAME                     string `json:"COMP_NAME"`
	ENTERCOMPANY                 string `json:"ENTER_COMPANY"`
	TAXPAYERCODE                 string `json:"TAXPAYER_CODE"`
	ENTERTAXPAYERCODE            string `json:"ENTER_TAXPAYER_CODE"`
	TAXPAYERNUMBER               string `json:"TAXPAYER_NUMBER"`
	REGISTERADDRESS              string `json:"REGISTER_ADDRESS"`
	ENTERREGADDR                 string `json:"ENTER_REG_ADDR"`
	REGISTERTEL                  string `json:"REGISTER_TEL"`
	ENTERREGTEL                  string `json:"ENTER_REG_TEL"`
	BANK                         string `json:"BANK"`
	ENTERBANK                    string `json:"ENTER_BANK"`
	ENTERBANKNAME                string `json:"ENTER_BANK_NAME"`
	BANKACC                      string `json:"BANK_ACC"`
	ENTERBANKACC                 string `json:"ENTER_BANK_ACC"`
	TOTAL                        string `json:"TOTAL"`
	PENDING                      string `json:"PENDING"`
	TOPAID                       string `json:"TO_PAID"`
	PAYMENTAMOUNT                string `json:"PAYMENT_AMOUNT"`
	FILESIZEEXCEED               string `json:"FILE_SIZE_EXCEED"`
	ONLYFILESTYPE                string `json:"ONLY_FILES_TYPE"`
	MAXUPLOAD                    string `json:"MAX_UPLOAD"`
	LOCALUPLOAD                  string `json:"LOCAL_UPLOAD"`
	SPACEUPLOAD                  string `json:"SPACE_UPLOAD"`
	NEWVP                        string `json:"NEW_VP"`
	SELECTVP                     string `json:"SELECT_VP"`
	ADDATT                       string `json:"ADD_ATT"`
	ADDCATALOGATT                string `json:"ADD_CATALOG_ATT"`
	ATTACHMENT                   string `json:"ATTACHMENT"`
	ATTUPLOADED                  string `json:"ATT_UPLOADED"`
	LOADMORE                     string `json:"LOAD_MORE"`
	CLICKLOADMORE                string `json:"CLICK_LOAD_MORE"`
	SLIDEUP                      string `json:"SLIDE_UP"`
	UPLOADFORMAT                 string `json:"UPLOAD_FORMAT"`
	ADDMEMBERS                   string `json:"ADD_MEMBERS"`
	ADMINADDMEMBERS              string `json:"ADMIN_ADD_MEMBERS"`
	SELECTEXECUTIVE              string `json:"SELECT_EXECUTIVE"`
	SELECTPARTICIPANTS           string `json:"SELECT_PARTICIPANTS"`
	SELECTAUDITOR                string `json:"SELECT_AUDITOR"`
	TASKVIEWPOINT                string `json:"TASK_VIEWPOINT"`
	TASKATTACHMENT               string `json:"TASK_ATTACHMENT"`
	TASKASSOCIATIONS             string `json:"TASK_ASSOCIATIONS"`
	DELETETASKCOMMENT            string `json:"DELETE_TASK_COMMENT"`
	DELETETASKCOMMENTTIP         string `json:"DELETE_TASK_COMMENT_TIP"`
	NOTVIEWTIP                   string `json:"NOT_VIEW_TIP"`
	NOTVIEWTIP2                  string `json:"NOT_VIEW_TIP2"`
	FILEPROCESSING               string `json:"FILE_PROCESSING"`
	FILEPROCESSINGFAIL           string `json:"FILE_PROCESSING_FAIL"`
	CANNOTPREVIEW                string `json:"CANNOT_PREVIEW"`
	EXECUTIVEMAX                 string `json:"EXECUTIVE_MAX"`
	PARTICIPATORSMAX             string `json:"PARTICIPATORS_MAX"`
	TASKTAG                      string `json:"TASK_TAG"`
	TASKTAGEMPTY                 string `json:"TASK_TAG_EMPTY"`
	REVIEWERSEMPTY               string `json:"REVIEWERS_EMPTY"`
	EXECUTIVEEMPTY               string `json:"EXECUTIVE_EMPTY"`
	CURRENTEXECUTOR              string `json:"CURRENT_EXECUTOR"`
	TASKDESCRIPTION              string `json:"TASK_DESCRIPTION"`
	TASKDESCRIPTIONEMPTY         string `json:"TASK_DESCRIPTION_EMPTY"`
	INPUTTASKDESCRIPTION         string `json:"INPUT_TASK_DESCRIPTION"`
	SAVECURRENTOPERATION         string `json:"SAVE_CURRENT_OPERATION"`
	INPUTALLINFO                 string `json:"INPUT_ALL_INFO"`
	TASKNAME                     string `json:"TASK_NAME"`
	CREATEUSER2                  string `json:"CREATE_USER2"`
	EXECUTOR                     string `json:"EXECUTOR"`
	REVIEWERS                    string `json:"REVIEWERS"`
	PARTICIPANTS                 string `json:"PARTICIPANTS"`
	NOPARTICIPANTS               string `json:"NO_PARTICIPANTS"`
	CREATETIME                   string `json:"CREATE_TIME"`
	ENDINGTIME                   string `json:"ENDING_TIME"`
	TASKNOTEMPTY                 string `json:"TASK_NOT_EMPTY"`
	TASKNOTSPACES                string `json:"TASK_NOT_SPACES"`
	TASKNOTSTARTS                string `json:"TASK_NOT_STARTS"`
	TASKEXISTS                   string `json:"TASK_EXISTS"`
	TASKEXECUTOR                 string `json:"TASK_EXECUTOR"`
	TASKREVIEWER                 string `json:"TASK_REVIEWER"`
	TASKDEADLINE                 string `json:"TASK_DEADLINE"`
	NEWTASK                      string `json:"NEW_TASK"`
	ENTERTASKNAME                string `json:"ENTER_TASK_NAME"`
	SELECTDEADLINE               string `json:"SELECT_DEADLINE"`
	SELECTTASKT                  string `json:"SELECT_TASKT"`
	ADDLABEL                     string `json:"ADD_LABEL"`
	ADMINADDLABEL                string `json:"ADMIN_ADD_LABEL"`
	FOUNDER                      string `json:"FOUNDER"`
	TASKDESC                     string `json:"TASKDESC"`
	TASKSTATUS                   string `json:"TASK_STATUS"`
	EXEC                         string `json:"EXEC"`
	REVIEWER                     string `json:"REVIEWER"`
	PARTIC                       string `json:"PARTIC"`
	TASKLIST                     string `json:"TASKLIST"`
	EXPORTFILE                   string `json:"EXPORT_FILE"`
	EXPORTTASK                   string `json:"EXPORT_TASK"`
	EXPORT                       string `json:"EXPORT"`
	SELECTEXPORT                 string `json:"SELECT_EXPORT"`
	FILE_NAME                    string `json:"FILE_NAME"`
	FILETYPE                     string `json:"FILE_TYPE"`
	TASKCREATED                  string `json:"TASK_CREATED"`
	TASKDELETED                  string `json:"TASK_DELETED"`
	UPDATETASKSTATUS             string `json:"UPDATE_TASK_STATUS"`
	NOTPERMISVIEW                string `json:"NOT_PERMIS_VIEW"`
	UPDATETASK                   string `json:"UPDATE_TASK"`
	NOVPFOUND                    string `json:"NO_VP_FOUND"`
	VPASSDELETED                 string `json:"VP_ASS_DELETED"`
	WILLPERMISS                  string `json:"WILL_PERMISS"`
	COMPANYNAME                  string `json:"COMPANY_NAME"`
	PLATFORMNAME                 string `json:"PLATFORM_NAME"`
	PRODUCTNAME                  string `json:"PRODUCT_NAME"`
	PRODUCTPLUGINNAME            string `json:"PRODUCT_PLUGIN_NAME"`
	ACCOUNTREGISTER              string `json:"ACCOUNT_REGISTER"`
	ACCOUNTNOTREGISTER           string `json:"ACCOUNT_NOT_REGISTER"`
	EMAILACCOUNTNOTEXIST         string `json:"EMAIL_ACCOUNT_NOT_EXIST"`
	USERACCOUNTED                string `json:"USER_ACCOUNTED"`
	BOUNDQQ                      string `json:"BOUNDQQ"`
	BOUNDWECHAT                  string `json:"BOUND_WECHAT"`
	VERIFYCODEERROR              string `json:"VERIFY_CODE_ERROR"`
	PASSWORDERRORINPUT           string `json:"PASSWORD_ERROR_INPUT"`
	CODEERRORINPUT               string `json:"CODE_ERROR_INPUT"`
	PHONEHASACCOUNTEDANDREBIND   string `json:"PHONE_HAS_ACCOUNTED_AND_REBIND"`
	ALLOWMERGEACCOUNT            string `json:"ALLOW_MERGE_ACCOUNT"`
	NETWORKLAWTIP                string `json:"NETWORK_LAW_TIP"`
	ACCOUNT                      string `json:"ACCOUNT"`
	LOGIN                        string `json:"LOGIN"`
	PARTICULARS                  string `json:"PARTICULARS"`
	OUTIMG                       string `json:"OUT_IMG"`
	OUTINGIMG                    string `json:"OUTING_IMG"`
	PLUGIN                       string `json:"PLUGIN"`
	MODEL                        string `json:"MODEL"`
	DOWNLOAD                     string `json:"DOWNLOAD"`
	PASSWORD                     string `json:"PASSWORD"`
	REGISTER                     string `json:"REGISTER"`
	RENAME                       string `json:"RENAME"`
	REMOVE                       string `json:"REMOVE"`
	INPUT                        string `json:"INPUT"`
	SAVE                         string `json:"SAVE"`
	RESET                        string `json:"RESET"`
	IMAGE                        string `json:"IMAGE"`
	COVER                        string `json:"COVER"`
	SEND                         string `json:"SEND"`
	PUBLIC                       string `json:"PUBLIC"`
	PRIVATE                      string `json:"PRIVATE"`
	TIP                          string `json:"TIP"`
	DATE                         string `json:"DATE"`
	YES                          string `json:"YES"`
	ASC                          string `json:"ASC"`
	DESC                         string `json:"DESC"`
	NO                           string `json:"NO"`
	SHARE                        string `json:"SHARE"`
	PREVIEW                      string `json:"PREVIEW"`
	VRCODE                       string `json:"VRCODE"`
	CHECKALL                     string `json:"CHECK_ALL"`
	CONFIRM                      string `json:"CONFIRM"`
	CANCEL                       string `json:"CANCEL"`
	FORMAT                       string `json:"FORMAT"`
	RESOLUTION                   string `json:"RESOLUTION"`
	PIXEL                        string `json:"PIXEL"`
	PLACE                        string `json:"PLACE"`
	BREADTH                      string `json:"BREADTH"`
	ALTITUDE                     string `json:"ALTITUDE"`
	BITDEPTH                     string `json:"BIT_DEPTH"`
	CREATEDATE                   string `json:"CREATE_DATE"`
	ISENCRYPT                    string `json:"IS_ENCRYPT"`
	PAGENUM                      string `json:"PAGE_NUM"`
	IJOIN                        string `json:"I_JOIN"`
	ICREATE                      string `json:"I_CREATE"`
	MEETING                      string `json:"MEETING"`
	VP                           string `json:"VP"`
	DISCUSS                      string `json:"DISCUSS"`
	WRITECOMMENTS                string `json:"WRITE_COMMENTS"`
	TASK                         string `json:"TASK"`
	COMMENT                      string `json:"COMMENT"`
	DOC                          string `json:"DOC"`
	PAGE                         string `json:"PAGE"`
	COMMENTNOTEMPTY              string `json:"COMMENT_NOT_EMPTY"`
	COMMENTWORDS                 string `json:"COMMENT_WORDS"`
	DETAILINFO                   string `json:"DETAIL_INFO"`
	CONFIRM2                     string `json:"CONFIRM2"`
	FINDPASSWORD                 string `json:"FIND_PASSWORD"`
	THIRDPARTY                   string `json:"THIRD_PARTY"`
	THIRDPARTYAUTHORIZE          string `json:"THIRD_PARTY_AUTHORIZE"`
	THIRDPARTYLOGIN              string `json:"THIRD_PARTY_LOGIN"`
	AUTHORIZE                    string `json:"AUTHORIZE"`
	READANDAGREE                 string `json:"READ_AND_AGREE"`
	AUTHORIZEAGREE               string `json:"AUTHORIZE_AGREE"`
	PLATFORMAGREEMENT            string `json:"PLATFORM_AGREEMENT"`
	WECHAT                       string `json:"WE_CHAT"`
	ACCOUNTLOGIN                 string `json:"ACCOUNT_LOGIN"`
	ACCOUNTPASSW                 string `json:"ACCOUNT_PASSW"`
	SMSLOGIN                     string `json:"SMS_LOGIN"`
	HAVENOACCOUNT                string `json:"HAVE_NO_ACCOUNT"`
	HAVEACCOUNT                  string `json:"HAVE_ACCOUNT"`
	IMMEDIATELYREGISTER          string `json:"IMMEDIATELY_REGISTER"`
	IMMEDIATELYLOGIN             string `json:"IMMEDIATELY_LOGIN"`
	FORGOTPASSWORD               string `json:"FORGOT_PASSWORD"`
	REMEMBERPASSWORD             string `json:"REMEMBER_PASSWORD"`
	SETPASSWORD                  string `json:"SET_PASSWORD"`
	CODELENGTHWRONG              string `json:"CODE_LENGTH_WRONG"`
	ENTERCORRECTCODE             string `json:"ENTER_CORRECT_CODE"`
	GETCODEAGAIN                 string `json:"GET_CODE_AGAIN"`
	RESETPASSWORD                string `json:"RESET_PASSWORD"`
	MOVEFILESUCCESS              string `json:"MOVE_FILE_SUCCESS"`
	MOVEFOLDERSUCCESS            string `json:"MOVE_FOLDER_SUCCESS"`
	CANNOTMOVE                   string `json:"CAN_NOT_MOVE"`
	CHOOSEMEMBER                 string `json:"CHOOSE_MEMBER"`
	MANAGERPERMISSION            string `json:"MANAGER_PERMISSION"`
	PHONEINVITE                  string `json:"PHONE_INVITE"`
	FRIENDINVITE                 string `json:"FRIEND_INVITE"`
	ALLRIGHTRESERVED             string `json:"ALL_RIGHT_RESERVED"`
	CANNOTBEEMPTY                string `json:"CAN_NOT_BE_EMPTY"`
	REINPUTPHONE                 string `json:"RE_INPUT_PHONE"`
	PASSWORDCANNOTBEEMPTY        string `json:"PASSWORD_CAN_NOT_BE_EMPTY"`
	PASSWORDLENGTHLIMIT          string `json:"PASSWORD_LENGTH_LIMIT"`
	PASSWORDNOTMATCH             string `json:"PASSWORD_NOT_MATCH"`
	HAVENOTCHOOSEAGREEMENT       string `json:"HAVE_NOT_CHOOSE_AGREEMENT"`
	INPUTCONFIRMPASSWORD         string `json:"INPUT_CONFIRM_PASSWORD"`
	CODECANNOTBEEMPTY            string `json:"CODE_CAN_NOT_BE_EMPTY"`
	CODELENGTHLIMIT              string `json:"CODE_LENGTH_LIMIT"`
	PHONENUM                     string `json:"PHONE_NUM"`
	TAG                          string `json:"TAG"`
	VALIDITY                     string `json:"VALIDITY"`
	INPUTMAIL                    string `json:"INPUT_MAIL"`
	ADDMAIL                      string `json:"ADD_MAIL"`
	BINDPHONE                    string `json:"BIND_PHONE"`
	INPUTPHONEORNAME             string `json:"INPUT_PHONE_OR_NAME"`
	ADDREMARK                    string `json:"ADD_REMARK"`
	CONTINUEINVITENEWMEMBER      string `json:"CONTINUE_INVITE_NEW_MEMBER"`
	PHONEREGINNCORRECT           string `json:"PHONE_REG_INNCORRECT"`
	CANNOTSHARETOSELF            string `json:"CAN_NOT_SHARE_TO_SELF"`
	PHONEONLYNUM                 string `json:"PHONE_ONLY_NUM"`
	EMAILREGINCORRECT            string `json:"EMAIL_REG_INCORRECT"`
	GETVERIFYCODE                string `json:"GET_VERIFY_CODE"`
	INPUTDYNAMICCODE             string `json:"INPUT_DYNAMIC_CODE"`
	INPUTPHONE                   string `json:"INPUT_PHONE"`
	INPUTPHONE2                  string `json:"INPUT_PHONE2"`
	INPUTPHONEANDREBIND          string `json:"INPUT_PHONE_AND_REBIND"`
	PHONEERRORANDREINPUT         string `json:"PHONE_ERROR_AND_REINPUT"`
	INPUTPASSWORD                string `json:"INPUT_PASSWORD"`
	INPUTCALL                    string `json:"INPUT_CALL"`
	INPUTPASSWORDAGAIN           string `json:"INPUT_PASSWORD_AGAIN"`
	INPUTCURRENTPASSWORD         string `json:"INPUT_CURRENT_PASSWORD"`
	CURRENTPASSWORD              string `json:"CURRENT_PASSWORD"`
	INPUTCODE                    string `json:"INPUT_CODE"`
	INPUTCOMPANY                 string `json:"INPUT_COMPANY"`
	COMPANYNOTEMPTY              string `json:"COMPANY_NOT_EMPTY"`
	NEWPASSWORD                  string `json:"NEW_PASSWORD"`
	INPUTNEWPASSWORD             string `json:"INPUT_NEW_PASSWORD"`
	INPUTNEWPASSWORDAGAIN        string `json:"INPUT_NEW_PASSWORD_AGAIN"`
	MEMBER                       string `json:"MEMBER"`
	CONFIRMPASSWORD              string `json:"CONFIRM_PASSWORD"`
	CONFIRMNEWPASSWORD           string `json:"CONFIRM_NEW_PASSWORD"`
	CONFIRMANDLOGIN              string `json:"CONFIRM_AND_LOGIN"`
	BINDANDLOGIN                 string `json:"BIND_AND_LOGIN"`
	PERSPACE                     string `json:"PER_SPACE"`
	MYPROJECT                    string `json:"MY_PROJECT"`
	ADDEDSERVER                  string `json:"ADDED_SERVER"`
	HELPCENTER                   string `json:"HELP_CENTER"`
	TIME                         string `json:"TIME"`
	NEWBUILD                     string `json:"NEW_BUILD"`
	SETTING                      string `json:"SETTING"`
	DELETE                       string `json:"DELETE"`
	CONTAINER                    string `json:"CONTAINER"`
	CONTAINERNAME                string `json:"CONTAINER_NAME"`
	CONTAINERDESC                string `json:"CONTAINER_DESC"`
	CONTAINERMEMBER              string `json:"CONTAINER_MEMBER"`
	DELETECONTAINER              string `json:"DELETE_CONTAINER"`
	DIRDELETED                   string `json:"DIR_DELETED"`
	DELETECONTAINERCONFIRM       string `json:"DELETE_CONTAINER_CONFIRM"`
	FILTRATE                     string `json:"FILTRATE"`
	SORT                         string `json:"SORT"`
	CLICKORDRAG2UPLOADCOVER      string `json:"CLICK_OR_DRAG2UPLOAD_COVER"`
	COVERSUPPROT                 string `json:"COVER_SUPPROT"`
	CHOOSEFRIEND                 string `json:"CHOOSE_FRIEND"`
	NODATA                       string `json:"NO_DATA"`
	SHOPPINGCART                 string `json:"SHOPPING_CART"`
	BYTE                         string `json:"BYTE"`
	GOBACK                       string `json:"GO_BACK"`
	BASEINFO                     string `json:"BASE_INFO"`
	SECURITYSET                  string `json:"SECURITY_SET"`
	SPACESTATIS                  string `json:"SPACE_STATIS"`
	WORKEXPER                    string `json:"WORK_EXPER"`
	TOTENCENT                    string `json:"TO_TENCENT"`
	COPYBUYINFO                  string `json:"COPY_BUY_INFO"`
	TRANSLATE                    string `json:"TRANSLATE"`
	ASSEMBLE                     string `json:"ASSEMBLE"`
	MODELSIZELIMIT               string `json:"MODEL_SIZE_LIMIT"`
	FAIL                         string `json:"FAIL"`
	SUCCESS                      string `json:"SUCCESS"`
	PROJECT                      string `json:"PROJECT"`
	EMPTY                        string `json:"EMPTY"`
	OPEN                         string `json:"OPEN"`
	INFORMATION                  string `json:"INFORMATION"`
	AVATAR                       string `json:"AVATAR"`
	PURCHASE                     string `json:"PURCHASE"`
	USERSETTING                  string `json:"USER_SETTING"`
	HOMEPAGE                     string `json:"HOMEPAGE"`
	LOGOUT                       string `json:"LOGOUT"`
	COPYSUCCESS                  string `json:"COPY_SUCCESS"`
	COPYFAILED                   string `json:"COPY_FAILED"`
	EDITCOVER                    string `json:"EDIT_COVER"`
	PROJECTCOVER                 string `json:"PROJECT_COVER"`
	CHANGECOVER                  string `json:"CHANGE_COVER"`
	YOURBROWSERNOTEXISTVIDEO     string `json:"YOUR_BROWSER_NOT_EXIST_VIDEO"`
	ISFIRSTFILE                  string `json:"IS_FIRST_FILE"`
	ISLASTFILE                   string `json:"IS_LAST_FILE"`
	VERSIONNO                    string `json:"VERSION_NO"`
	SHARETYPE                    string `json:"SHARE_TYPE"`
	LINKSHARE                    string `json:"LINK_SHARE"`
	SHAREHAVECODE                string `json:"SHARE_HAVE_CODE"`
	SMSSHARE                     string `json:"SMS_SHARE"`
	NOTCREATEVP                  string `json:"NOT_CREATE_VP"`
	VPTIP                        string `json:"VP_TIP"`
	CREATERIGHTNOW               string `json:"CREATE_RIGHT_NOW"`
	SHAREVP                      string `json:"SHARE_VP"`
	VPTIP1                       string `json:"VP_TIP1"`
	VPMODELDELETED               string `json:"VP_MODEL_DELETED"`
	VPNOPERMISSION               string `json:"VP_NO_PERMISSION"`
	VPWANT                       string `json:"VP_WANT"`
	VIEWONLYHAVECODE             string `json:"VIEW_ONLY_HAVE_CODE"`
	SHAREALLOWSAVE               string `json:"SHARE_ALLOW_SAVE"`
	EFFECTIVEFOREVER             string `json:"EFFECTIVE_FOREVER"`
	SHAREFILEORFOLDER            string `json:"SHARE_FILE_OR_FOLDER"`
	COPYLINKSUCCESS              string `json:"COPY_LINK_SUCCESS"`
	COPYLINKFAIL                 string `json:"COPY_LINK_FAIL"`
	COPYLINKANDCODE              string `json:"COPY_LINK_AND_CODE"`
	COPYLINK                     string `json:"COPY_LINK"`
	CHOOSESHAREPERSON            string `json:"CHOOSE_SHARE_PERSON"`
	COOSESHAREMAIL               string `json:"COOSE_SHARE_MAIL"`
	UPDATELINK                   string `json:"UPDATE_LINK"`
	CREATELINK                   string `json:"CREATE_LINK"`
	DOWNLOADVRCODE               string `json:"DOWNLOAD_VRCODE"`
	SOON                         string `json:"SOON"`
	CREATINGSHARE                string `json:"CREATING_SHARE"`
	INPUTEXECUTOR                string `json:"INPUT_EXECUTOR"`
	CURRENTPROJECTASSSIZE        string `json:"CURRENT_PROJECT_ASS_SIZE"`
	CURRENTPROJECTTOTALSIZE      string `json:"CURRENT_PROJECT_TOTAL_SIZE"`
	CHOOSE2MODEL                 string `json:"CHOOSE_2_MODEL"`
	INPUTASSEMBLE                string `json:"INPUT_ASSEMBLE"`
	CANNOTINPUTILLEGALCHAR       string `json:"CAN_NOT_INPUT_ILLEGAL_CHAR"`
	VIEWBIGIMG                   string `json:"VIEW_BIG_IMG"`
	EDITASSEMBLE                 string `json:"EDIT_ASSEMBLE"`
	ASSEMBLEMODEL                string `json:"ASSEMBLE_MODEL"`
	CREATEVP                     string `json:"CREATE_VP"`
	SEARCHVPBYDESC               string `json:"SEARCH_VP_BY_DESC"`
	SEARCHVPBYMAN                string `json:"SEARCH_VP_BY_MAN"`
	VPDESC                       string `json:"VP_DESC"`
	VPDESC1                      string `json:"VP_DESC1"`
	VPDESCNOTEMPTY               string `json:"VP_DESC_NOT_EMPTY"`
	VPDESCLIMIT                  string `json:"VP_DESC_LIMIT"`
	FILENOTSUPPORTVP             string `json:"FILE_NOT_SUPPORT_VP"`
	IMGUPLOADLIMIT               string `json:"IMG_UPLOAD_LIMIT"`
	DELVP                        string `json:"DEL_VP"`
	DELVPCONFIRM                 string `json:"DEL_VP_CONFIRM"`
	MODELASSPCIATE               string `json:"MODEL_ASSPCIATE"`
	GJASSOCIATE                  string `json:"GJ_ASSOCIATE"`
	ADDASSOCIATE                 string `json:"ADD_ASSOCIATE"`
	DELADOC                      string `json:"DEL_A_DOC"`
	DELADOCCONFIRM               string `json:"DEL_A_DOC_CONFIRM"`
	ADDAFTERCLICKMODEL           string `json:"ADD_AFTER_CLICK_MODEL"`
	ADDASSDOC                    string `json:"ADD_ASS_DOC"`
	ASSDOCADDED                  string `json:"ASS_DOC_ADDED"`
	DELETEASSDOC                 string `json:"DELETE_ASS_DOC"`
	ASSDOCDELETED                string `json:"ASSDOC_DELETED"`
	HAVENOPERMISSION             string `json:"HAVE_NO_PERMISSION"`
	LOADINGTRYLATER              string `json:"LOADING_TRY_LATER"`
	PROJECTNAME                  string `json:"PROJECT_NAME"`
	PROJECTDESC                  string `json:"PROJECT_DESC"`
	TAGPLACEHOLDER               string `json:"TAG_PLACEHOLDER"`
	ALLPROJECT                   string `json:"ALL_PROJECT"`
	PROJECTDESCRIBE              string `json:"PROJECT_DESCRIBE"`
	GOBACKINDEX                  string `json:"GOBACK_INDEX"`
	UPLOADINGFILE                string `json:"UPLOADING_FILE"`
	FILEMANAGE                   string `json:"FILE_MANAGE"`
	VPMANAGE                     string `json:"VP_MANAGE"`
	TASKMANAGE                   string `json:"TASK_MANAGE"`
	MEETINGMANAGEL               string `json:"MEETING_MANAGEL"`
	PROJECTSETTING               string `json:"PROJECT_SETTING"`
	PROJECTCANNOTDEL             string `json:"PROJECT_CAN_NOT_DEL"`
	DELPROJECT                   string `json:"DEL_PROJECT"`
	DELPROJECTCONFIRM            string `json:"DEL_PROJECT_CONFIRM"`
	NOTACTIVE                    string `json:"NOT_ACTIVE"`
	INPUTPROJECT1                string `json:"INPUT_PROJECT1"`
	INPUTPROJECTDESC             string `json:"INPUT_PROJECT_DESC"`
	INPUTPROJECTDESCLIMIT        string `json:"INPUT_PROJECT_DESC_LIMIT"`
	INPUTPROJECTTAG              string `json:"INPUT_PROJECT_TAG"`
	INPUTPROJECTTAGLIMIT         string `json:"INPUT_PROJECT_TAG_LIMIT"`
	INPUTPROJECTTAGSAME          string `json:"INPUT_PROJECT_TAG_SAME"`
	PROJECTEXISTED               string `json:"PROJECT_EXISTED"`
	INPUTPROJECT                 string `json:"INPUT_PROJECT"`
	NEWPROJECT                   string `json:"NEW_PROJECT"`
	COSTPROJECT                  string `json:"COST_PROJECT"`
	FREEPROJECT                  string `json:"FREE_PROJECT"`
	CURRENTPROJECT               string `json:"CURRENT_PROJECT"`
	SEARCHNOTFOUND               string `json:"SEARCH_NOT_FOUND"`
	GOBACKLIST                   string `json:"GOBACK_LIST"`
	CANNOTCREATE                 string `json:"CAN_NOT_CREATE"`
	PROJECTINFO                  string `json:"PROJECT_INFO"`
	PROJECTMEMBER                string `json:"PROJECT_MEMBER"`
	PROJECTROLE                  string `json:"PROJECT_ROLE"`
	PROJECTTAG                   string `json:"PROJECT_TAG"`
	DELTASK                      string `json:"DEL_TASK"`
	DELTASKCONFIRM               string `json:"DEL_TASK_CONFIRM"`
	LOCATEVP                     string `json:"LOCATE_VP"`
	REMEMBERVP                   string `json:"REMEMBER_VP"`
	INPUTCOMMENT                 string `json:"INPUT_COMMENT"`
	INPUTCOMMENT2                string `json:"INPUT_COMMENT2"`
	INPUTCOMMENT3                string `json:"INPUT_COMMENT3"`
	ALLPERMISSION                string `json:"ALL_PERMISSION"`
	MODELVIEWPOINT               string `json:"MODEL_VIEWPOINT"`
	DOING                        string `json:"DOING"`
	WILLEXAMER                   string `json:"WILL_EXAMER"`
	PASSED                       string `json:"PASSED"`
	NOTPASSED                    string `json:"NOT_PASSED"`
	FINISHED                     string `json:"FINISHED"`
	CANCELLED                    string `json:"CANCELLED"`
	REOPEN                       string `json:"REOPEN"`
	INPUTKEYWORD                 string `json:"INPUT_KEYWORD"`
	INPUTSEARCHDATA              string `json:"INPUT_SEARCH_DATA"`
	INPUTTASKSTATUS              string `json:"INPUT_TASK_STATUS"`
	EDITTASK                     string `json:"EDIT_TASK"`
	TASKRECORD                   string `json:"TASK_RECORD"`
	PUBLISHTASK                  string `json:"PUBLISH_TASK"`
	VPDETAIL                     string `json:"VP_DETAIL"`
	VIEWASSPCOATE                string `json:"VIEW_ASSPCOATE"`
	SWITCHVERSION                string `json:"SWITCH_VERSION"`
	UPDATEFILE                   string `json:"UPDATE_FILE"`
	SCANPERMISSION               string `json:"SCAN_PERMISSION"`
	DOWNLOADPERMISSION           string `json:"DOWNLOAD_PERMISSION"`
	CREATEPERMISSION             string `json:"CREATE_PERMISSION"`
	EDITPERMISSION               string `json:"EDIT_PERMISSION"`
	DELPERMISSION                string `json:"DEL_PERMISSION"`
	AUTHORIZEPERMISSION          string `json:"AUTHORIZE_PERMISSION"`
	MIXPERMISSION                string `json:"MIX_PERMISSION"`
	FIXPERMISSION                string `json:"FIX_PERMISSION"`
	TAGREPEAT                    string `json:"TAG_REPEAT"`
	TAGINPUT                     string `json:"TAG_INPUT"`
	NEWROLE                      string `json:"NEW_ROLE"`
	DELROLE                      string `json:"DEL_ROLE"`
	DELROLECONFIRM               string `json:"DEL_ROLE_CONFIRM"`
	ROLEHAVENOPERMISSION         string `json:"ROLE_HAVE_NO_PERMISSION"`
	ROLECANNOTBEEMPTY            string `json:"ROLE_CAN_NOT_BE_EMPTY"`
	ROLELENGTHLIMIT              string `json:"ROLE_LENGTH_LIMIT"`
	CHOOSEPERMISSION             string `json:"CHOOSE_PERMISSION"`
	ROLEEXISTED                  string `json:"ROLE_EXISTED"`
	ROLENAME                     string `json:"ROLE_NAME"`
	EDITROLE                     string `json:"EDIT_ROLE"`
	SERVERINFO                   string `json:"SERVER_INFO"`
	SERVERORDER                  string `json:"SERVER_ORDER"`
	ORDERNUM                     string `json:"ORDER_NUM"`
	ORDERCONTENT                 string `json:"ORDER_CONTENT"`
	ORDERDETAILS                 string `json:"ORDER_DETAILS"`
	GENERATEDORDER               string `json:"GENERATED_ORDER"`
	ORDERCOMDATE                 string `json:"ORDER_COM_DATE"`
	ORDERSUCCESS                 string `json:"ORDER_SUCCESS"`
	SERVERSTATUS                 string `json:"SERVER_STATUS"`
	INSERVICE                    string `json:"IN_SERVICE"`
	ENDSERVICE                   string `json:"END_SERVICE"`
	SERVICEDEADLINE              string `json:"SERVICE_DEADLINE"`
	RENDERSERVER                 string `json:"RENDER_SERVER"`
	MODELLIMIT                   string `json:"MODEL_LIMIT"`
	UPLOADMAX                    string `json:"UPLOAD_MAX"`
	TRANSFERMAX                  string `json:"TRANSFER_MAX"`
	ASSEMBLEMAX                  string `json:"ASSEMBLE_MAX"`
	VALIDITYTO                   string `json:"VALIDITY_TO"`
	MEMBEROCCUPY                 string `json:"MEMBER_OCCUPY"`
	SPACEOCCUPY                  string `json:"SPACE_OCCUPY"`
	INVITE                       string `json:"INVITE"`
	INVITEMEMBER                 string `json:"INVITE_MEMBER"`
	DELMEMBER                    string `json:"DEL_MEMBER"`
	DELMEMBERCONFIRM             string `json:"DEL_MEMBER_CONFIRM"`
	INPUTPHONECANNOTBEEMPTY      string `json:"INPUT_PHONE_CAN_NOT_BE_EMPTY"`
	CHOOSEINVITEFRIEND           string `json:"CHOOSE_INVITE_FRIEND"`
	CHANGEROLE                   string `json:"CHANGE_ROLE"`
	CHOOSEROLE                   string `json:"CHOOSE_ROLE"`
	OCCUPYPEOPLE                 string `json:"OCCUPY_PEOPLE"`
	REMAINPEOPLE                 string `json:"REMAIN_PEOPLE"`
	ADDMEMBER                    string `json:"ADD_MEMBER"`
	OPERATE                      string `json:"OPERATE"`
	ASSEMBLENAME                 string `json:"ASSEMBLE_NAME"`
	PROJECCTMEMBERFULL           string `json:"PROJECCT_MEMBER_FULL"`
	TOTALMEMBER                  string `json:"TOTAL_MEMBER"`
	DELVERSION                   string `json:"DEL_VERSION"`
	DELASSEMBLECONFIM            string `json:"DEL_ASSEMBLE_CONFIM"`
	DELDISSOLVEASS               string `json:"DEL_DISSOLVE_ASS"`
	DELAFFECTASS                 string `json:"DEL_AFFECT_ASS"`
	DELFILECONFIM                string `json:"DEL_FILE_CONFIM"`
	DELVERSIONCONFIRM            string `json:"DEL_VERSION_CONFIRM"`
	VERSIONDESC                  string `json:"VERSION_DESC"`
	UPDATEDATE                   string `json:"UPDATE_DATE"`
	UPDATEPEOPLE                 string `json:"UPDATE_PEOPLE"`
	EXECUTEOPERATE               string `json:"EXECUTE_OPERATE"`
	INPUTCHOOSEMEMBER            string `json:"INPUT_CHOOSE_MEMBER"`
	CHOOSEADDMEMBER              string `json:"CHOOSE_ADD_MEMBER"`
	SWITCH                       string `json:"SWITCH"`
	ALLMODEL                     string `json:"ALL_MODEL"`
	CHECKED                      string `json:"CHECKED"`
	MORE                         string `json:"MORE"`
	MYSPACE                      string `json:"MY_SPACE"`
	YOUHAVENOPERMISSION          string `json:"YOU_HAVE_NO_PERMISSION"`
	PERMISSION                   string `json:"PERMISSION"`
	UPLOADFILE                   string `json:"UPLOAD_FILE"`
	CHOOSENEEDUPLOADFILE         string `json:"CHOOSE_NEED_UPLOAD_FILE"`
	CHOOSECONTAINERFILE          string `json:"CHOOSE_CONTAINER_FILE"`
	SEARCHBYFILENAME             string `json:"SEARCHBY_FILENAME"`
	SEARCHBYUPLOADER             string `json:"SEARCHBY_UPLOADER"`
	DEMOTETOFREETOUSE            string `json:"DEMOTE_TO_FREE_TO_USE"`
	TASKCANCELREASON             string `json:"TASK_CANCEL_REASON"`
	CANCELREASON                 string `json:"CANCEL_REASON"`
	CANCELREASONLIMIT            string `json:"CANCEL_REASON_LIMIT"`
	SETPERMISSION                string `json:"SET_PERMISSION"`
	CHOOSECHANGEROLE             string `json:"CHOOSE_CHANGE_ROLE"`
	FOLDERALREADY                string `json:"FOLDER_ALREADY"`
	FOLDERNAMEEXISTED            string `json:"FOLDER_NAME_EXISTED"`
	FOLDERCANNOTBEEMPTY          string `json:"FOLDER_CANNOT_BE_EMPTY"`
	INPUTFOLDERNAME              string `json:"INPUT_FOLDER_NAME"`
	NEWFOLDER                    string `json:"NEW_FOLDER"`
	FILEEXIST                    string `json:"FILE_EXIST"`
	FILESAVED                    string `json:"FILE_SAVED"`
	DELFILEORFOLDER              string `json:"DEL_FILE_OR_FOLDER"`
	DELFILEORFOLDERCONFIRM       string `json:"DEL_FILE_OR_FOLDER_CONFIRM"`
	FOLDERNAME                   string `json:"FOLDER_NAME"`
	PROJECTFILE                  string `json:"PROJECT_FILE"`
	ASSOCIATEDFILE               string `json:"ASSOCIATED_FILE"`
	CREATOR                      string `json:"CREATOR"`
	UPLOADTIME                   string `json:"UPLOAD_TIME"`
	FILESTATUS                   string `json:"FILE_STATUS"`
	TRANSFERFILEAGAIN            string `json:"TRANSFER_FILE_AGAIN"`
	TRANSSUCESS                  string `json:"TRANS_SUCESS"`
	UPLOADSUCCESS                string `json:"UPLOAD_SUCCESS"`
	NOTCLICKWHENTRANSFER         string `json:"NOT_CLICK_WHEN_TRANSFER"`
	HANDLEFAIL                   string `json:"HANDLE_FAIL"`
	MOVEFILE                     string `json:"MOVE_FILE"`
	MOVEFILECONFIRM              string `json:"MOVE_FILE_CONFIRM"`
	FILEDEALANORMAL              string `json:"FILE_DEAL_ANORMAL"`
	FILELOADINGANDBEPATIENT      string `json:"FILE_LOADING_AND_BE_PATIENT"`
	FILEPREVIEW                  string `json:"FILE_PREVIEW"`
	FILECANNOTVIEW               string `json:"FILE_CAN_NOT_VIEW"`
	INPUTFILE                    string `json:"INPUT_FILE"`
	FILENAMEEXISTED              string `json:"FILE_NAME_EXISTED"`
	INPUTFILENAMEEXISTED         string `json:"INPUT_FILE_NAME_EXISTED"`
	FILEISPROCESSING             string `json:"FILE_IS_PROCESSING"`
	FILEISFAILED                 string `json:"FILE_IS_FAILED"`
	FILEISSAVING                 string `json:"FILE_IS_SAVING"`
	NOTSUPPROTFILETYPE           string `json:"NOT_SUPPROT_FILE_TYPE"`
	NOTSUPPROTFILE               string `json:"NOT_SUPPROT_FILE"`
	ONLYUPLOADPIC                string `json:"ONLY_UPLOAD_PIC"`
	UPLOADIMAGEEXCEED            string `json:"UPLOAD_IMAGE_EXCEED"`
	ONLYPICTURES                 string `json:"ONLY_PICTURES"`
	FILELIMIT5                   string `json:"FILE_LIMIT_5"`
	COMPRESSIMGLIMIT             string `json:"COMPRESS_IMG_LIMIT"`
	SAVING                       string `json:"SAVING"`
	NONE                         string `json:"NONE"`
	DEARUSER                     string `json:"DEAR_USER"`
	PROJECTNAMEMIN               string `json:"PROJECT_NAME_MIN"`
	PROJECTNAMEMAX               string `json:"PROJECT_NAME_MAX"`
	ADDNEWMEMBER                 string `json:"ADD_NEW_MEMBER"`
	NEWCONTAINER                 string `json:"NEW_CONTAINER"`
	SETCONTAINER                 string `json:"SET_CONTAINER"`
	MEMBERNAMECANNOTNEEMPTY      string `json:"MEMBER_NAME_CAN_NOT_NE_EMPTY"`
	INPUTCONTAINER               string `json:"INPUT_CONTAINER"`
	MYSPACENOTSUPPORT            string `json:"MYSPACE_NOT_SUPPORT"`
	CONTAINERNAMELIMIT           string `json:"CONTAINER_NAME_LIMIT"`
	CONTAINERDESCLIMIT           string `json:"CONTAINER_DESC_LIMIT"`
	NOTSUPPORTFILESHARE          string `json:"NOT_SUPPORT_FILE_SHARE"`
	FILENOTDEALSUCCESSCANOTSHARE string `json:"FILE_NOT_DEAL_SUCCESS_CANOT_SHARE"`
	CANNOTEDITMULTICONTAINER     string `json:"CAN_NOT_EDIT_MULTI_CONTAINER"`
	CHOOSEDELCONTAINER           string `json:"CHOOSE_DEL_CONTAINER"`
	CHOOSESETTINGCONTAINER       string `json:"CHOOSE_SETTING_CONTAINER"`
	CONTAINERNAMEEXISTED         string `json:"CONTAINER_NAME_EXISTED"`
	CONTAINERNAMENOTEMPTY        string `json:"CONTAINER_NAME_NOT_EMPTY"`
	NOTINVITESAMEMEMBER          string `json:"NOT_INVITE_SAME_MEMBER"`
	NOTACCESSFILE                string `json:"NOT_ACCESS_FILE"`
	NOTACCESSFILE2               string `json:"NOT_ACCESS_FILE2"`
	FREEPROJECTTOPAID            string `json:"FREE_PROJECT_TO_PAID"`
	FREEPROJECTADMIN             string `json:"FREE_PROJECT_ADMIN"`
	CURRENTPROJECTTOPAID         string `json:"CURRENT_PROJECT_TO_PAID"`
	CURRENTPROJECTADMIN          string `json:"CURRENT_PROJECT_ADMIN"`
	UPGRADEPROJECT               string `json:"UPGRADE_PROJECT"`
	ADMINUPGRADEPRO              string `json:"ADMIN_UPGRADE_PRO"`
	PASSWORDLIMIT1               string `json:"PASSWORD_LIMIT1"`
	PASSWORDLIMIT2               string `json:"PASSWORD_LIMIT2"`
	PASSWORDLIMIT3               string `json:"PASSWORD_LIMIT3"`
	INPUTCORRECTPAGENUM          string `json:"INPUT_CORRECT_PAGE_NUM"`
	JUMPTO                       string `json:"JUMP_TO"`
	TO                           string `json:"TO"`
	BACKMYSPACE                  string `json:"BACK_MYSPACE"`
	WEAK                         string `json:"WEAK"`
	STRONG                       string `json:"STRONG"`
	MEDIUM                       string `json:"MEDIUM"`
	BACKLOGIN                    string `json:"BACKLOGIN"`
	ICON                         string `json:"ICON"`
	CHINESE                      string `json:"CHINESE"`
	TRADCHINESE                  string `json:"TRAD_CHINESE"`
	ENGLISH                      string `json:"ENGLISH"`
	FUSE                         string `json:"FUSE"`
	OCCSPACE                     string `json:"OCC_SPACE"`
	REMAINSPACE                  string `json:"REMAIN_SPACE"`
	NODATAGOTO                   string `json:"NODATA_GOTO"`
	TOORDER                      string `json:"TOORDER"`
	MOVE                         string `json:"MOVE"`
	TOUPDATE                     string `json:"TO_UPDATE"`
	HISTORYV                     string `json:"HISTORY_V"`
	SIZE                         string `json:"SIZE"`
	LARGEVIEW                    string `json:"LARGE_VIEW"`
	FILESHAREING                 string `json:"FILE_SHAREING"`
	SEARCHDISCUSSION             string `json:"SEARCH_DISCUSSION"`
	SEARCHEA                     string `json:"SEARCH_EA"`
	NOMEMFOUND                   string `json:"NOMEM_FOUND"`
	ENTERNAME                    string `json:"ENTER_NAME"`
	DELETEORDER                  string `json:"DELETE_ORDER"`
	FAILEDDELORDER               string `json:"FAILED_DEL_ORDER"`
	NOTSTORAGE                   string `json:"NOT_STORAGE"`
	FAILEDSWITCHV                string `json:"FAILED_SWITCHV"`
	ASSEMBLEDFILE                string `json:"ASSEMBLED_FILE"`
	ASSFILEFAILED                string `json:"ASS_FILE_FAILED"`
	ASSDELETED                   string `json:"ASS_DELETED"`
	DELETEDFAILED                string `json:"DELETED_FAILED"`
	ASSMODELEDITED               string `json:"ASS_MODEL_EDITED"`
	SEARCHFAILED                 string `json:"SEARCH_FAILED"`
	FAILEDDELASSDOC              string `json:"FAILED_DEL_ASS_DOC"`
	FILEOPENED                   string `json:"FILE_OPENED"`
	BROWSE                       string `json:"BROWSE"`
	ESTABLISH                    string `json:"ESTABLISH"`
	INSUFUPLOAD                  string `json:"INSUF_UPLOAD"`
	NETWORKNOT                   string `json:"NETWORK_NOT"`
	NETWORKDIS                   string `json:"NETWORK_DIS"`
	NETWORKDOWN                  string `json:"NETWORK_DOWN"`
	PAYMENT                      string `json:"PAYMENT"`
	CANNOTINVITEYOU              string `json:"CANNOT_INVITE_YOU"`
	FAILEDUPDATE                 string `json:"FAILED_UPDATE"`
	FAILEDCREATELINK             string `json:"FAILED_CREATE_LINK"`
	MEMBERDELETED                string `json:"MEMBER_DELETED"`
	FAILEDDELMEM                 string `json:"FAILED_DEL_MEM"`
	CHANGEROLESUC                string `json:"CHANGE_ROLE_SUC"`
	FAILEDCHANGEROLE             string `json:"FAILED_CHANGE_ROLE"`
	SELECTMEMBERIMPORT           string `json:"SELECT_MEMBER_IMPORT"`
	ROLECREATED                  string `json:"ROLE_CREATED"`
	EDITROLESUC                  string `json:"EDIT_ROLE_SUC"`
	DELETEROLE                   string `json:"DELETE_ROLE"`
	FAILEDEDIT                   string `json:"FAILED_EDIT"`
	LABELCREATED                 string `json:"LABEL_CREATED"`
	FAILEDCREATEL                string `json:"FAILED_CREATEL"`
	EDITLABEL                    string `json:"EDIT_LABEL"`
	FAILEDEDITL                  string `json:"FAILED_EDITL"`
	EDITNOTES                    string `json:"EDIT_NOTES"`
	EDITNOTEFAILED               string `json:"EDIT_NOTE_FAILED"`
	MEMINVI                      string `json:"MEM_INVI"`
	MEMEXIST                     string `json:"MEM_EXIST"`
	USEREXIST                    string `json:"USER_EXIST"`
	ROLENOTEXIST                 string `json:"ROLE_NOT_EXIST"`
	MEMADDFAILED                 string `json:"MEM_ADD_FAILED"`
	PROMEMADDFAILED              string `json:"PRO_MEM_ADD_FAILED"`
	MODELNOTLOADED               string `json:"MODEL_NOT_LOADED"`
	NOPERMISCOMMENT              string `json:"NO_PERMIS_COMMENT"`
	MODELLOADING                 string `json:"MODEL_LOADING"`
	VIEWFILEDELETED              string `json:"VIEW_FILE_DELETED"`
	FIRSTFILE                    string `json:"FIRST_FILE"`
	LASTFILE                     string `json:"LAST_FILE"`
	CURRENTITEMNOT               string `json:"CURRENT_ITEM_NOT"`
	WANTEMPTYCART                string `json:"WANT_EMPTY_CART"`
	MULTIPLECART                 string `json:"MULTIPLE_CART"`
	WANTDELETEORDER              string `json:"WANT_DELETE_ORDER"`
	PRODUCT                      string `json:"PRODUCT"`
	UNITPRICE                    string `json:"UNIT_PRICE"`
	NUMBER                       string `json:"NUMBER"`
	TOTALPRICE                   string `json:"TOTAL_PRICE"`
	PLACEORDER                   string `json:"PLACE_ORDER"`
	EMPTYCART                    string `json:"EMPTY_CART"`
	TOTAL1                       string `json:"TOTAL1"`
	SETTLEMENT                   string `json:"SETTLEMENT"`
	PROJECTCREATED               string `json:"PROJECT_CREATED"`
	TENCENTTOPAY                 string `json:"TENCENT_TO_PAY"`
	COPYIDTENCENT                string `json:"COPY_ID_TENCENT"`
	CONFIGURE                    string `json:"CONFIGURE"`
	COLLECTION                   string `json:"COLLECTION"`
	PAIDITEMCREATED              string `json:"PAID_ITEM_CREATED"`
	FAILEDVIDEO                  string `json:"FAILED_VIDEO"`
	ONLYFORMATSUPPORTED          string `json:"ONLY_FORMAT_SUPPORTED"`
	NOMORESIZE                   string `json:"NOMORE_SIZE"`
	EDITAVATAR                   string `json:"EDIT_AVATAR"`
	ONLYCOMORDERDEL              string `json:"ONLY_COM_ORDER_DEL"`
	ORDERNAME                    string `json:"ORDER_NAME"`
	ORDERTIME                    string `json:"ORDER_TIME"`
	ORDERSTATUS                  string `json:"ORDER_STATUS"`
	ALLORDERS                    string `json:"ALL_ORDERS"`
	PERMANENT                    string `json:"PERMANENT"`
	DAYS7                        string `json:"DAYS7"`
	DAYONE                       string `json:"DAY_ONE"`
	RESTORELOCATION              string `json:"RESTORE_LOCATION"`
	DAY                          string `json:"DAY"`
	DAYS                         string `json:"DAYS"`
	HOUR                         string `json:"HOUR"`
	HOURS                        string `json:"HOURS"`
	DOCCON                       string `json:"DOC_CON"`
	FOLDER                       string `json:"FOLDER"`
	SEARCHCON                    string `json:"SEARCH_CON"`
	FILENOTEXIST                 string `json:"FILE_NOT_EXIST"`
	FILEDELETED                  string `json:"FILE_DELETED"`
	CONTACTSHAREAGAIN            string `json:"CONTACT_SHARE_AGAIN"`
	LINKEXPIRED                  string `json:"LINK_EXPIRED"`
	GOTOTENCENT                  string `json:"GOTO_TENCENT"`
	DOWNLOADCLIENT               string `json:"DOWNLOAD_CLIENT"`
	REGISTERVIEW                 string `json:"REGISTER_VIEW"`
	ALSOCONTACSHARE              string `json:"ALSO_CONTAC_SHARE"`
	FILEVIEWONCE                 string `json:"FILE_VIEW_ONCE"`
	WANTVIEWTIMES                string `json:"WANT_VIEW_TIMES"`
	SELECTFILESAVE               string `json:"SELECT_FILE_SAVE"`
	CANCELSHAREING               string `json:"CANCEL_SHAREING"`
	LINKWILLDELETED              string `json:"LINK_WILL_DELETED"`
	REQUESTDATA                  string `json:"REQUEST_DATA"`
	FILENOTPREVIEWED             string `json:"FILE_NOT_PREVIEWED"`
	NOTPREVIEWONLINE             string `json:"NOT_PREVIEW_ONLINE"`
	SAVEPERSPACE                 string `json:"SAVE_PER_SPACE"`
	SHARER                       string `json:"SHARER"`
	MODIFDATE                    string `json:"MODIF_DATE"`
	SURELOGOUT                   string `json:"SURE_LOGOUT"`
	NUMBERNOTCORRECT             string `json:"NUMBER_NOT_CORRECT"`
	INPUTPHONEFIRST              string `json:"INPUT_PHONE_FIRST"`
	ENTERACCOUNT                 string `json:"ENTER_ACCOUNT"`
	INPUTPASSW                   string `json:"INPUT_PASSW"`
	ENTERVERIFCODE               string `json:"ENTER_VERIF_CODE"`
	VERIFICODEINCORRECT          string `json:"VERIFI_CODE_INCORRECT"`
	LENGTHPASSWINCORRECT         string `json:"LENGTH_PASSW_INCORRECT"`
	LENGTHVERIFCODE              string `json:"LENGTH_VERIF_CODE"`
	USERLOGIN                    string `json:"USER_LOGIN"`
	ACCPASSWLOGIN                string `json:"ACC_PASSW_LOGIN"`
	SELECTDIRFOLDER              string `json:"SELECT_DIR_FOLDER"`
	LINKNOTEXIST                 string `json:"LINK_NOT_EXIST"`
	LINKERROR                    string `json:"LINK_ERROR"`
	GETYOURNAMEAVATARSEX         string `json:"GET_YOUR_NAME_AVATAR_SEX"`
	POWERBY                      string `json:"POWER_BY"`
	NOTDONELONGTIME              string `json:"NOT_DONE_LONG_TIME"`
	TOREFRESH                    string `json:"TO_REFRESH"`
	TOKENEXPIRED                 string `json:"TOKEN_EXPIRED"`
	SERVERSUC                    string `json:"SERVER_SUC"`
	NEWMODIFDATA                 string `json:"NEW_MODIF_DATA"`
	REQUESTQUEUE                 string `json:"REQUEST_QUEUE"`
	DATADELETED                  string `json:"DATA_DELETED"`
	ERRORREQUEST                 string `json:"ERROR_REQUEST"`
	USERNOTLOGIN                 string `json:"USER_NOT_LOGIN"`
	USERNOTACCESS                string `json:"USER_NOT_ACCESS"`
	RESOURCEDELETED              string `json:"RESOURCE_DELETED"`
	FORMATREQUEST                string `json:"FORMAT_REQUEST"`
	RESOURCENOLONGER             string `json:"RESOURCE_NO_LONGER"`
	VALIDATIONERROR              string `json:"VALIDATION_ERROR"`
	ERRORCHECKSERVER             string `json:"ERROR_CHECK_SERVER"`
	GATEWAYERROR                 string `json:"GATEWAY_ERROR"`
	SERVERNOTAVAILABLE           string `json:"SERVER_NOT_AVAILABLE"`
	GATEWAYTIMEOUT               string `json:"GATEWAY_TIMEOUT"`
	SHARELINKCOPY                string `json:"SHARE_LINK_COPY"`
	ONLINEPREVIEWNOT             string `json:"ONLINE_PREVIEW_NOT"`
	PROCESSFAILED                string `json:"PROCESS_FAILED"`
	SAVEOPENLATER                string `json:"SAVE_OPEN_LATER"`
	PROCESSOPENLATER             string `json:"PROCESS_OPEN_LATER"`
	BROWSERNOTWEBSOCKET          string `json:"BROWSER_NOT_WEBSOCKET"`
	LONGTIMENOT                  string `json:"LONG_TIME_NOT"`
}
