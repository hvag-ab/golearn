package data_structure

import (
	"fmt"
)

// 坐标 棋盘上一个点
type Point struct {
	x int
	y int
}

//当前坐标 和 已走坐标路径
type Val struct {
	point Point
	path  []Point
}

// 类
type DFS struct {
	q      []Val
	target int
}

// 判断将要走的坐标 是否 在已走坐标中
func In(next_point Point, val Val) bool{
	path := val.path
	for _,p :=range path {
		if next_point == p{
			return true
		}
	}
	return false
}

// 下一步走的路径集合
func (self *DFS) next(value Val) []Point {
	ps := []Point{}

	d := []Point{}
	d = append(d,Point{-2, 1})
	d = append(d,Point{-1, 2})  
	d = append(d,Point{1, 2})   
	d = append(d,Point{2, 1})   
	d = append(d,Point{2, -1})  
	d = append(d,Point{1, -2})  
	d = append(d,Point{-1, -2}) 
	d = append(d,Point{-2, -1}) 

	cur_point := value.point

	x := cur_point.x
	y := cur_point.y
	for _, p :=range d{
		xx := x+p.x
		yy := y+p.y		
		if xx >= 0 && xx < self.target{
			if yy >= 0 && yy < self.target{
				if !In(Point{xx,yy},value){
					ps = append(ps,Point{xx,yy})
				}
			}
		}
	}
	return ps
}

func (self *DFS) pop() Val {
	q0 := self.q[len(self.q)-1]
	self.q = self.q[0:len(self.q)-1]
	return q0
}


func (self *DFS) Solve(){
	for len(self.q) > 0 {
		val := self.pop()
		// fmt.Println("val_before",val)
		for _,next_point :=range self.next(val){
			var new_val Val
			new_val = val
			new_val.path = append(new_val.path,next_point)
			if len(new_val.path) == self.target*self.target{
				fmt.Println("val_after",val)
				return
			}else{
				new_val.point = next_point
				self.q = append(self.q,new_val)
			}
			
		}
	}
}

func NewDFS(x,y,target int) *DFS{
	d := &DFS{}
	d.target = target

	p := Point{x,y}
	ps := []Point{p}
	val := Val{p,ps}
	d.q = []Val{val}
	return d
}

