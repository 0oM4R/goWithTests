package main

const spain = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const helloSpanishPrefix = "Hola, "
const helloFrenchPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spain:
		prefix = helloSpanishPrefix
	case french:
		prefix = helloFrenchPrefix
	default:
		prefix = helloPrefix
	}
	return
}

// func main() {
// 	fmt.Println(Hello("world!"))
// }
