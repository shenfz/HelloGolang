package Graph

import (
	"fmt"
	"queue"
	"sync"
	"time"
)

/**
 * @Author: 方舟
 * @Email: 1328919715@qq.com
 * @Date: 2021/06/30 10:54
 * @Description:
 */

/*
  任务依赖 ,a任务执行的前提是b任务执行 ，ab任务构成依赖
  以有向图表示这种依赖关系，较为合适
  同时可以编排出 多依赖和多任务执行的先后关系

  图的遍历，和遍历树一样分为广度优先 BFS 和 深度优先 DFS

  图的两种存储方式：邻接矩阵、邻接表
  使用邻接矩阵可以更好地查询连通性，其原理也是用空间换时间

  graph 图
  vertex 顶点
  edge  边
*/

// 图结构
type DAG struct {
	Vertexs []*Vertex
}

// 顶点
type Vertex struct {
	Key      string
	Value    interface{}
	Parents  []*Vertex
	Children []*Vertex
}

// 添加顶点
func (dag *DAG) AddVertex(v *Vertex) {
	dag.Vertexs = append(dag.Vertexs, v)
}

// 添加边
func (dag *DAG) AddEdge(from, to *Vertex) {
	from.Children = append(from.Children, to)
	to.Parents = append(to.Parents, from)
}

func BFS(root *Vertex) []*Vertex {
	q := queue.New()
	q.Add(root)
	visited := make(map[string]*Vertex)
	all := make([]*Vertex, 0)
	for q.Length() > 0 {
		qSize := q.Length()
		for i := 0; i < qSize; i++ {
			//pop vertex
			currVert := q.Remove().(*Vertex)
			if _, ok := visited[currVert.Key]; ok {
				continue
			}
			visited[currVert.Key] = currVert
			all = append([]*Vertex{currVert}, all...)
			for _, val := range currVert.Children {
				if _, ok := visited[val.Key]; !ok {
					q.Add(val) //add child
				}
			}
		}
	}
	return all
}

func BFSNew(root *Vertex) [][]*Vertex {

}

// 单个执行
func doTask(vertexs *Vertex) {

}

// 并发执行
func doTasksNew(vertexs []*Vertex) {
	var wg sync.WaitGroup
	for _, v := range vertexs {
		wg.Add(1)
		go func(v *Vertex) {
			defer wg.Done()
			time.Sleep(5 * time.Second)
			fmt.Printf("do %v, result is %v \n", v.Key, v.Value)
		}(v) //notice
	}
	wg.Wait()
}
