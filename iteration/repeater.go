package iteration

import "strings"

func Repeater(s string) string {
	// var repeated string
	var repeated strings.Builder
	for i := 0; i < 5; i++ {
		// Plus And operator. Same as: repeated = repeated + s
		// repeated += s
		repeated.WriteString(s)
	}
	// return repeated
	return repeated.String()
}
