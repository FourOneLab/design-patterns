package design_principles

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

// IsValidIPAddressV1 使用正则表达式
func IsValidIPAddressV1(ip string) bool {
	if ip == "" {
		return false
	}

	validIP := regexp.MustCompile(
		"^(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|[1-9])\\." +
			"(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)\\." +
			"(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)\\." +
			"(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)$")

	return validIP.MatchString(ip)
}

// IsValidIPAddressV2 使用 net 库中的 ParseIP 函数
func IsValidIPAddressV2(ip string) bool {
	if ip == "" {
		return false
	}

	parseIP := net.ParseIP(ip)
	if parseIP == nil {
		return false
	}

	return true
}

// IsValidIPAddressV3 手写，只实现了解析以点分割的 IPv4
func IsValidIPAddressV3(ip string) bool {
	if ip == "" {
		return false
	}

	split := strings.Split(ip, ".")
	for i, s := range split {
		atoi, err := strconv.Atoi(s)
		if err != nil {
			return false
		}

		if i == 0 {
			if atoi <= 0 || atoi > 255 {
				return false
			}
		} else {
			if atoi < 0 || atoi > 255 {
				return false
			}
		}
	}

	return true
}
