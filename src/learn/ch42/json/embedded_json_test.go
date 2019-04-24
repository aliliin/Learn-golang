package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

//   利用反射实现，通过 FelidTag 来标识对应的 json 值
var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":20
		},
	"job_info":{
		"skills":["java","Go","C"]
		}
	} `

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}
