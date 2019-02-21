package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	runeMap := make(map[rune]int)
	longest := 0
	preLength := 0
	for i, v := range s {
		var length int
		if val, ok := runeMap[v]; !ok || val < i-preLength {
			length = preLength + 1
		} else {
			length = i - val
		}
		if length > longest {
			longest = length
		}
		preLength = length
		runeMap[v] = i
	}
	return longest
}

func lengthOfLongestSubstring2(s string) int {
	m := make(map[rune]int)
	maxLength := 0
	currentLength := 0
	previousDuplicatedIndex := 0

	for i, v := range s {

		newDuplicatedIndex := m[v]

		if newDuplicatedIndex > previousDuplicatedIndex {
			currentLength -= (newDuplicatedIndex - previousDuplicatedIndex)
			previousDuplicatedIndex = newDuplicatedIndex
		}

		currentLength++

		if currentLength > maxLength {
			maxLength = currentLength
		}

		m[v] = i + 1
	}

	return maxLength
}
