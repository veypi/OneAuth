package tools

import (
	"encoding/json"
	"github.com/veypi/OneBD"
	"net/http"
	"net/url"
	"reflect"
)

type Initer interface {
	Init(OneBD.Meta) error
}

func MultiIniter(m OneBD.Meta, is ...Initer) (err error) {
	for _, i := range is {
		err = i.Init(m)
		if err != nil {
			return
		}
	}
	return
}

func Query(addr string, query map[string]string, res interface{}) error {
	u, err := url.Parse(addr)
	if err != nil {
		return err
	}
	paras := &url.Values{}
	//设置请求参数
	for k, v := range query {
		paras.Set(k, v)
	}
	u.RawQuery = paras.Encode()
	resp, err := http.Get(u.String())
	//关闭资源
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	return json.NewDecoder(resp.Body).Decode(res)
}

func Struct2Map(obj interface{}) (data map[string]interface{}) {
	data = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	var item reflect.Value
	var k reflect.StructField
	for i := 0; i < objT.NumField(); i++ {
		k = objT.Field(i)
		item = objV.Field(i)
		if !item.IsNil() {
			data[k.Name] = item.Interface()
		}
	}
	return
}
