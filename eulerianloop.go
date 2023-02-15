package graph

func EulerLoop[K comparable, T any](g Graph[K, T]) bool {
	var cc *CC[K, T]
	cc = new(CC[K, T])
	if cc.Count > 1 {
		return false
	}
	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return false
	}
	for source, _ := range adjacencyMap {
		if len(adjacencyMap[source])%2 == 1 {
			return false
		}
	}
	return true
}
