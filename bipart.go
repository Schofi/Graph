package graph

func BipartiteDetection[K comparable, T any](g Graph[K, T]) bool {
	visited := make(map[K]bool)
	colors := make(map[K]int)
	isBipartite := false
	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return false
	}

	for source, _ := range adjacencyMap {
		colors[source] = -1
	}

	for source, _ := range adjacencyMap {
		if !visited[source] {
			if !dfsBipartiteDetection(g, source, 0, visited, colors) {
				isBipartite = true
				break
			}
		}
	}
	return isBipartite
}

func dfsBipartiteDetection[K comparable, T any](g Graph[K, T], start K, color int, visited map[K]bool, colors map[K]int) bool {
	visited[start] = true
	colors[start] = color
	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return false
	}

	for w, _ := range adjacencyMap[start] {
		if !visited[w] {
			if !dfsBipartiteDetection(g, w, 1-color, visited, colors) {
				return true
			} else if colors[w] == colors[start] {
				return true
			}
		}
	}
	return true
}
