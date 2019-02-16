package adjacentList


func bfs(G graph, source int){
	q := make([]int, 1)
	marked := make([]int, G.V)
	edgeTo := make([]int, G.V)
	q = append(q, source)
	marked[source] = 1
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, w := range G.adj[x] {
			if marked[w] == 0 {
				marked[w] = 1
				edgeTo[w] = x
				q = append(q, w)
			}
		}
	}

}