package interview

import (
	"fmt"
	"testing"
)

type path struct {
	x int
	y int
}

//// 9*9，找一条路径
//// 递归
//func find_one_path(graph [][]int, src, dst *path) []*path {
//	paths := make([]*path, 0)
//}
//
//func _find_one_path(graph [][]int, cur, dst *path, finds []*path) {
//	if graph[cur.x][cur.y] == 0 {
//		finds = append(finds, cur)
//	} else {
//		finds--
//	}
//	_find_one_path(graph, cur, dst, finds)
//}

// 类似A*做法，有open节点集合（可以走）和visited节点结合（已经走过的）
// 按广度找，一层一层找
// 注意：要先想清楚算法，再去动手写
// 把open当作队列，所以是广度优先
// 把open当做栈，所以是深度优先
// 狄克斯特拉和a*的不同在于，评估函数不同

// BFS实现
func find_path(graph [][]int, src, dst *path) []*path {
	open := make([]*path, 0)
	open = append(open, src)
	visited := make(map[int]struct{})
	findPath := make([]*path, 0)

	for len(open) > 0 {
		i := 0
		cur := open[i]
		open = open[0 : len(open)-1]
		if _, ok := visited[cur.x*100+cur.y]; !ok && graph[cur.x][cur.y] == 0 {
			findPath = append(findPath, cur)
			visited[cur.x*100+cur.y] = struct{}{}

			// 到终点了
			if cur.x == dst.x && cur.y == dst.y {
				break
			}

			// 格子可以走，入周围的格子，可以走&没走过
			if _, ok1 := visited[(cur.x+1)*100+cur.y]; !ok1 && graph[cur.x+1][cur.y] == 0 {
				open = append(open, &path{x: cur.x + 1, y: cur.y})
			}
			if _, ok1 := visited[(cur.x-1)*100+cur.y]; !ok1 && graph[cur.x-1][cur.y] == 0 {
				open = append(open, &path{x: cur.x - 1, y: cur.y})
			}
			if _, ok1 := visited[cur.x*100+cur.y+1]; !ok1 && graph[cur.x][cur.y+1] == 0 {
				open = append(open, &path{x: cur.x, y: cur.y + 1})
			}
			if _, ok1 := visited[cur.x*100+cur.y-1]; !ok1 && graph[cur.x][cur.y-1] == 0 {
				open = append(open, &path{x: cur.x, y: cur.y - 1})
			}
		} else {
			// 格子不能走

		}
		i++
	}

	return findPath
}

func Test_find_path(t *testing.T) {
	//6*6
	graph := [][]int{
		//0 1  2  3  4  5
		{1, 1, 1, 1, 1, 1}, //0
		{1, 0, 0, 0, 0, 1}, //1
		{1, 0, 1, 1, 1, 1}, //2
		{1, 0, 1, 0, 0, 1}, //3
		{1, 0, 1, 0, 0, 1}, //4
		{1, 1, 1, 1, 1, 1}, //5
	}
	src := &path{x: 4, y: 1}
	dst := &path{x: 1, y: 4}
	paths := find_path(graph, src, dst)
	printPath(paths)
}

func printPath(paths []*path) {
	for _, item := range paths {
		fmt.Println("x:", item.x, "y:", item.y)
	}
}
