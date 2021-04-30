package design_principles

import (
	"net"
	"regexp"
)

func IsValidIP(ip string) bool {
	if ip == "" {
		return false
	}

	validIP := regexp.MustCompile("^(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|[1-9])\\." +
		"(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)\\." +
		"(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)\\." +
		"(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)$")

	return validIP.MatchString(ip)
}

func CheckIfIPValid(ip string) bool {
	if ip == "" {
		return false
	}

	return net.ParseIP(ip) != nil
}
