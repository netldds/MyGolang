package signature_sdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	AppID  = "appid"
	AppKey = "appkey"
)

func TestGenerateSignature(t *testing.T) {
	//GET
	request, _ := http.NewRequest("GET", "https://bim.dxbim.com:8443/backend/api/v3/project?project_id=e9fd207040ef41caa6ba0fc4c766d29b", nil)
	GenerateSignature(request, "bim.dxbim.com", AppID, AppKey)
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		t.Errorf("%s%v", b, err)
	}
	request, resp, err = nil, nil, nil

	//POST
	body := struct {
		ProjectId string `json:"project_id"`
	}{ProjectId: "xxx"}
	b, err := json.Marshal(body)
	request, _ = http.NewRequest("POST", "http://localhost:8090/api/v1/checksum", bytes.NewBuffer(b))
	GenerateSignature(request, "bim.dxbim.com", AppID, AppKey)
	client = http.Client{}
	resp, err = client.Do(request)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		t.Errorf("%s%v", b, err)
	}
}
