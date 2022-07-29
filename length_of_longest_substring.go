package main

import (
	"fmt"
	"strings"
)

// 无重复字符的最长子串
/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

输入：s = "dvdf"
输出：3

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// ---------------------------------------------------------------------
/**
我的解题思路1（错误的，不够全面，s := "dvdf"，应该返回3,而不是2）：
从第二个开始，用该值去判断在不在前面的字符串里面，如果在则已经重复
如"abcabcbb", 从第二个开始也就是b, 根字符就为"a" b 不在a里面，则根字符变为为"ab"。
c 不在ab里面，根字符变成了"abc"
第四个字符a在abc里面，那么目前最长不重复为3, 根字符变成第四个字符 "a"
*/
type uu string

// -k乘n解法
func lengthOfLongestSubstring(s string) int {

	current := 1
	large := 1
	if len(s) == 1 {
		return 1
	}
	if len(s) == 0 {
		return 0
	}
	root := []byte(string(s[0]))
	for i := 1; i <= len(s)-1; i++ {
		currentS := s[i]
		e := string(currentS)
		fmt.Println(e)
		ok := strings.Contains(string(root), string(currentS))
		if ok {
			if current > large {
				large = current
			}
			current = 1
			root = []byte(string(currentS))

		} else {
			current++
			root = append(root, currentS)
		}
		r := string(root)
		fmt.Println(r)
	}
	if current > large {
		large = current
	}
	return large
}

func main() {
	s := "dvdf"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)

}
