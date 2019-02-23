package adjacentList

func FindDirectedCycle(G graph) bool {
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