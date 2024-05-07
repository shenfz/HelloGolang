package Graph

import "testing"

/**
 * @Author: 方舟
 * @Email: 1328919715@qq.com
 * @Date: 2021/06/30 10:56
 * @Description:
 */

// 对该图进行广度优先遍历，通过引入队列来减少时间复杂度，遍历后生成一个包含所有顶点的 slice

func Test_TaskDispatch(t *testing.T) {

	var dag = &DAG{}
	//添加顶点
	va := &Vertex{Key: "a", Value: "1"}
	vb := &Vertex{Key: "b", Value: "2"}
	vc := &Vertex{Key: "c", Value: "3"}
	vd := &Vertex{Key: "d", Value: "4"}
	ve := &Vertex{Key: "e", Value: "5"}
	vf := &Vertex{Key: "f", Value: "6"}
	vg := &Vertex{Key: "g", Value: "7"}
	vh := &Vertex{Key: "h", Value: "8"}
	vi := &Vertex{Key: "i", Value: "9"}
	//添加边
	dag.AddEdge(va, vb)
	dag.AddEdge(va, vc)
	dag.AddEdge(va, vd)
	dag.AddEdge(vb, ve)
	dag.AddEdge(vb, vh)
	dag.AddEdge(vb, vf)
	dag.AddEdge(vc, vf)
	dag.AddEdge(vc, vg)
	dag.AddEdge(vd, vg)
	dag.AddEdge(vh, vi)
	dag.AddEdge(ve, vi)
	dag.AddEdge(vf, vi)
	dag.AddEdge(vg, vi)

	// 1. 串行遍历

	// 2. 层级遍历

}
