package json_hlp

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
)

func Parse(data string) (map[string]interface{}, bool) {
	var root map[string]interface{}
	if err := json.Unmarshal([]byte(data), &root); err == nil {
		return root, true
	}

	return root, false
}

func Loadjson(path string, bufSize int) (map[string]interface{}, bool) {
	f, err := os.Open(path)
	if err != nil {
		return nil, false
	}

	data := make([]byte, bufSize)
	total, err2 := f.Read(data)
	if err2 != nil {
		return nil, false
	}

	var root map[string]interface{}
	err = json.Unmarshal(data[:total], &root)
	if err == nil {
		return root, true
	}

	return nil, false
}

func Savejson(path string, data interface{}) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}

	if bytes, err := json.Marshal(data); err == nil {
		if _, err := f.Write(bytes); err == nil {
			return true
		}
	}

	return false
}

func MapToJson(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func MergeMap(mp1 *map[string]interface{}, mp2 map[string]interface{}) {
	for k, v2 := range mp2 {
		if strings.Compare(reflect.TypeOf(v2).String(), "map[string]interface {}") == 0 {
			cmp2 := v2.(map[string]interface{})
			if v1, ok := (*mp1)[k]; ok {
				cmp1 := v1.(map[string]interface{})
				MergeMap(&cmp1, cmp2)
			}
		} else {
			(*mp1)[k] = v2
		}
	}
}
