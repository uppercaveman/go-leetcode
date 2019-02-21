package two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if v, ok := m[another]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return nil
}
