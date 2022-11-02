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
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchhHelloPrefix + name
	}
	return englishHelloPrefix + name
}