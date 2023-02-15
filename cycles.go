package graph

func CycleDetection[K comparable, T any](g Graph[K, T]) bool {
	visited := make(map[K]bool)
	cycleFlag := false
	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return false
	}
	for source, _ := range adjacencyMap {
		if !visited[source] {
			if dfsDetection(g, source, source, visited) {
				cycleFlag = true
				break
			}
		}
	}
	return cycleFlag
}

func dfsDetection[K comparable, T any](g Graph[K, T], start K, parent K, visited map[K]bool) bool {
	visited[start] = true
	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return false
	}

	for w, _ := range adjacencyMap[start] {
		if !visited[w] {
			if dfsDetection(g, w, start, visited) {
				return true
			} else if w != parent {
				return true
			}
		}
	}
	return false
}
