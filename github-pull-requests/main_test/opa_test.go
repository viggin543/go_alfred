package main_test

import (
	"encoding/json"
	"fmt"
	"github.com/yalp/jsonpath"
	"io/ioutil"
	"testing"
)

func TestJsonPath(t *testing.T) {
	content,_ :=ioutil.ReadFile("gtihub_response.json")
	fmt.Println(content)
	if len(content) == 0 {
		t.Fail()
	}



	var da interface{}
	_ = json.Unmarshal([]byte("body"), &da)
	authors, _ := jsonpath.Read(da, "$..[url,login]")
	fmt.Println(authors)

}

func Map(collection []interface{}, action func(interface{}) interface{}) []interface{} {
	var res = make([]interface{},len(collection));
	for idx,item := range collection {
		res[idx] = action(item)
	}
	return res
}

