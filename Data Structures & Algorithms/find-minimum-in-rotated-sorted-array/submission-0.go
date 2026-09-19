// if nums not rotated
// n[0] < n[last]

// n[i] < n[i+1]

// if nums is rotated
// n[0] > n[last]
// theres i where n[i] > n[i+1]

// track most min

// find smallest, start at middle 
// if n[0] > n[last], we know nums is rotated, so smallest is only between i and last
// bin search between i and last

// else, normal bin search

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if  nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}
