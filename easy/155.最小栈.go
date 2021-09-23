package easy

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start

type MinStack struct {
	e, m []int
	l    int
}

func MinConstructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), 0}
}

func (c *MinStack) Push(x int) {
	c.e = append(c.e, x)
	if c.l == 0 {
		c.m = append(c.m, x)
	} else {
		min := c.GetMin()
		if min > x {
			c.m = append(c.m, x)
		} else {
			c.m = append(c.m, min)
		}
	}
	c.l++
}

func (c *MinStack) Pop() {
	if c.l == 0 {
		return
	}
	c.l--
	c.m = c.m[:c.l]
	c.e = c.e[:c.l]
}

func (c *MinStack) Top() int {
	return c.e[c.l-1]
}

func (c *MinStack) GetMin() int {
	return c.m[c.l-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
