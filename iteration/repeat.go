package iteration

import "strings"

const repeatCount = 5

// BenchmarkRepeat-8   	30310506	        42.86 ns/op	       8 B/op	       1 allocs/op
func Repeat(s string) string {
	sb := strings.Builder{}

	for i := 0; i < repeatCount; i++ {
		sb.WriteString(s)
	}

	return sb.String()
}

// BenchmarkRepeat-8   	 5995449	       182.1 ns/op	      16 B/op	       4 allocs/op
// func Repeat(character string) string {
// 	var repeated string
// 	for i := 0; i < repeatCount; i++ {
// 		repeated += character
// 	}
// 	return repeated
// }
