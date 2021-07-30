package utils

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"net"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsInRange returns true if value lies between left and right border
func IsInRange(value, left, right float64) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// IsEmail is a constraint to do a simple validation for email addresses, it only check if the string contains "@"
// and that it is not in the first or last character of the string
// https://en.wikipedia.org/wiki/Email_address#Valid_email_addresses
func IsEmail(s string) bool {
	if !strings.Contains(s, "@") || s[0] == '@' || s[len(s)-1] == '@' {
		return false
	}
	return true
}

// IsURL check if the string is an URL.
func IsURL(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	return rxURL.MatchString(str)
}

// IsRequestURL check if the string rawurl, assuming
// it was received in an HTTP request, is a valid
// URL confirm to RFC 3986
func IsRequestURL(rawurl string) bool {
	url, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false //Couldn't even parse the rawurl
	}
	if len(url.Scheme) == 0 {
		return false //No Scheme found
	}
	return true
}

// IsRequestURI check if the string rawurl, assuming
// it was received in an HTTP request, is an
// absolute URI or an absolute path.
func IsRequestURI(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	return err == nil
}

// IsAlpha check if the string contains only letters (a-zA-Z). Empty string is valid.
func IsAlpha(s string) bool {
	for _, v := range s {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') {
			return false
		}
	}
	return true
}

// IsUTFLetter check if the string contains only unicode letter characters.
// Similar to IsAlpha but for all languages. Empty string is valid.
func IsUTFLetter(str string) bool {
	for _, v := range str {
		if !unicode.IsLetter(v) {
			return false
		}
	}
	return true

}

// IsAlphanumeric check if the string contains only letters and numbers. Empty string is valid.
func IsAlphanumeric(s string) bool {
	for _, v := range s {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') && ('9' < v || v < '0') {
			return false
		}
	}
	return true
}

// IsUTFLetterNumeric check if the string contains only unicode letters and numbers. Empty string is valid.
func IsUTFLetterNumeric(s string) bool {
	for _, v := range s {
		if !unicode.IsLetter(v) && !unicode.IsNumber(v) { //letters && numbers are ok
			return false
		}
	}
	return true
}

// IsNumeric check if the string contains only numbers. Empty string is valid.
func IsNumeric(s string) bool {
	for _, v := range s {
		if '9' < v || v < '0' {
			return false
		}
	}
	return true
}

