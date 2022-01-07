package kata

import (
	"net"
	"regexp"
)

func Is_valid_ip_easy_way(ip string) bool {
	if r := net.ParseIP(ip); r == nil {
		return false
	}
	return true
}

func Is_valid_ip_hard_way(ip string) bool {
	re := regexp.MustCompile(`^((\d|[1-9]\d|1\d{2}|2[0-4]\d|25[0-5])\.){3}(\d|[1-9]\d|(1\d{2}|2[0-4]\d|25[0-5]))$`)
	return re.Match([]byte(ip))
}
