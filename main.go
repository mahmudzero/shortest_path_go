package main
import "fmt"

func contains(arr []int, val int) bool {
	var _contains bool = false;
	for i, v := range arr {
		_contains = _contains || (v == val);
	}

	return _contains;
}

func find_shortest(graph [][2]int, start, end int) {
	var roots [][2]int;

	var roots_i int = 0
	for i, g := range graph {
		fmt.Println(i, g);
		if contains(g, start) {
			roots[roots_i] = g;
			roots_i++;
		}
	}
}

func main() {
	graph := [][2]int{ { 1, 2 }, { 2, 3 }, { 3, 1 } };
	find_shortest(graph, 2, 3);
}
