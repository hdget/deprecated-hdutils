package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"unicode"
	"unsafe"
)

// ToString convert the input to a string.
func ToString(obj interface{}) string {
	res := fmt.Sprintf("%v", obj)
	return string(res)
}

// ToJson convert the input to a valid JSON string
func ToJson(obj interface{}) (string, error) {
	switch reply := obj.(type) {
	case string:
		return reply, nil
	}

	bs, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return BytesToString(bs), nil
}

// ToFloat convert the input string to a float, or 0.0 if the input is not a float.
func ToFloat(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		res = 0.0
	}
	return res, err
}

// ToInt convert the input string to an int, or 0 if the input is not an integer.
func ToInt(str string) (int, error) {
	res, err := strconv.Atoi(str)
	if err != nil {
		res = 0
	}
	return res, err
}

// ToInt64 convert the input string to an 64bit integer, or 0 if the input is not an integer.
func ToInt64(str string) (int64, error) {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		res = 0
	}
	return res, err
}

// ToBoolean convert the input string to a boolean.
func ToBoolean(str string) (bool, error) {
	res, err := strconv.ParseBool(str)
	if err != nil {
		res = false
	}
	return res, err
}

// ToCamelCase converts from underscore separated form to camel case form.
func ToCamelCase(str string) string {
	byteSrc := []byte(str)
	chunks := rxCameling.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		chunks[idx] = bytes.Title(val)
	}
	return string(bytes.Join(chunks, nil))
}

// ToSnakeCase converts from camel case form to underscore separated form.
func ToSnakeCase(s string) string {
	s = ToCamelCase(s)
	runes := []rune(s)
	length := len(runes)
	var out []rune
	for i := 0; i < length; i++ {
		out = append(out, unicode.ToLower(runes[i]))
		if i+1 < length && (unicode.IsUpper(runes[i+1]) && unicode.IsLower(runes[i])) {
			out = append(out, '_')
		}
	}

	return string(out)
}

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
