package util

import (
	"code-comment-analyzer/protocol"
	"strconv"
)

func FormatUserIDStr(userID uint64) string {
	return strconv.FormatUint(userID, 10)
}

func FileSuffixToLanguage(suffix string) string {
	switch suffix {
	case protocol.FileSuffixPY:
		return protocol.LanguagePython
	case protocol.FileSuffixGO:
		return protocol.LanguageGo
	case protocol.FileSuffixJAVA:
		return protocol.LanguageJava
	case protocol.FileSuffixC:
		return protocol.LanguageC
	case protocol.FileSuffixCPP:
		return protocol.LanguageCpp
	case protocol.FileSuffixJS:
		return protocol.LanguageJs
	case protocol.FileSuffixHTML:
		return protocol.LanguageHtml
	case protocol.FileSuffixCSS:
		return protocol.LanguageCss
	default:
		return protocol.LanguageUnknown
	}
}
