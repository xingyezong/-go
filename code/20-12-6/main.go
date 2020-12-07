package main
import "fmt"
// leetcode -----1
// 思路当指定元素为i时，看看能否找到target-i的元素，可以用暴力法但是时间复杂度过高，因此用哈希
func twoSum(nums []int, target int) []int {
	hsmap := map[int]int{}
	for i, v := range nums {

		if p,ok := hsmap[target - v]; ok{
			fmt.Println(ok)
			return []int{i,p}
		}
		hsmap[v] = i // 要放在比较后，不然若target恰好等于首位数字的两倍时候会出错
	}
	return nil
}
func main()  {
	nums := []int{2,7,11,5}
	target := 16
	var a = twoSum(nums,target)
	fmt.Println(a)
}
