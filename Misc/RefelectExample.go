package Misc

import (
	"fmt"
	"reflect"
)

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

//insert into employee values("Naveen", 565, "Coimbatore", 90000, "India")
func createQuery(q interface{}) {

	q = employee{name: "zl",
		id:      5,
		address: "tianhong",
		salary:  18000,
		country: "CHN"}
	if reflect.TypeOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		query := fmt.Sprintf("Insert into %s values(", reflect.TypeOf(q).Name())
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s,%d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s,\"%s\"", query, v.Field(i).String())
				}
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
	}
}
func reflect_example() {
	method := func(i string) map[string]string {
		fmt.Println(i)
		return map[string]string{"id": "id1"}
	}
	var v reflect.Value

	v = reflect.ValueOf(method)
	vs := []reflect.Value{reflect.ValueOf("myStrings")}
	rvs := v.Call(vs)
	res := rvs[0].MapIndex(reflect.ValueOf("id")).String()
	fmt.Println(res)
}