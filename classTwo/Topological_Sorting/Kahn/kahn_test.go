package kahn

import "testing"

func TestKahn(t *testing.T) {
	g := InitGraph(5)
	g.AddEdge(2, 1)   //2先于1
	g.AddEdge(1, 0)   //1先于0
	g.AddEdge(0, 4)   //0先于4
	g.AddEdge(4, 3)   //4先于3
	TopoSortByKahn(g) //->2->1->0->4->3 结果符合预期
}
