package hello

const englishHelloPrefix = "hello, "

func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}
	return englishHelloPrefix + name
}