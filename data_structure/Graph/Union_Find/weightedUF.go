package uf

type weightedUF struct {
	rank []int
	parent []int
	count int
}

func New1(n int) *weightedUF {
	wuf := &weightedUF {
		count: n,
		rank: make([]int, n),
		parent: make([]int, n),
	}
	for i := range wuf.parent {
		wuf.parent[i] = i
	}
	return wuf
}

func (wuf *weightedUF) Find(i int) int{
	for wuf.parent[i] != i {
		i = wuf.parent[i]
	}
	return i
}


func (wuf *weightedUF) pcFind(i int) int{
	if wuf.parent[i] != i {
		wuf.parent[i] = wuf.pcFind(wuf.parent[i])
	}
	return wuf.parent[i]
}


func (wuf *weightedUF) Union(i, j int) {
	pi := wuf.Find(i)
	pj := wuf.Find(j)
	if pi == pj {
		return
	}

	if wuf.rank[pi] > wuf.rank[pj] {
		wuf.parent[pj] = pi
	} else {
		wuf.parent[pi] = pj
		if wuf.rank[pi] == wuf.rank[pj] {
			wuf.rank[pj]++
		}
	}
}


