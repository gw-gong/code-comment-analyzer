package protocol

const (
	LanguagePython = "Python"
	LanguageGo     = "Go"
	LanguageJava   = "Java"
	LanguageC      = "C"
	LanguageCpp    = "CPP"
	LanguageJs     = "JavaScript"
	LanguageHtml   = "HTML"
	LanguageCss    = "CSS"
)

func IsLanguageSupported(language string) bool {
	switch language {
	case LanguagePython, LanguageGo, LanguageJava, LanguageC, LanguageCpp, LanguageJs, LanguageHtml, LanguageCss:
		return true
	default:
		return false
	}
}
