package iteration

// Repeat receives a string and repeats it 5 times
func Repeat(character string, times int) string {
	var repeated string
	if times <= 0 {
		return ""
	}
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
