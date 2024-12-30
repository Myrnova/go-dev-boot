package CH8

func sum(nums ...int) int {
	value := 0
	for i := 0; i < len(nums); i++ {
		value += nums[i]
	}
	return value
}
