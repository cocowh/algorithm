package company

import (
	"fmt"
)

// 给定一个线段list，每个线段成员变量包括自身id， 起始点id， 结束点id。id均为string类型。
// 如果线段中某个线段的起点id等于另一个线段的终点id。则判定这两条线段可以连接。
// 请输出线段list所连接的最终的线段。大线段使用小线段的id表示。
// 例如
// Line1{"id": "hello1", "start_point": "test9", "end_point": "test2"}
// Line2{"id": "hello2", "start_point": "test3", "end_point": "test4"}
// Line3{"id": "hello3", "start_point": "test2", "end_point": "test3"}
// 输出为 ["hello1", "hello3", "hello2"]
// 成环或者有分支的情况抛出异常即可
// 时间复杂度要求O(n)
// 必须定义一个 包名为 `main` 的包，并实现 `main()` 函数。

type Point struct {
	Id        string
	LastPoint string
	NextPoint string
}

type Line struct {
	Points  []*Point
	SortIds []string
}

func (l *Line) Construct(points []*Point) {
	l.Points = points
}

func (l *Line) ToString() string {
	return fmt.Sprintf("%v", l.SortIds)
}

func (l *Line) JoinPoints() {
	if len(l.Points) == 0 {
		return
	}
	mlast, mnext := make(map[string]*Point), make(map[string]*Point)
	for _, v := range l.Points {
		if _, ok := mlast[v.Id]; ok { // 同一起点超过1个, 有分支或存在环
			panic("error, repeat last point")
		} else {
			mlast[v.LastPoint] = v
		}
		if _, ok := mnext[v.NextPoint]; ok {
			panic("error, repeat next point") //同一截止点超过1个, 有分支或环
		} else {
			mnext[v.NextPoint] = v
		}
	}
	l.SortIds = []string{l.Points[0].Id}
	lastStart, nextStart := l.Points[0].LastPoint, l.Points[0].NextPoint
	// 连接前面的
	for lastPoint, ok := mnext[l.Points[0].LastPoint]; ok; {
		l.SortIds = append([]string{lastPoint.Id}, l.SortIds...)
		lastPoint, ok = mnext[lastPoint.LastPoint]
		if ok && lastPoint.LastPoint == lastStart {
			panic("error, have ring")
		}
	}
	// 连接后面的
	for nextPoint, ok := mlast[l.Points[0].NextPoint]; ok; {
		l.SortIds = append(l.SortIds, []string{nextPoint.Id}...)
		nextPoint, ok = mlast[nextPoint.NextPoint]
		if ok && nextPoint.NextPoint == nextStart {
			panic("error, have ring")
		}
	}
}

func TestLine() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// 正常
	points := []*Point{
		{"hello1", "test9", "test2"},
		{"hello2", "test3", "test4"},
		{"hello3", "test2", "test3"},
	}
	l := &Line{points, []string{}}
	l.JoinPoints()
	fmt.Println(l.ToString())
	// --- 单环
	points = []*Point{
		{"hello1", "test1", "test2"},
		{"hello2", "test2", "test3"},
		{"hello3", "test3", "test1"},
	}
	l.Construct(points)
	l.JoinPoints()
	fmt.Println(l.ToString())
	// -- 分支
	points = []*Point{
		{"hello1", "test1", "test2"},
		{"hello2", "test3", "test2"},
		{"hello3", "test2", "test4"},
		{"hello4", "test4", "test1"},
	}
	l.Construct(points)
	l.JoinPoints()
	fmt.Println(l.ToString())
}
