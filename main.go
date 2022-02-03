package main
import "fmt"

func push(arr [][]int, val []int) [][]int {
	new_arr := make([][]int, len(arr) + 1)
	for i := 0; i < len(arr); i++ {
		new_arr[i] = arr[i]
	}
	new_arr[len(arr)] = val

	return new_arr
}

func contains(arr []int, val int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return true
		}
	}

	return false
}

func find_shortest(graph [][2]int, start, end int) int {
	var roots    [][2]int
	var roots_idx []int
	var ends     [][2]int
	var ends_idx []int

	for i := 0; i < len(graph); i++ {
		g := graph[i]
		// i wanna be able to use push, just for the funzies, but append is built in so
		// for now we are gonna use that
		if contains(g[:], start) { roots = append(roots, g); roots_idx = append(roots_idx, i) }
		if contains(g[:], end)   { ends  = append(ends, g);  ends_idx  = append(ends_idx, i)  }
	}

	if len(roots) == 0 || len(ends) == 0 { return -1 }

	cnt := 0

	for i := 0; i < len(roots); i++ {
		root := roots[i]
		if root == [2]int{ start, end } || root == [2]int{ end, start } { return 0 }

		new_graph := graph
		new_graph[roots_idx[i]] = new_graph[len(graph) - 1]
		new_graph = new_graph[:len(graph) - 1]

		cnt += (find_shortest(new_graph, root[1], end) + 1)
	}

	return cnt
}

func main() {
	graph := [][2]int{ { 1, 2 }, { 2, 3 }, { 3, 4 }, { 4, 5 } }
	fmt.Println(find_shortest(graph, 1, 5))
}
