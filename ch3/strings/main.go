package main

import (
	"fmt"
	"strconv"
	"strings"
)

// HasPrefix ...
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// HasSuffix ...
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Contains ...
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

// basename1
func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// basename2
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// comma
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(HasPrefix("apple", "app"))
	fmt.Println(HasSuffix("books", "oks"))
	fmt.Println(Contains("soup", "ou"))

	fmt.Println(basename1("a/b/cdef.go"))
	fmt.Println(basename2("a/b/cdef.go"))
	fmt.Println(comma("1234567"))

	// string contains an array of bytes.
	// once created it cannot be changed
	// ie. strings are immutable.
	// however they can be converted to byte slices and back
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s2)

	// Converting between Strings and Numbers
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	// FormatInt and FormatUint can be used to format numbers in a different base
	fmt.Println(strconv.FormatInt(int64(x), 2))

	s1 := fmt.Sprintf("x=%b", x)
	fmt.Println(s1)

	x1, err := strconv.Atoi("123")
	fmt.Println(x1, err)

	y2, err := strconv.ParseInt("123", 10, 64)
	fmt.Println(y2, err)
}
