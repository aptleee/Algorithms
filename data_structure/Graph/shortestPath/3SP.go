package shortestPath

// dijkstra is for non-negative edges, allow cycles
// DAG is for DAG, but allow negative edges
// BellMan-Ford can be used for all situations


//func relax(G Graph, v int) {
//	for w := range G.adj(v) {
//		if disTo[w] >= disTo[v] + e.weight {
//			disTo[w] = disTo[v] + e.weight
//			edgeTo[w] = v
//			if (pq.contains(w)) pq.change(w, distTo[w]);
//			else                pq.insert(w, distTo[w]);
//		}
//	}
//}

//func dijkstra(G Graph, s int) {
//	edgeTo := make([]int, G.V)
//	pq := MinPQ(G.V, 0);
//	for i := 0; i < G.V; i++ {
//		disTo[i] = math.MaxInt32
//	}
//	disTo[s] = 0
//
//	pq.insert(s, disTo[s])
//	for !pq.Empty() {
//		relax(pq.delMin())
//	}
//}

// func DAG(G graph, s int) {
//	edgeTo := make([]int, G.V)
//	for i := 0; i < G.V; i++ {
//		disTo[i] = math.MaxInt32
//	}
//	disTo[s] = 0
//	top := topologicalSort(G)
// 	for i := range top.order() {
//		relax(i)
//	}


// func BellmanFold(G graph, s int) {
//
