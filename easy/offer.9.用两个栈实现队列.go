package easy

import "fmt"

func TestStackToQueue() {
	obj := Constructor()
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	obj.AppendTail(1)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}

type CQueue struct {
	Stack1 []int
	Stack2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (c *CQueue) AppendTail(value int) {
	c.Stack1 = append(c.Stack1, value)
}

func (c *CQueue) DeleteHead() int {
	if len(c.Stack2) == 0 {
		for i := len(c.Stack1) - 1; i >= 0; i-- {
			c.Stack2 = append(c.Stack2, c.Stack1[i])
			c.Stack1 = c.Stack1[:len(c.Stack1)-1]
		}
	}
	if len(c.Stack2) == 0 {
		return -1
	} else {
		v := c.Stack2[len(c.Stack2)-1]
		c.Stack2 = c.Stack2[:len(c.Stack2)-1]
		return v
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
