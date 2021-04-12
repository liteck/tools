package tools

import (
	"fmt"
	"time"
)

func printInfo(v interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
}
