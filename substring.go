package stringsx

// Get the substring of a string, taking care of not breaking unicode characters
// For the optional arguments, first is the offset, second one is the length
func Substring(str string, limits ...int) string {
	var offset, length int
	var runes []rune
	switch {
	case len(limits) == 0:
		return str
	case len(limits) > 2:
		panic("the limits parameter should consist of 2 integers: offset and length")
	case len(limits) == 2:
		runes = []rune(str)
		offset = limits[0]
		length = limits[1]
		if length > len(runes) {
			length = len(runes)
		}
	case len(limits) == 1:
		runes = []rune(str)
		offset = limits[0]
		length = len(runes)
	}
	return string(runes[offset:length])
}
