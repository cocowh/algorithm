package company

import (
	"algorithm/common"
	"fmt"
)

// 题目大意
// 由两个坐标构建一个矩形，判断两个矩形是否相交，求相交的矩形
// class Rect {
//     public int minx;  //左下角x坐标
//     public int miny;  //左下角y坐标
//     public int maxx;  //右上角x坐标
//     public int maxy;  //右下角y坐标

//     public Rect(int minx, int miny, int maxx, int maxy){
//         this.minx = minx;
//         this.miny = miny;
//         this.maxx = maxx;
//         this.maxy = maxy;
//     }

//     public String toString(){
//         return String.format("[%d,%d,%d,%d]",this.minx,this.miny,this.maxx,this.maxy);
//     }
// }

// class MyCode {
//     // 不相交
//     public boolean judgeMidPart(int min1, int max1, int min2, int max2) {
//         // 最小大于最大，不相交 ,其余相交或包含
//         if (max1 <= min2 || min1 >= max2) {
//             return false;
//         }
//         return true;
//     }

//     // 取坐标。小取大，大取小
//     public int getRoordinate(int a, int b, boolean isMin) {

//         if (isMin) {
//             return Math.max(a, b);
//         } else {
//             return Math.min(a, b);
//         }
//     }

//     public Rect computeIntersect(Rect r1, Rect r2){
//         //todo
//         // 排除不相交状况
//         if (!this.judgeMidPart(r1.minx, r1.maxx, r2.minx, r2.maxx)
//                 || !this.judgeMidPart(r1.miny, r1.maxy, r2.miny, r2.miny)) {
//             return null;
//         }

//         // 取交点坐标
//         int newMinx = this.getRoordinate(r1.minx, r2.minx, true);
//         int newMiny = this.getRoordinate(r1.miny, r2.miny, true);
//         int newMaxx = this.getRoordinate(r1.maxx, r2.maxx, false);
//         int newMaxy = this.getRoordinate(r1.maxy, r2.maxy, false);

//         System.out.println(newMinx);
//         System.out.println(newMiny);
//         System.out.println(newMaxx);
//         System.out.println(newMaxy);
//         return new Rect(newMinx,newMiny, newMaxx, newMaxy);
//     }

//     public static void main (String[] args) {
//         MyCode test= new MyCode();
//         // 相交部分
//         Rect r1 = new Rect(1,1,3,3);
//         Rect r2 = new Rect(2,2,4,4);
//         Rect r3 = test.computeIntersect(r1,r2);
//         System.out.println(r3.toString());
//         // 包含
//         Rect r4 = new Rect(1,1, 4,4);
//         Rect r5 = new Rect(2,2, 3,3);
//         Rect r6 = test.computeIntersect(r4,r5);
//         System.out.println(r6.toString());
//         // 部分
//         Rect r7 = new Rect(1,1, 2,2);
//         Rect r8 = new Rect(1,2, 3,2);
//         Rect r9 = test.computeIntersect(r7,r8);
//         System.out.println(r9.toString());
//     }
// }

type Rect struct {
	Minx int //左下角x坐标
	Miny int //左下角y坐标
	Maxx int //右上角x坐标
	Maxy int //右上角y坐标
}

func (r *Rect) Construct(minx, miny, maxx, maxy int) {
	r.Minx = minx
	r.Miny = miny
	r.Maxx = maxx
	r.Maxy = maxy
}

func (r *Rect) ToString() string {
	return fmt.Sprintf("[%d,%d,%d,%d]", r.Minx, r.Miny, r.Maxx, r.Maxy)
}

type GenRect struct {
	R1 *Rect
	R2 *Rect
}

func (c *GenRect) Construct(r1, r2 *Rect) {
	c.R1 = r1
	c.R2 = r2
}

func (c *GenRect) CheckMidPart(min1, max1, min2, max2 int) bool {
	// 最小大于最大，不相交 ,其余相交或包含
	if max1 <= min2 || min1 >= max2 {
		return false
	}
	return true
}

// 取坐标。小取大，大取小
func (c *GenRect) GetRoordinate(a, b int, isMin bool) int {
	if isMin {
		return common.MaxInt(a, b)
	} else {
		return common.MinInt(a, b)
	}
}

func (c *GenRect) ComputeIntersect() *Rect {
	if !c.CheckMidPart(c.R1.Minx, c.R1.Maxx, c.R2.Minx, c.R2.Maxx) ||
		!c.CheckMidPart(c.R1.Miny, c.R1.Maxy, c.R2.Miny, c.R2.Maxy) {
		return &Rect{}
	}
	newMinx := c.GetRoordinate(c.R1.Minx, c.R2.Minx, true)
	newMiny := c.GetRoordinate(c.R1.Miny, c.R2.Miny, true)
	newMaxx := c.GetRoordinate(c.R1.Maxx, c.R2.Maxx, false)
	newMaxy := c.GetRoordinate(c.R1.Maxy, c.R2.Maxy, false)
	return &Rect{newMinx, newMiny, newMaxx, newMaxy}
}
func TestComputeIntersect() {
	ins := &GenRect{}
	res := &Rect{}
	r1 := &Rect{1, 1, 3, 3}
	r2 := &Rect{2, 2, 4, 4}
	ins.Construct(r1, r2)
	res = ins.ComputeIntersect()
	fmt.Println(res.ToString())
	// ---
	r1.Construct(1, 1, 4, 4)
	r2.Construct(2, 2, 3, 3)
	ins.Construct(r1, r2)
	res = ins.ComputeIntersect()
	fmt.Println(res.ToString())
	// --
	r1.Construct(1, 1, 2, 2)
	r2.Construct(1, 2, 3, 2)
	ins.Construct(r1, r2)
	res = ins.ComputeIntersect()
	fmt.Println(res.ToString())
}
