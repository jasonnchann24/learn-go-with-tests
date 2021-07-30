package iteration

// Repeat takes a string of characters and repeat it for 'n' of times
func Repeat(char string, n int) (repeated string) {
	count := 0
	for count < n {
		repeated += char
		count++
	}
	return
}
