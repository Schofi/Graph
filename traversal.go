package graph

import "fmt"

// DFS performs a depth-first search on the graph, starting from the given vertex. The visit
// function will be invoked with the hash of the vertex currently visited. If it returns false, DFS
// will continue traversing the graph, and if it returns true, the traversal will be stopped. In
// case the graph is disconnected, only the vertices joined with the starting vertex are visited.
//
// This example prints all vertices of the graph in DFS-order:
//
//	g := graph.New(graph.IntHash)
//
//	_ = g.AddVertex(1)
//	_ = g.AddVertex(2)
//	_ = g.AddVertex(3)
//
//	_ = g.AddEdge(1, 2)
//	_ = g.AddEdge(2, 3)
//	_ = g.AddEdge(3, 1)
//
//	_ = graph.DFS(g, 1, func(value int) bool {
//		fmt.Println(value)
//		return false
//	})
//
// Similarly, if you have a graph of City vertices and the traversal should stop at London, the
// visit function would look as follows:
//
//	func(c City) bool {
//		return c.Name == "London"
//	}
//
// DFS is non-recursive and maintains a stack instead.
func DFS[K comparable, T any](g Graph[K, T], start K, visit func(K) bool) error {
	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return fmt.Errorf("could not get adjacency map: %w", err)
	}

	if _, ok := adjacencyMap[start]; !ok {
		return fmt.Errorf("could not find start vertex with hash %v", start)
	}

	stack := make([]K, 0)
	visited := make(map[K]bool)

	stack = append(stack, start)

	for len(stack) > 0 {
		currentHash := stack[len(stack)-1]

		stack = stack[:len(stack)-1]

		if _, ok := visited[currentHash]; !ok {
			// Stop traversing the graph if the visit function returns true.
			if stop := visit(currentHash); stop {
				break
			}
			visited[currentHash] = true

			for adjacency := range adjacencyMap[currentHash] {
				stack = append(stack, adjacency)
			}
		}
	}

	return nil
}

// BFS performs a depth-first search on the graph, starting from the given vertex. The visit
// function will be invoked with the hash of the vertex currently visited. If it returns false, BFS
// will continue traversing the graph, and if it returns true, the traversal will be stopped. In
// case the graph is disconnected, only the vertices joined with the starting vertex are visited.
//
// This example prints all vertices of the graph in BFS-order:
//
//	g := graph.New(graph.IntHash)
//
//	_ = g.AddVertex(1)
//	_ = g.AddVertex(2)
//	_ = g.AddVertex(3)
//
//	_ = g.AddEdge(1, 2)
//	_ = g.AddEdge(2, 3)
//	_ = g.AddEdge(3, 1)
//
//	_ = graph.BFS(g, 1, func(value int) bool {
//		fmt.Println(value)
//		return false
//	})
//
// Similarly, if you have a graph of City vertices and the traversal should stop at London, the
// visit function would look as follows:
//
//	func(c City) bool {
//		return c.Name == "London"
//	}
//
// BFS is non-recursive and maintains a stack instead.
func BFS[K comparable, T any](g Graph[K, T], start K, count int, visit func(K) bool) error {
	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return fmt.Errorf("could not get adjacency map: %w", err)
	}

	if _, ok := adjacencyMap[start]; !ok {
		return fmt.Errorf("could not find start vertex with hash %v", start)
	}

	queue := make([]K, 0)
	visited := make(map[K]bool)

	visited[start] = true
	queue = append(queue, start)

	for len(queue) > 0 {
		currentHash := queue[0]

		queue = queue[1:]

		// Stop traversing the graph if the visit function returns true.
		if stop := visit(currentHash); stop {
			break
		}

		for adjacency := range adjacencyMap[currentHash] {
			if _, ok := visited[adjacency]; !ok {
				visited[adjacency] = true
				queue = append(queue, adjacency)
			}
		}

	}

	return nil
}

// BFSRecursive traverses g in breadth-first order starting at v.
// When the algorithm follows an edge (v, w) and finds a previously
// unvisited vertex w, it calls do(v, w, c) with c equal to
// the cost of the edge (v, w).

func BFSRecursive[K comparable, T any](g Graph[K, T], start K, do func(v, w K)) {
	visited := make(map[K]bool)
	visited[start] = true
	for queue := []K{start}; len(queue) > 0; {

		v := queue[0]
		queue = queue[1:]

		g.Visit(start, func(w K, c K) bool {
			if visited[w] {
				return false
			}
			do(v, w)

			visited[w] = true
			queue = append(queue, w)
			return false
		})

	}
}

func dfsHelper[K comparable, T any](g Graph[K, T], start K, visited map[K]bool) {
	visited[start] = true
	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return
	}

	edge, ok := adjacencyMap[start]
	if !ok {
		return
	}

	for i, _ := range edge {
		if visited[i] == false {
			dfsHelper(g, i, visited)
		}
	}
}

func DFSRecursive[K comparable, T any](g Graph[K, T], start K, do func(v, w K)) error {
	visited := make(map[K]bool)

	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return fmt.Errorf("could not get adjacency map: %w", err)
	}
	edge, ok := adjacencyMap[start]
	if !ok {
		return fmt.Errorf("could not find start vertex with hash %v", start)
	}

	for i, _ := range edge {
		if visited[i] == false {
			dfsHelper(g, i, visited)
			do(start, i)
		}
	}

	return nil
}

// https://faun.pub/implementing-recursive-and-iterative-dfs-on-a-binary-tree-golang-eda04949f4ee
func DFSTree() {

}

func BFSTree() {

}
