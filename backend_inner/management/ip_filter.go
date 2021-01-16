package management

import (
	"log"
	"net"
	"net/http"
	"strings"
)

func IpFilter(r *http.Request) bool {
	client_ip := GetIP(r)
	log.Println("get ip: ", client_ip)

	// return BlockAddr(net.ParseIP(client_ip))

	vpn_pool := []string{"10.84.192.0/19", "10.84.160.0/20", "10.84.224.0/19", "10.84.176.0/20"}
	ip := net.ParseIP(client_ip)
	for i := 0; i < len(vpn_pool); i++ {
		_, ipnet, _ := net.ParseCIDR(vpn_pool[i])
		if ipnet.Contains(ip) {
			return true
		}
	}
	return false
}

// GetIP gets a requests IP address by reading off the forwarded-for
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		log.Println("X-FORWARDED-FOR: ", forwarded)
		return forwarded
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}

func BlockAddr(ip net.IP) bool {
	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 && ip4[1] == 84 &&
		(ip4[2] >= 160 && ip4[2] <= 255)
}
