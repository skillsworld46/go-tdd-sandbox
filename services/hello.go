package services

const (
	Spanish = "Spanish"
	French = "French"
   
	englishHello = "Hello, "
	spanishHello = "Hola, "
	frenchHello = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (helloPrefix string) {
	switch language {
	case Spanish:
		helloPrefix = spanishHello
	case French:
		helloPrefix = frenchHello
	default:
		helloPrefix = englishHello
	}
	return
}