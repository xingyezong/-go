package main
import "fmt"
// leetcode -----1
// 思路当指定元素为i时，看看能否找到target-i的元素，可以用暴力法但是时间复杂度过高，因此用哈希
func twoSum(nums []int, target int) []int {
	hsmap := map[int]int{}
	for i, v := range nums {
		hsmap[v] = i
		if p,ok := hsmap[target - v]; ok{
			fmt.Println(ok)
			return []int{i,p}
		}

	}
	return nil
}
func main()  {
	nums := []int{2,7,11,5}
	target := 16
	var a = twoSum(nums,target)
	fmt.Println(a)
}
