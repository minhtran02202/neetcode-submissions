import "slices"
func findKthLargest(nums []int, k int) int {
	slices.Sort(nums)
	return nums[len(nums) - k]
}
