package iteration

func Repeat(char string)  string{
	for i := 0; i < 2; i++{
		char = char + char
	}
	return char
}