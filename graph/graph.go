/*
 * @Description: 最小路径 
 * 这里用map是因为更加直观地看到结果，正常应该用数组，效率会更好。这里与《数据结构和算法分析》给出的伪代码实现有些出入， 这里实现需保证队列的唯一性。其实队列不在乎是否唯一，这里是为实现算法而做的限制。
  其实这里分无权重和有权重查找最小路径，如果无权重把VW结构体的Weight设置为 1 即可，而且不用判断权重大小。
 * @Author: longerQiu
 * @Date: 2019-11-23 09:53:13
 * @LastEditTime: 2019-11-24 17:28:45
 */
package graph

import (
	"fmt"
	"go-algorithm/structure"
)

type Vertex struct {
	Know int //标识
	D    int //最小值
	P    string // 父向量
}

//VW 相对权重
type VW struct {
	Vertex string //顶点
	Weight int //权重
}

//Dijkstra算法获取最小路径 有权重
type DijkstraGraph struct {
	table map[string]*Vertex //邻接表
	graph map[string][]*VW //图
}

func NewDijkstraGraph() *DijkstraGraph  {
	return new(DijkstraGraph)
}

//GetTable 获取邻接表
func (g *DijkstraGraph) GetGraph() map[string][]*VW {
	return g.graph
}

//GetTable 获取邻接表
func (g *DijkstraGraph) GetTable() map[string]*Vertex {
	return g.table
}

//MakeGraph 创建一张图和邻接表
func (g *DijkstraGraph) CreateGraphAndTable()  {
	g.graph = make(map[string][]*VW)
	g.table = make(map[string]*Vertex)
}

//AddVW 添加一个顶点
func (g *DijkstraGraph) AddVW(v string, args ...*VW) {
	for _, vw := range args {
		g.graph[v] = append(g.graph[v], vw)
	}

	g.table[v] = &Vertex{0, 0, "0"}
}

//GetVw 获取一个顶点
func (g *DijkstraGraph) GetVw(vw string) []*VW {
	if vw, ok := g.graph[vw]; ok {
		return vw
	} else {
		return nil
	}
}

func (g *DijkstraGraph) Dijkstra(v string) {
	q := structure.NewQueue()
	q.EnQueue(v)
	for !q.IsEmpty() {
		fmt.Printf("queue: %v \n", q)
		e, _ := q.DeQueue()
		pv := e.(string)
		fmt.Println("DeQueue: ", pv)
		graph := g.GetVw(pv)  //从图里获取他的出度
		if graph == nil {
			continue
		}
		
		g.table[pv].Know = 1
		for _, vw := range graph {
			if g.table[vw.Vertex].Know == 0 { 
				g.table[vw.Vertex].Know = 1
				g.table[vw.Vertex].D = vw.Weight + g.table[pv].D
				g.table[vw.Vertex].P = pv
				q.EnQueue(vw.Vertex)
			} else {
				if g.table[vw.Vertex].D > vw.Weight + g.table[pv].D {
					g.table[vw.Vertex].D = vw.Weight + g.table[pv].D
					g.table[vw.Vertex].P = pv
				}
			}
		}
	}
}


 