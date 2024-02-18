package mixed

import "net"

func isValidIp(ip string) bool {
	res := net.ParseIP(ip)
	return res != nil
}
