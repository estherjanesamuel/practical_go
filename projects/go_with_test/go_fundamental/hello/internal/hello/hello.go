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
		return englishHelloPrefix + "world"
	}
	prefix := englishHelloPrefix
	switch language {
	case french:
		prefix = frenchhHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return prefix + name
}