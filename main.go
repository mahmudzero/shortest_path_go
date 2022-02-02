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

func find_shortest(graph [][2]int, start, end int) {
	var roots [][2]int

	for i := 0; i < len(graph); i++ {
		g := graph[i]
		if contains(g[:], start) {
			// i wanna be able to use push, just for the funzies, but append is built in so
			// for now we are gonna use that
			roots = append(roots, g)
		}
	}

	fmt.Println(roots)
}

func main() {
	graph := [][2]int{ { 1, 2 }, { 2, 3 }, { 3, 4 } }
	find_shortest(graph, 1, 3)
}
