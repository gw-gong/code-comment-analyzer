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
	case ".py":
		return protocol.LanguagePython
	case ".go":
		return protocol.LanguageGo
	case ".java":
		return protocol.LanguageJava
	case ".c":
		return protocol.LanguageC
	case ".cpp":
		return protocol.LanguageCpp
	case ".js":
		return protocol.LanguageJs
	case ".html":
		return protocol.LanguageHtml
	case ".css":
		return protocol.LanguageCss
	default:
		return protocol.LanguageUnknown
	}
}
