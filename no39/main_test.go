package no39_test

import (
	"strings"
	"testing"
)

func BenchmarkSliceConcatination(b *testing.B) {
	abc := "abcdefghijklmnopqrstuvwxyz"

	testcases := map[string]func(*testing.B){
		"simple concat": func(b *testing.B) {
			res := ""
			for _, rne := range abc {
				res += string(rne)
			}
		},
		"string builder": func(b *testing.B) {
			sb := strings.Builder{}
			for _, rne := range abc {
				_, _ = sb.WriteRune(rne)
			}
			_ = sb.String()
		},
		"string builder with pre-alloc": func(b *testing.B) {
			sb := strings.Builder{}
			sb.Grow(len(abc))
			for _, rne := range abc {
				_, _ = sb.WriteRune(rne)
			}
			_ = sb.String()
		},
	}

	for range b.N {
		for name, bfunc := range testcases {
			b.Run(name, bfunc)
		}
	}
}
