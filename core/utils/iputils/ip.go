package iputils

import (
	"encoding/json"
	"github.com/bitxx/logger/logbase"
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"net/http"
	"strings"
)

// GetLocation 获取外网ip地址
func GetLocation(ip, key string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "inner ip"
	}
	url := "https://restapi.amap.com/v5/ip?ip=" + ip + "&type=4&key=" + key
	resp, err := http.Get(url)
	if err != nil {
		logbase.Errorf("restapi.amap.com failed: %s", err)
		return "unknown ip"
	}
	defer resp.Body.Close()
	s, err := io.ReadAll(resp.Body)

	m := make(map[string]string)

	err = json.Unmarshal(s, &m)
	if err != nil {
		logbase.Errorf("Umarshal failed: %s", err)
	}
	//if m["province"] == "" {
	//	return "未知位置"
	//}
	return m["country"] + "-" + m["province"] + "-" + m["city"] + "-" + m["district"] + "-" + m["isp"]
}

// GetLocaHonst 获取局域网ip地址
func GetLocaHonst() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		logbase.Errorf("net.Interfaces failed, err: %s", err)
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}

	}
	return ""
}

func GetClientIP(c *gin.Context) string {
	ClientIP := c.ClientIP()
	//fmt.Println("ClientIP:", ClientIP)
	RemoteIP := c.RemoteIP()
	//fmt.Println("RemoteIP:", RemoteIP)
	ip := c.Request.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = c.Request.Header.Get("X-real-ip")
	}
	if ip == "" {
		ip = "127.0.0.1"
	}
	if RemoteIP != "127.0.0.1" {
		ip = RemoteIP
	}
	if ClientIP != "127.0.0.1" {
		ip = ClientIP
	}
	return ip
}
