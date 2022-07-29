package main

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。
示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
链接：https://leetcode.cn/problems/two-sum
*/

// twoSum1 简单的双层循环, 时间复杂度 O(n^2)
// 执行用时：28 ms, 在所有 Go 提交中击败了23.27%的用户
// 内存消耗：3.4 MB, 在所有 Go 提交中击败了78.95%的用户
func twoSum1(nums []int, target int) []int {
	result := make([]int, 2, 2)
	for i, num := range nums {
		for j := i + 1; j <= len(nums)-1; j++ {
			if num+nums[j] == target {
				result[0], result[1] = i, j
				return result
			}
		}
	}
	return result
}

// twoSum2 哈希表解法, 时间复杂度 O(n)
// 执行用时：4 ms, 在所有 Go 提交中击败了94.53% 的用户
// 内存消耗：4 MB, 在所有 Go 提交中击败了63.97%的用户
func twoSum2(nums []int, target int) []int {
	result := make([]int, 2, 2)
	myMap := make(map[int]int, len(nums))
	// map的 key为切片的值，value为索引
	for i, num := range nums {
		idx, ok := myMap[target-num]
		if ok {
			result[0], result[1] = i, idx
			return result
		}
		myMap[num] = i
	}

	return result
}

//func main() {
//	nums := []int{2, 5, 5, 11}
//	target := 10
//	res := twoSum2(nums, target)
//	fmt.Println(res)
//}
