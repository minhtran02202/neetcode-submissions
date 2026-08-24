// Input: nums = nums

// at ith:
// actions:

// for each nums at [i+1:],
// check if there's a increase subsequence

// there's none, so move to i+1

// for each nums at [i+1],
// check if there's a increase subsequence

// i+x is increase subsequent of i+1, find i+x subsequence
// ...

// result increase subsequent of i is max of increase subsequent of x, y, etc + 1 (1 to include count i)


func lengthOfLIS(nums []int) int {
	finalRes := 0
	memo := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		res := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i]{
				res = max(res, 1 + memo[j])
			}
		}
		memo[i] = res
		finalRes = max(finalRes, res)
	}

	return finalRes
}
