package util

import "net"

// IsCorrectIP is a method
// that checks if a given IP is correct.
func IsCorrectIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
