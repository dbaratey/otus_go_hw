package hw02_unpack_string //nolint:golint,stylecheck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abccd",
			expected: "abccd",
		},
		{
			input:    "3abc",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa0b",
			expected: "aab",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "d\n5abc",
			expected: "d\n\n\n\n\nabc",
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

func TestUnpackWithEscape(t *testing.T) {
	//t.Skip() // Remove if task with asterisk completed

	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
		{
			input:    `qwe\\5a`,
			expected: `qwe\\\\\a`,
		},
		{
			input:    "qw\\ne",
			expected: ``,
			err:      ErrInvalidString,
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

func BenchmarkUnpack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%d", 42); x != "42" {
			_, err := Unpack("asd0asda0sd0ada0sd0a0d0a0d0as0d0a0d0a0d0a0d0a0d0a0d0a0d0a0d0a0dhjhkjhshfjsdbwkjbebkfjhwuefljwfnuewnfekwnfuwefnwkejfnkewjndxc.,m.sdlk;flksdjlkgjdskljgkljfnms,mdnfnbkdjfhskdnsdjkvndjkvsn")
			if err != nil {
				b.Fatal(err)
			}

		}
	}
}
