package easy

func isStraight(nums []int) bool {
	m := make(map[int]int, 5)
	min, max := 14, 0
	for _, v := range nums {
		if v == 0 { // 跳过大小王
			continue
		}
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = 1
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return max-min < 5
}