// IsUTFNumeric check if the string contains only unicode numbers of any kind.
// Numbers can be 0-9 but also Fractions ¾,Roman Ⅸ and Hangzhou 〩. Empty string is valid.
func IsUTFNumeric(s string) bool {
	for _, v := range s {
		if !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}

// IsWhole returns true if value is whole number
func IsWhole(value float64) bool {
	return math.Abs(math.Remainder(value, 1)) == 0
}

// IsNatural returns true if value is natural number (positive and whole)
func IsNatural(value float64) bool {
	return IsWhole(value) && value > 0
}

// IsUTFDigit check if the string contains only unicode radix-10 decimal digits. Empty string is valid.
func IsUTFDigit(s string) bool {
	for _, v := range s {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}

// IsHexadecimal check if the string is a hexadecimal number.
func IsHexadecimal(str string) bool {
	_, err := strconv.ParseInt(str, 16, 0)
	return err == nil
}

// IsLowerCase check if the string is lowercase. Empty string is valid.
func IsLowerCase(str string) bool {
	if len(str) == 0 {
		return true
	}
	return str == strings.ToLower(str)
}

// IsUpperCase check if the string is uppercase. Empty string is valid.
func IsUpperCase(str string) bool {
	if len(str) == 0 {
		return true
	}
	return str == strings.ToUpper(str)
}

// IsInt check if the string is an integer. Empty string is valid.
func IsInt(str string) bool {
	if len(str) == 0 {
		return true
	}
	_, err := strconv.Atoi(str)

	return err == nil
}

// IsFloat check if the string is a float.
func IsFloat(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

// IsByteLength check if the string's length (in bytes) falls in a range.
func IsByteLength(str string, min, max int) bool {
	return len(str) >= min && len(str) <= max
}

// CreditCard check if the string is a credit card.
func CreditCard(str string) bool {
	r, _ := regexp.Compile("[^0-9]+")
	sanitized := r.ReplaceAll([]byte(str), []byte(""))
	if !rxCreditCard.MatchString(string(sanitized)) {
		return false
	}
	var sum int64
	var digit string
	var tmpNum int64
	var shouldDouble bool
	for i := len(sanitized) - 1; i >= 0; i-- {
		digit = string(sanitized[i:(i + 1)])
		tmpNum, _ = ToInt64(digit)
		if shouldDouble {
			tmpNum *= 2
			if tmpNum >= 10 {
				sum += ((tmpNum % 10) + 1)
			} else {
				sum += tmpNum
			}
		} else {
			sum += tmpNum
		}
		shouldDouble = !shouldDouble
	}

	return sum%10 == 0
}

// IsJSON check if the string is valid JSON (note: uses json.Unmarshal).
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsMultibyte check if the string contains one or more multibyte chars. Empty string is valid.
func IsMultibyte(s string) bool {
	for _, v := range s {
		if v >= utf8.RuneSelf {
			return true
		}
	}

	return len(s) == 0
}

// IsASCII check if the string contains ASCII chars only. Empty string is valid.
func IsASCII(s string) bool {
	for _, v := range s {
		if v >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

// PrintableASCII check if the string contains printable ASCII chars only. Empty string is valid.
func PrintableASCII(s string) bool {
	for _, v := range s {
		if v < ' ' || v > '~' {
			return false
		}
	}
	return true
}

// IsBase64 check if a string is base64 encoded.
func IsBase64(s string) bool {
	if len(s) == 0 {
		return false
	}
	_, err := base64.StdEncoding.DecodeString(s)

	return err == nil
}

// IsFilePath check is a string is Win or Unix file path and returns it's type.
func IsFilePath(str string) (bool, int) {
	if rxWinPath.MatchString(str) {
		// check windows path limit see:
		// http://msdn.microsoft.com/en-us/library/aa365247(VS.85).aspx#maxpath
		if len(str[3:]) > 32767 {
			return false, Win
		}
		return true, Win
	} else if rxUnixPath.MatchString(str) {
		return true, Unix
	}
	return false, Unknown
}

// IsDataURI checks if a string is base64 encoded data URI such as an image
func IsDataURI(str string) bool {
	dataURI := strings.Split(str, ",")
	if !rxDataURI.MatchString(dataURI[0]) {
		return false
	}
	return IsBase64(dataURI[1])
}

// IsDNSName will validate the given string as a DNS name
func IsDNSName(str string) bool {
	if str == "" || len(strings.Replace(str, ".", "", -1)) > 255 {
		// constraints already violated
		return false
	}
	return rxDNSName.MatchString(str)
}

// IsDialString validates the given string for usage with the various Dial() functions
func IsDialString(str string) bool {
	if h, p, err := net.SplitHostPort(str); err == nil && h != "" && p != "" && (IsDNSName(h) || IsIP(h)) && IsPort(p) {
		return true
	}

	return false
}

// IsIP checks if a string is either IP version 4 or 6.
func IsIP(str string) bool {
	return net.ParseIP(str) != nil
}

// IsPort checks if a string represents a valid port
func IsPort(str string) bool {
	if i, err := strconv.Atoi(str); err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ".")
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ":")
}

// IsMAC check if a string is valid MAC address.
// Possible MAC formats:
// 01:23:45:67:89:ab
// 01:23:45:67:89:ab:cd:ef
// 01-23-45-67-89-ab
// 01-23-45-67-89-ab-cd-ef
// 0123.4567.89ab
// 0123.4567.89ab.cdef
func IsMAC(str string) bool {
	_, err := net.ParseMAC(str)
	return err == nil
}

// IsLatitude check if a string is valid latitude.
func IsLatitude(str string) bool {
	if str == "" {
		return false
	}

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}

	if 90.0 < f || f < -90.0 {
		return false
	}

	return true
}

// IsLongitude check if a string is valid longitude.
func IsLongitude(str string) bool {
	if str == "" {
		return false
	}

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}

	if 180.0 < f || f < -180.0 {
		return false
	}

	return true
}

// IsSemver check if string is valid semantic version
func IsSemver(str string) bool {
	return rxSemver.MatchString(str)
}

// IsStringLength check string's length (including multi byte strings)
func IsStringLength(str string, min int, max int) bool {
	slen := utf8.RuneCountInString(str)
	return slen >= min && slen <= max
}

// Exists returns whether the given file or directory exists or not
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// IsIntContains check if int slice contains checkItem
func IsIntContains(list []int, checkItem int) bool {
	if len(list) == 0 {
		return false
	}

	for _, item := range list {
		if item == checkItem {
			return true
		}
	}

	return false
}

// IsInt64Contains check if int64 slice contains checkItem
func IsInt64Contains(list []int64, checkItem int64) bool {
	if len(list) == 0 {
		return false
	}

	for _, item := range list {
		if item == checkItem {
			return true
		}
	}

	return false
}

// IsStringContains check if string slice contains checkItem
func IsStringContains(list []string, checkItem string) bool {
	if len(list) == 0 {
		return false
	}

	for _, item := range list {
		if item == checkItem {
			return true
		}
	}

	return false
}
