package iteration

import "strings"

func Repeater(s string) string {
	var repeated strings.Builder
	for i := 0; i < 5; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
