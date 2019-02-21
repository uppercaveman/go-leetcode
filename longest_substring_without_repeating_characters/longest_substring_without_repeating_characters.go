package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	runeMap := make(map[rune]int)
	maxLength := 0
	preLength := 0
	for k, v := range s {
		if val, ok := runeMap[v]; !ok || val < k-preLength {
			preLength++
		} else {
			preLength = k - val
		}
		if preLength > maxLength {
			maxLength = preLength
		}

		runeMap[v] = k
	}
	return maxLength
}

func lengthOfLongestSubstring2(s string) int {
	m := make(map[rune]int)
	maxLength := 0
	currentLength := 0
	previousDuplicatedIndex := 0

	for k, v := range s {

		newDuplicatedIndex := m[v]

		if newDuplicatedIndex > previousDuplicatedIndex {
			currentLength -= (newDuplicatedIndex - previousDuplicatedIndex)
			previousDuplicatedIndex = newDuplicatedIndex
		}

		currentLength++

		if currentLength > maxLength {
			maxLength = currentLength
		}

		m[v] = k + 1
	}

	return maxLength
}
