package First_Search

import (
	"testing"
)

func TestBFS(t *testing.T) {
	graph = Constructor(8)
	graph.addEdge(0, 1)
	graph.addEdge(1, 2)
	graph.addEdge(0, 3)
	graph.addEdge(3, 4)
	graph.addEdge(1, 4)
	graph.addEdge(2, 5)
	graph.addEdge(4, 5)
	graph.addEdge(4, 6)
	graph.addEdge(5, 7)
	graph.addEdge(6, 7)

	bfs(0, 6) //终端执行命令看打印信息 go test -v  ./*.go
}
