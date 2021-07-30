package utils

import "regexp"

// Basic regular expressions for validating strings
const (
	pCreditCard string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
	// pISBN10     string = "^(?:[0-9]{9}X|[0-9]{10})$"
	// pISBN13     string = "^(?:[0-9]{13})$"
	// pAlpha       string = "^[a-zA-Z]+$"
	// pAlphanumeric string = "^[a-zA-Z0-9]+$"
	// pNumeric      string = "^[-+]?[0-9]+$"
	// pInt         string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	// pFloat       string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	// pHexadecimal string = "^[0-9a-fA-F]+$"
	// pASCII       string = "^[\x00-\x7F]+$"
	// pMultibyte string = "[^\x00-\x7F]"
	// pFullWidth string = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	// pHalfWidth string = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	// pBase64    string = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	// pPrintableASCII string = "^[\x20-\x7E]+$"
	pDataURI string = "^data:.+\\/(.+);base64$"
	pDNSName string = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{1,62}){1}(.[a-zA-Z0-9]{1}[a-zA-Z0-9_-]{1,62})*$`
	pURL     string = `^((ftp|https?):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(([a-zA-Z0-9]+([-\.][a-zA-Z0-9]+)*)|((www\.)?))?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?))(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`
	// pSSN      string = `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
	pWinPath  string = `^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	pUnixPath string = `^((?:\/[a-zA-Z0-9\.\:]+(?:_[a-zA-Z0-9\:\.]+)*(?:\-[\:a-zA-Z0-9\.]+)*)+\/?)$`
	pSemver   string = "^v?(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"
	Cameling  string = `[\p{L}\p{N}]+`
)

// Used by IsFilePath func
const (
	// Unknown is unresolved OS type
	Unknown = iota
	// Win is Windows type
	Win
	// Unix is *nix OS types
	Unix
)

// Regular expressions patterns
var (
	// rxEmail          = regexp.MustCompile(Email)
	rxCreditCard = regexp.MustCompile(pCreditCard)
	// rxAlpha          = regexp.MustCompile(Alpha)
	// rxAlphanumeric   = regexp.MustCompile(Alphanumeric)
	// rxNumeric        = regexp.MustCompile(Numeric)
	// rxInt         = regexp.MustCompile(pInt)
	// rxFloat       = regexp.MustCompile(pFloat)
	// rxHexadecimal = regexp.MustCompile(pHexadecimal)
	// rxASCII          = regexp.MustCompile(ASCII)
	// rxPrintableASCII = regexp.MustCompile(PrintableASCII)
	// rxMultibyte = regexp.MustCompile(pMultibyte)
	// rxFullWidth = regexp.MustCompile(pFullWidth)
	// rxHalfWidth = regexp.MustCompile(pHalfWidth)
	// rxBase64         = regexp.MustCompile(Base64)
	rxDataURI  = regexp.MustCompile(pDataURI)
	rxDNSName  = regexp.MustCompile(pDNSName)
	rxURL      = regexp.MustCompile(pURL)
	rxWinPath  = regexp.MustCompile(pWinPath)
	rxUnixPath = regexp.MustCompile(pUnixPath)
	rxSemver   = regexp.MustCompile(pSemver)
	rxCameling = regexp.MustCompile(Cameling)
)
