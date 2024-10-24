package pov

type Tree struct {
	value    string
	children []*Tree
	par      *Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	me := &Tree{value, children, nil}
	for _, ch := range children {
		ch.par = me
	}
	return me
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	cur := find(from, tr)
	if cur == nil {
		return nil
	}
	curCopy := cur
	for cur.par != nil {
		cur.children = append(cur.children, cur.par)
		cur.par.children = removeNode(cur.par.children, cur)
		cur = cur.par
	}
	return curCopy
}

func removeNode(nodes []*Tree, node *Tree) []*Tree {
	for i, n := range nodes {
		if n == node {
			return append(nodes[:i], nodes[i+1:]...)
		}
	}
	return nodes
}

func find(target string, cur *Tree) *Tree {
	if cur.value == target {
		return cur
	}
	for _, ch := range cur.children {
		ans := find(target, ch)
		if ans != nil {
			return ans
		}
	}
	return nil
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	pov := tr.FromPov(from)
	if pov == nil {
		return nil
	}
	var f func(*Tree, []string) bool
	pathFinal := []string{pov.value}
	f = func(cur *Tree, path []string) bool {
		if cur == nil {
			return false
		}
		if cur.value == to {
			pathFinal = append(pathFinal, path...)
			return true
		}
		for _, ch := range cur.children {
			path = append(path, ch.value)
			if f(ch, path) {
				return true
			}
			path = path[:len(path)-1]
		}
		return false
	}
	if !f(pov, []string{}) {
		return nil
	}
	return pathFinal
}
