package main

const upperToLowerDiff = 32

func normalize(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + upperToLowerDiff
	}
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return c
	}
	return ' '
}

func isAlphanumeric(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for !isAlphanumeric(s[l]) && l < len(s)-1 {
			l++
		}
		for !isAlphanumeric(s[r]) && r > 0 {
			r--
		}
		nL, nR := normalize(s[l]), normalize(s[r])
		if nL != nR {
			return false
		}
		l++
		r--
	}
	return true
}
