package tools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"hash/crc32"
	"io"
	"strings"
	"time"
)

var gmtLoc = time.FixedZone("GMT", 0)

func GenUUID() string {
	return strings.ToUpper(strings.Replace(uuid.Must(uuid.NewV4(), nil).String(), "-", "", -1))
}

func MD4(data string) string {
	m := md5.New()
	_, _ = io.WriteString(m, data)
	x := hex.EncodeToString(m.Sum(nil))
	return strings.ToUpper(x)
}

func String(v interface{}) string {
	x, _ := json.Marshal(v)
	return string(x)
}

func GetNow() string {
	return time.Now().In(gmtLoc).Format("2006-01-02 15:04:05")
}

func GenNo(prefix, suffix string) string {
	//这里根据自身业务生成单号
	orderNo := prefix + time.Now().Format("20060102150405")
	// 3.生成uuid的hashCode值
	hashCodeV := int(crc32.ChecksumIEEE([]byte(uuid.Must(uuid.NewV4(), nil).String())))
	// 4.可能为负数
	if hashCodeV < 0 {
		hashCodeV = -hashCodeV
	}
	orderNo += fmt.Sprintf("%s%010d", suffix, hashCodeV)
	return orderNo
}
