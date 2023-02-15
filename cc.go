package graph

//connected components

type CC[K comparable, T any] struct {
	Graph[K, T]
	Visited map[K]int
	Count   int
}

func dfsCc[K comparable, T any](g Graph[K, T], start K, ccNumber int, visited map[K]int) {
	queue := make([]K, 0)
	queue = append(queue, start)
	visited[start] = ccNumber
	for len(queue) > 0 {
		currentHash := queue[0]

		queue = queue[1:]
		adjacencyMap, err := g.AdjacencyMap()
		if err != nil {
			return
		}
		for adjacency := range adjacencyMap[currentHash] {
			if _, ok := visited[adjacency]; !ok {
				visited[adjacency] = ccNumber
				queue = append(queue, adjacency)
			}
		}

	}
}

func NewCC[K comparable, T any](g Graph[K, T]) (CC[K, T], error) {
	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return CC[K, T]{}, err
	}

	var ccNumber int

	visited := make(map[K]int)

	for start, _ := range adjacencyMap {
		dfsCc(g, start, ccNumber, visited)
		ccNumber++
	}
	if err != nil {
		return CC[K, T]{}, err
	}
	return CC[K, T]{Count: ccNumber, Graph: g, Visited: visited}, err
}

func (c *CC[K, T]) IsConnected(source, target K) bool {
	return c.Visited[source] == c.Visited[target]
}

func (c *CC[K, T]) Components() [][]K {
	components := make([][]K, 10)
	adjacencyMap, err := c.Graph.AdjacencyMap()
	if err != nil {
		return nil
	}
	for v, _ := range adjacencyMap {
		cNumber, ok := c.Visited[v]
		if !ok {
			continue
		}
		components[cNumber] = append(components[cNumber], v)
	}
	return components
}
