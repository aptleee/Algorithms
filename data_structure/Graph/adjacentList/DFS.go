package adjacentList

import (
	"container/list"
)

type graph struct {
	V int
	E int
	adj []list.List
	id []int
}

func New(V, E int) graph {
	return graph{
		V, E, make([]list.List, V), make([]int, V),
	}
}

func dfsOrder() ([]int, []int, []int) {
	pre := make([]int, 1)  //  queue
	post := make([]int, 1) // queue
	G := graph{}
	marked := make([]int, G.V)
	var dfs func(v int)
	dfs = func(v int) {
		pre = append(pre, v)
		marked[v] = 1
		for w := range G.adj {
			if marked[w] == 0 {
				dfs(w)
			}
		}
		post = append(post, v)
	}

	for i := 0; i < G.V; i++ {
		if marked[i] == 0 {
			dfs(i)
		}
	}

	reversePost := make([]int, len(post))
	copy(reversePost, post)
	reverse(reversePost)
	return pre, post, reversePost

}

func reverse(A []int) {
	i, j := 0, len(A)-1
	for i < j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
}

func CC(G graph) int {
	count := 0
	marked := make([]int, G.V)
	var dfs func(i int)
	dfs = func(i int) {
		marked[i] = 1
		G.id[i] = count
		for _, x := range G.adj[i] {
			if marked[x] == 0 {
				dfs(x)
			}
		}
	}
	for i := 0; i < G.V; i++ {
		if marked[i] == 0 {
			dfs(i)
			count++
		}
	}
	return count
}

func (G *graph) connected(v, w int) bool {
	return G.id[v] == G.id[w]
}

func twoColor(G graph) bool {
	color := make([]bool, G.V)
	marked := make([]int, G.V)
	isTwoColor := true
	var dfs func(i int)
	dfs = func(i int)  {
		marked[i] = 1
		for _, x := range G.adj[i] {
			if marked[x] == 0 {
				color[x] = !color[i]
				dfs(x)
			} else if color[i] == color[x] {
				isTwoColor = false
			}
		}
	}
	for i := 0; i < G.V; i++ {
		if marked[i] == 0 {
			dfs(i)
		}
	}
	return isTwoColor
}


func findDirectedCycle(G graph) bool {
	marked := make([]bool, G.V)
	onStack := make([]bool, G.V)
	hasCycle := false
	var dfs func(s int)
	dfs = func(s int) {
		onStack[s] = true
		marked[s] = true
		for _, x := range G.adj[s] {
			if !marked[x]{
				dfs(x)
			} else if onStack[x] {
				hasCycle = true
			}
		}
		onStack[s] = false
	}

	for i := 0; i < G.V; i++ {
		if !marked[i] {
			dfs(i)
		}
	}
	return hasCycle

}