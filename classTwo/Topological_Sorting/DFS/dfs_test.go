package dfs

import "testing"

func TestDFS(t *testing.T) {
	g := InitGraph(5)
	g.AddEdge(2, 1)  //2先于1
	g.AddEdge(0, 1)  //0先于1
	g.AddEdge(0, 4)  //0先于4
	g.AddEdge(4, 3)  //4先于3
	TopoSortByDFS(g) //->0->2->1->4->3 结果符合预期
}
