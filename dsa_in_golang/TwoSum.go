func twoSum(nums []int, target int) []int {
    var hash = make(map[int]int)
    for i, num := range nums{
        if index, ok := hash[target-num]; ok{
            return []int{i,index}
        }
        hash[nums[i]] = i
    }
    return nil
}
