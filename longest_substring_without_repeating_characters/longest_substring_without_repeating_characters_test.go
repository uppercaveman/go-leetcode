package longest_substring_without_repeating_characters

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []string{
		"abcabcbb",
		"aaaaaaaa",
		"abcde",
		"",
		"pwwkew",
		"wpwwwwdcawwxyzk",
		"wpwwwpwpwp",
		"wpwwwwwwpsdwd",
	}
	results := []int{
		3,
		1,
		5,
		0,
		3,
		5,
		2,
		4,
	}

	for i := 0; i < len(results); i++ {
		if ret := lengthOfLongestSubstring(tests[i]); ret != results[i] {
			t.Fatalf("case %d failed\nactual: %d, expect: %d\n", i, ret, results[i])
		}
	}
}
