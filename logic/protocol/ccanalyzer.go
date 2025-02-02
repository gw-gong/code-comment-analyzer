package protocol

const (
	LanguagePython  = "Python"
	LanguageGo      = "Go"
	LanguageJava    = "Java"
	LanguageC       = "C"
	LanguageCpp     = "CPP"
	LanguageJs      = "JavaScript"
	LanguageHtml    = "HTML"
	LanguageCss     = "CSS"
	LanguageUnknown = "Unknown"
)

func IsLanguageSupported(language string) bool {
	switch language {
	case LanguagePython, LanguageGo, LanguageJava, LanguageC, LanguageCpp, LanguageJs, LanguageHtml, LanguageCss:
		return true
	default:
		return false
	}
}

func FileSuffixToLanguage(suffix string) string {
	switch suffix {
	case ".py":
		return LanguagePython
	case ".go":
		return LanguageGo
	case ".java":
		return LanguageJava
	case ".c":
		return LanguageC
	case ".cpp":
		return LanguageCpp
	case ".js":
		return LanguageJs
	case ".html":
		return LanguageHtml
	case ".css":
		return LanguageCss
	default:
		return LanguageUnknown
	}
}
