package graph

type UnionFind[K comparable, T any] struct {
	count   int
	parents map[K]K
}

func (u *UnionFind[K, T]) Find(x K) K {
	if v, ok := u.parents[x]; !ok {
		return v
	}
	if x == u.parents[x] {
		return x
	}
	u.parents[x] = u.Find(u.parents[x])
	return u.parents[x]
}

func (u *UnionFind[K, T]) Union(x, y K) {
	rx, ry := u.Find(x), u.Find(y)
	if ry != rx {
		u.parents[rx] = ry
		u.count--
	}
}

func (u *UnionFind[K, T]) isConnected(x, y K) bool {
	return u.Find(x) == u.Find(y)
}
