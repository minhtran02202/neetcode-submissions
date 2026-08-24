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
    memo := make([]int, len(nums))

	for i, _ := range memo{
		memo[i] = -1
	}

	var findAt func(int) int
	findAt = func(at int) int {
		if at == len(nums) {
			return 0
		}

		if memo[at] != -1 {
			return memo[at]
		}

		res := 1

		for i := at + 1; i < len(nums); i++ {
			if nums[i] > nums[at] {
				res = max(res, 1 + findAt(i))
			}
		}

		memo[at] = res
		return memo[at]
	}

	res := 0
	for i := len(nums) - 1; i >= 0; i--{
		res = max(res, findAt(i))
	}

	return res
}
