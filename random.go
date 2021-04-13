package tools

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func RandomString(num int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(int(math.Pow10(num)))
	nStr := fmt.Sprintf("%d", n)
	if len(nStr) < num {
		for i := 0; i < num-len(nStr); i++ {
			nStr += "0"
		}
	}
	return nStr
}
