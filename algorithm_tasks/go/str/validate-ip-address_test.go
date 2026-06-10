//

package str

import (
	"strconv"
	"strings"
	"testing"
)

func validIPAddress(queryIP string) string {
	nums := strings.Split(queryIP, ":")

	isIPv6 := len(nums) > 1
	var (
		res       string
		checkFunc func(nums string) bool
		maxOctets = 8
	)

	if isIPv6 {
		checkFunc = isValidIPv6
		res = "IPv6"
	} else {
		checkFunc = isValidIPv4
		res = "IPv4"
		maxOctets = 4
		nums = strings.Split(queryIP, ".")
	}
	if len(nums) != maxOctets {
		return "Neither"
	}

	for _, num := range nums {

		if !checkFunc(num) {
			return "Neither"
		}
	}
	return res

}

func isValidIPv4(num string) bool {
	i, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return false
	}

	if num == "" ||
		num[0] == '0' && len(num) > 1 ||
		0 > i || i > 255 {
		return false
	}
	return true
}

func isValidIPv6(num string) bool {
	i, err := strconv.ParseInt(num, 16, 64)
	if err != nil {
		return false
	}

	if num == "" ||
		len(num) > 4 ||
		0 <= i && i > 65535 {
		return false
	}
	return true
}

func TestValidIPAddress(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// --- IPv4 valid ---
		{"ipv4 basic", "172.16.254.1", "IPv4"},
		{"ipv4 zeros", "0.0.0.0", "IPv4"},
		{"ipv4 max", "255.255.255.255", "IPv4"},
		{"ipv4 mixed", "192.168.1.0", "IPv4"},

		// --- IPv4 invalid ---
		{"ipv4 basic", "12.12.12.12.12", "Neither"},
		{"ipv4 leading zero", "192.168.01.1", "Neither"},
		{"ipv4 leading double zero", "192.168.1.00", "Neither"},
		{"ipv4 out of range", "256.256.256.256", "Neither"},
		{"ipv4 too few parts", "1.1.1", "Neither"},
		{"ipv4 too many parts", "1.1.1.1.1", "Neither"},
		{"ipv4 empty part", "1.1.1.", "Neither"},
		{"ipv4 leading dot", ".1.1.1", "Neither"},
		{"ipv4 letters", "192.168.@.1", "Neither"},
		{"ipv4 negative", "192.168.-1.1", "Neither"},

		// --- IPv6 valid ---
		{"ipv6 full", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "IPv6"},
		{"ipv6 full with 0", "2001:0db8:85a3:0:0:8A2E:0370:7334", "IPv6"},
		{"ipv6 short segments", "2001:db8:85a3:0:0:8A2E:0370:7334", "IPv6"},
		{"ipv6 mixed case", "2001:0db8:85a3:0000:0000:8A2E:0370:7334", "IPv6"},
		{"ipv6 min segment length", "0:0:0:0:0:0:0:0", "IPv6"},
		{"ipv6 uppercase", "FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF", "IPv6"},

		// --- IPv6 invalid ---
		{"ipv6 double colon", "2001:0db8:85a3::8A2E:0370:7334", "Neither"},
		{"ipv6 invalid char", "2001:0db8:85a3:0000:0000:8a2e:037j:7334", "Neither"},
		{"ipv6 too long segment", "02001:0db8:85a3:0000:0000:8a2e:0370:7334", "Neither"},
		{"ipv6 too few parts", "2001:0db8:85a3:0000:0000:8a2e:0370", "Neither"},
		{"ipv6 too many parts", "2001:0db8:85a3:0000:0000:8a2e:0370:7334:1234", "Neither"},
		{"ipv6 empty segment", "2001:0db8:85a3::0000:8a2e:0370:7334", "Neither"},

		// --- Neither ---
		{"empty string", "", "Neither"},
		{"just dots", "...", "Neither"},
		{"just colons", ":::", "Neither"},
		{"random string", "foobar", "Neither"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validIPAddress(tt.input)
			if got != tt.expected {
				t.Errorf("validIPAddress(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
