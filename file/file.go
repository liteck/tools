package tools

import (
	"errors"
	"os"
)

/**
文件类相关
 */

/**
创建文件夹.
*/
func FileCreateFolder(filePath string ) (err error){
	if filePath == ""{
		err = errors.New("filePath can not be nil")
		return
	}
	if _, err = os.Stat(filePath);err!=nil{
		if os.IsNotExist(err) {
			if err = os.Mkdir(filePath, os.FileMode(0777))   ;err!=nil{
				return
			}
		}
	}
	return nil
}
