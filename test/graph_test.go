/*
 * @Description: 
 * @Author: longerQiu
 * @Date: 2019-11-23 11:04:24
 * @LastEditTime: 2019-11-24 17:29:03
 */
package test

import (
	"fmt"
	"testing"
	"go-algorithm/graph"
)

func TestDijkstra(t *testing.T) {
	g := graph.NewDijkstraGraph()
	g.CreateGraphAndTable()
	g.AddVW("v1", &graph.VW{"v2", 2}, &graph.VW{"v4", 1})
	g.AddVW("v2", &graph.VW{"v4", 3}, &graph.VW{"v5", 10})
	g.AddVW("v3", &graph.VW{"v1", 4}, &graph.VW{"v6", 5})
	g.AddVW("v4", &graph.VW{"v3", 2}, &graph.VW{"v5", 2}, &graph.VW{"v6", 8}, &graph.VW{"v7", 4})
	g.AddVW("v5", &graph.VW{"v7", 6})
	g.AddVW("v6")
	g.AddVW("v7", &graph.VW{"v6", 1})
	g.Dijkstra("v1")
	tb := g.GetTable()
	for k, v := range tb {
		fmt.Printf("key: %s, vertex: %+v \n", k, v)
	}
}
