package main 

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name,language string) string{
	if (name == ""){
		return englishHelloPrefix + "World"
	}else if (language == spanish){
		return spanishHelloPrefix + name
	}else if(language == "French"){
		return "Bonjour, "+name
	}
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("world",""))
}