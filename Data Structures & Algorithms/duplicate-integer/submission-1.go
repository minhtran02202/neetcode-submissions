import "slices"

func hasDuplicate(nums []int) bool {
    if len(nums)==0 || len(nums)==1{return false}
    slices.Sort(nums)
    //if empty or 1 element, false
    for i:=1; i<len(nums); i++{
        if nums[i]==nums[i-1]{return true}
    }

    return false
}
