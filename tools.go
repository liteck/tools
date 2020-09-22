package tools

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"reflect"
	"strconv"
	"strings"
)

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

func GetLocalIp() string {
	if addrs, err := net.InterfaceAddrs(); err != nil {
		return "127.0.0.1"
	} else {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}
	return "127.0.0.1"
}

func ConvertYTOF(yuan string) (fen string, err error) {
	if err = CheckPriceFormat(yuan); err != nil {
		return "", err
	}
	if f_yuan, err := strconv.ParseFloat(yuan, 32); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("%.f", f_yuan*100), nil
	}
}

func ConvertFTOY(fen string) (yuan string, err error) {
	if f_fen, err := strconv.ParseFloat(fen, 32); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("%.02f", f_fen/100), nil
	}
}

func CheckPriceFormat(fee string) error {
	if len(fee) == 0 {
		return errors.New("金额不能为空")
	}
	if strings.Contains(fee, ".") {
		arr := strings.Split(fee, ".")
		if len(arr[1]) <= 2 {
			return nil
		} else {
			return errors.New("金额只支持两位小数")
		}
	} else {
		return nil
	}
}
