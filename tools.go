package tools

import (
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
