package github.com/hhhanafubukiii/Golang-practice/leetcode

type Slice []int

var def int

func TwoSum(slice Slice, target int) []int {
	hashMap := make(map[int]int)
	for idx, val := range slice {
		def = target - val
		if _, ok := hashMap[def]; ok {
			return []int{hashMap[def], idx}
		}
		hashMap[val] = idx
	}
	return []int{0, 0}
}
