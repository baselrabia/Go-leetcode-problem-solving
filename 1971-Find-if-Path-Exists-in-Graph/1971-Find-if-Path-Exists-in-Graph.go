func validPath(n int, edges [][]int, source int, destination int) bool {

    graph := make(map[int][]int)

    for _, edge := range edges {
        u,v := edge[0],edge[1]
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }

    visited := make(map[int]bool)

    var dfs func(int)

    dfs = func (node int) {
        visited[node] = true

        for _,neighbor := range graph[node] {

            if !visited[neighbor]{
                dfs(neighbor)
            }
        }

    }

    dfs(source)
    return visited[destination]
    
}