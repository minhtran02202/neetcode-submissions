func hasDuplicate(nums []int) bool {
    hash := make(map[int]bool)

    for _,val:=range(nums){
        if hash[val]{return true}
        hash[val]=true
    }

    return false
}
