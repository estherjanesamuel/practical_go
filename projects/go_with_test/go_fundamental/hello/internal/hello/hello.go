package hello

const englishHelloPrefix = "hello, "

func Hello(name string, language string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}