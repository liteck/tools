package tools

import (
	"net"
	"testing"

	"github.com/liteck/logs"
)

func Test_GetLocalIp(t *testing.T) {
	if addrs, err := net.InterfaceAddrs(); err != nil {
		logs.Error(err)
	} else {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					t.Log(ipnet.IP.String())
					logs.Debug(ipnet.IP.String())
				}
			}
		}
	}
}
