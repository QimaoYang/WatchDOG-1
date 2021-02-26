package management

import (
	"log"
	"net"
	"net/http"
	"strings"
)

func IpFilter(r *http.Request) (bool, string) {
	client_ip := GetIP(r)

	ipFilter := false

	// black list
	// vpn_pool := []string{"10.84.192.0/19", "10.84.160.0/20", "10.84.224.0/19", "10.84.176.0/20", "10.135.128.0/18", "10.134.0.0/16"}
	// ip := net.ParseIP(client_ip)
	// for i := 0; i < len(vpn_pool); i++ {
	// 	_, ipnet, _ := net.ParseCIDR(vpn_pool[i])
	// 	if ipnet.Contains(ip) {
	// 		ipFilter = false
	// 	}
	// }

	// white list
	ip4 := net.ParseIP(client_ip).To4()
	if ip4 != nil && ip4[0] == 10 && ip4[1] == 84 &&
		(ip4[2] == 84 || ip4[2] == 85 || ip4[2] == 86 || ip4[2] == 87 || ip4[2] == 97 || ip4[2] == 98 || ip4[2] == 111) {
		ipFilter = true
	}
	return ipFilter, client_ip
}

// GetIP gets a requests IP address by reading off the forwarded-for
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		log.Println("get X-FORWARDED-FOR: ", forwarded)
		return forwarded
	}

	realIp := r.Header.Get("X-Real-IP")
	if realIp != "" {
		log.Println("get X-Real-IP: ", realIp)
		return realIp
	}

	remoteAddr := strings.Split(r.RemoteAddr, ":")[0]
	log.Println("get RemoteAddr: ", remoteAddr)
	return remoteAddr
}
