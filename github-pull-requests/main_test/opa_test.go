package main_test

import (
	"encoding/json"
	"fmt"
	"github.com/yalp/jsonpath"
	"testing"
)

func TestJsonPath(t *testing.T) {
	var body = []byte(`{"html_url":456,"x":{"nested":{"html_url":111,"very_nested":{"html_url":222}}}}`)
	var da interface{}
	_ = json.Unmarshal(body, &da)
	authors, _ := jsonpath.Read(da, "$..html_url")
	authors = Map(authors.([]interface{}), func(item interface{}) interface{} {
		return item.(float64)
	})
	fmt.Println("b4 float",authors)
	authors = Map(authors.([]interface{}), func(i interface{}) interface{} {
		return i.(float64) + 1
	})
	fmt.Println("plus one",authors)
}

func Map(collection []interface{}, action func(interface{}) interface{}) []interface{} {
	var res = make([]interface{},len(collection));
	for idx,item := range collection {
		res[idx] = action(item)
	}
	return res
}

