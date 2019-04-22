/*
1.搞定OAUTH2.0
2.搞定weibo 3rd application
3.部署application
4.重构支持公开版本
5.上weibo SDK
*/
package main

import (
	"flag"
	"fmt"
	"github.com/xiocode/weigo"
	"reflect"
)

func TT() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	test()
}

func test() {
	api := weigo.NewAPIClient("3417104247", "f318153f6a80329f06c1d20842ee6e91", "http://127.0.0.1/callback", "code")
	authorize_url, err := api.GetAuthorizeUrl(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(authorize_url)
	var result map[string]interface{}
	err=api.RequestAccessToken("1fdaa295b73d2a9568e284383ced5e9e",&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	access_token := result["access_token"]
	fmt.Println(reflect.TypeOf(access_token), access_token)
	expires_in := result["expires_in"]
	fmt.Println(reflect.TypeOf(expires_in), expires_in)
}
