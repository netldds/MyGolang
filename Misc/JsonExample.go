package Misc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type response1 struct {
	Name string
	Age  int
	Time time.Time
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
		//panic(e)
	}
}
func ExampleJson() {
	//The JSON package can automatically encode your custom data types.It will only include export fields in the encode output
	stru1 := response1{Name: "zl", Age: 999, Time: time.Now()}
	databytes, err := json.Marshal(stru1)
	checkError(err)
	fi, err := os.OpenFile("data.json", os.O_CREATE|os.O_RDWR, 0666)
	defer fi.Close()
	fmt.Println(fi.Write(databytes))

	var dat map[string]interface{}

	json.Unmarshal(databytes, &dat)
	AgeFloat := dat["Age"].(float64)
	fmt.Println("Age:", AgeFloat)

	NameStr := dat["Name"]
	fmt.Println(NameStr)

}

type Thorax struct {
	Name string `json:"name"`
}

type Soma struct {
	Data   interface{} `json:"data"`
	Length int         `json:"length"`
}

func EscapeJson() {
	//"{\"data\":\"{\\\"name\\\":\\\"thorax_one\\\"}\",\"length\":180}"
	thorax := Thorax{
		Name: "thorax_one",
	}
	t, _ := json.Marshal(thorax)
	body := Soma{
		Data:   string(t),
		Length: 180,
	}
	b, _ := json.Marshal(body)
	b, _ = json.Marshal(string(b))
	fmt.Println(string(b))
}
