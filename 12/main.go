package main

import (
	"fmt"
	"os"
	"strings"
	//"unicode/utf8"
)

type Node struct {
	Neighbors []Node
	Height int
	Pos Pos
	Dist int
	End bool
}

type Pos struct {
	x int
	y int
}

func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/12/input.txt")

	//nodes := make([]Node, 0)

	grid, n_rows, n_cols, start, stop := make_grid(file)

	fmt.Printf("Rows: %v, Cols: %v, start: %v, stop: %v\n", n_rows, n_cols, start, stop)

	fill_neighbors(grid, n_cols, n_rows)

	// for _, r := range grid {
	// 	for _, c := range r {
	// 		fmt.Printf("pos: %v, ns: %v\n",c.Pos, len(c.Neighbors))
	// 	}
	// }
	start_node := grid[start.x][start.y]
	//fmt.Println(start_node)
	fmt.Println(grid[stop.x][stop.y].Height)
	path := find_path(grid, start_node)
	
	//print_grid(grid, false)
	fmt.Println(len(path))
	
	fmt.Println(grid[stop.x][stop.y].Dist - 2)
	

}

// breadth first search to find the shortest path to stop_node
func find_path(grid [][]Node, start_node Node) []Node {
	q := make([]Node, 0)
	q = append(q, start_node)
	grid[start_node.Pos.x][start_node.Pos.y].Dist = 0

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, n := range curr.Neighbors {
			if grid[n.Pos.x][n.Pos.y].Dist == -1 {
				grid[n.Pos.x][n.Pos.y].Dist = grid[curr.Pos.x][curr.Pos.y].Dist + 1
				q = append(q, grid[n.Pos.x][n.Pos.y])
			}
		}
	}

	return q
}

func make_grid(file []uint8) (grid [][]Node, n_rows int, n_cols int, start Pos, stop Pos) {
	for _, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {

		n_cols = len(row)
		n_rows += 1
	}

	grid = make([][]Node, n_rows)
	for i := range grid {
		grid[i] = make([]Node, n_cols)
	}

	for i, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		for j, c := range row { 
			grid[i][j] = Node{}
			if string(c) == "S" { 
				start = Pos{x: i, y: j} 
				//grid[i][j] = 0
				grid[i][j].Height = 0
			} else if string(c) == "E" { 
				stop = Pos{x: i, y: j} 
				grid[i][j].Height = 26
				grid[i][j].End = true
			} else {
				//grid[i][j] = int(c-97)
				grid[i][j].Height = int(c-97)
			}
			grid[i][j].Pos = Pos{x: i, y: j}
			grid[i][j].Dist = -1

		}
	}

	return grid, n_rows, n_cols, start, stop
}

func print_grid(grid [][]Node, do_string bool) {
	for _, r := range grid {
		for _, c := range r {
			if do_string { 
				fmt.Printf("%v ", string(uint8(c.Height+97))) 
			} else {
				//fmt.Printf("%v ", c.Height)
				//fmt.Printf("%v ", c.Dist)
				if !c.End {
					fmt.Printf("%v ", "X")
				} else {
					fmt.Printf("%v ", ".")
				}
			}
		}
		fmt.Println()
		
	}
}

func fill_neighbors(grid [][]Node, n_cols int, n_rows int) {
	for i, r := range grid {
		for j, _ := range r {
			grid[i][j].Neighbors = append(grid[i][j].Neighbors, get_neighbors(grid, grid[i][j], n_cols, n_rows)...)
		}
	}
}

func get_neighbors(grid [][]Node, p Node, n_cols int, n_rows int) (ns []Node) {
	x, y := p.Pos.x, p.Pos.y

	//up
	if x - 1 >= 0 && grid[x-1][y].Height - 1 <= p.Height {
		ns = append(ns, grid[x-1][y])
	}
	//right
	if y + 1 < n_cols && grid[x][y+1].Height - 1 <= p.Height {
		ns = append(ns, grid[x][y+1])
	}
	//down
	if x + 1 < n_rows && grid[x+1][y].Height - 1 <= p.Height {
		ns = append(ns, grid[x+1][y])
	}
	//left
	if y - 1 >= 0 && grid[x][y-1].Height - 1 <= p.Height {
		ns = append(ns, grid[x][y-1])
	}
	return ns
}