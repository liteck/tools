package tools

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"reflect"
	"strings"

	"code.google.com/p/mahonia"
)

/**
utf8转换为 gbk
**/
func ConvertUTF2GBK(utf string) (gbk string) {
	gbk_enc := mahonia.NewEncoder("GBK")
	return gbk_enc.ConvertString(utf)
}

/**
gbk转换为 utf8
**/
func ConvertGBK2UTF(gbk string) (utf string) {
	enc := mahonia.NewDecoder("GBK")
	return enc.ConvertString(gbk)
}

func MD5(data string) string {
	m := md5.New()
	io.WriteString(m, data)
	return hex.EncodeToString(m.Sum(nil))
}

/**
把json 的struct 转换为 map
*/
func JsonToMap(a interface{}) map[string]interface{} {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Name
		value := v.Field(i).Interface()
		tag := t.Field(i).Tag.Get("json")
		if tag != "" {
			if strings.Contains(tag, ",") {
				ps := strings.Split(tag, ",")
				key = ps[0]
			} else {
				key = tag
			}
		}
		data[key] = value
	}
	return data
}
