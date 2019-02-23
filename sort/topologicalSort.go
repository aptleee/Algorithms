package sort

import "container/list"

type graph struct {
	V int
	E int
	adj []list.List

}

func dfs(G graph) {
	if findDirectedCycle(G) {
		return
	}
	marked := make([]bool, G.V)
	ret := make([]int, 0)
	var dfs func(s int)
	dfs = func(s int) {
		marked[s] = true
		for _, x := range G.adj[s] {
			if !marked[x] {
				dfs(x)
			}
		}
		ret = append(ret, s)
	}
	for  i := 0; i < G.V; i++ {
		if !marked[i] {
			dfs(i)
		}
	}
}

func bfs(G graph) {}

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