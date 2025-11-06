package util

import (
	"net"
	"net/http"
	"strings"
)

// GetClientIP extracts the client IP address from the http.Request.
func GetClientIP(r *http.Request) string {
	// Check the X-Forwarded-For header (used by proxies)
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// X-Forwarded-For can contain multiple IPs, split by commas.
		ips := strings.Split(xForwardedFor, ",")
		// Return the first IP in the list (the original client IP).
		return strings.TrimSpace(ips[0])
	}

	// Check the X-Real-IP header (used by some proxies)
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	// Fallback to RemoteAddr (direct client connection)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // Return as is if unable to split
	}
	return ip
}
