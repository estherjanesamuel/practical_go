package hello


const (
	french = "French"
	spanish = "Spanish"
	frenchhHelloPrefix = "bonjour, "
	spanishHelloPrefix = "hola, "
	englishHelloPrefix = "hello, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchhHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return 
}