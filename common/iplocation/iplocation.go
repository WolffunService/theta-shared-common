package iplocation

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func ReadUserIp(r *http.Request) (string, error) {

	//Get IP from the X-Real-IP header
	ip := r.Header.Get("X-Real-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from the cf-connecting-ip header
	ip = r.Header.Get("cf-connecting-ip")
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-Forwarded-For header
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP = net.ParseIP(ip)
		if netIP != nil {
			if isPublicIP(netIP) {
				return ip, nil
			}
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		if isPublicIP(netIP) {
			return ip, nil
		}
	}

	return "", fmt.Errorf("[ReadUserIp] IP not found")
}

func isPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return true
}
