package common

import (
	"encoding/json"
	"fmt"
	"github.com/yalp/jsonpath"
)

func ParseToSplitStr(body []byte, jsonPath string) []string {
	parsed := parseResp(body, jsonPath)
	return toSplitStr(parsed)
}


func ParseToSplitInt(body []byte, jsonPath string) []int {
	parsed := parseResp(body, jsonPath)
	return toSplitInt(parsed)
}

func ParseToSting(body []byte, jsonPath string) string {
	parsed := parseResp(body, jsonPath)
	return parsed.(string)
}


func parseResp(body []byte, jsonPath string) interface{} {
	var response interface{}
	err := json.Unmarshal(body, &response)
	PanicIfNonEmpty(err)
	parsed, err := jsonpath.Read(response, jsonPath)
	fmt.Println(response)
	if err != nil {
		fmt.Println(response)
		panic(err)
	}
	return parsed
}



func toSplitStr(parsed interface{}) []string {
	result := []string{}
	for _, v := range parsed.([]interface{}) {
		result = append(result, v.(string))
	}
	return result
}

func toSplitInt(parsed interface{}) []int {
	result := []int{}
	for _, v := range parsed.([]interface{}) {
		result = append(result, int(v.(float64)))
	}
	return result
}

