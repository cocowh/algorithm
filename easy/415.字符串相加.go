package easy

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	// 确保 n1 <= n2
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	n1, n2 := len(num1), len(num2)
	a1, a2 := []byte(num1), []byte(num2)

	// 进位
	carry := byte(0)

	// buf 保存 []byte 格式的答案 +1保存可能存在的进位
	buf := make([]byte, n2+1)
	// 默认结果进位
	buf[0] = '1'

	i := 1
	for i <= n2 {
		// a1 和 a2 相加
		if i <= n1 {
			buf[n2+1-i] = a1[n1-i] - '0'
		}
		buf[n2+1-i] += a2[n2-i] + carry

		// 处理进位问题
		if buf[n2+1-i] > '9' {
			buf[n2+1-i] -= 10
			carry = byte(1)
		} else {
			carry = byte(0)
		}
		i++
	}

	if carry == 1 {
		return string(buf)
	}
	return string(buf[1:])
}

// @lc code=end
