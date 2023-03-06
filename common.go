package BitrixGo

import (
	"fmt"
	"net/http"
	"reflect"
)

func AddParam(req *http.Request, param string, value string) {
	q := req.URL.Query()
	q.Add(param, value)
	req.URL.RawQuery = q.Encode()
}

func AddParamsFromStruct(req *http.Request, params interface{}) {
	t := reflect.TypeOf(params)
	v := reflect.ValueOf(params)
	q := req.URL.Query()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Tag.Get("bitrix")
		value := v.Field(i)
		if value.Interface() == "" || value.Interface() == nil || value.Interface() == 0 {
			continue
		}
		q.Add(field, fmt.Sprintf("%v", value))
	}
	req.URL.RawQuery = q.Encode()
}
