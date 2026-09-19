import "slices"

func findMin(nums []int) int { slices.Sort(nums); return nums[0] }
