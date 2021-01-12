package management

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func IpFilter(r *http.Request) bool {
	client_ip := GetIP(r)
	fmt.Println("get ip: ", client_ip)

	return BlockAddr(net.ParseIP(client_ip))

	// vpn_pool := []string{"10.84.192.0/19", "10.84.160.0/20", "10.84.224.0/19", "10.84.176.0/20"}
	// for _, network := range vpn_pool {
	// 	if network.Contains(net.ParseIP(client_ip)) {
	// 		return true
	// 	}
	// }
	// return false
}

// GetIP gets a requests IP address by reading off the forwarded-for
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		fmt.Println("X-FORWARDED-FOR ", forwarded)
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
		(ip4[2] == 192 || ip4[2] == 160 || ip4[2] == 224 || ip4[2] == 176)
}
