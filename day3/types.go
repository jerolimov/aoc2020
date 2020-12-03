package main

type slope struct {
	x int
	y int
}

type area struct {
	tree [][]bool
}

func (a *area) addRow() {
	a.tree = append(a.tree, make([]bool, 0))
}

func (a *area) addFree() {
	a.tree[len(a.tree)-1] = append(a.tree[len(a.tree)-1], false)
}

func (a *area) addTree() {
	a.tree[len(a.tree)-1] = append(a.tree[len(a.tree)-1], true)
}

func (a *area) rows() int {
	return len(a.tree)
}

func (a *area) isTree(x int, y int) bool {
	return a.tree[y][x%len(a.tree[y])]
}

func (a area) String() string {
	output := ""
	for _, row := range a.tree {
		for _, tree := range row {
			if tree {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output
}
